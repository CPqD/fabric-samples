# Farm chaincode
Farm chaincode represents the negotiation between farmers and the slaughterhouse.

This chaincode provides the following operations:
* *registerPrivateSale*: registers a private sale on the ledger, visible only for some organizations.
* *queryPrivateSale*: queries private sales registered using the earlier method.
* *registerSale*: registers a public sale on the ledger.
* *querySale*: queries public sales registered using the earlier method.
* *getHistoryByKey*: get all transactions related to the given key. 

## Install Farm

Farm chaincode will be installed in the following organizations: farm1, farm2, slaughterhouse and inspection1.

To install it in both peers 0 and 1 of the organizations mentioned above, execute:

```
./installChaincode.sh -l <language> -n farm -v 1.0 -p 0 -o "farm1 farm2 slaughterhouse inspection1"

./installChaincode.sh -l <language> -n farm -v 1.0 -p 1 -o "farm1 farm2 slaughterhouse inspection1"
```

Since the same chaincode is developed in both Go and Node.js, you can use `<language>` argument to specify the one of your preference: `go`** or `node`.

If the arguments are not provided, default values will be used: `language:node`, `version:1.0`, `peer:0`, and will be installed in all organizations.

**Important**: make sure organization names are between double quotes.

**Important2**: before installing `go` version of chaincode it is necessary to download some dependencies. Run the following commands:
```
go get -u github.com/hyperledger/fabric/protos/msp
go get -u github.com/hyperledger/fabric/core/chaincode
go get -u github.com/gogo/protobuf/proto
```

## Instantiate Farm on the channel

Farm chaincode will be instantiated using a collection definition in order to allow private data on the channel. The endorsement policies for this chaincode are defined in [collections_config.json](collections_config.json) file. There are two policies defined:

* `collectionFarm1PD`: defines that Farm1, Slaughterhouse and Inspection1 will be authorized to see private data recorded using this collection.
* `collectionFarm2PD`: defines that Farm2, Slaughterhouse and Inspection1 will be authorized to see private data recorded using this collection.

To instantiate the chaincode execute: 

```
./instantiateChaincode.sh -l <language> -n farm -v 1.0 -p 0 -o farm1 -P "OR ('SlaughterhouseMSP.member', 'Farm1MSP.member','Farm2MSP.member')" -c collections_config.json
```

`<language>` must be the same used in the installation step.

## Testing chaincode from command line

### Executing private transactions

Access `cli` container:
```
docker exec -it cli bash
```

Change the target organization to farm1:

```
cd scripts
. ./changeOrgAndPeer.sh farm1
```

Register a private sale using the following command:

```
peer chaincode invoke -C kingbeefcattlechannel -n farm -c '{"Args": ["36FEEAB3-0E17-4F6B-8E7C-5E052C552E90","{\"property\":\"farm1.kingbeefcattle.com\",\"slaughterhouse\":\"slaughterhouse.kingbeefcattle.com\",\"cattle\":[{\"id\":\"52A0BD52-C6C7-4764-B670-1261930C565A\",\"weight\":720,\"breed\":\"Nelore\",\"dtLastBrucellosisVaccine\":\"2018-07-05\",\"dtLastFootAndMouthDeseaseVaccine\":\"2018-07-05\",\"age\":30,\"classe\":\"ox\",\"productionSystem\":\"confinement\"},{\"id\":\"B771DB6C-6AC7-4150-A7D8-C3F08CEEDF0E\",\"weight\":920,\"breed\":\"Angus\",\"dtLastBrucellosisVaccine\":\"2018-04-05\",\"dtLastFootAndMouthDeseaseVaccine\":\"2018-04-05\",\"age\":12,\"classe\":\"calf\",\"productionSystem\":\"semi-confinement\"}]}"],"Function": "registerPrivateSale"}'
```

Query the registered private sale running the command:
```
peer chaincode query -n farm -C kingbeefcattlechannel -c '{"Args":["36FEEAB3-0E17-4F6B-8E7C-5E052C552E90"],"Function":"queryPrivateSale"}'
```

