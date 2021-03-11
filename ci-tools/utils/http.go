package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var Client = http.Client{
	Transport: &http.Transport{
		TLSHandshakeTimeout: 10 * time.Second,
	},
	Timeout: 30 * time.Second,
}

func HttpDo(method, url string, reader io.Reader, headers *map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, reader)

	if err != nil {
		return nil, fmt.Errorf("failed to create the request: %w", err)
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}

	resp, err := Client.Do(req)
	return resp, err
}

func SendUploadRequest(url string, reader io.Reader, headers *map[string]string) (*http.Response, error) {
	resp, err := HttpDo("POST", url, reader, headers)

	if err != nil {
		return nil, fmt.Errorf("failed to upload: %w", err)
	}

	if resp != nil {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("got '%s': %s", resp.Status, body)
	}

	return nil, errors.New("we got no response and no error")
}
