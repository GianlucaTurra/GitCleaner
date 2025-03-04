#!/bin/bash
git branch -vv | tr "*" " " | cut -d " " -f 3
