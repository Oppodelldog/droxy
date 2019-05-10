#!/usr/bin/env bash

export DROXY_FILE_SIZE=$( stat --printf="%s" droxy)


venom run tests.yml --env true 
