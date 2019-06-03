package models

import (
	"errors"
	"github.com/guregu/dynamo"
	"pytest-queries-bot/pytestqueries/db"
	"pytest-queries-bot/pytestqueries/generated"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/consts"
	"time"
)

type Event struct {
	// HashSHA1 contains the git pull request
	// event's SHA1 commit hash
	HashSHA1 string

	// EntryCreationDate is the date where this entry was created/updated.
	// Not to be confused with the GitHub event date.
	EntryDate time.Time `dynamodbav:"_date"`

	// HasRapport is true if we already have a
	// JSON report uploaded for that hash
	HasRapport bool `dynamodbav:"has_rapport"`

	// DiffUploaded is true if we already have a
	// diff report sent for that hash
	DiffUploaded bool `dynamodbav:"diff_was_uploaded"`

	// PullRequestID is the pull request ID in GitHub.
	// This will be used to comment with results into GitHub.
	PullRequestID int64 `dynamodbav:"pr_id"`
}

func EventTable() dynamo.Table {
	return db.Get().Table(generated.DynamoEventsTableName)
}

// RetrieveEvent gets the event associated to a given hash
// to ensure the request is correct and expected
func RetrieveEvent(commitHash string) (*Event, error) {
	event := &Event{}
	return event, EventTable().
		Get("HashSHA1", commitHash).
		Filter("HasRapport = ?", false).
		One(event)
}

func CheckEvent(request *awstypes.Request) (*Event, *PullRequest, error) {
	// Check if the received body is not too large
	if len(request.Body) > consts.MaxUploadSize {
		return nil, nil, errors.New("body is too big")
	}

	commitHash, ok := request.Headers[consts.CommitHashHeaderName]
	if !ok {
		commitHash = request.Headers[consts.CommitHashHeaderNameLower]
	}
	if len(commitHash) != consts.SHA1Length {
		return nil, nil, errors.New("invalid or missing SHA1 commit hash")
	}

	event, err := RetrieveEvent(commitHash)
	if err != nil {
		return nil, nil, err
	}

	pr := &PullRequest{}
	err = PullRequestTable().Get("PullRequestID", event.PullRequestID).One(pr)

	if err != nil {
		return nil, nil, err
	}

	return event, pr, nil
}
