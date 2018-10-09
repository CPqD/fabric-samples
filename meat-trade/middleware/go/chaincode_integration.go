/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// Invoke method that makes integration with CC
func (setup *FabricSetup) Invoke(value InvokeTO) (string, error) {

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{
		ChaincodeID: setup.ChainCodeID,
		Fcn:         value.Function,
		Args:        value.Args,
	}) //, channel.WithRetry(retry.DefaultChannelOpts))

	if err != nil {
		fmt.Println(err.Error())
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	return string(response.TransactionID), nil
}

// QueryFarm method that makes integration with CC
func (setup *FabricSetup) QueryFarm(value InvokeTO) (BeefCattleSale, error) {
	response, err := setup.client.Query(channel.Request{
		ChaincodeID: setup.ChainCodeID,
		Fcn:         value.Function,
		Args:        value.Args,
	})

	if err != nil {
		return BeefCattleSale{}, fmt.Errorf("failed to query: %v", err)
	}

	var sale BeefCattleSale

	if err := json.Unmarshal(response.Payload, &sale); err != nil {
		return BeefCattleSale{}, fmt.Errorf("failed to query: %v", err)
	}

	return sale, nil
}

// QuerySlaughterhouse method that makes integration with CC
func (setup *FabricSetup) QuerySlaughterhouse(value InvokeTO) (CutsBeefCattle, error) {
	response, err := setup.client.Query(channel.Request{
		ChaincodeID: setup.ChainCodeID,
		Fcn:         value.Function,
		Args:        value.Args,
	})

	if err != nil {
		return CutsBeefCattle{}, fmt.Errorf("failed to query: %v", err)
	}

	var cuts CutsBeefCattle

	if err := json.Unmarshal(response.Payload, &cuts); err != nil {
		return CutsBeefCattle{}, fmt.Errorf("failed to query: %v", err)
	}

	return cuts, nil
}
