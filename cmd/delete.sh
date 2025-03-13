#!/bin/bash
if [ "$#" -eq 0]; then
    echo "You need to provide at least one branch name"
    exit 1
fi

for BRANCH in "$@"; do
    git branch -d $BRANCH
done
