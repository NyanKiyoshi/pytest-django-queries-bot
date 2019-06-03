package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
	"log"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/client"
	"pytest-queries-bot/pytestqueries/github/models"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request awstypes.Request) (Response, error) {
	log.Printf("%v", request.Headers)
	event, err := models.CheckEvent(&request)

	if err != nil {
		return Response{StatusCode: 400, Body: err.Error()}, err
	}

	if event.DiffUploaded {
		return Response{StatusCode: 403, Body: "A diff was already uploaded for this revision"}, nil
	}

	ghClient, ctx := client.GetClient()
	comment := github.IssueComment{
		Body: &request.Body,
	}

	if event.GitHubCommentID != 0 {
		_, _, err := ghClient.Issues.EditComment(
			*ctx, event.OwnerName, event.RepoName, event.GitHubCommentID, &comment,
		)

		if err != nil {
			return Response{
				StatusCode: 500,
				Body:       fmt.Sprintf("Failed to create comment %d...", event.GitHubCommentID),
			}, err
		}

		event.DiffUploaded = true
	}

	if !event.DiffUploaded {

		newComment, _, err := ghClient.Issues.CreateComment(
			*ctx, event.OwnerName, event.RepoName, event.PullRequestNumber, &comment,
		)

		if err != nil {
			return Response{
				StatusCode: 500,
				Body:       fmt.Sprintf("Failed to create comment %d...", event.GitHubCommentID),
			}, err
		}

		event.GitHubCommentID = *newComment.ID
	}

	err = models.EventTable().Put(event).Run()
	return Response{StatusCode: 201, Body: "Comment created."}, err
}

func main() {
	lambda.Start(Handler)
}
