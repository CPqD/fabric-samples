/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// CutsBeefCattle struct
type CutsBeefCattle struct {
	Slaughterhouse string           `json:"slaughterhouse"`
	Market         string           `json:"market"`
	Carcases       []CarcasesStruct `json:"carcases"`
}
