#!/usr/bin/env bash

HERE=`readlink -f "$(dirname $0)"`
has_error=0

function ensureenv() {
    local key=$1
    local value=${!key}

    [[ -z "${value}" ]] && {
        echo "Missing environment variable: '${key}'" >&2
        has_error=1
    }
}

ensureenv TRAVIS_REPO_SLUG
ensureenv TRAVIS_BRANCH
ensureenv GITHUB_GQL_TOKEN
ensureenv DIFF_RESULTS_BASE_URL
ensureenv QUERIES_RESULTS_PATH

[[ ${has_error} -eq 1 ]] && {
    echo "Found errors. Exiting..." >&2
    exit 1
}

user=`echo ${TRAVIS_REPO_SLUG} | cut -d/ -f1`
repo=`echo ${TRAVIS_REPO_SLUG} | cut -d/ -f2`
ref_name=${TRAVIS_BRANCH}

base_ref_hash=$($HERE/tools/queries-get-gh-base-ref -u "${user}" -n "${repo}" -r "${ref_name}") || {
    echo "Failed to get the base HEAD commit..." >&2
    exit 1
}

head_results_path="${mktemp}"
curl -X GET "${DIFF_RESULTS_BASE_URL}/${base_ref_hash}" -Lo "${head_results_path}" || {
    echo "[Warning] Did not find a HEAD base results for ${base_ref_hash}" >&2
    echo '{}' > ${head_results_path}
}

django-queries diff "${head_results_path}" "${QUERIES_RESULTS_PATH}"
