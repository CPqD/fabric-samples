# Meat Trade
This is a network and application sample that demonstrate the use of Go and Node.js chaincodes.
Using Fabric-SDK-Go, a REST API for integration of external applications is provided and also a Web Application that consumes such API.

## Scenario

The explored scenario is a marketplace of beef cattle. The consortium is composed by:
* 2 farms (*farm1* and *farm2*);
* 1 slaughterhouse (*slaughterhouse*);
* 2 inspection agencies (*inspection1* and *inspection2*);
* 1 supermarket (*supermarket1*).

In total there are 6 organizations plus an orderer service using SOLO consensus. Each organization, in turn, has two peers (*peer0* and *peer1*). All participants communicate through the same channel (*kingbeefcattlechannel*).

In a real simplified scenario, farmers are responsible for cattle raising and fattening, and subsequent sale to the slaughterhouse. The slaughterhouse split the cattle in different pieces (carcases) which are then sold to the supermarket.

This way this sample provides chaincodes to represent two transactions: from farmes to the slaughterhouse; and from the slaughterhouse to the supermarket.
Chaincodes also provide methods to use private data, hidding the whole transaction from unauthorized parties. 

## Requirements and setup

This sample works with version 1.2.0 of Fabric.

The following softwares must be installed before proceeding:
* [Docker](https://www.docker.com/get-started) version 18.06.0 or higher
* [Docker Compose](https://docs.docker.com/compose/install/) version 1.22.0 or higher
* [Go](https://golang.org/dl/) version 1.10.0 or higher
* [Node.js](https://nodejs.org/en/download/) version 8.11.x
* NPM version 5.6.0
* cURL version 7.54.0 (to download binaries)

In addition there must be downloaded some Fabric binaries to generate crypto material, channel settings and the genesis block. 

To download the required binaries, execute:
```
./bootstrap.sh
```

By the end this script will create a `bin` folder in this level.
## For the impatient

To start everything for the example, execute the following script and access the application at http://localhost:8080 using your favorite browser.
```
./runAll.sh
```

To stop everything, execute:
```
./stopAll.sh
```
The next sections will explain all steps in detail.
## Creating and launching the network

1. Navigate to `network` folder.
2. To generate the crypto material of all network participants, run ``./generate.sh``.
3. To launch the containers, run ``./up.sh``.
4. To create the channel, run ``./createChannel.sh``.
5. To join the participants, run ``./joinChannel.sh``.
6. To update anchor peers, run ``./updateAnchors.sh``.

**Optional commands**:
  * To stop the network, run ``./down.sh``
  * **To stop the network, remove all containers, volumes and images, run `./cleanup.sh`.
  * To restart the network, run ``./restart.sh`` (this script will run `./down.sh` and `./up.sh` sequentially).

** If you have other images you don't want to remove be careful running `cleanup.sh` script.

## Chaincode

To install **Farm** chaincode follow the instructions in [README.md](chaincode/farm/README.md).

To install **Slaughterhouse** chaincode follow the instructions in [README.md](chaincode/slaughterhouse/README.md).

## Middleware
In order to start the middleware follow the instructions described in [Middleware README.md](middleware/README.md).

## Run Web application

To start the Web application and perform operations on the ledger follow the instructions described in [Web App README.md](web-application/README.md).













