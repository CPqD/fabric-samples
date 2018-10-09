#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Stops the network.
# Usage example: ./down.sh

# Source env
. ./env.sh

docker-compose -f ${COMPOSE_FILE} down
