package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名称
	Balance   float64 `json:"balance"`   //余额
}

// QueryAccountList 查询账户列表
func (s *SmartContract) QueryAccountInfor(ctx contractapi.TransactionContextInterface, accountId string) (string, error) {
	var accountList []Account
	var account Account

	if accountId == "" {
		resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(AccountKey, []string{})
		if err != nil {
			return "", fmt.Errorf("Get accountList error: %v", err)
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return "", errors.New(fmt.Sprintf("get next hash error: %v", err))
			}
			err = json.Unmarshal(val.GetValue(), &account)
			if err != nil {
				return "", errors.New(fmt.Sprintf("unmarshal error: %v", err))
			}
			accountList = append(accountList, account)
		}
	} else {
		resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(AccountKey, []string{accountId})
		if err != nil {
			return "", fmt.Errorf("Get accountList error: %v", err)
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return "", errors.New(fmt.Sprintf("get next hash error: %v", err))
			}
			err = json.Unmarshal(val.GetValue(), &account)
			if err != nil {
				return "", errors.New(fmt.Sprintf("unmarshal error: %v", err))
			}
			accountList = append(accountList, account)
		}
	}

	return fmt.Sprintf("%+v", accountList), nil
}
