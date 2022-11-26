/*
 SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	"yh/api"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&api.SmartContract{})
	if err != nil {
		log.Panicf("Error create ticket chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting ticket chaincode: %v", err)
	}
}
