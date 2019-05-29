#!/usr/bin/env bash

function ensureenv() {
    local key=$1
    local default=$2
    local value=${!key:-${default}}

    [[ -z "${value}" ]] && {
        echo "Missing environment variable: '${key}'" >&2
        exit 1
    }

    export "${key}=${value}"
}

ensureenv AWS_DYNAMO_REGION us-east-2
ensureenv S3_BUCKET gh-reports
ensureenv GITHUB_SECRET_KEY
ensureenv GITHUB_WEBHOOK_URL
ensureenv REQUIRED_SECRET_KEY
ensureenv S3_AWS_ACCESS_KEY_ID
ensureenv S3_AWS_SECRET_KEY
ensureenv S3_AWS_SESSION_TOKEN
