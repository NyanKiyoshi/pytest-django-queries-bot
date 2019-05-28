package ghevents

import (
	"pytest-queries-bot/pytestqueries/github/awstypes"
)

func Dispatch(request *awstypes.Request) (awstypes.Response, error) {
	event := request.Headers["X-Github-Event"]

	switch event {
	case "pull_request":
		return pullrequest(request)
	case "ping":
		return ping()
	default:
		return unknown()
	}
}
