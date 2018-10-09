#!/bin/bash
#
# Copyright 2018 CPqD. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a convenience script to prepare and start the network,
# install and instantiate the chaincodes, start the go middleware
# and start the web appplication.
# Usage: ./runAll.sh

BASE_DIR=${PWD}
CC_LANG="go"

# Ask user to choose a language for chaincode.
# Default is Go as defined in env.sh
function getLang() {
    read -p "Choose a language for chaincode (Go or Node): [G/n] " lang
    case $lang in
        g | G | "")
            echo "- Using Go as chaincode language."
            ;;
        n | N)
            echo "- Using Node as chaincode language."
            CC_LANG="node"
            ;;
    *)
            echo "- Invalid response. Try again."
            getLang
            ;;
  esac
}

# ask user to choose a language for chaincode
echo
getLang

# start network
cd ${BASE_DIR}/network
./start.sh -l ${CC_LANG} && \

# start middleware
cd ${BASE_DIR}/middleware && \
docker-compose up -d && \

# start application
cd ${BASE_DIR}/web-application && \
./startApp.sh
