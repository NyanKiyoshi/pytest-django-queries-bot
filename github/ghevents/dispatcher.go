package ghevents

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
)

const GithubEventHeader string = "x-github-event"
const GithubEventCanonicalMIMEHeader string = "X-Github-Event"

// Dispatch handles supported events. See https://developer.github.com/v3/activity/events/types/
// for the list of available events and data structures--or https://godoc.org/github.com/google/go-github/v32/github.
func Dispatch(request *awstypes.Request) (awstypes.Response, error) {
	event, ok := request.Headers[GithubEventHeader]
	if !ok {
		event = request.Headers[GithubEventCanonicalMIMEHeader]
	}

	switch event {
	case "pull_request":
		return pullrequest(request)
	case "ping":
		return ping()
	case "push":
		return push(request)
	default:
		return unknown("event", &event)
	}
}
