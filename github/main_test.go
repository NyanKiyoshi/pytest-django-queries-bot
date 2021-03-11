package main

import (
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/config"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/awstypes"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/ghevents"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/github/security"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fakeSecretKey = []byte("secret")

func init() {
	config.WebhookSecretKey = fakeSecretKey
}

var invalidRequests = []struct {
	testName     string
	request      awstypes.Request
	expectedBody string
}{
	{
		testName:     "missing header",
		request:      awstypes.Request{},
		expectedBody: "{\"message\": \"authentication failed: missing header or empty signature\"}",
	},
	{
		testName:     "invalid secret key",
		request:      awstypes.Request{Headers: map[string]string{HmacHeader: "invalid"}},
		expectedBody: "{\"message\": \"authentication failed\"}",
	},
}

func TestHandler_unauthorized_request_returns_denied(t *testing.T) {
	expected := awstypes.Response{
		StatusCode:      401,
		Headers:         map[string]string{"Content-Type": "application/json"},
		IsBase64Encoded: false,
	}
	for _, tt := range invalidRequests {
		expected.Body = tt.expectedBody
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
