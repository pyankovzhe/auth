.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: test
test:
	go test -v -race -timeout 10s ./...

.PHONY: check
check:
	golangci-lint run

.DEFAULT_GOAL := build
