#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Creates a channel named $CHANNEL_NAME
# This script invokes scripts/createChannel.sh which is executed inside CLI container.

# Usage example: ./createChannel.sh

# Source env.sh
. ./env.sh

docker exec cli scripts/createChannel.sh ${CHANNEL_NAME} ${CLI_DELAY} ${CC_LANG} ${CLI_TIMEOUT} ${VERBOSE} ${DOMAIN} ${FOUNDER_ORG}
