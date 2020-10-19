package upstream

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

// RawReportInput contains the input for uploading a raw JSON report
type RawReportInput struct {
	PostCommentUrl string
	SHA1Revision   string
	SecretKey      string
}

func (input *RawReportInput) GetUpstreamUrl() string {
	return input.PostCommentUrl
}

func (input *RawReportInput) GetContentType() string {
	return "application/json"
}

func (input *RawReportInput) GetRevision() string {
	return input.SHA1Revision
}

func (input *RawReportInput) GetHeaders() *Headers {
	headers := *GetBaseHeaders(input)
	headers["X-Secret-Key"] = input.SecretKey
	return &headers
}

func (input *RawReportInput) PostFromStdin() (*http.Response, error) {
	if err := ValidateInput(input); err != nil {
		return nil, err
	}
	return PostFromReader(input, bufio.NewReader(os.Stdin))
}

func (input *RawReportInput) PostFromReader(r io.Reader) (*http.Response, error) {
	if err := ValidateInput(input); err != nil {
		return nil, err
	}
	return PostFromReader(input, r)
}
