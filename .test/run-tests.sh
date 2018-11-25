#!/usr/bin/env bash

set -e

docker version
go version

go get ./...
buildFilePath=".test/droxy"
go build -o ${buildFilePath}
trap 'rm -f ${buildFilePath}' ERR EXIT


.test/bats-tests.sh

