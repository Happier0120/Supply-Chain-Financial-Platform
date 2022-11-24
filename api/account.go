package api

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名称
	Balance   float64 `json:"balance"`   //余额
}

// QueryAccountList 查询账户列表
func (s *SmartContract) QueryAccountList(ctx contractapi.TransactionContextInterface) (string, error) {
	return "", nil
}

func (s *SmartContract) QueryAccountInformation(ctx contractapi.TransactionContextInterface, args []string) (string, error) {
	return "", nil
}
