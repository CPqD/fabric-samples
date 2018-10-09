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
	msp "github.com/hyperledger/fabric/protos/msp"
	sc "github.com/hyperledger/fabric/protos/peer"
)

/** Collections **/

//Farm1FarmCollection id of Farm1
const Farm1FarmCollection string = "collectionFarm1PD"

//Farm2Collection id of Farm2
const Farm2Collection string = "collectionFarm2PD"

/** MSPID **/

//MspFarm1 MSP of Farm1
const MspFarm1 string = "Farm1MSP"

//MspFarm2 MSP da fazenda Farm2
const MspFarm2 string = "Farm2MSP"

//MspSlaughterhouse MSP of slaughterhouse
const MspSlaughterhouse string = "SlaughterhouseMSP"

//MspInspection1 MSP of inspection agent 1
const MspInspection1 string = "Inspection1MSP"

//NotAllowedOrg error message
const NotAllowedOrg string = "Organization not authorized to perform this action."

//WrongNumberOfArgs error message
const WrongNumberOfArgs string = "Wrong number of arguments, please provide %v."

//BeefCattleSale struct
type BeefCattleSale struct {
	Property       string `json:"property"`
	Slaughterhouse string `json:"slaughterhouse"`
	Cattle         []struct {
		ID                               string `json:"id"`
		Weight                           string `json:"weight"`
		Breed                            string `json:"breed"`
		DtLastBrucellosisVaccine         string `json:"dtLastBrucellosisVaccine"`
		DtLastFootAndMouthDeseaseVaccine string `json:"dtLastFootAndMouthDeseaseVaccine"`
		Age                              string `json:"age"`
		Classe                           string `json:"classe"`
		ProductionSystem                 string `json:"productionSystem"`
	} `json:"cattle"`
}

//Invoke method used to invoke all allowed methods
func (sale *BeefCattleSale) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately

	switch function {
	case "registerSale":
		return sale.registerSale(APIstub, args)
	case "querySale":
		return sale.querySale(APIstub, args)
	case "registerPrivateSale":
		return sale.registerPrivateSale(APIstub, args)
	case "queryPrivateSale":
		return sale.queryPrivateSale(APIstub, args)
	case "getHistoryForKey":
		return sale.getHistoryForKey(APIstub, args)
	}

	return shim.Error("method with name :" + function + " does not exist")
}

//Init method used to initialize blockchain
func (sale *BeefCattleSale) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (sale *BeefCattleSale) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/**registerSale method which make a sale
 */
func (sale *BeefCattleSale) registerSale(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 2 {
		return shim.Error(fmt.Sprintf(WrongNumberOfArgs, 2))
	}

	invoiceNumber := args[0]
	bytes := []byte(args[1])

	var beefCatle BeefCattleSale

	if err := json.Unmarshal(bytes, &beefCatle); err != nil {
		return shim.Error(err.Error())
	}

	beefCatlesAsByte, err := json.Marshal(beefCatle)

	if err != nil {
		return shim.Error(err.Error())
	}

	mspRequestID, err := sale.getMspid(APIstub)

	if mspRequestID == MspFarm1 || mspRequestID == MspFarm2 {
		if err := APIstub.PutState(invoiceNumber, beefCatlesAsByte); err != nil {
			return shim.Error(err.Error())
		}
	} else {
		return shim.Error(NotAllowedOrg)
	}

	return shim.Success(beefCatlesAsByte)
}

func (sale *BeefCattleSale) querySale(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
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

func (sale *BeefCattleSale) getHistoryForKey(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
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

func (sale *BeefCattleSale) registerPrivateSale(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 2 {
		return shim.Error(fmt.Sprintf(WrongNumberOfArgs, 2))
	}

	danfe := args[0]
	bytes := []byte(args[1])

	var beefCatle BeefCattleSale

	if err := json.Unmarshal(bytes, &beefCatle); err != nil {
		return shim.Error(err.Error())
	}

	beefCatlesAsByte, err := json.Marshal(beefCatle)

	if err != nil {
		return shim.Error(err.Error())
	}

	collectionName, err := sale.getCollectionName(APIstub)

	if err != nil {
		return shim.Error(err.Error())
	}

	if err := APIstub.PutPrivateData(collectionName, danfe, beefCatlesAsByte); err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(beefCatlesAsByte)
}

func (sale *BeefCattleSale) queryPrivateSale(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error("Número DANFE é obrigatório")
	}

	danfe := args[0]

	mspid, err := sale.getMspid(APIstub)

	if err != nil {
		return shim.Error(err.Error())
	}

	if mspid == MspSlaughterhouse || mspid == MspInspection1 {
		var beefCatlesAsByte []byte

		beefCatlesAsByte, err = APIstub.GetPrivateData(Farm1FarmCollection, danfe)

		if err != nil {
			return shim.Error(err.Error())
		} else if beefCatlesAsByte != nil {
			return shim.Success(beefCatlesAsByte)
		}

		beefCatlesAsByte, err = APIstub.GetPrivateData(Farm2Collection, danfe)

		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(beefCatlesAsByte)
	}

	collectionName, err := sale.getCollectionNameByMspid(mspid)

	if err != nil {
		return shim.Error(err.Error())
	}

	beefCatlesAsByte, err := APIstub.GetPrivateData(collectionName, danfe)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(beefCatlesAsByte)
}

/** Métodos auxiliares: **/

func (*BeefCattleSale) getMspid(APIstub shim.ChaincodeStubInterface) (string, error) {
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

func (sale *BeefCattleSale) getCollectionName(APIstub shim.ChaincodeStubInterface) (string, error) {
	mspid, err := sale.getMspid(APIstub)

	if err != nil {
		return "", errors.New(err.Error())
	}

	return sale.getCollectionNameByMspid(mspid)
}

func (*BeefCattleSale) getCollectionNameByMspid(Mspid string) (string, error) {

	switch Mspid {
	case MspFarm1:
		return Farm1FarmCollection, nil
	case MspFarm2:
		return Farm2Collection, nil
	default:
		return "", errors.New(NotAllowedOrg)
	}
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(BeefCattleSale))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
