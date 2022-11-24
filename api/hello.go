package api

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Hello 测试
func (s *SmartContract) Hello(ctx contractapi.TransactionContextInterface, name string) (string, error) {

	err := ctx.GetStub().PutState("hello"+name, []byte(name))
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
