# !/bin/bash
if [ "$#" -eq 0 ]; then
    echo "You need to provide at least one branch name"
    exit 1
fi

REMOTE="$(git remote -v | head -n1 | cut -f1)"

for BRANCH in "$@"; do
    git push -u $REMOTE $BRANCH
done
