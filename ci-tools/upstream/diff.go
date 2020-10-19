package upstream

import "net/http"

// DiffInput contains the input for uploading a pre-formatted plain-text difference report
type DiffInput struct {
	PostCommentUrl string
	SHA1Revision string
}

func (input *DiffInput) GetUpstreamUrl() string {
	return input.PostCommentUrl
}

func (input *DiffInput) GetContentType() string {
	return "text/plain"
}

func (input *DiffInput) GetRevision() string {
	return input.SHA1Revision
}

func (input *DiffInput) GetHeaders() *Headers {
	return GetBaseHeaders(input)
}

func (input *DiffInput) PostFromStdin() (*http.Response, error) {
	if err := ValidateInput(input); err != nil {
		return nil, err
	}
	return PostFromStdin(input)
}
