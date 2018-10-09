#!/bin/bash
#
# Copyright 2018 CPqD. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a convenience script to start webapp.
# Usage: ./startApp.sh

docker build -t webapp . && \
docker run --name webapp1 -p 8080:8080 -d webapp
