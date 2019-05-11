#!/usr/bin/env bash

set -euxo pipefail

export DROXY_FILE_SIZE=$(stat -c "%s" droxy)

docker version
env

venom run tests.yml --env true --strict