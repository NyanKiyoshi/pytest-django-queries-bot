package ghevents

import "pytest-queries-bot/pytestqueries/github/awstypes"

func unknown() (awstypes.Response, error) {
	return awstypes.Response{
		StatusCode: 405,
		Body: "Unknown or unsupported event",
	}, nil
}
