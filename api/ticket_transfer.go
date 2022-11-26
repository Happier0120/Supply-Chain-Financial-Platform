package api

import (
	"encoding/json"
	"fmt"
	"yh/model"
	"yh/pkg/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const ticketCollection = "ticketCollection"

// 默认设定核心企业为初始化环境的组织
var coreOrgMSPID string

type SmartContract struct {
	contractapi.Contract
}

// 初始化操作
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) (string, error) {
	//预设核心企业为初始化环境org
	var err error
	coreOrgMSPID, err = ctx.GetClientIdentity().GetMSPID() // Org1MSP
	if err != nil {
		return "", fmt.Errorf("init coreOrgID failed: %v", err)
	}
	if err = ctx.GetStub().PutState(coreOrgMSPID, []byte(coreOrgMSPID)); err != nil {
		return "", fmt.Errorf("failed to put core org MSPID")
	}
	//添加账户列表
	var accountIds = [4]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
		"4e07408562be",
	}
	var userNames = [4]string{"company_core", "company_a", "company_b", "company_c"}
	var balances = [4]float64{10000, 0, 0, 0}
	//初始化账号数据
	var accountList []model.Account
	for i, val := range accountIds {
		account := model.Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		accountList = append(accountList, account)
		// 写入账本
		accountBytes, err := json.Marshal(account)
		if err != nil {
			return "", fmt.Errorf("failed to marshal account")
		}

		accountKey, err := ctx.GetStub().CreateCompositeKey(model.AccountKey, []string{account.AccountId})
		if err != nil {
			return "", fmt.Errorf("Failed to create composite key: %v", err)
		}

		err = ctx.GetStub().PutState(accountKey, accountBytes)
		if err != nil {
			return "", fmt.Errorf("Failed to write accountlists into ledger: %v", err)
		}
	}

	return fmt.Sprintf("CoreOrgMSPID: %v, Account list: %+v", coreOrgMSPID, accountList), nil
}

// 只有核心企业可以发行票据
// CreateTicket只有核心企业可以调用
func (s *SmartContract) CreateTicket(ctx contractapi.TransactionContextInterface, ticketID string, description string) (string, error) {

	transientMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", fmt.Errorf("error getting transient: %v", err)
	}
	transientJSON, ok := transientMap["ticket_properties"]
	if !ok {
		return "", fmt.Errorf("ticket_properties key not found in the transient map")
	}
	// Get MSPID of submitting client identity
	clientOrgMSPID, err := utils.GetClientOrgMSPID(ctx, false)
	if err != nil {
		return "", fmt.Errorf("failed to get clientOrgMSPID: %v", err)
	}
	// check if requestor is coreOrg, only core org could create tickets
	if coreOrgMSPID != clientOrgMSPID {
		return "", fmt.Errorf("Only core org could create tickets: %v", err)
	}
	// check if ticket already exists
	ticketAsBytes, err := ctx.GetStub().GetState(ticketID)
	if err != nil {
		return "", fmt.Errorf("failed to get ticket: %v", err)
	} else if ticketAsBytes != nil {
		fmt.Println("Ticket already exists: " + ticketID)
		return "", fmt.Errorf("this ticket already exists: " + ticketID)
	}

	// submitting ticket
	ticket := model.Ticket{
		Type:          ticketCollection,
		ID:            ticketID,
		OwnerOrgMSPID: coreOrgMSPID,
		Description:   description,
	}
	ticketJSON, err := json.Marshal(ticket)
	if err != nil {
		return "", fmt.Errorf("failed to marshal ticket into JSON: %v", err)
	}

	if err = ctx.GetStub().PutState(ticketID, ticketJSON); err != nil {
		return "", fmt.Errorf("failed to put ticket state: %v", err)
	}
	// setting endorsement, only owner can update or query private data
	if err = utils.SetTicketStateBasedEndorsement(ctx, ticket.ID, coreOrgMSPID); err != nil {
		return "", fmt.Errorf("failed to set ticket endorsement: %v", err)
	}

	collection := utils.BuildCollectionName(coreOrgMSPID)
	if err = ctx.GetStub().PutPrivateData(collection, ticketID, transientJSON); err != nil {
		return "", fmt.Errorf("failed to put transientJSON to private data: %v", err)
	}

	return fmt.Sprintf("%v have been created: %+v, \n private detial: %v", ticketID, ticket, transientJSON), nil
}

// transfer ticket
func (s *SmartContract) TransferTicket(ctx contractapi.TransactionContextInterface, ticketID string, toOrgMSPID string) (string, error) {
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
		return "", fmt.Errorf("Only owner %v could transfet this ticket: %v", ticket.OwnerOrgMSPID, ticket.ID)
	}
	// process changes
	ticket.OwnerOrgMSPID = toOrgMSPID
	updateTicketJSON, err := json.Marshal(ticket)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal ticket: %v", err)
	}

	err = ctx.GetStub().PutState(ticketID, updateTicketJSON)
	if err != nil {
		return "", fmt.Errorf("When transfer ticket, falied to write new ticket into ledger: %v", err)
	}

	// setting endorsement, only owner can update or query private data
	err = utils.SetTicketStateBasedEndorsement(ctx, ticketID, toOrgMSPID)
	if err != nil {
		return "", fmt.Errorf("failed to set ticket endorsement: %v", err)
	}
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", fmt.Errorf("error getting transient data: %v", err)
	}

	immutablePropertiesJSON, ok := transMap["ticket_properties"]
	if !ok {
		return "", fmt.Errorf("ticket_properties key not found in the transient map")
	}

	collectionBuyer := utils.BuildCollectionName(toOrgMSPID)
	err = ctx.GetStub().PutPrivateData(collectionBuyer, ticket.ID, immutablePropertiesJSON)
	if err != nil {
		return "", fmt.Errorf("failed to put Asset private properties for buyer: %v", err)
	}

	return fmt.Sprintf("Succeed! Transfer %v from %v to %v \n", ticket.ID, clientOrgID, toOrgMSPID), nil
}

// 更改票据公开信息备注
func (s *SmartContract) ChangeDescription(ctx contractapi.TransactionContextInterface, ticketID, newDescription string) (string, error) {
	// check that only owner can change description
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
		return "", fmt.Errorf("Only owner can change description")
	}

	// process changes
	ticket.Description = newDescription
	updateTicketJSON, err := json.Marshal(ticket)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal ticket: %v", err)
	}

	err = ctx.GetStub().PutState(ticketID, updateTicketJSON)
	if err != nil {
		return "", fmt.Errorf("When change tickets' public description, falied to write new ticket into ledger: %v", err)
	}

	return fmt.Sprintf("Succeed! Change %v's descirption to %v \n", ticket.ID, ticket.Description), nil
}
