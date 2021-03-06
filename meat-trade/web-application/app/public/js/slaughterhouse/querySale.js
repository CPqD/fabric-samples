/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

var formQuerySale, queryResultTable;

(function () {
    queryResultTable = document.querySelector("#tableResult");
    formQuerySale = document.querySelector('#formQuerySale');
}());


function addQueryResulToTable(result) {
    let tBody = queryResultTable.querySelector("tbody");
    if (tBody) {
        tBody.innerHTML = '';
        for (var i = 0; i < result.carcases.length; i++) {
            let carcase = result.carcases[i];
            let register = "<tr>"
                + "<th>" + carcase.id + "</th>"
                + "<th>" + carcase.idAnimal + "</th>"
                + "<td>" + carcase.weight + "</td>"
                + "<td>" + carcase.type + "</td>"
                + "</tr>";
            tBody.insertAdjacentHTML('beforeend', register)
        }
    }
}

function clearResultTable() {
    let tBody = queryResultTable.querySelector("tbody");    
    if (tBody) {
        tBody.innerHTML = '';
    } 
}

function resetResultTable() {
    let tBody = queryResultTable.querySelector("tbody");
    if (tBody) {
        tBody.innerHTML = '';
        let register = '<tr>'
        + '<th colspan="5">No records found</th>'
        + "</tr>";
        tBody.insertAdjacentHTML('beforeend', register)
    }
}


function querySale() {
  
    var fcn = 'querySale';
    let saleTypes = document.getElementsByName("saleType");
    if (saleTypes) {
        for (var i = 0; i < saleTypes.length; i++) {
            var item = saleTypes[i];
            if (item && item.checked) {
                fcn = item.value;
            }
        }
    }

    let key = formQuerySale.querySelector('#invoiceNumber').value;
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

    let request = {};
    request.orgName = org.toLocaleLowerCase();
    request.key = key;
    request.fcn = fcn;
    request.mspID = org + "MSP";
    
    clearResultTable()
    showAlert('Querying sale...', MESSAGE_TYPE.INFO);
    var xhr = new XMLHttpRequest();
    xhr.timeout = 300000;
    xhr.onreadystatechange = function () {        
        if (this.readyState === 4) {
            console.log(xhr.responseText);
            if (this.status === 200) {                
                resetResultTable();
                
                let response = this.responseText;
                if (response) {
                    addQueryResulToTable(JSON.parse(response));
                }
                
                showAlert('Query succeeded.', MESSAGE_TYPE.SUCCESS);
            } else if (this.status == 404) {
                resetResultTable();
                showAlert('No records found.', MESSAGE_TYPE.WARNING);
            } else {
                showAlert('Failed querying sale. Try again.', MESSAGE_TYPE.ERROR);
            }
        }
    };

    xhr.open('POST', BASE_URL + '/queryslaughterhouse', true);
    xhr.send(JSON.stringify(request));
    return false;
}
