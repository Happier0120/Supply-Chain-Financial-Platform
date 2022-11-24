/*
 SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	"yh/api"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)



// Invoke 实现Invoke接口调用智能合约
// func (s *SmartContract) Invoke(ctx contractapi.TransactionContextInterface) pb.Response {
// 	funcName, args := ctx.GetStub().GetFunctionAndParameters()

// 	switch funcName {
// 	case "Hello":
// 		return api.Hello(ctx, args)
// 	case "QueryAccountList":
// 		return api.QueryAccountList(ctx, args)
// 	case "QueryAccountInformation":
// 		return api.QueryAccountInformation(ctx, args)
// 	case "CreateTicket":
// 		return api.CreateTicket(ctx, args)
// 	case "TransferTicket":
// 		return api.TransferTicket(ctx, args)
// 	case "ChangePublicDescription":
// 		return api.ChangePublicDescription(ctx, args)
// 	case "QueryTicketPublicProperties":
// 		return api.QueryTicketPublicProperties(ctx, args)
// 	case "QueryTicketPrivateProperties":
// 		return api.QueryTicketPrivateProperties(ctx, args)
// 	case "QueryTicketHistory":
// 		return api.QueryTicketHistory(ctx, args)
// 	case "GetAllTicket":
// 		return api.GetAllTicket(ctx)
// 	default:
// 		return shim.Error(fmt.Sprintf("未找到指定功能: %s", funcName))
// 	}
// }

func main() {
	chaincode, err := contractapi.NewChaincode(&api.SmartContract{})
	if err != nil {
		log.Panicf("Error create ticket chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting ticket chaincode: %v", err)
	}
}
