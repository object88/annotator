#!/usr/bin/env bash
set -e

export GO111MODULE=on

if [ $(uname) == "Darwin" ]; then
  # Cannot do a static compilation on Darwin.
  go build -ldflags "-s -w" -o bin/db main/main.go
else
  export CGO_ENABLED=0
  go build -tags "netgo" -ldflags "-s -w -extldflags \"-static\"" -o bin/db main/main.go
fi
