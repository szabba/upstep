#!/usr/bin/env bash

tmpFile=$(mktemp)
git diff | tee "$tmpFile"
if [[ $(wc -c "$tmpFile") -ne 0 ]]
then
    "git diff was not empty"
    exit 1
fi
