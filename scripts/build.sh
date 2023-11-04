#!/bin/bash

# set -x

TAGS=""
for i in "$@"; do
    # debug option
    [[ $1 == "debug" ]] && TAGS="${TAGS}debug," && continue
done

GitCommit=$(git rev-parse --short HEAD || echo "Unknow")

rm -rf build
mkdir -p build

go build -v -o build -trimpath -ldflags \
    "-w -s -X main.VERSION=${GitCommit}" -tags "$TAGS" ./cmd/scribe
