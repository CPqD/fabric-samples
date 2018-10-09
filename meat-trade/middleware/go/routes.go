/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import "net/http"

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes define collection of routes
type Routes []Route

// Create the collection of routes
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"InvokeFarm",
		"POST",
		"/invokefarm",
		InvokeFarm,
	},
	Route{
		"QueryFarm",
		"POST",
		"/queryfarm",
		QueryFarm,
	},
	Route{
		"InvokeSlaughterhouse",
		"POST",
		"/invokeslaughterhouse",
		InvokeSlaughterhouse,
	},
	Route{
		"QuerySlaughterhouse",
		"POST",
		"/queryslaughterhouse",
		QuerySlaughterhouse,
	},
}
