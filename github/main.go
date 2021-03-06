package main

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/config"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/ghevents"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/security"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/logging"
	"github.com/aws/aws-lambda-go/lambda"
)

const HmacHeader string = "x-hub-signature"
const CanonicalMIMEHmacHeader string = "X-Hub-Signature"

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request awstypes.Request) (awstypes.Response, error) {
	logging.Logger.Debugf("%+v", request)

	signature, ok := request.Headers[HmacHeader]
	if !ok {
		signature, ok = request.Headers[CanonicalMIMEHmacHeader]
	}

	if !ok {
		return awstypes.Response{
			StatusCode: 401,
			Body:       `{"message": "authentication failed: missing header or empty signature"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	ok = security.VerifySignature(config.WebhookSecretKey, signature, []byte(request.Body))
	if !ok {
		return awstypes.Response{
			StatusCode: 401,
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
