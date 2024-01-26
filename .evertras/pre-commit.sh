#!/usr/bin/env bash

set -e

# Go to root of repository from .git/hooks
cd "${0%/*}/../.."

make fmt
