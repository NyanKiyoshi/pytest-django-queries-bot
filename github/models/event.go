package models

import (
	"github.com/guregu/dynamo"
	"pytest-queries-bot/pytestqueries/db"
)

type Event struct {
	HashSHA1 string `dynamodbav:"sha1_hash"`
}

func (ev *Event) Table() dynamo.Table {
	return db.Get().Table("gh_hooks_events")
}
