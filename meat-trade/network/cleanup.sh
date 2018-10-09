#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Remove containers, images, volumes and the created network.
# Usage example: ./cleanup.sh

# Source env
. ./env.sh

# Get CONTAINER_IDS and remove them all.
function clearContainers() {
  CONTAINER_IDS=$(docker ps -a | awk -v domain="dev-peer.*.${DOMAIN}.*" '($2 ~ domain) {print $1}')
  if [ -z "${CONTAINER_IDS}" -o "${CONTAINER_IDS}" == " " ]; then
    echo "---- No containers available for deletion ----"
  else
    docker rm -f ${CONTAINER_IDS}
  fi
}

# Remove all created images.
function removeUnwantedImages() {
  DOCKER_IMAGE_IDS=$(docker images | awk -v domain="dev-peer.*.${DOMAIN}.*" '($1 ~ domain) {print $3}')
  if [ -z "${DOCKER_IMAGE_IDS}" -o "${DOCKER_IMAGE_IDS}" == " " ]; then
    echo "---- No images available for deletion ----"
  else
    docker rmi -f ${DOCKER_IMAGE_IDS}
  fi
}

# Stop the network and remove the volumes.
function cleanUp() {
  # stop containers, just in case
  docker-compose -f ${COMPOSE_FILE} down --volumes --remove-orphans

  # Bring down the network, deleting the volumes
  # Delete any ledger backups
  docker run -v $PWD:/tmp/meat-trade --rm hyperledger/fabric-tools:$IMAGE_TAG rm -Rf /tmp/meat-trade/ledgers-backup
  # Cleanup the chaincode containers
  clearContainers
  # Cleanup images
  removeUnwantedImages
  # Remove orderer block and other channel configuration transactions and certs
  rm -rf channel-artifacts crypto-config
  # Remove the docker-compose yaml file that was customized to the kingbeefcattle
  rm -f docker-compose-e2e.yaml
}

cleanUp
