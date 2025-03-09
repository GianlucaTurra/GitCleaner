#!/bin/bash
git branch --format "%(refname:short) %(upstream)" | awk '{if (!$2) print $1}'
