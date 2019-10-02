# BEGIN __DO_NOT_INCLUDE__
GO_ARCHETYPE_VERSION := 0.1.2
GO_ARCHETYPE_DIR := .tmp/go-archetype-$(GO_ARCHETYPE_VERSION)
GO_ARCHETYPE := $(GO_ARCHETYPE_DIR)/go-archetype
# END __DO_NOT_INCLUDE__

# BEGIN __INCLUDE_GRPC__
PROTOC_VERSION := 3.8.0
RELEASE_OS :=
PROTOC_DIR := .tmp/protoc-$(PROTOC_VERSION)
PROTOC_BIN := $(PROTOC_DIR)/bin/protoc
# END __INCLUDE_GRPC__

GO := go
ifdef GO_BIN
	GO = $(GO_BIN)
endif

GOLANGCI_LINT_VERSION := v1.18.0
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

GIT_COMMIT := $(shell git rev-parse --short HEAD 2> /dev/null || echo "no-revision")
GIT_COMMIT_MESSAGE := $(shell git show -s --format='%s' 2> /dev/null | tr ' ' _ | tr -d "'")
GIT_TAG := $(shell git describe --tags 2> /dev/null || echo "no-tag")
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null || echo "no-branch")
BUILD_TIME := $(shell date +%FT%T%z)
VERSION_PACKAGE := github.com/rantav/go-template/version

# BEGIN __INCLUDE_GRPC__
ifeq ($(OS),Windows_NT)
	echo "Windows not supported yet, sorry...."
	exit 1
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		RELEASE_OS = linux
	endif
	ifeq ($(UNAME_S),Darwin)
		RELEASE_OS = osx
	endif
endif
# END __INCLUDE_GRPC__

all: test lint

tidy:
	$(GO) mod tidy -v

fmt:
	gofmt -s -w .

build: protoc
	$(GO) build ./...

# Use this target when building in CI so that you get all the lovely version information in the resulting executable
ci-build: protoc
	$(GO) build -v -ldflags '-X $(VERSION_PACKAGE).GitHash=$(GIT_COMMIT) -X $(VERSION_PACKAGE).GitTag=$(GIT_TAG) -X $(VERSION_PACKAGE).GitBranch=$(GIT_BRANCH) -X $(VERSION_PACKAGE).BuildTime=$(BUILD_TIME) -X $(VERSION_PACKAGE).GitCommitMessage=$(GIT_COMMIT_MESSAGE)'

test: build
	$(GO) test -cover -race -v ./...

test-coverage:
	$(GO) test ./... -race -coverprofile=.testCoverage.txt && $(GO) tool cover -html=.testCoverage.txt

ci-test: ci-build
	$(GO) test -race $$($(GO) list ./...) -v -coverprofile .testCoverage.txt

# BEGIN __DO_NOT_INCLUDE__
# Template generation and test section, do not include it in the generated project

GITHUB_BASE := github.com
TMP_PROJECT_NAME=my-go-project
TMP_DIR=.tmp/go/$(TMP_PROJECT_NAME)

help:
	@echo TODO: Add help here

init: $(GO_ARCHETYPE) guard-PROJECT_NAME guard-GROUP_NAME guard-DESTINATION
	# DESTINATION, PROJECT_NAME and GROUP_NAME are provided externally like so:
	# make init PROJECT_NAME=rantav GROUP_NAME=my-go-project DESTINATION=/Users/ran/dev/src/github.com/rantav/my-go-project
	LOG_LEVEL=info $(GO_ARCHETYPE) \
		transform \
		--transformations transformations.yml \
		--source=. \
		--destination $(DESTINATION)  \
		-- \
		--name $(PROJECT_NAME) \
		--repo_path $(GITHUB_BASE)/$(GROUP_NAME)

cleanup-test-dir:
	rm -rf $(TMP_DIR)
	mkdir -p $(TMP_DIR)

test-template:
	make test-template-no-grpc
	make test-template-with-grpc

