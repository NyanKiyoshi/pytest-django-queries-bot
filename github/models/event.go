package models

import (
	"crypto/sha1"
	"errors"
	"github.com/guregu/dynamo"
	"pytest-queries-bot/pytestqueries/db"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/consts"
)

const tableName string = "gh_hooks_events"

type Event struct {
	// HashSHA1 contains the git pull request
	// event's SHA1 commit hash
	HashSHA1 string `dynamodbav:"sha1_hash"`

	// HasRapport is true if we already have a
	// JSON report uploaded for that hash
	HasRapport bool `dynamodbav:"has_rapport"`

	// DiffUploaded is true if we already have a
	// diff report sent for that hash
	DiffUploaded bool `dynamodbav:"diff_was_uploaded"`
}

func EventTable() dynamo.Table {
	return db.Get().Table(tableName)
}

// RetrieveEvent gets the event associated to a given hash
// to ensure the request is correct and expected
func RetrieveEvent(commitHash string) (*Event, error) {
	event := &Event{}
	return event, EventTable().
		Get("HashSHA1", commitHash).
		Filter("HashSHA1", commitHash, "HasRapport", false).
		One(event)
}

func CheckEvent(request *awstypes.Request) (*Event, error) {
	// Check if the received body is not too large
	if len(request.Body) > consts.MaxUploadSize {
		return nil, errors.New("body is too big")
	}

	commitHash := request.Headers[consts.CommitHashHeaderName]
	if len(commitHash) != sha1.Size {
		return nil, errors.New("invalid or missing SHA1 commit hash")
	}

	return RetrieveEvent(commitHash)
}
