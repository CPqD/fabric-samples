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

updateAnchors () {
	for org in "${ORG_LIST[@]}"; do	    
		updateAnchorPeers 0 ${org}
		echo "===================== peer${peer}.${org} anchor updated ===================== "
		sleep ${DELAY}
		echo	    
	done
}

updateAnchors
