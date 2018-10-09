#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This creates the channel that will be used by the organizations

# Source utils
. ./scripts/utils.sh

CHANNEL_NAME="$1"
DELAY="$2"
CC_LANG="$3"
TIMEOUT="$4"
VERBOSE="$5"
DOMAIN="$6"
ORG="$7"

createChannel() {
	setGlobals 0 ${ORG}

	if [ -z "${CORE_PEER_TLS_ENABLED}" -o "${CORE_PEER_TLS_ENABLED}" = "false" ]; then
                set -x
		peer channel create -o orderer.${DOMAIN}:7050 -c ${CHANNEL_NAME} -f ./channel-artifacts/channel.tx >&log.txt
		res=$?
                set +x
	else
				set -x
		peer channel create -o orderer.${DOMAIN}:7050 -c ${CHANNEL_NAME} -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
		res=$?
				set +x
	fi
	cat log.txt
	verifyResult $res "Channel creation failed"
	echo "===================== Channel '${CHANNEL_NAME}' created ===================== "
	echo
}

echo "Creating channel ${CHANNEL_NAME}"
createChannel

