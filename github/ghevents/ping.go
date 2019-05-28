package ghevents

import (
	"pytest-queries-bot/pytestqueries/github/awstypes"
)

func ping() (awstypes.Response, error) {
	return awstypes.Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "{\"status\": 200}",
		Headers: map[string]string{
			"Content-Type":           "application/json",
		},
	}, nil
}
