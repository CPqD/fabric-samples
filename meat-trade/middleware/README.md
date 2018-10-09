# Middleware
The middleware is integrated with hyperledger network (using fabric-sdk-go) and provides a REST interface to allow integration of external applications.

## Starting the middleware
The middleware is provided as a Docker container. It runs on port `7500` and it belongs to `network_kingbeefcattle` network.

In order to start the middleware navigate to `middleware` folder and run the command:
```
docker-compose up
```
**Note 1:** for the correction execution of this command, make sure that the network_kingbeefcattle exists.

**Note 2:** the execution of the aforementioned command may take a while specially in the first time, since all go dependencies will be downloaded.

## Services
The services described in the following sections are provided by the middleware.

### /invokefarm
* Full path: `http://localhost:7500/invokefarm`
* HTTP method: `POST`
* Description: used to register a new sale from Farm to Slaughterhouse on the ledger. Both public and private sales can be registered using this method.
* Payload sample:
```
{
	"fcn":"registerSale",
	"key": "69525EE7-5E49-4319-90DB-56660237C550",
	"orgName": "farm1",
	"mspID": "Farm1MSP",
	"sale": {
		"property": "Farm1",
		"slaughterhouse": "Slaughterhouse",
		"cattle": [{
	        "id": "Bovine033",
	        "weight": "300",
	        "breed": "Nelore",
	        "dtLastBrucellosisVaccine": "10/11/2017",
	        "dtLastFootAndMouthDeseaseVaccine": "10/11/2017",
	        "age": "23",
	        "classe": "Calf",
	        "productionSystem":"Confinement"
	    }]
	}
}
```

### /queryfarm
* Full path: `http://localhost:7500/queryfarm`
* HTTP method: `POST`
* Description: used to query a sale from Farm to Slaughterhouse registered on the ledger. Both public and private sales can be queried using this method.
* Payload sample:
```
{
	"fcn":"querySale",
	"key": "69525EE7-5E49-4319-90DB-56660237C550",
	"orgName": "farm1"
}
```

### /invokeslaughterhouse
* Full path: `http://localhost:7500/invokeslaughterhouse`
* HTTP method: `POST`
* Description: used to register a new sale from Slaughterhouse to Market on the ledger. Both public and private sales can be registered using this method.
* Payload sample:
```
{
   "fcn":"registerSale",
   "key":"89FFA2CF-5296-4855-BC20-F1F38C7206B2",
   "orgName":"slaughterhouse",
   "mspID":"SlaughterhouseMSP",
   "cutsBeefCattle": {
	    "slaughterhouse": "slaughterhouse.kingbeefcattle.com",
	    "market": "supermarket1.kingbeefcattle.com",
	    "carcases": [
	        {
	            "idAnimal": "69525EE7-5E49-4319-90DB-56660237C550",
	            "id": "746267CE-4EFB-496D-933F-0E71F7734E99",
	            "weight": "330",
	            "type": "right"
	        },
	        {
	            "idAnimal": "69525EE7-5E49-4319-90DB-56660237C550",
	            "id": "D1EEF6CD-FC70-4CC0-8108-0BD508B09AC5",
	            "weight": "330",
	            "type": "left"
	        }
	    ]
	}
}
```

### /queryslaughterhouse
* Full path: `http://localhost:7500/queryslaughterhouse`
* HTTP method: `POST`
* Description: used to query a sale from Slaughterhouse to Market registered on the ledger. Both public and private sales can be queried using this method.
* Payload sample:
```
{
	"fcn":"querySale",
	"key": "69525EE7-5E49-4319-90DB-56660237C550",
	"orgName": "farm1"
}
```

## Stopping the middleware
To stop the middleware run the following command:
```
docker-compose down
```
