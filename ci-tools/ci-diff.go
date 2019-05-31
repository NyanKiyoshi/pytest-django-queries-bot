// +build diff

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
	DiffEndpoint url.URL `env:"DIFF_ENDPOINT,required"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	targetURL := cfg.DiffEndpoint.String()
	contentType := "text/plain"

	utils.SendUploadRequest(
		targetURL,
		contentType,
		bufio.NewReader(os.Stdin),
		nil,
	)
}
