#!/usr/bin/env bash

set -euxo pipefail

export VENOM_VAR_DROXY_FILE_SIZE=$(stat -c "%s" droxy)

docker version

venom run tests.yml