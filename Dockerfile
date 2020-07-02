
FROM golang:alpine AS gobuild
COPY . /src
WORKDIR /src
RUN set -ex; \
    CGO_ENABLED=0 \
    go build \
	-ldflags '-s -extldflags=-static' \
	-tags 'netgo osusergo static_build' \
	-o /caddygen \
	.

FROM scratch
COPY --from=gobuild /caddygen /bin/caddygen
ENTRYPOINT ["/bin/caddygen"]
