package main

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/ghevents"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/security"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

const HmacHeader string = "X-Hub-Signature"

var secretKey = []byte(os.Getenv("GITHUB_SECRET_KEY"))

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request awstypes.Request) (awstypes.Response, error) {
	signature, ok := request.Headers[HmacHeader]
	if !ok {
		return awstypes.Response{
			StatusCode: 403,
			Body:       `{"message": "authentication failed: missing header"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	ok = security.VerifySignature(secretKey, signature, []byte(request.Body))
	if !ok {
		return awstypes.Response{
			StatusCode: 403,
			Body:       `{"message": "authentication failed"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	return ghevents.Dispatch(&request)
}

func main() {
	lambda.Start(Handler)
}
