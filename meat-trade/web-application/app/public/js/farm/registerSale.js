/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

var formAddCattle, cattleTable, formRegisterCattle;
var allCattle = {};
allCattle.sale = {};
allCattle.sale.cattle = [];

(function () {
    formAddCattle = document.querySelector("#formAddCattle");
    formRegisterCattle = document.querySelector("#formRegisterCattle");
    cattleTable = document.querySelector("#cattleTable");
}());


function addCattle() {
    console.log("validateAddCattle called");

    let cattle = {};
    cattle.id = formAddCattle.querySelector("#cattleId").value;
    cattle.weight = formAddCattle.querySelector("#weight").value;
    cattle.breed = formAddCattle.querySelector("#breed").value;
    cattle.dtLastBrucellosisVaccine = formAddCattle.querySelector("#brucellosisVaccine").value;
    cattle.dtLastFootAndMouthDeseaseVaccine = formAddCattle.querySelector("#footAndMountVaccine").value;
    cattle.age = formAddCattle.querySelector("#age").value;
    cattle.classe = formAddCattle.querySelector("#class").value;
    let productionSystemOptions = document.getElementsByName("productionSystem");
    if (productionSystemOptions) {
        for (var i = 0; i < productionSystemOptions.length; i++) {
            var item = productionSystemOptions[i];
            if (item && item.checked) {
                cattle.productionSystem = item.value;
            }
        }
    }

    allCattle.sale.cattle.push(cattle);
    appendCattleToTable(cattle);
    formAddCattle.reset()
    return false;
}


function appendCattleToTable(cattle) {
    let tBody = addedCattle.querySelector("tbody");
    if (tBody) {
        let register = "<tr>"
            + "<th>" + cattle.id + "</th>"
            + "<td>" + cattle.weight + "</td>"
            + "<td>" + cattle.breed + "</td>"
            + "<td>" + cattle.dtLastBrucellosisVaccine + "</td>"
            + "<td>" + cattle.dtLastFootAndMouthDeseaseVaccine + "</td>"
            + "<td>" + cattle.age + "</td>"
            + "<td>" + cattle.classe + "</td>"
            + "<td>" + cattle.productionSystem + "</td>"
            + "</tr>";

        tBody.insertAdjacentHTML('beforeend', register)
    }
}

function resetCattleTable() {
    let tBody = addedCattle.querySelector("tbody");
    if (tBody) {
        tBody.innerHTML = '';
    }
}


function registerSale() {
    if (allCattle.sale.cattle.length == 0) {
        showAlert('At least one cattle must be added', MESSAGE_TYPE.WARNING);
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
    let key = formRegisterCattle.querySelector('#invoiceNumber').value;
    var org;
    let orgs = document.getElementsByName("farm");
    if (orgs) {
        for (var i = 0; i < orgs.length; i++) {
            var item = orgs[i];
            if (item && item.checked) {
                org = item.value;
            }
        }
    }

    allCattle.orgName = org.toLocaleLowerCase();
    allCattle.key = key;
    allCattle.fcn = fcn;
    allCattle.mspID = org + "MSP";
    allCattle.sale.property = allCattle.orgName + DOMAIN;
    allCattle.sale.slaughterhouse = 'slaughterhouse' + DOMAIN;

    showAlert('Registering sale...', MESSAGE_TYPE.INFO);
    var xhr = new XMLHttpRequest();
    xhr.timeout = 300000;
    
    xhr.onreadystatechange = function() {
        console.log(xhr.status);
        if (this.readyState === 4) {
            if (this.status === 200) {
                allCattle.sale.cattle = [];
                resetCattleTable();
                showAlert('Sale registered successfully.', MESSAGE_TYPE.SUCCESS);
            } else {
                showAlert('Failed registering sale. Try again.', MESSAGE_TYPE.ERROR);
            }
        }
    };

    xhr.open('POST', BASE_URL + '/invokefarm', true);
    xhr.send(JSON.stringify(allCattle));
    return false;
}
