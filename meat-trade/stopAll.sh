#!/bin/bash
#
# Copyright 2018 CPqD. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a convenience script to stop and remove all containers.
# Usage: ./stopAll.sh

BASE_DIR=$PWD

# stop application
cd ${BASE_DIR}/web-application && \
./stopApp.sh && \

# stop middleware
cd ${BASE_DIR}/middleware && \
docker-compose down && \

# stop network
cd ${BASE_DIR}/network && \
./cleanup.sh
