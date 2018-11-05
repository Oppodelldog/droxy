#!/usr/bin/env bash

# run functional tests locally
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -w /test -v "$(pwd)":/test oppodelldog/docker-bats test.sh
