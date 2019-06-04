package ghevents

import (
	"encoding/json"
	"github.com/google/go-github/github"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/models"
	"time"
)

func push(request *awstypes.Request) (awstypes.Response, error) {
	var payload github.PushEvent
	var err error

	if err = json.Unmarshal([]byte(request.Body), &payload); err != nil {
		return awstypes.Response{StatusCode: 400}, err
	}

	commitHash := *payload.HeadCommit.ID
	event, err := models.RetrieveEvent(commitHash)

	if err != nil {
		event.HashSHA1 = commitHash
		event.EntryDate = time.Now()
		err := models.EventTable().Put(event).Run()

		if err != nil {
			return awstypes.Response{
				StatusCode: 500,
				Body:       `{"status": 500, "message": "creation failed"}`,
			}, err
		}

		return awstypes.Response{
			StatusCode: 201,
			Body:       `{"status": 201}`,
		}, nil
	}

	return awstypes.Response{
		StatusCode: 200,
		Body:       `{"status": 200}`,
	}, nil
}
