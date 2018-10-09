#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Adds all peers 0 of all organizations to the channel $CHANNEL_NAME.
# This script invokes scripts/joinChannel.sh which is executed inside CLI container.

# Usage: ./joinChannel.sh

# Source env.sh
. env.sh

# Waits 5 minutes for channel creation and propagation
sleep 5
docker exec cli scripts/joinChannel.sh ${CHANNEL_NAME} ${CLI_DELAY} ${CC_LANG} ${CLI_TIMEOUT} ${VERBOSE} ${DOMAIN} ${ORG_LIST[@]}
