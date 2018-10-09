/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

/**
 * Renders the view for query slaughterhouse sales.
 * 
 * @param {Object} application object
 * @param {Object} req request object
 * @param {Object} res response object
 */
module.exports.showQuery = function (application, req, res) {
	res.render("slaughterhouse/query");
};

/**
 * Renders the view for registering slaughterhouse sales.
 * 
 * @param {Object} application object
 * @param {Object} req request object
 * @param {Object} res response object
 */
module.exports.showRegister = function (application, req, res) {
	res.render("slaughterhouse/register");
};



