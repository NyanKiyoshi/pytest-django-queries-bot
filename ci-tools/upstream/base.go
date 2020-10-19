package upstream

import (
	"fmt"
	"github.com/NyanKiyoshi/pytest-django-queries-bot/ci-tools/utils"
	"io"
	"net/http"
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

	// PostFromReader sends Stdin as body to the specified upstream
	PostFromStdin() (*http.Response, error)

	// PostFromReader sends giving reader as body to the specified upstream
	PostFromReader(r io.Reader) (*http.Response, error)
}

func GetBaseHeaders(input Input) *Headers {
	headers := Headers{"X-Commit-Rev": input.GetRevision()}
	headers["Content-Type"] = input.GetContentType()
	return &headers
}

func ValidateInput(input Input) (err error) {
	rev := input.GetRevision()
	if len(rev) != 40 {
		err = fmt.Errorf("sha1 revision length is invalid, expected: %d, got: %d", 40, len(rev))
	}
	return
}

func PostFromReader(input Input, reader io.Reader) (*http.Response, error) {
	url := input.GetUpstreamUrl()
	headers := input.GetHeaders()

	return utils.SendUploadRequest(
		url,
		reader,
		(*map[string]string)(headers),
	)
}
