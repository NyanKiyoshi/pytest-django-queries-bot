package utils

import (
	"io"
	"log"
	"net/http"
	"time"
)

var Client = http.Client{
	Transport: &http.Transport{
		TLSHandshakeTimeout: 10 * time.Second,
	},
	Timeout: 30 * time.Second,
}

func SendUploadRequest(url, contentType string, reader io.Reader, headers *map[string]string) *http.Response {

	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		log.Fatalf("Failed to create the request: %s", err.Error())
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}

	req.Header.Add("Content-Type", contentType)

	resp, err := Client.Do(req)

	if err != nil {
		log.Fatalf("Failed to upload: %s", err.Error())
	}

	if resp != nil {
		log.Printf("Uploaded and got HTTP %d (%s)", resp.StatusCode, resp.Status)
		return resp
	}

	log.Fatal("We got no response and no error... Feels ignored.")
	return nil
}
