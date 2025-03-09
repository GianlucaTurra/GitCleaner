#!/bin/bash
git branch -r | sed 's/origin\///' | cut -d " " -f 3 | tr -d "HEAD"
