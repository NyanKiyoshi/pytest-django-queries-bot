package ghevents

import (
	"encoding/json"
	"errors"
	"github.com/google/go-github/github"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/models"
)

func synchronizePR(data *github.PullRequest) error {
	if data.Head.SHA == nil {
		return errors.New("pull request hash missing")
	}

	event := models.Event{HashSHA1: *data.Head.SHA}
	return models.EventTable().Put(event).Run()
}

func pullrequest(request *awstypes.Request) (awstypes.Response, error) {
	var payload github.PullRequestEvent
	var err error

	if err = json.Unmarshal([]byte(request.Body), &payload); err != nil ||
		payload.Action == nil {
		return awstypes.Response{StatusCode: 400}, err
	}

	switch *payload.Action {
	case "opened":
	case "synchronize":
		err = synchronizePR(payload.PullRequest)
		break
	default:
		return unknown("pull_request action", payload.Action)
	}

	if err != nil {
		return awstypes.Response{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}

	return awstypes.Response{
		StatusCode: 200,
		Body:       `{"status": 200}`,
	}, nil
}
