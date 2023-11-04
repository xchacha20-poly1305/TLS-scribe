#!/bin/bash

if [[ -z "$TAG_NAME" ]]; then
    TAG_NAME=$(git rev-parse --short HEAD || echo "Unknow")
fi

export CGO_ENABLED=0
platforms=("darwin" "linux" "windows" "freebsd" "android")
architectures=("amd64" "386" "arm" "arm64")

rm -rf ./build/releases
mkdir -p ./build/releases

for os in "${platforms[@]}"; do

    for arch in "${architectures[@]}"; do
        output_name="scribe-${os}-${arch}"
        if [ "$os" == "windows" ]; then
            output_name="${output_name}.exe"
        fi

        echo "Building ${output_name}..."
        GOOS="$os" GOARCH="$arch" go build -o "build/releases/${output_name}" \
            -trimpath -ldflags "-w -s -X main.VERSION=${TAG_NAME}" ./cmd/scribe
    done

done
