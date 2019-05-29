package ghevents

import "pytest-queries-bot/pytestqueries/github/awstypes"

func unknown() (awstypes.Response, error) {
	return awstypes.Response{
		StatusCode: 405,
		Body:       `{"message": "Unknown or unsupported event"}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
