package main

import (
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/ghevents"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/security"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fakeSecretKey = []byte("secret")

func init() {
	secretKey = fakeSecretKey
}

var invalidRequests = []struct {
	testName string
	request  awstypes.Request
}{
	{testName: "missing header", request: awstypes.Request{}},
	{testName: "invalid secret key", request: awstypes.Request{Headers: map[string]string{HmacHeader: "invalid"}}},
}

func TestHandler_unauthorized_request_returns_denied(t *testing.T) {
	expected := awstypes.Response{
		StatusCode: 403,
	}
	for _, tt := range invalidRequests {
		t.Run(tt.testName, func(t *testing.T) {
			response, err := Handler(tt.request)
			assert.Nil(t, err, "Should not have an error")
			assert.Equal(t, expected, response)
		})
	}
}

func TestHandler_validRequest(t *testing.T) {
	body := "no body needed"
	request := awstypes.Request{
		Headers: map[string]string{
			HmacHeader:                 fmt.Sprint(string(security.SignaturePrefix), security.NewHMAC([]byte(body), fakeSecretKey)),
			ghevents.GithubEventHeader: "ping",
		},
		Body: body,
	}
	expected := awstypes.Response{
		StatusCode: 200,
		Body:       `{"status": 200}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	response, err := Handler(request)
	assert.Nil(t, err, "Should not have an error")
	assert.Equal(t, expected, response)
}
