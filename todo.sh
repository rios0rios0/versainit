#!/bin/bash
set -euxo pipefail

# update all git repositories in the current directory
for repo in $(find . -mindepth 1 -maxdepth 1 -type d); do
    pushd "$repo"

    # determine the main branch name
    if command git show-ref -q --verify refs/heads/main; then
        default_branch=main
    else
        default_branch=master
    fi

    # switch to the main branch and pull the latest changes
    git switch "$default_branch"
    git pull

    # remove local branches that have been deleted remotely
    # i.e., branches that have been merged
    set +e
    git fetch --all -p && git branch -vv | awk "/^[^*]/ && /: gone]/{print \$1}" | xargs git branch -d
    set -e

    popd
done
