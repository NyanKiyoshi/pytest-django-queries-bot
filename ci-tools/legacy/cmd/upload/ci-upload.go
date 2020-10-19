// +build upload

package main

import (
	"flag"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/upstream"
	"github.com/caarlos0/env"
	"log"
	"net/url"
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

	input := upstream.RawReportInput{
		PostCommentUrl: cfg.UploadEndpoint.String(),
		SecretKey: cfg.SecretKey,
	}

	flag.StringVar(&input.SHA1Revision, "rev", "", "The SHA1 commit revision")
	flag.Parse()

	if _, err := input.PostFromStdin(); err != nil {
		log.Fatal(err)
	}
}
