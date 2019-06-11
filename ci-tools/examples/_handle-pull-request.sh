#!/usr/bin/env bash

HERE=`readlink -f "$(dirname $0)"`
has_error=0

. ${HERE}/_utils.sh

ensureenv TRAVIS_PULL_REQUEST_SHA
ensureenv TRAVIS_COMMIT_RANGE
ensureenv TRAVIS_REPO_SLUG
ensureenv TRAVIS_BRANCH
ensureenv DIFF_RESULTS_BASE_URL
ensureenv QUERIES_RESULTS_PATH
ensureenv DIFF_ENDPOINT

[[ ${has_error} -eq 1 ]] && {
    echo "Found errors. Exiting..." >&2
    exit 1
}

user=`echo ${TRAVIS_PULL_REQUEST_SLUG} | cut -d/ -f1`
repo=`echo ${TRAVIS_PULL_REQUEST_SLUG} | cut -d/ -f2`
ref_name=${TRAVIS_PULL_REQUEST_BRANCH}

base_ref_hash=$(echo ${TRAVIS_COMMIT_RANGE} | sed -E 's/\.\.\..+//')
head_results_path="/tmp/base-results.json"
missing_head=0

curl --fail -X GET "${DIFF_RESULTS_BASE_URL}/${base_ref_hash}" -L -o "${head_results_path}" || {
    echo "[Warning] Did not find a HEAD base results for ${base_ref_hash}" >&2
    cp -v "${QUERIES_RESULTS_PATH}" "${head_results_path}"
    missing_head=1
}

django-queries diff "${head_results_path}" "${QUERIES_RESULTS_PATH}" > /tmp/diff || {
    echo "Failed to generate the diff. Aborting..." >&2
    exit 1
}

difference_count=$(grep -e  '^-' -e '^+' -c /tmp/diff)
echo -n "Differences count (+/-): ${difference_count}"

echo Uploading...

${HERE}/tools/queries-diff --rev ${TRAVIS_PULL_REQUEST_SHA} <<EOF

Here is the report for ${TRAVIS_PULL_REQUEST_SHA} (${TRAVIS_PULL_REQUEST_SLUG} @ ${TRAVIS_PULL_REQUEST_BRANCH})
$([[ ${missing_head} -eq 1 ]] && echo "Missing base report (${base_ref_hash}). The results couldn't be compared." || echo "Base comparison is ${base_ref_hash}.")



<details><summary>$([[ ${difference_count} -eq 0 ]] && echo "No differences were found." || echo "**Found ${difference_count} differences!**") (click me)</summary>
<p>

\`\`\`diff
$(cat /tmp/diff)
\`\`\`

</p>
</details>
EOF
