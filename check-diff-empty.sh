#!/usr/bin/env bash

tmpFile=$(mktemp)
git diff | tee "$tmpFile"
if [[ -s "$tmpFile" ]]
then
    echo "git diff was not empty"
    exit 1
fi
