#!/bin/bash
git branch -vv | grep origin | tr "*" " " | cut -d " " -f 3
