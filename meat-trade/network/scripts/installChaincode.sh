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

# Linguagem do chaincode
CC_LANG="$1"
# Nome e versão do chaincode a ser instanciado, versão default é 1.0.
CC_NAME="$2"
CC_VERSION=${3:-1.0}
PEER="$4"
shift 4
ORG_LIST=( "$@" )

LANG_OPT="-l"

# Chaincode path
CC_SRC_PATH=/opt/gopath/src/github.com/chaincode/${CC_NAME}/node/

if [ "$CC_LANG" = "go" ]; then
    CC_SRC_PATH=github.com/chaincode/${CC_NAME}/go/
    LANG_OPT=""
    CC_LANG=""
fi

installChaincode() {

  setGlobals ${PEER} ${ORG}

  set -x
  peer chaincode install -n ${CC_NAME} -v ${CC_VERSION} ${LANG_OPT} ${CC_LANG} -p ${CC_SRC_PATH} >&log.txt
  res=$?
  set +x
  cat log.txt
  verifyResult $res "Chaincode installation on peer${PEER}.${ORG} has failed"
  echo "===================== Chaincode is installed on peer${PEER}.${ORG} ===================== "
  echo
}

# Installs $CC_NAME in peer$PEER of all organizations
for ORG in "${ORG_LIST[@]}"; do
    installChaincode ${PEER} ${ORG}
done

