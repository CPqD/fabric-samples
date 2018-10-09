#
# Copyright IBM Corp All Rights Reserved
# Modifications Copyright CPqD All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

COUNTER=1
MAX_RETRY=5

# Verify the result of the execution
verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to execute ==========="
    echo
    exit 1
  fi
}

# Set peers globals
setGlobals() {
  PEER=$1
  ORG=$2

  ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/${DOMAIN}/orderers/orderer.${DOMAIN}/msp/tlscacerts/tlsca.${DOMAIN}-cert.pem
  CORE_PEER_LOCALMSPID="${ORG[@]^}MSP"
  CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/${ORG}.${DOMAIN}/peers/peer${PEER}.${ORG}.${DOMAIN}/tls/ca.crt
  CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/${ORG}.${DOMAIN}/users/Admin@${ORG}.${DOMAIN}/msp
  CORE_PEER_ADDRESS=peer${PEER}.${ORG}.${DOMAIN}:7051
  
  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

updateAnchorPeers() {
  PEER=$1
  ORG=$2
  setGlobals ${PEER} ${ORG}

  if [ -z "${CORE_PEER_TLS_ENABLED}" -o "${CORE_PEER_TLS_ENABLED}" = "false" ]; then
    set -x
    peer channel update -o orderer.${DOMAIN}:7050 -c ${CHANNEL_NAME} -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx >&log.txt
    res=$?
    set +x
  else
    set -x
    peer channel update -o orderer.${DOMAIN}:7050 -c ${CHANNEL_NAME} -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls ${CORE_PEER_TLS_ENABLED} --cafile ${ORDERER_CA} >&log.txt
    res=$?
    set +x
  fi
  cat log.txt
  verifyResult $res "Anchor peer update failed"
  echo "===================== Anchor peers updated for org '${CORE_PEER_LOCALMSPID}' on channel '${CHANNEL_NAME}' ===================== "
  sleep $DELAY
  echo
}

# Sometimes Join takes time hence RETRY at least 5 times
joinChannelWithRetry() {
  PEER=$1
  ORG=$2
  setGlobals ${PEER} ${ORG}

  set -x
  peer channel join -b ${CHANNEL_NAME}.block >&log.txt
  res=$?
  set +x
  cat log.txt
  if [ $res -ne 0 -a ${COUNTER} -lt ${MAX_RETRY} ]; then
    COUNTER=$(expr $COUNTER + 1)
    echo "peer${PEER}.${ORG} failed to join the channel, Retry after ${DELAY} seconds"
    sleep $DELAY
    joinChannelWithRetry $PEER $ORG
  else
    COUNTER=1
  fi
  verifyResult $res "After ${MAX_RETRY} attempts, peer${PEER}.${ORG} has failed to join channel '${CHANNEL_NAME}' "
}
