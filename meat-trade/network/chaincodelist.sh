#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Lists installed (default) ou instantiated chaincodes.
# Usage examples:
#    ./chaincodelist.sh
#    ./chaincodelist.sh instantiated

# Source env
. ./env.sh

OPTION=${1:-"installed"}
CMD1=""
CMD2="peer chaincode list --$OPTION --channelID $CHANNEL_NAME"

if [ "$#" -gt 1 ]; then
    ORG=$2
    PEER=""
    if [ "$#" -gt 2 ]; then
       PEER=$3
    fi
    CMD1=". ./scripts/changeOrgAndPeer.sh $ORG $PEER;"
fi

docker exec -it cli bash -c "$CMD1 $CMD2"
