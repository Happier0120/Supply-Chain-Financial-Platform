package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"yh/model"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 通过AccountId获取账户
func (s *SmartContract) QueryAccountaById(ctx contractapi.TransactionContextInterface, accountId string) ([]*model.Account, error) {
	var accountList []*model.Account
	var account *model.Account

	if accountId == "" {
		resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(model.AccountKey, []string{})
		if err != nil {
			return nil, fmt.Errorf("Get accountList error: %v", err)
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("get next hash error: %v", err))
			}
			err = json.Unmarshal(val.GetValue(), &account)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("unmarshal error: %v", err))
			}
			accountList = append(accountList, account)
		}
	} else {
		resultIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(model.AccountKey, []string{accountId})
		if err != nil {
			return nil, fmt.Errorf("Get accountList error: %v", err)
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("get next hash error: %v", err))
			}
			err = json.Unmarshal(val.GetValue(), &account)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("unmarshal error: %v", err))
			}
			accountList = append(accountList, account)
		}
	}

	return accountList, nil
}

// 更新账户余额
func (s *SmartContract) UpdateAccountBalance(ctx contractapi.TransactionContextInterface, from, to, ticketID string) error {
	privateDataString, err := s.QueryTicketPrivate(ctx, ticketID)
	if err != nil {
		return fmt.Errorf("Query ticket private data failed, err:%v", err)
	}
	var privateDataJSON, _ = json.Marshal(privateDataString)

	var privateData model.TicketPrivate
	if err = json.Unmarshal(privateDataJSON, &privateData); err != nil {
		return err
	}
	value, err := strconv.ParseFloat(privateData.Value, 64)
	if err != nil {
		return err
	}

	// change balances
	fromAccountList, err := s.QueryAccountaById(ctx, from)
	if err != nil {
		return err
	}
	fromAccount := fromAccountList[0]
	toAccountList, err := s.QueryAccountaById(ctx, to)
	if err != nil {
		return err
	}
	toAccount := toAccountList[0]

	fromAccount.Balance -= value
	toAccount.Balance += value

	fromAccountJson, err := json.Marshal(fromAccount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(fromAccount.AccountId, fromAccountJson)
	if err != nil {
		return err
	}

	toAccountJson, err := json.Marshal(toAccount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(fromAccount.AccountId, toAccountJson)
	if err != nil {
		return err
	}
	return nil
}

// 查询账户余额
func (s *SmartContract) getAccountBalance(ctx contractapi.TransactionContextInterface, accountID string) (float64, error) {
	accountList, err := s.QueryAccountaById(ctx, accountID)
	if err != nil {
		return 0, err
	}
	return accountList[0].Balance, nil
}
