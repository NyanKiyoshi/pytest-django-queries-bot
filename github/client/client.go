package client

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"pytest-queries-bot/pytestqueries/generated"
)

func GetClient() (*github.Client, *context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: generated.GithubAccessToken,
	})
	tc := oauth2.NewClient(ctx, ts)
	gh := github.NewClient(tc)
	return gh, &ctx
}