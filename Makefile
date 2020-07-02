
IMAGE ?= docker.io/robertgzr/caddygen:latest
GO_SRCS = $(shell find . -name '*.go')

all: caddygen

caddygen: $(GO_SRCS)
	go build -o $@ .

.PHONY: clean
clean:
	go clean .

.PHONY: container
container:
	docker build -t $(IMAGE) .
