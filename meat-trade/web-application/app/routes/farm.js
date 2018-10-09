/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

/**
 * All routes farm related must to be placed in this file. 
 * From here you can call a given method in the controller layer, for example, to render an specific page or to handle form data.
 * @param {Object} application 
 */
module.exports = function (application) {
	
	application.get('/farm/query', function (req, res) {
		application.app.controllers.farm.showQuery(application, req, res);
	});

	application.get('/farm', function (req, res) {
		application.app.controllers.farm.showRegister(application, req, res);
    });
    
    application.get('/farm/register', function (req, res) {
		application.app.controllers.farm.showRegister(application, req, res);
	});
}