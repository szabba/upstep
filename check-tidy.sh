#!/usr/bin/env bash

function noDiff {
    local tmpFile
    tmpFile=$(mktemp)
    git diff | tee "$tmpFile"
    if [[ $(wc -c "$tmpFile") -ne 0 ]]
    then
        "git diff was not empty"
        exit 1
    fi
}

go mod tidy
noDiff
