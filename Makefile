BINARY=go-refresh

build:
	go build -o $(BINARY) ./cmd/go-refresh

test:
	go test ./...

fmt:
	go fmt ./...

.PHONY: build run test fmt lang keywords clean