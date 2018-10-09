/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

var alertElement;

let BASE_URL = 'http://localhost:7500';
let DOMAIN = '.kingbeefcattle.com';

(function(){
    console.log('loaded')
    alertElement = document.querySelector('#alert');
}());

let MESSAGE_TYPE = {
    WARNING : 0,     
    ERROR: 1,
    SUCCESS: 2,
    INFO : 3
}


function showAlert(message, type) {
    hideAlert();
    resetAlertClass();
    var alertclass = "alert-info";
    switch(type) {
        case MESSAGE_TYPE.ERROR:
        alertclass = 'alert-danger';
        break;
        case MESSAGE_TYPE.WARNING:
        alertclass = 'alert-warning';
        break;
        case MESSAGE_TYPE.SUCCESS:
        alertclass = 'alert-success';
        break;        
    }
    alertElement.classList.add(alertclass);
    alertElement.querySelector("h5").innerHTML = message;

    alertElement.style.display = 'block'

}

function hideAlert() {
    alertElement.style.display = 'none';
}


function resetAlertClass() {
    for(var i=0; i<alertElement.classList.length; i++) {
        let c = alertElement.classList[i];
        alertElement.classList.remove(c);
    }
    alertElement.classList.add('alert');
}