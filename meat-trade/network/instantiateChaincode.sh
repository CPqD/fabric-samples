#!/bin/bash
#
# Copyright 2018. CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Instantiates a chaincode in a peer of an organization.
# 
# This script invokes scripts/instantiateChaincode.sh which is executed inside CLI container.
# The chaincode name and version can be specified as script options.
# There are default values for all options.

# Usage example:
#     ./instantiateChaincode.sh -l <language> -n <chaincodeName> -v <chaincodeVersion> -p <peerNumber> -o <org>

# Source env.sh
. ./env.sh

# Default chaincode language
CC_LANG="go"

# Default chaincode name
CC_NAME="farm"

# Default chaincode version
CC_VERSION="1.0"

# Peer number where the chaincode will be instantiated
PEER="0"

# Organization where the chaincode will be instantiated
ORG=slaughterhouse

# Default init arguments
ARGS='{"Args":[]}'

# Default policy (none)
POLICY=""

# Collections config
COLLECTIONS_CONFIG=""

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
    ARGS="'"${OPTARG}"'"    
    ;;       
  P)    
    POLICY=${OPTARG}    
    ;;
  c)    
    COLLECTIONS_CONFIG="-c ${OPTARG}"
    ;; 
  \?)
    echo "Invalid option: -${OPTARG}." >&2
    ;;
  esac
done

echo ""
echo "Instantiating chaincode ${CC_NAME} version ${CC_VERSION} in peer${PEER}.${ORG}"
echo ""
docker exec cli scripts/instantiateChaincode.sh -l ${CC_LANG} -n ${CC_NAME} -v ${CC_VERSION} -p ${PEER} -o ${ORG} -a ${ARGS} -P "${POLICY}" ${COLLECTIONS_CONFIG}

