#!/bin/bash
#
# Copyright CPqD All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a convenience script to restart the network.
# Usage: ./restart.sh

./down.sh
sleep 2
./up.sh
