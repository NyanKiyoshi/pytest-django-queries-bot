package utils

import (
	"io"
	"log"
	"net/http"
)

func SendUploadRequest(url, contentType string, reader io.Reader, headers *map[string]string) {
	req, err := http.NewRequest(url, contentType, reader)

	if err != nil {
		log.Fatalf("Failed to create the request: %s", err.Error())
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Failed to upload: %s", err.Error())
	}

	if resp != nil {
		log.Printf("Uploaded and got HTTP %d (%s)", resp.StatusCode, resp.Status)
	}

	log.Fatal("We got no response and no error... Feels ignored.")
}
