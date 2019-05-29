package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

// Get retrieves the AWS Dynamo DB
func Get() *dynamo.DB {
	region := os.Getenv("DYNAMO_AWS_REGION")
	db := dynamo.New(session.New(), &aws.Config{Region: &region})
	return db
}
