/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 * 
 * SPDX-License-Identifier: Apache-2.0
 */

const shim = require('fabric-shim')

const FARM2_COLLECTION = 'collectionFarm2PD';
const FARM1_COLLECTION = 'collectionFarm1PD';

const ALLOWED_MSPS_REGISTER = ['Farm1MSP', 'Farm2MSP'];
const ALLOWED_MSPS_QUERY = ['SlaughterhouseMSP', 'Inspection1MSP'];

var FarmChaincode = class {
  
  async Init() {    
    return shim.success("FarmChaincode successfully instantiated");
  }

  async Invoke(stub) {
    let ret = stub.getFunctionAndParameters();
    console.info(ret);
    let method = this[ret.fcn];
    if (!method) {
      console.log('method with name :' + ret.fcn + ' does not exist');
      shim.error(new Error('method with name :' + ret.fcn + ' does not exist'));
    }
    try {
      let payload = await method(stub, ret.params);
      return shim.success(payload);
    } catch (err) {
      console.log(err);
      return shim.error(err);
    }
  }

  /**
   * Registers public sale on the ledger
   * @param {ChaincodeStub} stub 
   * @param {Array} args 
   */
  async registerSale(stub, args) {

    if (args.length < 2) {
      throw new Error('Wrong number of arguments, please provide 2.');
    }

    let invoiceNumber = args[0];  
    let sale = args[1];

    let creator = stub.getCreator();

    //Only Farm1 and Farm2 are allowed to perform this operation.
    if (ALLOWED_MSPS_REGISTER.indexOf(creator.mspid) == -1) {
      throw new Error('Organization not authorized to perform this action.');
    }

    try {
      var cattle = JSON.parse(sale);
    } catch (convertException) {
      console.log('JSON parse error');
      throw new Error('JSON parse error');
    }

    var transaction = cattle;

    await stub.putState(invoiceNumber, Buffer.from(JSON.stringify(transaction)));

  }

  /**
   * Queries public sale on ledger
   * @param {ChaincodeStub} stub 
   * @param {Array} args 
   */
  async querySale(stub, args) {

    let invoiceNumber = args[0];
    if (!invoiceNumber) {
      throw new Error('Wrong number of arguments, please provide 1.');
    }
    let transactionAsBytes = await stub.getState(invoiceNumber);
    if (!transactionAsBytes || transactionAsBytes.toString().length <= 0) {
      console.log('Transaction not found');
      throw new Error('Transaction not found');
    }
    return transactionAsBytes;
  }

  /**
   * Registers private sale on ledger.
   * @param {ChaincodeStub} stub
   * @param {Array} args 
   */
  async registerPrivateSale(stub, args) {
    if (args.length < 2) {
      throw new Error('Wrong number of arguments, please provide 2.');
    }

    let invoiceNumber = args[0];
    let sale = args[1];
    try {
      var cattle = JSON.parse(sale);
    } catch (convertException) {
      console.log('JSON parse error');
      throw new Error('JSON parse error');
    }

    var transaction = cattle;    
    
    let creator = stub.getCreator();
    if (!creator) {
      throw new Error('Failed getting request creator');
    }
    
    //Only Farm1 and Farm2 are allowed to perform this operation.
    if (ALLOWED_MSPS_REGISTER.indexOf(creator.mspid) == -1) {
      throw new Error('Organization not authorized to perform this action.');
    }

    //Collection name is defined in collections_config.json file.
    //Here we are defining the name of the collection with which private data will be save according to requester MSP Id.    
    var collectionName = FARM2_COLLECTION;    

    //Farm1MSP
    if (creator.mspid == ALLOWED_MSPS_REGISTER[0]) {
      collectionName = FARM1_COLLECTION;
    }    
    await stub.putPrivateData(collectionName, invoiceNumber, Buffer.from(JSON.stringify(transaction)));
  }

  /**
   * Queries private sale on ledger
   * @param {ChaincodeStub} stub 
   * @param {Array} args 
   */
  async queryPrivateSale(stub, args) {
    let invoiceNumber = args[0];
    if (!invoiceNumber) {
      throw new Error('Wrong number of arguments, please provide 1.');
    }

    //Collection name is defined in collections_config.json file.
    //Here we are defining the name of the collection with which private data will be save according to requester MSP id.
    var collectionName = FARM2_COLLECTION;
    let creator = stub.getCreator();
    if (!creator) {
      throw new Error('Failed getting request creator');
    }

    if (creator.mspid == ALLOWED_MSPS_REGISTER[0]) {
      collectionName = FARM1_COLLECTION;
    }

    var transactionAsBytes = await stub.getPrivateData(collectionName, invoiceNumber);
    if (!transactionAsBytes || transactionAsBytes.toString().length <= 0) {       
      if (ALLOWED_MSPS_QUERY.indexOf(creator.mspid) != -1) {        
        if (collectionName == FARM1_COLLECTION) {
          collectionName = FARM2_COLLECTION;
        } else {
          collectionName = FARM1_COLLECTION;
        }
        transactionAsBytes = await stub.getPrivateData(collectionName, invoiceNumber);
        if (!transactionAsBytes || transactionAsBytes.toString().length <= 0) {
          throw new Error('Transaction not found');
        }
        return transactionAsBytes;
      }
      console.log('Transaction not found');
      throw new Error('Transaction not found');
    }
    return transactionAsBytes;
  }

  /**
   * Get history for a given on the ledger
   * @param {ChaincodeStub} stub 
   * @param {Array} args 
   */
  async getHistoryForKey(stub, args) {

    let key = args[0];
    if (!key) {
      throw new Error('Wrong number of arguments, please provide 1');
    }

    let iterator = await stub.getHistoryForKey(key);

    let allResults = [];

    while (true) {
      let res = await iterator.next();

      if (res.value && res.value.value.toString()) {
        let jsonRes = {};
        jsonRes.key = res.value.key;
        jsonRes.isDeleted = res.value.is_delete;
        jsonRes.transactionId = res.value.tx_id;
        try {
          jsonRes.record = JSON.parse(res.value.value.toString('utf8'));
        } catch (err) {
          console.log(err);
          jsonRes.record = res.value.value.toString('utf8');
        }
        allResults.push(jsonRes);
      }

      if (res.done) {
        console.log('end of data');
        await iterator.close();
        console.info(allResults);
        return Buffer.from(JSON.stringify(allResults));
      }
    }
  }
};

let chaincode = new FarmChaincode();

shim.start(chaincode);