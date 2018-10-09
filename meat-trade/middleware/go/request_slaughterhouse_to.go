/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// RequestSlaughterhouseTO struct to make invoke request to Slaughterhouse CC
type RequestSlaughterhouseTO struct {
	Function       string         `json:"fcn"`
	Key            string         `json:"key"`
	CutsBeefCattle CutsBeefCattle `json:"cutsBeefCattle"`
	OrgName        string         `json:"orgName"`
	MspID          string         `json:"mspID"`
}
