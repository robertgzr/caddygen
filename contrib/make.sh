#!/bin/sh

set -e

EXTLD_FLAGS='-extldflags="-fno-PIC -static"'
LD_FLAGS="-ldflags='${EXTLD_FLAGS} -s'"
GO_FLAGS="${LD_FLAGS} -buildmode pie -tags 'netgo static_build'"

# static build
CGO_ENABLED=0 go build "${GO_FLAGS}" -o caddygen

# container image
docker build -f contrib/Dockerfile -t ${IMAGE_TAG:-caddygen} .
