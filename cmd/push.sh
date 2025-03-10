# !/bin/bash
if [ "$#" -eq 0 ]; then
    echo "You need to provide at least one branch name"
    exit 1
fi

REMOTE="$(git remote -v | head -n1 | cut -f1)"

for VAR in "$@"; do
    # git push $REMOTE $VAR
    echo "git push ${REMOTE} ${VAR}"
done
