#!/usr/bin/env bash

set -euxo pipefail

SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${SCRIPT_PATH}

export DROXY_UID=$(id -u)
export DROXY_GID=$(id -g)
export DROXY_TEST_VAR="this tests variable substitution in config plus passing variables to the container"
export DROXY_GLOBAL_OVERWRITE="this is supposed to be overwritten by global config"
export DROXY_LOCAL_OVERWRITE="this is supposed to be overwritten by command config"
export DROXY_MOUNT_DIR="/tmp"
export DROXY_HOST_TEST_DIR="/tmp/droxy-functional-tests"



docker version

./droxy clones -f

venom run tests.local.yml --env true --strict