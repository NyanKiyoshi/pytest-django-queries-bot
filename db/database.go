package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
	"pytest-queries-bot/pytestqueries/generated"
)

// Get retrieves the AWS Dynamo DB
func Get() *dynamo.DB {
	region := os.Getenv("DYNAMO_AWS_REGION")
	db := dynamo.New(session.New(), &aws.Config{
		Region: &region,
		Credentials: credentials.NewStaticCredentials(
			generated.DynamoAwsAccessKeyId,
			generated.DynamoAwsSecretKey,
			generated.DynamoAwsSessionToken,
		),
	})
	return db
}
