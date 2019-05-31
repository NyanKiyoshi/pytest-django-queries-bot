package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"pytest-queries-bot/pytestqueries/generated"
)

// Get retrieves the AWS Dynamo DB
func Get() *dynamo.DB {
	cfg := aws.NewConfig().
		WithRegion(generated.DynamoAwsRegion).
		WithCredentials(credentials.NewStaticCredentials(
			generated.DynamoAwsAccessKeyId,
			generated.DynamoAwsSecretKey,
			generated.DynamoAwsSessionToken)).
		WithLogLevel(aws.LogDebugWithHTTPBody)
	db := dynamo.New(session.New(), cfg)
	return db
}
