#!/bin/bash -e
GO_MAJOR_VERSION="go1"
GO_MINOR_VERSION="12"
GO_VERSION="$GO_MAJOR_VERSION.$GO_MINOR_VERSION" # Minimal required go version
NETRC="$HOME/.netrc"

die() { echo "$*" 1>&2 ; exit 1; }

echo
echo "Running preflight checks..."
# Validate go exists and version >= 1.12
echo
echo "Checking go binary..."
which go || die "go is not installed or not found in \$PATH. Please install from here with minimal version of $GO_VERSION https://golang.org/doc/install"
echo "go binary ok"
echo

echo "Checking go version is at least >= $GO_VERSION"
major=$(go version | grep "version go" | cut -d' ' -f3 | cut -d. -f1)
minor=$(go version | grep "version go" | cut -d' ' -f3 | cut -d. -f2)
echo "major: $major"
echo "minor: $minor"
[ "$major" == "$GO_MAJOR_VERSION" ] || die "Go major version should be $GO_MAJOR_VERSION"
[ "$minor" -ge "$GO_MINOR_VERSION" ] || die "Go minor version should be at least $GO_MINOR_VERSION"
echo "OK, version check passed, found $major.$minor"
echo

echo "Making sure gofmt is installed"
which gofmt || die "gofmt is not installed or not found in the \$PATH. Please install from here with minimal version of $GO_VERSION https://golang.org/doc/install"
echo "OK, gofmt validated"
echo

echo "Making sure goimports is installed"
which gofmt || die "goimports is not installed or not found in the \$PATH. Please install: go get golang.org/x/tools/cmd/goimports"
echo "OK, gofmt validated"
echo

echo "Checking that \$GOPATH is defined"
[ -z "${GOPATH}" ] && die "\$GOPATH is not defined. Please set it up https://github.com/golang/go/wiki/SettingGOPATH"
echo "OK, \$GOPATH is defined: \$GOPATH=$GOPATH"
echo

echo "Done all preflight checks - PASSED"
echo
