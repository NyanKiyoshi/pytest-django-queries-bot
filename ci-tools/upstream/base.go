package upstream

import (
	"bufio"
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/utils"
	"net/http"
	"os"
)

type Headers map[string]string

type Input interface {
	// GetUpstreamUrl returns the URL to post the body to
	GetUpstreamUrl() string

	// GetContentType returns the content-type of the body to send
	GetContentType() string

	// GetRevision returns the commit sha1 hash
	GetRevision() string

	// GetHeaders returns the base headers from GetBaseHeaders on top of additional headers to send
	GetHeaders() *Headers

	// PostFromStdin sends Stdin as body to the specified upstream
	PostFromStdin() (*http.Response, error)
}

func GetBaseHeaders(input Input) *Headers {
	return &Headers{"X-Commit-Rev": input.GetRevision()}
}

func ValidateInput(input Input) (err error) {
	rev := input.GetRevision()
	if len(rev) != 40 {
		err = fmt.Errorf("sha1 revision length is invalid, expected: %d, got: %d", 40, len(rev))
	}
	return
}

func PostFromStdin(input Input) (*http.Response, error) {
	url := input.GetUpstreamUrl()
	ct := input.GetContentType()
	headers := input.GetHeaders()

	return utils.SendUploadRequest(
		url,
		ct,
		bufio.NewReader(os.Stdin),
		(*map[string]string)(headers),
	)
}
