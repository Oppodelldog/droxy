#!/usr/bin/env bash

set -euxo pipefail

export DROXY_FILE_SIZE=$(stat -c "%s" droxy)
export DROXY_UID=$(id -u)
export DROXY_GID=$(id -g)
export DROXY_TEST_VAR="this tests variable substitution in config plus passing variables to the container"
docker version

venom run tests.yml --env true --strict