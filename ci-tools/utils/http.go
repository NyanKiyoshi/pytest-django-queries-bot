package utils

import (
	"io"
	"log"
	"net/http"
	"time"
)

func SendUploadRequest(url, contentType string, reader io.Reader, headers *map[string]string) {
	client := http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest(url, contentType, reader)

	if err != nil {
		log.Fatalf("Failed to create the request: %s", err.Error())
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Failed to upload: %s", err.Error())
	}

	if resp != nil {
		log.Printf("Uploaded and got HTTP %d (%s)", resp.StatusCode, resp.Status)
	}

	log.Fatal("We got no response and no error... Feels ignored.")
}
