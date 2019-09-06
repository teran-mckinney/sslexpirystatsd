#!/usr/bin/env bash

set -eE

shellcheck "$0"

# Before we build...
go fmt
go doc
go test

go build

# strip -s decensor

cleanup() {
    echo "Cleaning up."
}

trap fail $(seq 1 64)

fail() {
    echo "FAIL: $1"
    cleanup
    exit 1
}

echo Success
