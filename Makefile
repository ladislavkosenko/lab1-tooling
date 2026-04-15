.PHONY: fmt lint test build all

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...

build:
	go build -o bin/app.exe ./cmd/app/main.go

all: fmt lint test build