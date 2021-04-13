package db

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// Get retrieves the AWS Dynamo DB
func Get() *dynamo.DB {
	cfg := aws.NewConfig().
		WithRegion(config.DynamoAwsRegion).
		WithCredentials(credentials.NewEnvCredentials()).
		WithLogLevel(aws.LogDebugWithRequestErrors)
	db := dynamo.New(session.New(), cfg)
	return db
}
