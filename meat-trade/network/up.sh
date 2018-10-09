#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Starts the containers defined in yaml config files.
# Usage: ./up.sh

# Source env
. ./env.sh

docker-compose -f ${COMPOSE_FILE} up -d 2>&1

if [ $? -ne 0 ]; then
    echo "ERROR! Unable to start network."
    exit 1
fi
