#!/bin/bash
git fetch -p
{
    git branch --format "%(refname:short) %(upstream)" | awk '{if (!$2) print $1;}'
    git branch -vv | grep ": gone]" | awk '{print $3}' | tr -d "[:" | tr "/" " " | awk '{print $2}'
}
