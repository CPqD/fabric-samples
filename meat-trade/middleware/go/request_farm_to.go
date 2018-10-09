/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// RequestFarmTO struct to make invoke request to Farm CC
type RequestFarmTO struct {
	Function    string         `json:"fcn"`
	Key         string         `json:"key"`
	Sale        BeefCattleSale `json:"sale"`
	OrgName     string         `json:"orgName"`
	MspID       string         `json:"mspID"`
	ChaincodeID string         `json:"chaincodeID"`
}
