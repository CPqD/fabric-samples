/*******************************************************************************
 * Copyright 2018 CPqD. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// BeefCattle struct to register cattle
type BeefCattle struct {
	ID                               string `json:"id"`
	Weight                           string `json:"weight"`
	Breed                            string `json:"breed"`
	DtLastBrucellosisVaccine         string `json:"dtLastBrucellosisVaccine"`
	DtLastFootAndMouthDeseaseVaccine string `json:"dtLastFootAndMouthDeseaseVaccine"`
	Age                              string `json:"age"`
	Classe                           string `json:"classe"`
	ProductionSystem                 string `json:"productionSystem"`
}
