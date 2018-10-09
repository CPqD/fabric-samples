#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Source environment variables
. /opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/.env
# Source utils
. ./scripts/utils.sh

CHANNEL_NAME="$1"
DELAY="$2"
CC_LANG="$3"
TIMEOUT="$4"
VERBOSE="$5"
DOMAIN="$6"
shift 6
ORG_LIST=( "$@" )

joinChannel () {
	for org in "${ORG_LIST[@]}"; do
	    for peer in 0 1; do
		joinChannelWithRetry $peer $org
		echo "===================== peer${peer}.${org} joined channel '${CHANNEL_NAME}' ===================== "
		sleep $DELAY
		echo
	    done
	done
}

joinChannel
