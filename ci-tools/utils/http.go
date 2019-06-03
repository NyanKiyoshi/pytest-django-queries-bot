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

func SendUploadRequest(url, contentType string, reader io.Reader, headers *map[string]string) (*http.Response, error) {

	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		return nil, fmt.Errorf("failed to create the request: %s", err)
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}

	req.Header.Add("Content-Type", contentType)

	resp, err := Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to upload: %s", err)
	}

	if resp != nil {
		if resp.StatusCode > 299 || resp.StatusCode < 200 {
			return resp, nil
		}
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("got %s: %s", resp.Status, body)
	}

	return nil, errors.New("we got no response and no error... Feels ignored")
}
