package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func Get() *dynamo.DB {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})
	return db
}
