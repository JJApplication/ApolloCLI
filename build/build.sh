#!/usr/bin/env bash

# build for app
echo "GOPROXY: $(go env GOPROXY)"
echo "GOOS: $(go env GOOS)"
echo "GOARCH: $(go env GOARCH)"

date -u

echo "start to build"
build_version=$(date "+%Y-%m-%d")
go build -ldflags "-w -s -X github.com/JJApplication/ApolloCLI._version=$build_version" -trimpath -o apollocli ./cmd
echo "done"
