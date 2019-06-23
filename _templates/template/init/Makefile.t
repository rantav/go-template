---
to: Makefile
---

all: test

build:
	go build ./...

test: build
	go test -race -v ./...

run-server:
	go run main.go serve