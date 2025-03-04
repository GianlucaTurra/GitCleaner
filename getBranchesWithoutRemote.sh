#!/bin/bash
git branch -vv | grep -v "origin" | tr "*" " " | cut -d " " -f 3
