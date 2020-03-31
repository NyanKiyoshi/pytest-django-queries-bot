package ghevents

import (
	"encoding/json"
	"errors"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/models"
	"github.com/google/go-github/github"
	"time"
)

func synchronizePR(data *github.PullRequest) (*awstypes.Response, error) {
	if data.Head.SHA == nil {
		return nil, errors.New("pull request hash missing")
	}

	event := models.Event{
		HashSHA1:      *data.Head.SHA,
		EntryDate:     time.Now(),
		PullRequestID: *data.ID,
	}
	err := models.EventTable().Put(event).Run()

	if err != nil {
		return nil, err
	}

	pr := &models.PullRequest{}
	err = models.PullRequestTable().Get("PullRequestID", *data.ID).One(pr)

	if pr.PullRequestNumber == 0 {
		pr := models.PullRequest{
			PullRequestID:     *data.ID,
			PullRequestNumber: *data.Number,
			OwnerName:         *data.Base.Repo.Owner.Login,
			RepoName:          *data.Head.Repo.Name,
			EntryDate:         time.Now(),
		}
		err = models.PullRequestTable().Put(pr).Run()
	}

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
