package upstream

import "net/http"

// RawReportInput contains the input for uploading a raw JSON report
type RawReportInput struct {
	PostCommentUrl string
	SHA1Revision string
	SecretKey string
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
	return PostFromStdin(input)
}
