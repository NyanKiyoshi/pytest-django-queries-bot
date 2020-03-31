package ghevents

import (
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
)

func unknown(expectedEventType string, eventname *string) (awstypes.Response, error) {
	return awstypes.Response{
		StatusCode: 405,
		Body:       fmt.Sprintf(`{"message": "Unknown or unsupported %s: %s"}`, expectedEventType, *eventname),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
