#!/usr/bin/env bash

set -e

# Go to root of repository from .git/hooks
cd "${0%/*}/../.."

files=$(git diff --cached --name-only --diff-filter=ACMR | sed 's| |\\ |g')

[ -z "$files" ] && exit 0

echo "$files" | xargs npx prettier --ignore-unknown --write
