/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// BeefCattleSale struct that contains cattle sail details
type BeefCattleSale struct {
	Property       string       `json:"property"`
	Slaughterhouse string       `json:"slaughterhouse"`
	Cattle         []BeefCattle `json:"cattle"`
}
