package main

import (
	"bufio"
	"github.com/caarlos0/env"
	"log"
	"net/http"
	"net/url"
	"os"
)

type config struct {
	RapportBaseUrl url.URL `env:"RAPPORT_BASE_URL,required"`
	UploadEndpoint url.URL `env:"UPLOAD_ENDPOINT,required"`
	DiffEndpoint   url.URL `env:"DIFF_ENDPOINT,required"`
	SecretKey      string  `env:"SECRET_UPLOAD_KEY"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	var targetURL string
	var contentType string

	if cfg.SecretKey != "" {
		targetURL = cfg.UploadEndpoint.String()
		contentType = "application/json"
	} else {
		targetURL = cfg.DiffEndpoint.String()
		contentType = "text/plain"
	}

	reader := bufio.NewReader(os.Stdin)
	resp, err := http.Post(targetURL, contentType, reader)

	if err != nil {
		log.Fatalf("failed to upload: %s", err.Error())
		return
	}

	if resp != nil {
		log.Printf("Uploaded and got HTTP %d (%s)", resp.StatusCode, resp.Status)
		return
	}

	log.Fatal("We got no response and no error... Feels ignored.")
}
