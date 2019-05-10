#!/usr/bin/env bash

set -euxo pipefail

export DROXY_FILE_SIZE=$( stat --printf="%s" droxy)

venom run tests.yml --env true 
