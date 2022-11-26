package api

import (
	"encoding/json"
	"fmt"
	"yh/model"
	"yh/pkg/utils"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 查询票据公开数据
func (s *SmartContract) QueryTicket(ctx contractapi.TransactionContextInterface, ticketID string) ([]*model.Ticket, error) {
	var ticketList []*model.Ticket
	ticket, err := s.ReadTicket(ctx, ticketID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get ticket state: %v", err)
	}
	ticketList = append(ticketList, ticket)

	return ticketList, nil
}

// 查询票据隐私数据
func (s *SmartContract) QueryTicketPrivate(ctx contractapi.TransactionContextInterface, ticketID string) (string, error) {

	ticketprivate, err := s.ReadTicketPrivateProperties(ctx, ticketID)
	if err != nil {
		return "", fmt.Errorf("Failed to get private data: %v", err)
	}

	return fmt.Sprintf("Succeed! %v's private field: %+v", ticketID, ticketprivate), nil
}

func (s *SmartContract) ReadTicket(ctx contractapi.TransactionContextInterface, ticketID string) (*model.Ticket, error) {
	// Since only public data is accessed in this function, no access control is required
	ticketJSON, err := ctx.GetStub().GetState(ticketID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if ticketJSON == nil {
		return nil, fmt.Errorf("%s does not exist", ticketID)
	}

	var ticket *model.Ticket
	err = json.Unmarshal(ticketJSON, &ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

// 查询票据私有数据
func (s *SmartContract) ReadTicketPrivateProperties(ctx contractapi.TransactionContextInterface, ticketID string) (string, error) {
	// check if request is sent by owner
	ticketBytes, err := ctx.GetStub().GetState(ticketID)
	if err != nil {
		return "", fmt.Errorf("Failed to get ticket state: %v", err)
	}
	var ticket model.Ticket
	if err = json.Unmarshal(ticketBytes, &ticket); err != nil {
		return "", fmt.Errorf("Failed to unmarshal ticket: %v", err)
	}
	clientOrgID, err := utils.GetClientOrgMSPID(ctx, false)
	if err != nil {
		return "", fmt.Errorf("Failed to get client Org MSPID: %v", err)
	}
	if clientOrgID != ticket.OwnerOrgMSPID {
		return "", fmt.Errorf("Only owner can query private data")
	}

	// get private data
	collectionSender := utils.BuildCollectionName(clientOrgID)
	privatedata, err := ctx.GetStub().GetPrivateData(collectionSender, ticketID)
	if err != nil {
		return "", fmt.Errorf("Failed to get ticket private data: %v", err)
	}

	result := string(privatedata)
	return result, nil
}

// 溯源，查询指定票据的从发行到结算的流通全过程
func (s *SmartContract) QueryTicketHistory(ctx contractapi.TransactionContextInterface, ticketID string) ([]*model.QueryResult, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(ticketID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get history for key: %v", err)
	}
	defer resultsIterator.Close()

	var queryResultList []*model.QueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Failed to get iterator next: %v", err)
		}

		var ticket model.Ticket
		err = json.Unmarshal(response.Value, &ticket)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal ticket: %v", err)
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("Failed to get timestamp: %v", err)
		}
		record := &model.QueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    ticket,
		}
		queryResultList = append(queryResultList, record)
	}

	return queryResultList, nil
}

// 查询所有公开数据
func (s *SmartContract) GetAllTicket(ctx contractapi.TransactionContextInterface) ([]*model.Ticket, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("Failed to get all tickets: %v", err)

	}
	defer resultsIterator.Close()

	var ticketList []*model.Ticket
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Failed to get Iterator next: %v", err)
		}

		var ticket *model.Ticket
		err = json.Unmarshal(queryResponse.Value, &ticket)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal ticket: %v", err)
		}
		ticketList = append(ticketList, ticket)
	}

	return ticketList, nil
}
