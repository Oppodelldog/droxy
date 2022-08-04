#!/usr/bin/env bash

set -euxo pipefail

SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${SCRIPT_PATH}

export DROXY_MOUNT_DIR="/tmp"
export DROXY_HOST_TEST_DIR="/tmp/droxy-functional-tests"
export DROXY_TEST_VAR="this tests variable substitution in config plus passing variables to the container"
export VENOM_VAR_DROXY_UID=$(id -u)
export VENOM_VAR_DROXY_GID=$(id -g)
export VENOM_VAR_DROXY_TEST_VAR=$DROXY_TEST_VAR
export VENOM_VAR_DROXY_MOUNT_DIR=$DROXY_MOUNT_DIR
export VENOM_VAR_DROXY_HOST_TEST_DIR=$DROXY_HOST_TEST_DIR


docker version

./droxy clones -f

venom run tests.local.yml