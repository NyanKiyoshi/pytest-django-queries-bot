package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"pytest-queries-bot/pytestqueries/github/awstypes"
	"pytest-queries-bot/pytestqueries/github/ghevents"
	"pytest-queries-bot/pytestqueries/github/security"
)

func getGHSecretKey() []byte {
	return []byte(os.Getenv("GITHUB_SECRET_KEY"))
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request awstypes.Request) (awstypes.Response, error) {
	signature, ok := request.Headers["HTTP_X_HUB_SIGNATURE"]
	if !(ok && security.CheckHMAC([]byte(signature), []byte(request.Body), getGHSecretKey())) {
		return awstypes.Response{
			StatusCode: 400,
		}, nil
	}

	return ghevents.Dispatch(&request)
}

func main() {
	lambda.Start(Handler)
}
