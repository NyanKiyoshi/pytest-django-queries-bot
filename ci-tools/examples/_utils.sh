#!/usr/bin/env bash

function ensureenv() {
    local key=$1
    local value=${!key}

    [[ -z "${value}" ]] && {
        echo "Missing environment variable: '${key}'" >&2
        has_error=1
    }
}
