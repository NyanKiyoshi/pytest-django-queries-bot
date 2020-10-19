package client

import (
	"context"
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/config"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v32/github"
	"net/http"
)

func GetClient(installationID int64) (*github.Client, *context.Context) {
	ctx := context.Background()
	ts, err := ghinstallation.NewKeyFromFile(
		http.DefaultTransport, config.GithubAppId, installationID, "private-key.pem",
	)

	if err != nil {
		panic(fmt.Errorf("failed to create JWT token: %v", err))
	}

	gh := github.NewClient(&http.Client{
		Transport: ts,
	})
	return gh, &ctx
}
