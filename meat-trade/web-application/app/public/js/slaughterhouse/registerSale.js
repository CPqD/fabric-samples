/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

var formAddCarcase, carcasesTable, formRegisterCarcase;
var allCarcases = {};
allCarcases.cutsBeefCattle = {};
allCarcases.cutsBeefCattle.carcases = [];

(function () {
    formAddCarcase = document.querySelector("#formAddCarcase");
    formRegisterCarcase = document.querySelector("#formRegisterCarcase");
    carcasesTable = document.querySelector("#carcasesTable");
}());


function addCarcase() {
    console.log("validate addCarcase called");

    let carcase = {};
    carcase.idAnimal = formAddCarcase.querySelector("#cattleId").value;
    carcase.id = formAddCarcase.querySelector("#carcaseId").value;
    carcase.weight = formAddCarcase.querySelector("#weight").value;
    
    let types = document.getElementsByName('carcaseType');    
    if (types) {
        for (var i = 0; i < types.length; i++) {
            var item = types[i];
            if (item && item.checked) {
                carcase.type = item.value;
            }
        }
    }
   
     allCarcases.cutsBeefCattle.carcases.push(carcase);
    appendCattleToTable(carcase);
    formAddCarcase.reset()
    return false;
}


function appendCattleToTable(carcase) {
    let tBody = carcasesTable.querySelector("tbody");
    if (tBody) {
        let register = "<tr>"
            + "<th>" + carcase.id + "</th>"
            + "<th>" + carcase.idAnimal + "</th>"
            + "<td>" + carcase.weight + "</td>"
            + "<td>" + carcase.type + "</td>"            
            + "</tr>";

        tBody.insertAdjacentHTML('beforeend', register)
    }
}

function resetCarcaseTable() {
    let tBody = carcasesTable.querySelector("tbody");
    if (tBody) {
        tBody.innerHTML = '';
    }
}


function registerSale() {
    if (allCarcases.cutsBeefCattle.carcases.length == 0) {
        showAlert('At least one carcase must be added', MESSAGE_TYPE.WARNING);
        return false;
    }

    var fcn = 'registerSale';
    let saleTypes = document.getElementsByName("saleType");
    if (saleTypes) {
        for (var i = 0; i < saleTypes.length; i++) {
            var item = saleTypes[i];
            if (item && item.checked) {
                fcn = item.value;
            }
        }
    }
    let key = formRegisterCarcase.querySelector('#invoiceNumber').value;
    var org;
    let orgs = document.getElementsByName("slaughterhouse");
    if (orgs) {
        for (var i = 0; i < orgs.length; i++) {
            var item = orgs[i];
            if (item && item.checked) {
                org = item.value;
            }
        }
    }

    allCarcases.orgName = org.toLocaleLowerCase();
    allCarcases.key = key;
    allCarcases.fcn = fcn;
    allCarcases.mspID = org + "MSP";
    allCarcases.cutsBeefCattle.property = allCarcases.orgName + DOMAIN;
    allCarcases.cutsBeefCattle.slaughterhouse = 'slaughterhouse' + DOMAIN;
    allCarcases.cutsBeefCattle.market = 'market' + DOMAIN;

    showAlert('Registering sale...', MESSAGE_TYPE.INFO);
    var xhr = new XMLHttpRequest();
    xhr.timeout = 300000;
    xhr.onreadystatechange = function() {
        console.log(xhr.status);
        if (this.readyState === 4) {
            if (this.status === 200) {
                 allCarcases.cutsBeefCattle.carcases = [];
                resetCarcaseTable();
                showAlert('Sale registered successfully.', MESSAGE_TYPE.SUCCESS);
            } else {
                showAlert('Failed registering sale. Try again.', MESSAGE_TYPE.ERROR);
            }
        }
    };

    xhr.open('POST', BASE_URL + '/invokeslaughterhouse', true);
    xhr.send(JSON.stringify(allCarcases));
    return false;
}
