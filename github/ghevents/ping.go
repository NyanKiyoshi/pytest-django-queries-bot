package ghevents

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
)

func ping() (awstypes.Response, error) {
	return awstypes.Response{
		StatusCode: 200,
		Body:       "{\"status\": 200}",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
