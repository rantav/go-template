#!/bin/bash -e

TMP="${TMPDIR:=/tmp}"

die () {
    echo >&2 "$@"
    echo >&2 "Usage: $0 github-username project-name"
    exit 1
}

[ "$#" -eq 2 ] || die "2 arguments required, only $# provided"


username=$1
project_name=$2

mkdir $project_name
DESTINATION=`pwd`/$project_name

rm -rf $TMP/go-template
git clone --depth 1 https://github.com/rantav/go-template.git $TMP/go-template
cd $TMP/go-template

make setup-validations

# Templatize
echo Creating project template
echo =========================
make init GROUP_NAME=$username PROJECT_NAME=$project_name DESTINATION=$DESTINATION

cd $DESTINATION
git init
# Build to validate
make
make setup-git-hooks

git add .
