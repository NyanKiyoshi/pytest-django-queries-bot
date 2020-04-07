// +build upload

package main

import (
	"bufio"
	"flag"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/utils"
	"github.com/caarlos0/env"
	"log"
	"net/url"
	"os"
)

type config struct {
	UploadEndpoint url.URL `env:"UPLOAD_ENDPOINT,required"`
	SecretKey      string  `env:"SECRET_UPLOAD_KEY,required"`
	SHA1Revision   string
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	flag.StringVar(&cfg.SHA1Revision, "rev", "", "The SHA1 commit revision")
	flag.Parse()

	if len(cfg.SHA1Revision) != 40 {
		flag.Usage()
		os.Exit(1)
	}

	targetURL := cfg.UploadEndpoint.String()
	contentType := "application/json"

	if _, err := utils.SendUploadRequest(
		targetURL,
		contentType,
		bufio.NewReader(os.Stdin),
		&map[string]string{
			"X-Secret-Key": cfg.SecretKey,
			"X-Commit-Rev": cfg.SHA1Revision},
	); err != nil {
		log.Fatal(err)
	}
}
