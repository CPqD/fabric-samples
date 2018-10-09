# Slaughterhouse chaincode
Slaughterhouse chaincode represents the negotiation between slaughterhouse and the supermarket.

This chaincode provides the following operations:
* *registerPrivateSale*: registers a private sale on the ledger, visible only for some organizations.
* *queryPrivateSale*: queries private sales registered using the earlier method.
* *registerSale*: registers a public sale on the ledger.
* *querySale*: queries public sales registered using the earlier method.
* *getHistoryByKey*: get all transactions related to the given key. 

## Install Slaughterhouse

Slaughterhouse chaincode will be installed in the following organizations: slaughterhouse, supermarket1, inspection1 and inspection2.

To install it in both peers 0 and 1 of the organizations mentioned above, execute:

```
./installChaincode.sh -l <language> -n slaughterhouse -v 1.0 -p 0 -o "slaughterhouse inspection1 inspection2 supermarket1"

./installChaincode.sh -l <language> -n slaughterhouse -v 1.0 -p 1 -o "slaughterhouse inspection1 inspection2 supermarket1"
```

Since the same chaincode is developed in Go and Node.js, you can use `<language>` argument to specify the one of your preference: `go`** or `node`.

If the arguments are not provided, default values will be used: `language:node`, `version:1.0`, `peer:0`, and will be installed in all organizations.

**Important**: make sure organization names are between double quotes.

**Important 2**: before installing `go` version of chaincode it is necessary to download some dependencies. Run the following commands:
```
go get -u github.com/hyperledger/fabric/protos/msp
go get -u github.com/hyperledger/fabric/core/chaincode
go get -u github.com/gogo/protobuf/proto
```

## Instantiate Slaughterhouse on the channel

Slaughterhouse chaincode  will be instantiated using a collection definition in order to allow private data on the channel. The endorsement policies for this chaincode are defined in [collections_config.json](collections_config.json) file. There is only one policy defined:

* `collectionSupemarket1PD`: defines that Slaughterhouse, Inspection1 and Inspection2 will be authorized to see private data recorded using this collection.

To instantiate the chaincode execute: 

```
./instantiateChaincode.sh -l <language> -n slaughterhouse -v 1.0 -p 0 -o slaughterhouse -P "OR ('SlaughterhouseMSP.member', 'Inspection1MSP.member','Inspection2MSP.member')" -c collections_config.json
```

## Testing chaincode from command line

### Executing private transactions

Access `cli` container:

```
docker exec -it cli bash
```

Change the organization to slaughterhouse:
```
cd scripts
. ./changeOrgAndPeer.sh slaughterhouse 0
```

Register a private sale using the following command:
```
peer chaincode invoke -C kingbeefcattlechannel -n slaughterhouse -c '{"Args":["C4AFC245-E650-4010-9E88-63CA65BE5A53","{\"slaughterhouse\":\"slaughterhouse.kingbeefcattle.com\",\"market\":\"supermarket1.kingbeefcattle.com\",\"carcases\":[{\"id\":\"746267CE-4EFB-496D-933F-0E71F7734E99\",\"idAnimal\":\"52A0BD52-C6C7-4764-B670-1261930C565A\",\"weight\":300,\"type\":\"right\"},{\"id\":\"D1EEF6CD-FC70-4CC0-8108-0BD508B09AC5\",\"idAnimal\":\"52A0BD52-C6C7-4764-B670-1261930C565A\",\"weight\":330,\"type\":\"left\"}]}"],"Function":"registerPrivateSale"}'
```

Query the registered private sale running the command:
```
peer chaincode query -n slaughterhouse -C kingbeefcattlechannel -c '{"Args":["C4AFC245-E650-4010-9E88-63CA65BE5A53"],"Function":"queryPrivateSale"}'
```

Switch target peer to inspection2:
```
. ./changeOrgAndPeer.sh inspection2 0
peer chaincode query -n slaughterhouse -C kingbeefcattlechannel -c '{"Args":["C4AFC245-E650-4010-9E88-63CA65BE5A53"],"Function":"queryPrivateSale"}'
```

Switch target peer to inspection1:
```
. ./changeOrgAndPeer.sh inspection1 0
peer chaincode query -n slaughterhouse -C kingbeefcattlechannel -c '{"Args":["C4AFC245-E650-4010-9E88-63CA65BE5A53"],"Function":"queryPrivateSale"}'
```

## Inspect chaincode in Node.js

Once dev-mode is enabled it is possible to inspect chaincodes developed in Node.js.

Before running the commands below, make sure to install the chaincode dependencies:

```
npm install
```

Navigate to Node.js chaincode folder of slaughterhouse and run the following command:

```
CORE_CHAINCODE_ID_NAME="slaughterhouse:1.0" node --inspect slaughterhouse.js --peer.address localhost:7052
```

If the command above succeed, it is possible to inspect chaincode execution using some remote debugging tool, such as `Inspect` of Google Chrome.

Open the Web browser and navigate to `chrome://inspect/#devices`. In `Remote Target` section should be possible to see name of inspected chaincode file, e.g, `slaughterhouse.js`.

Click `inspect`. Chrome console should be opened and from this moment, chaincode operations can be inspected in the console.