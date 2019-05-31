// +build upload

package main

import (
	"bufio"
	"github.com/caarlos0/env"
	"log"
	"net/url"
	"os"
	"pytest-queries-bot/pytestqueries/ci-tools/utils"
)

type config struct {
	UploadEndpoint url.URL `env:"UPLOAD_ENDPOINT,required"`
	SecretKey      string  `env:"SECRET_UPLOAD_KEY,required"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	targetURL := cfg.UploadEndpoint.String()
	contentType := "application/json"

	utils.SendUploadRequest(
		targetURL,
		contentType,
		bufio.NewReader(os.Stdin),
		&map[string]string{"X-Secret-Key": cfg.SecretKey},
	)
}
