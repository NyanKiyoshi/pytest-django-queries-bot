package main

import (
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/generated"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var Client = http.Client{
	Transport: &http.Transport{
		TLSHandshakeTimeout: 10 * time.Second,
	},
	Timeout: 30 * time.Second,
}

func sendReq(meth string, url, contentType string, reader io.Reader, headers *map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(meth, url, reader)

	if err != nil {
		return nil, fmt.Errorf("failed to create the request: %s", err)
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}

	req.Header.Add("Content-Type", contentType)
	return Client.Do(req)
}

func getAccessToken(accessCode string) {
	values := url.Values{
		"client_id":     []string{generated.GithubClientId},
		"client_secret": []string{generated.GithubClientSecret},
		"code":          []string{accessCode},
	}
	resp, err := sendReq(
		"POST",
		"https://github.com/login/oauth/access_token?"+values.Encode(),
		"text/plain",
		nil,
		&map[string]string{
			"Accept": "application/xml",
		})
	if err != nil {
		log.Fatal(err)
	}

	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		log.Print(resp.Status)
		log.Printf("%s", body)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	url_ := "https://github.com/login/oauth/authorize?" + url.Values{
		"client_id":    []string{generated.GithubClientId},
		"redirect_url": []string{"http://127.0.0.1:5999/installed"},
	}.Encode()
	http.Redirect(w, r, url_, http.StatusTemporaryRedirect)
}

func handleAuthorized(w http.ResponseWriter, r *http.Request) {
	url_, err := url.Parse(r.RequestURI)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	code := url_.Query().Get("code")

	if code == "" {
		http.Error(w, "Missing auth code", http.StatusBadRequest)
	}

	getAccessToken(code)
	http.Error(w, "OK", http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/installed", handleAuthorized)
	log.Print("Serving on http://127.0.0.1:5999")
	panic(http.ListenAndServe(":5999", mux))
}
