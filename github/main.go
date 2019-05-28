package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/ghevents"
	"pytest-queries-bot/pytestqueries/github/security"
)

const HmacHeader string = "HTTP_X_HUB_SIGNATURE"

var secretKey = []byte(os.Getenv("GITHUB_SECRET_KEY"))

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request awstypes.Request) (awstypes.Response, error) {
	signature, ok := request.Headers[HmacHeader]
	if !(ok && security.CheckHMAC([]byte(signature), []byte(request.Body), secretKey)) {
		return awstypes.Response{
			StatusCode: 403,
		}, nil
	}

	return ghevents.Dispatch(&request)
}

func main() {
	lambda.Start(Handler)
}
