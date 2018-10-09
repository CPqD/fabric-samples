#!/bin/bash
#
# Copyright 2018. CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Installs a chaincode in peers 0 of all organizations.
# 
# This script invokes scripts/installChaincode.sh which is executed inside CLI container.
# The chaincode name and version can be specified as script options.
# There are default values for all options.
# The list of organizations can be specified as option too.
# The default for this option is the list of all organizations supplied in env.sh.
# The chaincode default dir is fabric-samples/meat-trade/chaincode.

# Usage example:
#     ./installChaincode.sh -l <language> -n <chaincodeName> -v <chaincodeVersion> -p <peerNumber> -o <org1 org2>

# Source env.sh
. ./env.sh

# Default chaincode language
CC_LANG="go"

# Default chaincode name
CC_NAME="farm"

# Default chaincode version
CC_VERSION="1.0"

# Peer number
PEER="0"

# Organizations list where the chaincode will be installed
ORGS=${ORG_LIST[@]}

while getopts ":n:v:l:p:o:" opt; do  
  case "$opt" in  
  l)
    CC_LANG=$OPTARG
    ;;
  n)
    CC_NAME=$OPTARG
    ;;
  v)
    CC_VERSION=$OPTARG
    ;;
  p)
    PEER=$OPTARG
    ;;
  o)
    ORGS=()
    for var in $OPTARG
    do     
     ORGS+=( "$var" )
    done
    ;;   
  \?)
    echo "Invalid option: -${OPTARG}." >&2
    ;;
  esac
done

echo ""
echo "Installing chaincode $CC_NAME version $CC_VERSION in peer$PEER em ${ORGS[@]}."
echo ""
docker exec cli scripts/installChaincode.sh ${CC_LANG} ${CC_NAME} ${CC_VERSION} $PEER ${ORGS[@]}

