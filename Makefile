ifeq ($(CI), true)
	GO = go
else
	GO = go1.21.0
endif


all: build

build:
	$(GO) build -ldflags="-s -w" ./cmd/server

generate:
	$(GO) generate ./...


style:
	$(GO) fmt ./...
	golangci-lint run

test:
	$(GO) test ./...

