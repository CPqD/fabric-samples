/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
	sc "github.com/hyperledger/fabric/protos/peer"
)

/** Collections **/

//Supermarket1Collection private data
const Supermarket1Collection string = "collectionSupermarket1PD"

//MspSupermarket1 MSP of supermarket
const MspSupermarket1 string = "Supermarket1MSP"

//MspInspection2 MSP of inspection2 org
const MspInspection2 string = "Inspection2MSP"

//MspInspection1 MSP of inspection1 org
const MspInspection1 string = "Inspection1MSP"

//MspSlaughterhouse MSP of slaughterhouse org
const MspSlaughterhouse string = "SlaughterhouseMSP"

//NotAllowedOrg error message
const NotAllowedOrg string = "Organization not authorized to perform this action."

//WrongNumberOfArgs error message
const WrongNumberOfArgs string = "Wrong number of arguments, please provide %v."

//ErrorMessagePermissionDenied Error Message
const ErrorMessagePermissionDenied string = "Organization not authorized to perform this action."

//CutsOfBeef structure
type CutsOfBeef struct {
	Slaughterhouse string `json:"slaughterhouse"`
	Market         string `json:"market"`
	Carcases       []struct {
		IDAnimal string `json:"idAnimal"`
		ID       string `json:"id"`
		Weight   string `json:"weight"`
		Type     string `json:"type"`
	} `json:"carcases"`
}

//Init method used to initialize blockchain
func (p *CutsOfBeef) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

//Invoke methods
func (p *CutsOfBeef) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "registerSale" {
		return p.registerSale(APIstub, args)
	} else if function == "querySale" {
		return p.querySale(APIstub, args)
	} else if function == "getHistoryForKey" {
		return p.getHistoryForKey(APIstub, args)
	} else if function == "registerPrivateSale" {
		return p.registerPrivateSale(APIstub, args)
	} else if function == "queryPrivateSale" {
		return p.queryPrivateSale(APIstub, args)
	}

	return shim.Error("method with name :" + function + " does not exist")
}

func (p *CutsOfBeef) registerSale(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 2 {
		return shim.Error(fmt.Sprintf(WrongNumberOfArgs, 2))
	}

	danfe := args[0]
	bytes := []byte(args[1])

	var beefCatle CutsOfBeef

	if err := json.Unmarshal(bytes, &beefCatle); err != nil {
		return shim.Error(err.Error())
	}

	beefCatlesAsByte, err := json.Marshal(beefCatle)

	if err != nil {
		return shim.Error(err.Error())
	}
	mspRequestID, err := p.getMspid(APIstub)

	if mspRequestID == MspSlaughterhouse {
		if err := APIstub.PutState(danfe, beefCatlesAsByte); err != nil {
			return shim.Error(err.Error())
		}
	} else {
		return shim.Error(NotAllowedOrg)
	}

	return shim.Success(nil)
}

func (p *CutsOfBeef) querySale(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error(fmt.Sprintf(WrongNumberOfArgs, 1))
	}

	danfe := args[0]

	beefCatlesAsByte, err := APIstub.GetState(danfe)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(beefCatlesAsByte)
}

func (p *CutsOfBeef) getHistoryForKey(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf(WrongNumberOfArgs, 1))
	}

	danfe := args[0]

	historyIterator, err := APIstub.GetHistoryForKey(danfe)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer historyIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for historyIterator.HasNext() {
		keyModification, err := historyIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		json.Marshal(keyModification)
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(keyModification.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(keyModification.Value))

		buffer.WriteString(", \"IsDelete\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(strconv.FormatBool(keyModification.IsDelete))
		buffer.WriteString("}")

		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (p *CutsOfBeef) registerPrivateSale(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 2 {
		return shim.Error(fmt.Sprintf(WrongNumberOfArgs, 2))
	}

	danfe := args[0]
	bytes := []byte(args[1])

	mspid, err := p.getMspid(APIstub)

	if err != nil {
		return shim.Error(err.Error())
	}

	if mspid != MspSlaughterhouse {
		return shim.Error(ErrorMessagePermissionDenied)
	}

	var beefCatle CutsOfBeef

	if err := json.Unmarshal(bytes, &beefCatle); err != nil {
		return shim.Error(err.Error())
	}

	beefCatlesAsByte, err := json.Marshal(beefCatle)

	if err != nil {
		return shim.Error(err.Error())
	}

	if err := APIstub.PutPrivateData(Supermarket1Collection, danfe, beefCatlesAsByte); err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (p *CutsOfBeef) queryPrivateSale(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error(fmt.Sprintf(WrongNumberOfArgs, 1))
	}

	danfe := args[0]

	mspid, err := p.getMspid(APIstub)

	if err != nil {
		return shim.Error(err.Error())
	}

	collection, err := p.getCollectionNameByMspid(mspid)

	if err != nil {
		return shim.Error(err.Error())
	}

	beefCatlesAsByte, err := APIstub.GetPrivateData(collection, danfe)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(beefCatlesAsByte)
}

func (p *CutsOfBeef) getMspid(APIstub shim.ChaincodeStubInterface) (string, error) {
	serializedCreator, err := APIstub.GetCreator()

	if err != nil {
		return "", errors.New(err.Error())
	}

	creator := &msp.SerializedIdentity{}

	if err := proto.Unmarshal(serializedCreator, creator); err != nil {
		return "", errors.New(err.Error())
	}

	return creator.Mspid, nil
}

func (p *CutsOfBeef) getCollectionName(APIstub shim.ChaincodeStubInterface) (string, error) {
	mspid, err := p.getMspid(APIstub)

	if err != nil {
		return "", errors.New(err.Error())
	}

	return p.getCollectionNameByMspid(mspid)
}

func (p *CutsOfBeef) getCollectionNameByMspid(Mspid string) (string, error) {

	if Mspid == MspSlaughterhouse || Mspid == MspSupermarket1 || Mspid == MspInspection1 || Mspid == MspInspection2 {
		return Supermarket1Collection, nil
	} else {
		return "", errors.New(ErrorMessagePermissionDenied)
	}
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(CutsOfBeef))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
