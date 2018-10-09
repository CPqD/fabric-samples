#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a convenience script to change org and peer inside CLI container.
# Default values are FOUNDER_ORG and peer0.

# Source environment variables
. /opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/.env
# Source utils
. /opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/utils.sh

ORG=${1:-$FOUNDER_ORG}
PEER=${2:-0}

setGlobals ${PEER} ${ORG}
export CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/${ORG}.${DOMAIN}/peers/peer${PEER}.${ORG}.${DOMAIN}/tls/server.key
export CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/${ORG}.${DOMAIN}/peers/peer${PEER}.${ORG}.${DOMAIN}/tls/server.crt

echo "### Peer and Organization changed to peer${PEER}.${ORG} ###"

