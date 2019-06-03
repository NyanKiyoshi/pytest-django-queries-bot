#!/usr/bin/env bash

export PATH="$PATH:$(npm bin)"
has_error=0

[[ -f production.env ]] && . ./production.env || { [[ -f development.env ]] && . ./development.env ; }

function ensuredep() {
    which "$1" > /dev/null || {
        echo "Missing required command: '$1'" >&2
        has_error=1
    }
}

function ensureenv() {
    local key=$1
    local default=$2
    local value=${!key:-${default}}

    [[ -z "${value}" ]] && {
        echo "Missing environment variable: '${key}'" >&2
        has_error=1
    }

    export "${key}=${value}"
}

ensureenv DYNAMO_AWS_REGION
ensureenv S3_BUCKET gh-reports
ensureenv GITHUB_SECRET_KEY
ensureenv GITHUB_WEBHOOK_URL
ensureenv REQUIRED_SECRET_KEY
ensureenv S3_AWS_ACCESS_KEY_ID
ensureenv S3_AWS_SECRET_KEY
ensureenv AWS_ACCESS_KEY_ID
ensureenv AWS_SECRET_ACCESS_KEY

ensuredep go
ensuredep bash
ensuredep make
ensuredep npm
ensuredep serverless

[[ ${has_error} -eq 1 ]] && {
    echo "Found errors. Exiting..." >&2
    exit 1
}
