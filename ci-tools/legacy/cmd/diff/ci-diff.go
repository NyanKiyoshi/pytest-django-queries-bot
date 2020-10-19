// +build diff

package main

import (
	"flag"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/upstream"
	"github.com/caarlos0/env"
	"log"
	"net/url"
)

type config struct {
	DiffEndpoint url.URL `env:"DIFF_ENDPOINT,required"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	input := upstream.DiffInput{
		PostCommentUrl: cfg.DiffEndpoint.String(),
	}

	flag.StringVar(&input.SHA1Revision, "rev", "", "The SHA1 commit revision")
	flag.Parse()

	if _, err := input.PostFromStdin(); err != nil {
		log.Fatal(err)
	}
}
