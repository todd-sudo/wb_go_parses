.PHONY: build

run:
	./cmd/bin/main

build:
	go build -v ./cmd/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build