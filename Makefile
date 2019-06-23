TMP_PROJECT_NAME=my-go-project
TMP_DIR=/tmp/go/$(TMP_PROJECT_NAME)

all: test

build:
	go build ./...

test: build
	go test -race -v ./...

run-server:
	go run main.go serve

test-tempalte:
	rm -rf $(TMP_DIR)
	mkdir $(TMP_DIR); \
	cp -r * $(TMP_DIR); \
	mkdir -p $(TMP_DIR); \
	cd $(TMP_DIR); \
	HYGEN_OVERWRITE=1 hygen template init $(TMP_PROJECT_NAME) \
		--repo_path=gitlab.appsflyer.com/rantav \
		--long_description="My awesome go project" \
		;\
	make
