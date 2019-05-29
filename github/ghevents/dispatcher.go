package ghevents

import (
	"pytest-queries-bot/pytestqueries/github/awstypes"
)

const GithubEventHeader string = "X-GitHub-Event"

func Dispatch(request *awstypes.Request) (awstypes.Response, error) {
	event := request.Headers[GithubEventHeader]

	switch event {
	case "pull_request":
		return pullrequest(request)
	case "ping":
		return ping()
	default:
		return unknown()
	}
}
