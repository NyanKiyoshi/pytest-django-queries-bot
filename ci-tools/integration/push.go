package integration

import (
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/upstream"
	"github.com/google/go-github/v32/github"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
)

type PushHandlerLazySecrets struct {
	PostPayloadEndpointUrl *url.URL
	PostKey                string
}

type PushHandler struct {
	event              github.PushEvent
	Secrets            *PushHandlerLazySecrets
	HeadReportLocation string
	Logger             *logrus.Logger

	// GetLazySecretsFunc is called in order to defer the configuration and validation
	// of the handler's configuration as those values can be missing depending on how the caller
	// configured their CI, the value could be empty or not the mapping not set at all.
	//
	// For this reason, we only configure the secrets when we the caller is actually invoking that
	// specific handler.
	GetLazySecretsFunc func() *PushHandlerLazySecrets
}

func (h *PushHandler) GetEventPtr() interface{} {
	return &h.event
}

func (h *PushHandler) Invoke() (err error) {
	// Populate secrets if not already populated
	if h.Secrets == nil {
		h.Secrets = h.GetLazySecretsFunc()
	}

	event := h.event
	head := event.HeadCommit.GetSHA()
	if head == "" {
		head = event.HeadCommit.GetID()
	}

	input := upstream.RawReportInput{
		PostCommentUrl: h.Secrets.PostPayloadEndpointUrl.String(),
		SHA1Revision:   head,
		SecretKey:      h.Secrets.PostKey,
	}

	fp, err := os.OpenFile(h.HeadReportLocation, os.O_RDONLY, 0)
	if err != nil {
		err = fmt.Errorf("failed to open head report ('%s'): %w", h.HeadReportLocation, err)
		return
	}

	if _, err = input.PostFromReader(fp); err != nil {
		err = fmt.Errorf("failed to post raw results from file: %w", err)
	}
	return
}
