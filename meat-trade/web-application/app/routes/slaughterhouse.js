/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */


/**
 * All routes slaughterhouse related must to be placed in this file. 
 * From here you can call a given method in the controller layer, for example, to render an specific page or to handle form data.
 * @param {Object} application 
 */
module.exports = function (application) {
	
	application.get('/slaughterhouse/query', function (req, res) {
		application.app.controllers.slaughterhouse.showQuery(application, req, res);
	});

	application.get('/slaughterhouse', function (req, res) {
		application.app.controllers.slaughterhouse.showRegister(application, req, res);
    });
    
    application.get('/slaughterhouse/register', function (req, res) {
		application.app.controllers.slaughterhouse.showRegister(application, req, res);
	});
}