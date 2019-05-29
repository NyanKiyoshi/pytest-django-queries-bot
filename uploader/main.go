package main

import (
	"crypto/hmac"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/models"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// SecretKeyHeaderName defines the header that contains
// the secret key for uploading a file.
const SecretKeyHeaderName = "X-Upload-Secret-Key"

// ExpectedSecretKey contains the expected secret key to receive
// that will allow the request to be handled.
var ExpectedSecretKey = []byte(os.Getenv("REQUIRED_SECRET_KEY"))

// Various S3 information.
var s3BucketRegion = aws.String(os.Getenv("S3_AWS_REGION"))
var s3BucketName = os.Getenv("S3_BUCKET")

const jsonContentType = "application/json"

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx awstypes.Request) (Response, error) {
	secretKey, found := ctx.Headers[SecretKeyHeaderName]

	if !found {
		return Response{StatusCode: 400}, nil
	}

	// Time based comparison of the received key to compare with the received key
	if !hmac.Equal([]byte(secretKey), ExpectedSecretKey) {
		return Response{StatusCode: 403, Body: "Bad credentials"}, nil
	}

	// Check that the body might be JSON
	if ctx.Headers["Content-Type"] != jsonContentType {
		return Response{StatusCode: 400, Body: "Expected JSON"}, nil
	}

	// Retrieve the event to ensure the request is correct and expected
	event, err := models.CheckEvent(&ctx)

	if err != nil {
		return Response{StatusCode: 400, Body: err.Error()}, err
	}

	// Create a session to AWS to upload to the S3 bucket
	awsSession, err := session.NewSession(&aws.Config{
		Region: s3BucketRegion,
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("S3_AWS_ACCESS_KEY_ID"),
			os.Getenv("S3_AWS_SECRET_KEY"),
			os.Getenv("S3_AWS_SESSION_TOKEN"),
		),
	})

	if err != nil {
		return Response{StatusCode: 500, Body: "Failed to start uploader"}, err
	}

	// Start a new uploader and upload the request body to the S3 bucket
	uploader := s3manager.NewUploader(awsSession)
	s3ContentType := string(jsonContentType) // copy the string because AWS wants a pointer

	if _, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      &s3BucketName,
		Key:         &event.HashSHA1,
		ContentType: &s3ContentType,
	}); err != nil {
		return Response{StatusCode: 500, Body: "Failed to upload"}, err
	}

	if err := models.EventTable().Update("HashSHA1", event.HashSHA1).
		Set("HasRapport", true).
		Run(); err != nil {
		return Response{StatusCode: 500, Body: "Failed to update event data"}, nil
	}

	return Response{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
