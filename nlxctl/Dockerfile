# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building nlxctl
COPY nlxctl                                         /go/src/go.nlx.io/nlx/nlxctl
COPY config-api                                     /go/src/go.nlx.io/nlx/config-api
COPY common                                         /go/src/go.nlx.io/nlx/common
COPY .git                                           /go/src/go.nlx.io/nlx/nlxctl/.git

WORKDIR /go/src/go.nlx.io/nlx/nlxctl

RUN GIT_COMMIT_HASH=$(git rev-parse HEAD) && \
        GIT_TAG_NAME=$(git describe --tags) && \
        go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME' " \
        -o dist/bin/nlxctl ./

# Release binary on latest alpine image.
FROM alpine:latest

RUN apk add --no-cache ca-certificates bash

COPY --from=build /go/src/go.nlx.io/nlx/nlxctl/dist/bin/nlxctl /usr/local/bin/nlxctl

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser