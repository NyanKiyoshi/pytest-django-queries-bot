package ghevents

import (
	"pytest-queries-bot/pytestqueries/github/awstypes"
)

const GithubEventHeader string = "X-GitHub-Event"

// Dispatch handles supported events. See https://developer.github.com/v3/activity/events/types/
// for the list of available events and data structures--or https://godoc.org/github.com/google/go-github/github.
func Dispatch(request *awstypes.Request) (awstypes.Response, error) {
	event := request.Headers[GithubEventHeader]

	switch event {
	case "pull_request":
		return pullrequest(request)
	case "ping":
		return ping()
	case "push":
		return ping()
	default:
		return unknown()
	}
}