test-template-no-grpc: cleanup-test-dir $(GO_ARCHETYPE)
	LOG_LEVEL=info \
	$(GO_ARCHETYPE) \
		transform \
		--transformations transformations.yml \
		--source=. \
		--destination $(TMP_DIR)  \
		-- \
		--name $(TMP_PROJECT_NAME) \
		--repo_path github.com/rantav \
		--description "My awesome go project" \
		--include_grpc no
	cd $(TMP_DIR) &&\
		make

test-template-with-grpc: cleanup-test-dir $(GO_ARCHETYPE)
	LOG_LEVEL=info \
	$(GO_ARCHETYPE) \
		transform \
		--transformations transformations.yml \
		--source=. \
		--destination $(TMP_DIR)  \
		-- \
		--name $(TMP_PROJECT_NAME) \
		--repo_path github.com/rantav \
		--description "My awesome go project" \
		--include_grpc yes
	cd $(TMP_DIR) &&\
		make

$(GO_ARCHETYPE):
	rm -rf $(GO_ARCHETYPE_DIR)
	mkdir -p $(GO_ARCHETYPE_DIR)
	cd $(GO_ARCHETYPE_DIR) &&\
		curl -sL https://github.com/rantav/go-archetype/releases/download/v$(GO_ARCHETYPE_VERSION)/go-archetype_$(GO_ARCHETYPE_VERSION)_$(RELEASE_OS)_x86_64.tar.gz | tar xz
	# Test installation
	$(GO_ARCHETYPE) --help

# END __DO_NOT_INCLUDE__

setup: setup-validations setup-git-hooks $(GO_ARCHETYPE)

setup-git-hooks:
	git config core.hooksPath .githooks

setup-validations:
	# A set of validations to help the first time use kick in
	scripts/preflight-checks.sh

# BEGIN __INCLUDE_GRPC__
$(PROTOC_BIN): protoc-gen-go
	@echo "Installing unzip (if required)"
	@which unzip || apt-get update || sudo apt-get update
	@which unzip || apt-get install unzip || sudo apt-get install unzip
	@echo Installing protoc
	rm -rf $(PROTOC_DIR)
	mkdir -p $(PROTOC_DIR)
	cd $(PROTOC_DIR) &&\
		curl -OL https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(RELEASE_OS)-x86_64.zip &&\
		unzip protoc-$(PROTOC_VERSION)-$(RELEASE_OS)-x86_64.zip
	chmod +x $(PROTOC_BIN)

protoc-gen-go:
	@echo "Installing protoc-gen-go (if required)"
	@which protoc-gen-go > /dev/null || GO111MODULE=on $(GO) get -u github.com/golang/protobuf/protoc-gen-go

# END __INCLUDE_GRPC__

lint: $(GOLANGCI_LINT)
	$(GOPATH)/bin/golangci-lint run --fast --enable-all -D gochecknoglobals -D gochecknoinits

$(GOLANGCI_LINT):
	GO111MODULE=on $(GO) get github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

run-server:
	# After running, try:
	#	curl localhost:8081/_/health.json
	#	curl localhost:8081/_/metrics.json
	deployments/dev.sh

run-client:
	# BEGIN __INCLUDE_GRPC__
	# Run a simple client that connects to the server above and runs a single healthcheck
	$(GO) run main.go test-grpc-client --server-address 127.0.01:8080
	@echo
	# END __INCLUDE_GRPC__
	@echo Testing health...
	curl localhost:8081/_/health.json
	@echo
	@echo Checking metrics...
	curl localhost:8081/_/metrics.json

# BEGIN __INCLUDE_GRPC__
protoc: $(PROTOC_BIN) protoc-gen-go
	mkdir -p internal/generated/grpc
	$(PROTOC_BIN) --proto_path=grpc/idl/ --go_out=plugins=grpc:internal/generated/grpc grpc/idl/*.proto
# END __INCLUDE_GRPC__

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		exit 1; \
	fi
