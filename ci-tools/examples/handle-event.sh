#!/usr/bin/env bash

HERE=`readlink -f "$(dirname $0)"`

if [[ -n "$TRAVIS_PULL_REQUEST_SLUG" ]]; then
    echo "Uploading diff results from pull request..." >&2
    ${HERE}/_handle-pull-request.sh
else
    echo "Uploading raw results from push..." >&2
    ${HERE}/_handle-branch-push.sh
fi
