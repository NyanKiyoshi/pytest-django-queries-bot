package ghevents

import (
	"encoding/json"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"github.com/google/go-github/github"
)

func synchronizePR() {

}

func pullrequest(request *awstypes.Request) (awstypes.Response, error) {
	var payload github.PullRequestEvent

	if err := json.Unmarshal([]byte(request.Body), payload); err != nil || payload.Action == nil {
		return awstypes.Response{StatusCode: 400}, err
	}

	switch *payload.Action {
	case "synchronize":
		synchronizePR()
		break
	default:
		return unknown()
	}

	return awstypes.Response{
		StatusCode: 200,
	}, nil
}
