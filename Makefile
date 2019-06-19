all: test

build:
	go build ./...

test: build
	go test -v ./...

run-server:
	go run main.go serve