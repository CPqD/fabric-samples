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

CC_LANG=""
CC_NAME=""
CC_VERSION=""
PEER=""
ORG=""
ARGS=""
POLICY=""
POL_OPT=""
LANG_OPT="-l"

while getopts ":n:v:l:p:o:a:P:c:" opt; do  
  case "$opt" in  
  l)
    CC_LANG=${OPTARG}
    ;;
  n)
    CC_NAME=${OPTARG}
    ;;
  v)
    CC_VERSION=${OPTARG}
    ;;
  p)
    PEER=${OPTARG}
    ;;
  o)
    ORG=${OPTARG}
    ;;
  a)
    ARGS=${OPTARG}    
    ;;       
  P)
    POLICY=${OPTARG}
    POL_OPT="-P"
    ;;  
  c)            
    COLLECTIONS_CONFIG=${OPTARG}
    ;; 
  \?)
    echo "Invalid option: -${OPTARG}." >&2
    ;;
  esac
done

if [ "${CC_LANG}" = "go" ]; then

    LANG_OPT=""
    CC_LANG=""
fi

if [ ! -z "${COLLECTIONS_CONFIG}" ]
then  
  COLLECTIONS_CONFIG="--collections-config /opt/gopath/src/github.com/chaincode/${CC_NAME}/${COLLECTIONS_CONFIG}"
fi

setGlobals ${PEER} ${ORG}

# Tests if the chaincode was installed in peer${PEER}
INSTALLED=`peer chaincode list --installed | grep ${CC_NAME}`
if [ -z "${INSTALLED}" ]; then
   echo "===================== Chaincode is not installed on peer${PEER}.${ORG} on channel '${CHANNEL_NAME}' ===================== "
   exit 1
fi

if [ -z "${CORE_PEER_TLS_ENABLED}" -o "${CORE_PEER_TLS_ENABLED}" = "false" ]; then
   set -x
   peer chaincode instantiate -o orderer.${DOMAIN}:7050 -C ${CHANNEL_NAME} -n ${CC_NAME} ${LANG_OPT} ${CC_LANG} -v ${CC_VERSION} -c ${ARGS} ${POL_OPT} "${POLICY}" ${COLLECTIONS_CONFIG} >&log.txt
   res=$?
   set +x
else
  set -x
  peer chaincode instantiate -o orderer.${DOMAIN}:7050 --tls ${CORE_PEER_TLS_ENABLED} --cafile ${ORDERER_CA} -C ${CHANNEL_NAME} -n ${CC_NAME} ${LANG_OPT} ${CC_LANG} -v ${CC_VERSION} -c ${ARGS} ${POL_OPT} "${POLICY}" ${COLLECTIONS_CONFIG} >&log.txt
  res=$?
  set +x
fi
cat log.txt
verifyResult $res "Chaincode instantiation on peer${PEER}.${ORG} on channel '${CHANNEL_NAME}' failed"
echo "===================== Chaincode is instantiated on peer${PEER}.${ORG} on channel '${CHANNEL_NAME}' ===================== "
echo
