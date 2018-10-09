/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// SlaughterhouseAnchor relative path to anchor settings file
const SlaughterhouseAnchor string = "/channel-artifacts/SlaughterhouseMSPanchors.tx"

// Farm1Anchor relative path to anchor settings file
const Farm1Anchor string = "/channel-artifacts/Farm1MSPanchors.tx"

// Farm2Anchor relative path to anchor settings file
const Farm2Anchor string = "/channel-artifacts/Farm2MSPanchors.tx"

// Inspection1Anchor relative path to anchor settings file
const Inspection1Anchor string = "/channel-artifacts/Inspection1MSPanchors.tx"

// Inspection2Anchor relative path to anchor settings file
const Inspection2Anchor string = "/channel-artifacts/Inspection2MSPanchors.tx"

// Supermarket1Anchor relative path to anchor settings file
const Supermarket1Anchor string = "/channel-artifacts/Supermarket1MSPanchors.tx"

//NotAllowedOrg error message
const NotAllowedOrg string = "Organization not authorized to perform this action."

// Index function relative path to anchor settings file
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to middleware sample of hyperledger integration!\n")
}

// InvokeFarm Post method to invoke Farm CC
func InvokeFarm(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var requestTO RequestFarmTO

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.Unmarshal(body, &requestTO); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	sale, err := json.Marshal(requestTO.Sale)

	if err != nil {
		fmt.Println("Erro ao descerializar a venda. Error: " + err.Error())
	}

	fabricSetup, err := createFabricSetup(requestTO.OrgName, "farm")

	if err != nil {
		panic(err)
	}

	defer fabricSetup.CloseSDK()

	invoketo := InvokeTO{
		Function: requestTO.Function,
		Key:      requestTO.Key,
		Args:     [][]byte{[]byte(requestTO.Key), sale},
	}

	fabricSetup.Invoke(invoketo)
}

// QueryFarm Post method to make a query to Farm CC
func QueryFarm(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var requestTO RequestFarmTO

	if err := json.Unmarshal(body, &requestTO); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	fabricSetup, err := createFabricSetup(requestTO.OrgName, "farm")

	if err != nil {
		panic(err)
	}

	defer fabricSetup.CloseSDK()

	invoketo := InvokeTO{
		Function: requestTO.Function,
		Key:      requestTO.Key,
		Args:     [][]byte{[]byte(requestTO.Key)},
	}

	sale, err := fabricSetup.QueryFarm(invoketo)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // unprocessable entity
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(sale); err != nil {
		panic(err)
	}
}

// InvokeSlaughterhouse Post method to invoke Slaughterhouse CC
func InvokeSlaughterhouse(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var requestTO RequestSlaughterhouseTO

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.Unmarshal(body, &requestTO); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	sale, err := json.Marshal(requestTO.CutsBeefCattle)

	if err != nil {
		fmt.Println("Erro ao descerializar a venda. Error: " + err.Error())
	}

	fabricSetup, err := createFabricSetup(requestTO.OrgName, "slaughterhouse")

	if err != nil {
		panic(err)
	}

	defer fabricSetup.CloseSDK()

	invoketo := InvokeTO{
		Function: requestTO.Function,
		Key:      requestTO.Key,
		Args:     [][]byte{[]byte(requestTO.Key), sale},
	}

	fabricSetup.Invoke(invoketo)
}

// QuerySlaughterhouse Post method to make a query to Slaughterhouse CC
func QuerySlaughterhouse(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var requestTO RequestSlaughterhouseTO

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.Unmarshal(body, &requestTO); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	fabricSetup, err := createFabricSetup(requestTO.OrgName, "slaughterhouse")

	if err != nil {
		panic(err)
	}

	defer fabricSetup.CloseSDK()

	invoketo := InvokeTO{
		Function: requestTO.Function,
		Key:      requestTO.Key,
		Args:     [][]byte{[]byte(requestTO.Key)},
	}

	sale, err := fabricSetup.QuerySlaughterhouse(invoketo)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // unprocessable entity
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(sale); err != nil {
		panic(err)
	}
}

// createFabricSetup Set up Fabric SDK for org thats made the request
func createFabricSetup(orgName string, chaincodeID string) (FabricSetup, error) {
	var channelConfig string

	switch strings.ToUpper(orgName) {
	case "FARM1":
		channelConfig = Farm1Anchor
	case "FARM2":
		channelConfig = Farm2Anchor
	case "SLAUGHTERHOUSE":
		channelConfig = SlaughterhouseAnchor
	case "INSPECTION1":
		channelConfig = Inspection1Anchor
	case "INSPECTION2":
		channelConfig = Inspection2Anchor
	case "SUPERMARKET1":
		channelConfig = Supermarket1Anchor
	default:
		return FabricSetup{}, fmt.Errorf(NotAllowedOrg)

	}

	fmt.Println(orgName)

	fabricSetup := FabricSetup{
		// Network parameters
		OrdererID: "orderer.kingbeefcattle.com",

		// Channel parameters
		ChannelID:     "kingbeefcattlechannel",
		ChannelConfig: channelConfig,

		// Chaincode parameters
		ChainCodeID:     chaincodeID,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		OrgAdmin:        "Admin",
		OrgName:         orgName,
		ConfigFile:      orgName + ".yaml",

		// User parameters
		UserName: "User1",
	}

	if err := fabricSetup.Initialize(); err != nil {
		return fabricSetup, err
	}

	return fabricSetup, nil
}
