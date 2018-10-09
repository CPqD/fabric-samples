#!/bin/bash
#
# Copyright 2018 CPqD. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a convenience script to prepare and start the network
# and install and instantiate the chaincodes.
# Usage: ./start.sh

# Source env.sh
. ./env.sh

while getopts ":l:" opt; do  
  case "$opt" in  
  l)
    CC_LANG=${OPTARG}
    ;;
  \?)
    echo "Invalid option: -${OPTARG}." >&2
    ;;
  esac
done

if [ "${CC_LANG,,}" = "go" ]; then
    echo
    echo "Installing Go dependencies…"
    set -x
    go get -u github.com/gogo/protobuf/proto
    go get -u github.com/hyperledger/fabric/protos/msp
    go get -u github.com/hyperledger/fabric/core/chaincode
    set +x
fi
./generate.sh && \
./up.sh && \
./createChannel.sh && \
./joinChannel.sh && \
./updateAnchors.sh && \
./installChaincode.sh -l ${CC_LANG} -n farm -v 1.0 -p 0 -o "slaughterhouse farm1 farm2 inspection1" && \
./installChaincode.sh -l ${CC_LANG} -n farm -v 1.0 -p 1 -o "slaughterhouse farm1 farm2 inspection1" && \
./instantiateChaincode.sh -l ${CC_LANG} -n farm -v 1.0 -p 0 -o farm1 -P "OR ('SlaughterhouseMSP.member', 'Farm1MSP.member','Farm2MSP.member', 'Inspection1MSP.member')" -c collections_config.json && \
./installChaincode.sh -l ${CC_LANG} -n slaughterhouse -v 1.0 -p 0 -o "slaughterhouse inspection1 supermarket1 inspection2" && \
./installChaincode.sh -l ${CC_LANG} -n slaughterhouse -v 1.0 -p 1 -o "slaughterhouse inspection1 supermarket1 inspection2" && \
./instantiateChaincode.sh -l ${CC_LANG} -n slaughterhouse -v 1.0 -p 0 -o slaughterhouse -P "OR ('SlaughterhouseMSP.member', 'Inspection1MSP.member','Inspection2MSP.member', 'Supermerket1MSP.member')" -c collections_config.json
