/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 * 
 * SPDX-License-Identifier: Apache-2.0
 */

const shim = require('fabric-shim')

const SUPERMARKET_COLLECTION = 'collectionSupermarket1PD';

const ALLOWED_MSPS_REGISTER = ['SlaughterhouseMSP'];
const ALLOWED_MSPS_QUERY = ['SlaughterhouseMSP', 'Inspection1MSP', 'Inspection2MSP'];

var SHouseChaincode = class {

  // Initialize the chaincode
  async Init(stub) {    
    return shim.success("SHouseChaincode successfully instantiated");
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

    let creator = stub.getCreator();
    if (!creator) {
      throw new Error('Failed getting request creator');
    }

    if (ALLOWED_MSPS_REGISTER.indexOf(creator.mspid) == -1) {
      throw new Error('Organization not authorized to perform this action.');
    }

    let invoiceNumber = args[0];    
    let sale = args[1];

    let queryTransaction = await stub.getState(invoiceNumber);
    if (queryTransaction && queryTransaction.toString().length > 0) {
      throw new Error('Unable to perform action. Key already exists.');
    }


    try {
      var carcases = JSON.parse(sale);
    } catch (convertException) {
      console.log('JSON parse error');
      throw new Error('JSON parse error');
    }

    var transaction = carcases;

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

    let queryTransaction = await stub.getState(invoiceNumber);
    if (queryTransaction && queryTransaction.toString().length > 0) {
      throw new Error('Unable to perform action. Key already exists.');
    }


    try {
      var carcases = JSON.parse(sale);
    } catch (convertException) {
      console.log('JSON parse error');
      throw new Error('JSON parse error');
    }

    var transaction = carcases;
    
    let creator = stub.getCreator();
    if (!creator) {
      throw new Error('Failed getting request creator');
    }

    if (ALLOWED_MSPS_REGISTER.indexOf(creator.mspid) == -1) {
      throw new Error('Organization not authorized to perform this action.');
    }
    await stub.putPrivateData(SUPERMARKET_COLLECTION, invoiceNumber, Buffer.from(JSON.stringify(transaction)));
    
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
    
    let creator = stub.getCreator();
    if (!creator) {
      throw new Error('Failed getting request creator');
    }

    if (ALLOWED_MSPS_QUERY.indexOf(creator.mspid) == -1) {
      throw new Error('Organization not authorized to perform this action.');
    }

    let transactionAsBytes = await stub.getPrivateData(SUPERMARKET_COLLECTION, invoiceNumber);
    if (!transactionAsBytes || transactionAsBytes.toString().length <= 0) {
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

let chaincode = new SHouseChaincode();

shim.start(chaincode);