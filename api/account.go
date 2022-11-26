package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"yh/model"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// QueryAccountList 查询账户列表
func (s *SmartContract) QueryAccountInfor(ctx contractapi.TransactionContextInterface, accountId string) (string, error) {
	var accountList []model.Account
	var account model.Account

	if accountId == "" {
		resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(model.AccountKey, []string{})
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
		resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(model.AccountKey, []string{accountId})
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
