#!/usr/bin/env bash

# run functional tests locally
docker run \
    --rm \
    -it \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -w /go/src/github.com/Oppodelldog/droxy  \
    -v "$(pwd)":/go/src/github.com/Oppodelldog/droxy \
    --entrypoint .test/run-tests.sh \
    oppodelldog/docker-bats 
