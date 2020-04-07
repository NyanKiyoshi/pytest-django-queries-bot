package main

import (
	"crypto/hmac"
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/config"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"strings"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// SecretKeyHeaderName defines the header that contains
// the secret key for uploading a file.
const SecretKeyHeaderName = "X-Secret-Key"
const SecretKeyHeaderNameLower = "x-secret-key"
const ContentTypeHeaderName = "Content-Type"
const ContentTypeHeaderNameLower = "content-type"

// ExpectedSecretKey contains the expected secret key to receive
// that will allow the request to be handled.
var ExpectedSecretKey = []byte(config.WebhookSecretKey)

const jsonContentType = "application/json"

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx awstypes.Request) (Response, error) {
	secretKey, found := ctx.Headers[SecretKeyHeaderName]

	if !found {
		secretKey, found = ctx.Headers[SecretKeyHeaderNameLower]
	}

	if !found {
		return Response{StatusCode: 400}, fmt.Errorf("no creds: %v", ctx.Headers)
	}

	// Time based comparison of the received key to compare with the received key
	if !hmac.Equal([]byte(secretKey), ExpectedSecretKey) {
		return Response{StatusCode: 403, Body: "Bad credentials"}, fmt.Errorf("bad creds: %s", secretKey)
	}

	// Check that the body might be JSON
	contentType, ok := ctx.Headers[ContentTypeHeaderName]
	if !ok {
		contentType = ctx.Headers[ContentTypeHeaderNameLower]
	}
	if contentType != jsonContentType {
		return Response{StatusCode: 400, Body: "Expected JSON"}, fmt.Errorf("invalid content type: %s", ctx.Headers["Content-Type"])
	}

	// Retrieve the event to ensure the request is correct and expected
	event, err := models.CheckEvent(&ctx)

	if err != nil {
		return Response{StatusCode: 400, Body: err.Error()}, err
	}

	// Create a session to AWS to upload to the S3 bucket
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.S3AwsRegion),
		Credentials: credentials.NewEnvCredentials(),
	})

	if err != nil {
		return Response{StatusCode: 500, Body: "Failed to start uploader"}, err
	}

	// Start a new uploader and upload the request body to the S3 bucket
	uploader := s3manager.NewUploader(awsSession)
	s3ContentType := string(jsonContentType) // copy the string because AWS wants a pointer

	if _, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(config.S3Bucket),
		Key:         &event.HashSHA1,
		ContentType: &s3ContentType,
		Body:        strings.NewReader(ctx.Body),
	}); err != nil {
		return Response{StatusCode: 500, Body: "Failed to upload"}, err
	}

	if err := models.EventTable().Update("HashSHA1", event.HashSHA1).
		Set("HasRapport", true).
		Run(); err != nil {
		return Response{StatusCode: 500, Body: "Failed to update event data"}, err
	}

	return Response{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
