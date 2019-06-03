// +build diff

package main

import (
	"bufio"
	"flag"
	"github.com/caarlos0/env"
	"log"
	"net/url"
	"os"
	"pytest-queries-bot/pytestqueries/ci-tools/utils"
)

type config struct {
	DiffEndpoint url.URL `env:"DIFF_ENDPOINT,required"`
	SHA1Revision string
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

	targetURL := cfg.DiffEndpoint.String()
	contentType := "text/plain"

	if _, err := utils.SendUploadRequest(
		targetURL,
		contentType,
		bufio.NewReader(os.Stdin),
		&map[string]string{
			"X-Commit-Rev": cfg.SHA1Revision},
	); err != nil {
		print(err)
		os.Exit(1)
	}
}
