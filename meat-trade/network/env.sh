#!/bin/bash
#
# Copyright 2018 CPqD. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
# Defines environment variables (default values)

# Obtain the OS and Architecture string that will be used to select the correct
# native binaries for your platform, e.g., darwin-amd64 or linux-amd64
OS_ARCH=$(echo "$(uname -s | tr '[:upper:]' '[:lower:]' | sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
# How many seconds the CLI should wait for a response from another container before giving up
export CLI_TIMEOUT=50
# Delay in seconds between commands
export CLI_DELAY=3
# Channel name
export CHANNEL_NAME="kingbeefcattlechannel"
# Domain name
export DOMAIN="kingbeefcattle.com"
# Founder organization name
export FOUNDER_ORG="slaughterhouse"
# List of organizations
export ORG_LIST=("slaughterhouse" "farm1" "farm2" "inspection2" "inspection1" "supermarket1")
# Default docker-compose yaml definition
export COMPOSE_FILE=docker-compose-cli.yaml
# Default language for chaincodes
export CC_LANG=go
# Default image tag
export IMAGE_TAG="latest"

export COMPOSE_PROJECT_NAME=network
export PATH=${PWD}/../bin:${PWD}:${PATH}
export FABRIC_CFG_PATH=${PWD}
export VERBOSE=true
export DEVMODE=""
# To work in devmode, uncomment the next line below
#export DEVMODE="--peer-chaincodedev"

# Define GOPATH if not defined or empty (just for shure)
if [ -z "$GOPATH" ]; then
    export GOPATH=${HOME}/go
fi

# Copy some environment variables to scripts dir making them accesible inside the containers
if [ -e ./scripts/.env ]; then
    rm ./scripts/.env
fi

echo "CLI_TIMEOUT=${CLI_TIMEOUT}" >> ./scripts/.env
echo "CLI_DELAY=${CLI_DELAY}" >> ./scripts/.env
echo "CHANNEL_NAME=${CHANNEL_NAME}" >> ./scripts/.env
echo "DOMAIN=${DOMAIN}" >> ./scripts/.env
echo "FOUNDER_ORG=${FOUNDER_ORG}" >> ./scripts/.env
