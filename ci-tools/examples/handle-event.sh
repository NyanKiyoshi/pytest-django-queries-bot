#!/usr/bin/env bash

HERE=`readlink -f "$(dirname $0)"`

if [[ -n "$TRAVIS_COMMIT_RANGE" ]]; then
    ${HERE}/_handle-pull-request.sh
else
    ${HERE}/_handle-branch-push.sh
fi
