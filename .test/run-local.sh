#!/usr/bin/env bash

# run functional tests locally
docker run \
    --rm \
    -it \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v "${PWD}":/go/src/github.com/Oppodelldog/droxy \
    -w /go/src/github.com/Oppodelldog/droxy \
    --entrypoint make \
    oppodelldog/droxy:functional-tests \
    functional-tests 
