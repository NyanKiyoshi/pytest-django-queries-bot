package ghevents

import (
	"encoding/json"
	"errors"
	"github.com/google/go-github/github"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/models"
	"time"
)

func synchronizePR(data *github.PullRequest) (*awstypes.Response, error) {
	if data.Head.SHA == nil {
		return nil, errors.New("pull request hash missing")
	}

	event := models.Event{HashSHA1: *data.Head.SHA, EntryDate: time.Now()}
	err := models.EventTable().Put(event).Run()

	if err != nil {
		return nil, err
	}

	return &awstypes.Response{
		StatusCode: 201,
		Body:       `{"status": 201, "message": "entry successfully created"}`,
	}, nil
}

func pullrequest(request *awstypes.Request) (awstypes.Response, error) {
	var payload github.PullRequestEvent
	var response *awstypes.Response
	var err error

	if err = json.Unmarshal([]byte(request.Body), &payload); err != nil ||
		payload.Action == nil {
		return awstypes.Response{StatusCode: 400}, err
	}

	switch *payload.Action {
	case "opened", "synchronize":
		response, err = synchronizePR(payload.PullRequest)
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

	if response != nil {
		return *response, nil
	}

	return awstypes.Response{
		StatusCode: 200,
		Body:       `{"status": 200}`,
	}, nil
}
