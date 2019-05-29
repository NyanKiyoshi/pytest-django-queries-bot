package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"pytest-queries-bot/pytestqueries/github/awstypes"
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
		return Response{StatusCode: 403, Body: "A diff was already uploaded"}, nil
	}

	return Response{StatusCode: 404}, nil
}

func main() {
	lambda.Start(Handler)
}
