#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Updates anchor peers of all organizations.
# This means that aditional configurations will be done in the genesis block.
# This script invokes scripts/updateAnchors.sh which is executed inside CLI container.

# Usage: ./updateAnchors.sh

# Source env.sh
. env.sh

docker exec cli scripts/updateAnchors.sh $CHANNEL_NAME ${CLI_DELAY} ${CC_LANG} ${CLI_TIMEOUT} ${VERBOSE} ${DOMAIN} ${ORG_LIST[@]}


