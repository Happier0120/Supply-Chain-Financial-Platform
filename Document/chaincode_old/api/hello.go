package api

import (
	"fmt"
	"yh/model"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Hello 测试
func (s *SmartContract) Hello(ctx contractapi.TransactionContextInterface, name string) (string, error) {
	testKey, err := ctx.GetStub().CreateCompositeKey(model.TestKey, []string{name})
	if err != nil {
		return "", fmt.Errorf("Failed to create composite key: %v", err)
	}

	err = ctx.GetStub().PutState(testKey, []byte(name))
	if err != nil {
		return "", fmt.Errorf("failed to put hello test: %v", err)
	}
	helloJSON, err := ctx.GetStub().GetState("hello" + name)
	if err != nil {
		return "", fmt.Errorf("failed to get hello test: %v", err)
	}
	result := "Succeed! " + string(helloJSON)
	return result, nil
}