To check if private data worked, switch target container of `cli` to `farm2`: 

```
. ./changeOrgAndPeer.sh farm2 0
```

Execute the same query of private sale as did before:
```
peer chaincode query -n farm -C kingbeefcattlechannel -c '{"Args":["36FEEAB3-0E17-4F6B-8E7C-5E052C552E90"],"Function":"queryPrivateSale"}'
```

This time the query should return a 500 status code with the message "Transaction not found". This happens because Farm2 is not included in the private data collection `collectionSantannafarmPD` that was used to register the private sale.

Switch target container of `cli` to `slaughterhouse`: 

```
. ./changeOrgAndPeer.sh slaughterhouse 0
```

Run the query again:
```
peer chaincode query -n farm -C kingbeefcattlechannel -c '{"Args":["36FEEAB3-0E17-4F6B-8E7C-5E052C552E90"],"Function":"queryPrivateSale"}'
```

Since Slaughterhouse is included in `collectionSantannafarmPD` private data collection it would be able to see the transaction performed by farm1.

The same is true for the organization Inspection1:

```
. ./changeOrgAndPeer.sh inspection1
peer chaincode query -n farm -C kingbeefcattlechannel -c '{"Args":["36FEEAB3-0E17-4F6B-8E7C-5E052C552E90"],"Function":"queryPrivateSale"}'
```

### Executing public transactions

Access `cli` container:
```
docker exec -it cli bash
```

Only Farm1 and Farm2 are allowed to register public sales using farm chaincode. To register a public transaction using Farm1, execute:

```
. ./changeOrgAndPeer.sh farm1

peer chaincode invoke -C kingbeefcattlechannel -n farm -c '{"Args": ["BFE0897E-3F42-48F9-8387-38F296BED017","{\"property\":\"farm1.kingbeefcattle.com\",\"slaughterhouse\":\"slaughterhouse.kingbeefcattle.com\",\"cattle\":[{\"id\":\"709F5FC1-5C95-40A6-BE10-CC6886D43BF6\",\"weight\":720,\"breed\":\"Nelore\",\"dtLastBrucellosisVaccine\":\"2018-07-05\",\"dtLastFootAndMouthDeseaseVaccine\":\"2018-07-05\",\"age\":30,\"classe\":\"ox\",\"productionSystem\":\"confinement\"},{\"id\":\"BE0FB9FD-3EFA-4E4E-9C82-A91816E109EB\",\"weight\":920,\"breed\":\"Angus\",\"dtLastBrucellosisVaccine\":\"2018-04-05\",\"dtLastFootAndMouthDeseaseVaccine\":\"2018-04-05\",\"age\":12,\"classe\":\"calf\",\"productionSystem\":\"semi-confinement\"}]}"],"Function": "registerSale"}'
```

To query the public transaction, execute:

```
peer chaincode query -n farm -C kingbeefcattlechannel -c '{"Args":["BFE0897E-3F42-48F9-8387-38F296BED017"],"Function":"querySale"}'
```

To get the history of the key, execute: 

```
peer chaincode invoke -n farm -C kingbeefcattlechannel -c '{"Args":["BFE0897E-3F42-48F9-8387-38F296BED017"],"Function":"getHistoryForKey"}'
```

## Inspect chaincode in Node.js

Once dev-mode is enabled it is possible to inspect chaincodes developed in Node.js.

Before running the commands below, make sure to install the chaincode dependencies:

```
npm install
```

Navigate to Node.js chaincode folder of farm and run the following command:

```
CORE_CHAINCODE_ID_NAME="farm:1.0" node --inspect farm.js --peer.address localhost:7052
```

If the command above succeed it is possible to inspect chaincode execution using some remote debugging tool, such as `Inspect` of Google Chrome.

Open the Web browser and navigate to `chrome://inspect/#devices`. In `Remote Target` section should be possible to see name of inspected chaincode file, e.g, `farm.js`.

Click `inspect`. Chrome console should be opened and from this moment, chaincode operations can be inspected in the console.
