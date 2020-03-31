package client

import (
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
	"net/http"
)

const GitHubAppID = int64(59286)

func GetClient(installationID int64) (*github.Client, *context.Context) {
	ctx := context.Background()
	ts, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, GitHubAppID, installationID, "private-key.pem")
	if err != nil {
		panic(fmt.Errorf("failed to create JWT token: %v", err))
	}
	gh := github.NewClient(&http.Client{
		Transport: ts,
	})
	return gh, &ctx
}
