#!/usr/bin/env bash

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

[[ ${has_error} -eq 1 ]] && {
    echo "Found errors. Exiting..." >&2
    exit 1
}

user=`echo ${TRAVIS_REPO_SLUG} | cut / -f0`
repo=`echo ${TRAVIS_REPO_SLUG} | cut / -f1`

(
    curl https://api.github.com/graphql -H "Authorization: bearer ${GITHUB_GQL_TOKEN}" -X POST -d @- <<EOF
{
  repository(owner: "${user}", name: "${repo}") {
    ref(qualifiedName: "${TRAVIS_BRANCH}") {
      associatedPullRequests(last: 1) {
        edges {
          node {
            baseRefOid
          }
        }
      }
    }
  }
}
EOF
) | grep "baseRefOid:" | cut -d: -f1 | tr '" ' -f1 || {
    err_code=$?
    echo "Failed to retrieve repo data"
    exit ${err_code}
}
