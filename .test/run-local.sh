#!/usr/bin/env bash

set -euxo pipefail

export DROXY_UID=$(id -u)
export DROXY_GID=$(id -g)
export DROXY_TEST_VAR="this tests variable substitution in config plus passing variables to the container"
export DROXY_MOUNT_DIR="/tmp"

docker version

venom run tests.local.yml --env true --strict