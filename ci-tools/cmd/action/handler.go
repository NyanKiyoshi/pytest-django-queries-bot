package main

import (
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/integration"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
)

func fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func getFromEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		fatalf("failed to find '%s' in environment", key)
	}
	return value
}

func getUrlFromEnv(key string) *url.URL {
	rawValue := getFromEnv(key)
	urlValue, err := url.Parse(rawValue)

	if err != nil {
		fatalf("invalid URL for '%s': %w", key, err)
	}
	return urlValue
}

func SetUpHandlers() *integration.HandlerMap {
	logger := logrus.StandardLogger()
	handlers := integration.NewHandlerMap()
	handlers.Register(
		"pull_request", &integration.PullRequestHandler{
			Logger:                    logger,
			RetrieveQueriesPayloadUrl: getUrlFromEnv("DIFF_RESULTS_BASE_URL"),
			PostDiffCommentUrl:        getUrlFromEnv("DIFF_ENDPOINT"),
			HeadReportLocation:        getFromEnv("QUERIES_RESULTS_PATH"),
		})
	handlers.Register(
		"push", &integration.PushHandler{
			Logger: logger,
			GetLazySecretsFunc: func() *integration.PushHandlerLazySecrets {
				return &integration.PushHandlerLazySecrets{
					PostPayloadEndpointUrl: getUrlFromEnv("QUERIES_UPLOAD_ENDPOINT_URL"),
					PostKey:                getFromEnv("QUERIES_UPLOAD_SECRET"),
				}
			},
			HeadReportLocation: getFromEnv("QUERIES_RESULTS_PATH"),
		})
	return handlers
}

func main() {
	handlers := SetUpHandlers()
	eventName := getFromEnv("GITHUB_EVENT_NAME")
	payloadPath := getFromEnv("GITHUB_EVENT_PATH")

	handler, err := handlers.Lookup(eventName)
	if err != nil {
		fatalf("lookup for event '%s' failed: %w", eventName, err)
	}

	rawPayload, err := ioutil.ReadFile(payloadPath)
	if err != nil {
		fatalf("payload reading the payload from file ('%s'): %w", payloadPath, err)
	}

	err = integration.ParseEvent(rawPayload, handler)
	if err != nil {
		fatalf("payload parsing for event '%s' failed: %w", eventName, err)
	}

	err = handler.Invoke()
	if err != nil {
		fatalf("invocation for event '%s' failed: %w", err)
	}
}
