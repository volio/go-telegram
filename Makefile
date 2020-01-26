BUILD_DIR := build
BINARY_NAME := go-common

#environment
export GOPRIVATE=github.com/volio
export GOPROXY=https://goproxy.io

.PHONY: lint build

fmt:
	goimports -local -l -w .

lint:
	golangci-lint run --timeout=3m ./...

build:
	go build ./...
