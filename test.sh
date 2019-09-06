#!/usr/bin/env bash

set -eE

shellcheck "$0"

# Before we build...
go fmt
go doc
go test

go build

strip -s sslexpirystatsd

cleanup() {
    echo "Cleaning up."
}

trap fail $(seq 1 64)

fail() {
    echo "FAIL: $1"
    cleanup
    exit 1
}

./sslexpirystatsd validate_configuration samples/valid_configuration.json
./sslexpirystatsd validate_configuration samples/invalid_configuration.json && fail "We thought an invalid configuration was valid."

echo Success
