#!/bin/bash

set -e
set -x

GOLDFLAGS="-s -w"

# dependencies: rsrc (go get github.com/akavel/rsrc)
rsrc -manifest execute-elevated.exe.manifest
GOOS=windows GOARCH=386 go build -ldflags="$GOLDFLAGS" -o execute-elevated.exe

GOOS=linux GOARCH=amd64 go build -ldflags="$GOLDFLAGS" -o execute-elevated
