package models

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/config"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/db"
	"github.com/guregu/dynamo"
	"time"
)

type PullRequest struct {
	// PullRequestID is the pull request ID in GitHub.
	// This will be used to comment with results into GitHub.
	PullRequestID int64

	// OwnerName is the owner of the target repository.
	OwnerName string `dynamodbav:"owner_name"`

	// RepoName is the name of the target repository.
	RepoName string `dynamodbav:"repo_name"`

	// InstallationId is ID of the GitHub Application installation.
	InstallationId int64 `dynamodbav:"installation_id"`

	// PullRequestNumber is the pull request number in GitHub.
	// This will be used to comment with results into GitHub.
	PullRequestNumber int `dynamodbav:"pr_number"`

	// GitHubCommentID is the bot's diff comment over the pull request,
	// if any. This will be used to update any existing comment.
	GitHubCommentID int64 `dynamodbav:"github_comment_id"`

	// EntryCreationDate is the date where this entry was created/updated.
	// Not to be confused with the GitHub event date.
	EntryDate time.Time `dynamodbav:"_date"`
}

func PullRequestTable() dynamo.Table {
	return db.Get().Table(config.DynamoPullReqTableName)
}

func GetPullRequest(pullID int64) (*PullRequest, error) {
	pr := &PullRequest{}
	err := PullRequestTable().Get("PullRequestID", pullID).One(pr)
	return pr, err
}
