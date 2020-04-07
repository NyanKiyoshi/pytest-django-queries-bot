package config

import (
	"os"
	"strconv"
)

var DynamoAwsRegion = os.Getenv("DYNAMO_AWS_REGION")
var DynamoEventsTableName = os.Getenv("DYNAMO_EVENTS_TABLE_NAME")
var DynamoPullReqTableName = os.Getenv("DYNAMO_PULL_REQ_TABLE_NAME")

var GithubAppId, _ = strconv.ParseInt(os.Getenv("GITHUB_APP_ID"), 10, 64)
var UploadSecretKey = []byte(os.Getenv("UPLOAD_SECRET_KEY"))
var WebhookSecretKey = []byte(os.Getenv("WEBHOOK_SECRET_KEY"))

var S3AwsRegion = os.Getenv("S3_AWS_REGION")
var S3Bucket = os.Getenv("S3_BUCKET")
