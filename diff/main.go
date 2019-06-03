package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
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
	event, err := models.CheckEvent(&request)

	if err != nil {
		return Response{StatusCode: 400, Body: err.Error()}, err
	}

	if event.DiffUploaded {
		return Response{StatusCode: 403, Body: "A diff was already uploaded for this revision"}, nil
	}

	ghClient, ctx := client.GetClient()
	comment := github.PullRequestComment{
		Body: &request.Body,
	}

	if event.GitHubCommentID != 0 {
		_, _, err := ghClient.PullRequests.EditComment(
			*ctx, event.OwnerName, event.RepoName, event.GitHubCommentID, &comment,
		)
		if err == nil {
			return Response{StatusCode: 201, Body: "Comment created."}, nil
		}
	}

	_, _, err = ghClient.PullRequests.CreateComment(
		*ctx, event.OwnerName, event.RepoName, event.PullRequestNumber, &comment,
	)

	if err == nil {
		return Response{StatusCode: 201, Body: "Comment created."}, nil
	}

	return Response{StatusCode: 500, Body: "Failed to create a comment..."}, err
}

func main() {
	lambda.Start(Handler)
}
