#!/usr/bin/env bash

set -ex

APP_ROOT="$( cd "$(dirname "$0")"; cd .. ; pwd -P )"

docker version

testImage="oppodelldog/droxy:latest"
containerName="droxy-ci-${RANDOM}"
containerDir="/go/src/github.com/Oppodelldog/droxy"

docker run --rm  --name ${containerName} -v"${APP_ROOT}":${containerDir} -w${containerDir} ${testImage} make ci

testImage="oppodelldog/docker-bats:latest"
containerName="droxy-ci-functional-${RANDOM}"
docker run --rm  --name ${containerName} -v"/var/run/docker.sock:/var/run/docker.sock" -v"${APP_ROOT}:${containerDir}" -w${containerDir} ${testImage} .test/run-tests.sh

