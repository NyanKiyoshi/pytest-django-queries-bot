package main

import (
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/client"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/v32/github"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request awstypes.Request) (Response, error) {
	event, err := models.CheckEvent(&request)
	if err != nil {
		return Response{StatusCode: 400, Body: err.Error()}, err
	}

	pr, err := models.GetPullRequest(event.PullRequestID)
	if err != nil {
		return Response{StatusCode: 400, Body: err.Error()}, err
	}

	if event.DiffUploaded {
		return Response{StatusCode: 403, Body: "A diff was already uploaded for this revision"}, nil
	}

	ghClient, ctx := client.GetClient(pr.InstallationId)
	comment := github.IssueComment{
		Body: &request.Body,
	}

	if pr.GitHubCommentID != 0 {
		_, _, err := ghClient.Issues.EditComment(
			*ctx, pr.OwnerName, pr.RepoName, pr.GitHubCommentID, &comment,
		)

		if err != nil {
			return Response{
				StatusCode: 500,
				Body:       fmt.Sprintf("Failed to create comment %d...", pr.GitHubCommentID),
			}, err
		}

		event.DiffUploaded = true
	}

	if !event.DiffUploaded {

		newComment, _, err := ghClient.Issues.CreateComment(
			*ctx, pr.OwnerName, pr.RepoName, pr.PullRequestNumber, &comment,
		)

		if err != nil {
			return Response{
				StatusCode: 500,
				Body:       fmt.Sprintf("Failed to create comment %d...", pr.GitHubCommentID),
			}, err
		}

		pr.GitHubCommentID = *newComment.ID
	}

	err = models.EventTable().Put(event).Run()
	err = models.PullRequestTable().Put(pr).Run()
	return Response{StatusCode: 201, Body: "Comment created."}, err
}

func main() {
	lambda.Start(Handler)
}
