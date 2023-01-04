package api

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) ([]*model.Account, error) {
	//预设核心企业为初始化环境org
	var err error
	coreOrgMSPID, err = ctx.GetClientIdentity().GetMSPID() // Org1MSP
	if err != nil {
		return nil, fmt.Errorf("init coreOrgID failed: %v", err)
	}

	//添加账户列表
	var accountIds = [4]string{
		"Org1MSP",
		"Org2MSP",
		"Org3MSP",
		"Org4MSP",
	}
	var userNames = [4]string{"Org1MSP", "Org2MSP", "Org3MSP", "Org4MSP"} // Org1MSP为核心组织
	var balances = [4]float64{10000, 0, 0, 0}                             // 银行初始为核心企业授信 $10000

	//初始化账号数据
	var accountList []*model.Account
	for i, val := range accountIds {
		account := &model.Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		accountList = append(accountList, account)

		// 写入账本
		accountBytes, err := json.Marshal(account)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal account: %v", err)
		}
		accountKey, err := ctx.GetStub().CreateCompositeKey(model.AccountKey, []string{account.AccountId})
		if err != nil {
			return nil, fmt.Errorf("Failed to create composite key: %v", err)
		}
		err = ctx.GetStub().PutState(accountKey, accountBytes)
		if err != nil {
			return nil, fmt.Errorf("Failed to write accountlists into ledger: %v", err)
		}
	}
	return accountList, nil
}

// 只有核心企业可以发行票据
// CreateTicket只有核心企业可以调用
func (s *SmartContract) CreateTicket(ctx contractapi.TransactionContextInterface, toOrgMSPID, description string) ([]*model.Ticket, error) {

	transientMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return nil, fmt.Errorf("error getting transient: %v", err)
	}
	transientJSON, ok := transientMap["ticket_properties"]
	if !ok {
		return nil, fmt.Errorf("ticket_properties key not found in the transient map")
	}
	// Get MSPID of submitting client identity
	clientOrgMSPID, err := utils.GetClientOrgMSPID(ctx, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get clientOrgMSPID: %v", err)
	}
	// check if requestor is coreOrg, only core org could create tickets
	if coreOrgMSPID != clientOrgMSPID {
		return nil, fmt.Errorf("Only core org could create tickets: %v", err)
	}

	// submitting ticket
	ticketID := utils.RandomString(10)

	channelID := ctx.GetStub().GetChannelID()
	txID := ctx.GetStub().GetTxID()
	txTimestamp, _ := ctx.GetStub().GetTxTimestamp()

	var ticketList []*model.Ticket
	ticket := &model.Ticket{
		Type:          ticketCollection,
		ID:            ticketID,
		ChannelID:     channelID,
		TxID:          txID,
		TxTimestamp:   txTimestamp.String(),
		OwnerOrgMSPID: toOrgMSPID,
		Guarantor:     "ICBC",
		CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
		DueDate:       "2023-12-12",
		Description:   description,
	}

	ticketList = append(ticketList, ticket)

	ticketJSON, err := json.Marshal(ticket)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ticket into JSON: %v", err)
	}

	if err = ctx.GetStub().PutState(ticketID, ticketJSON); err != nil {
		return nil, fmt.Errorf("failed to put ticket state: %v", err)
	}
	// setting endorsement, only owner can update or query private data
	if err = utils.SetTicketStateBasedEndorsement(ctx, ticket.ID, coreOrgMSPID); err != nil {
		return nil, fmt.Errorf("failed to set ticket endorsement: %v", err)
	}

	collection := utils.BuildCollectionName(toOrgMSPID)
	if err = ctx.GetStub().PutPrivateData(collection, ticketID, transientJSON); err != nil {
		return nil, fmt.Errorf("failed to put transientJSON to private data: %v", err)
	}
	log.Printf("Submitting clientOrg is :%v, collection is: %v", coreOrgMSPID, collection)

	return ticketList, nil
}

// transfer ticket
func (s *SmartContract) TransferTicket(ctx contractapi.TransactionContextInterface, ticketID, toOrgMSPID string) ([]*model.Ticket, error) {
	ticketBytes, err := ctx.GetStub().GetState(ticketID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get ticket state: %v", err)
	}
	var ticket model.Ticket
	if err = json.Unmarshal(ticketBytes, &ticket); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal ticket: %v", err)
	}

	clientOrgID, err := utils.GetClientOrgMSPID(ctx, false)
	if err != nil {
		return nil, fmt.Errorf("Failed to get client Org MSPID: %v", err)
	}
	if clientOrgID != ticket.OwnerOrgMSPID {
		return nil, fmt.Errorf("Only owner %v could transfet this ticket: %v", ticket.OwnerOrgMSPID, ticket.ID)
	}
	// process changes
	var ticketList []*model.Ticket
	ticket.OwnerOrgMSPID = toOrgMSPID
	ticketList = append(ticketList, &ticket)

	updateTicketJSON, err := json.Marshal(ticket)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal ticket: %v", err)
	}

	err = ctx.GetStub().PutState(ticketID, updateTicketJSON)
	if err != nil {
		return nil, fmt.Errorf("When transfer ticket, falied to write new ticket into ledger: %v", err)
	}

	err = utils.SetTicketStateBasedEndorsement(ctx, ticketID, toOrgMSPID)
	if err != nil {
		return nil, fmt.Errorf("failed to set ticket endorsement: %v", err)
	}
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return nil, fmt.Errorf("error getting transient data: %v", err)
	}

	immutablePropertiesJSON, ok := transMap["ticket_properties"]
	if !ok {
		return nil, fmt.Errorf("ticket_properties key not found in the transient map")
	}

	collectionBuyer := utils.BuildCollectionName(toOrgMSPID)
	err = ctx.GetStub().PutPrivateData(collectionBuyer, ticket.ID, immutablePropertiesJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put Asset private properties for buyer: %v", err)
	}

	return ticketList, nil
}

// 更改票据公开信息备注
func (s *SmartContract) ChangeDescription(ctx contractapi.TransactionContextInterface, ticketID, newDescription string) ([]*model.Ticket, error) {
	// check that only owner can change description
	ticketBytes, err := ctx.GetStub().GetState(ticketID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get ticket state: %v", err)
	}
	var ticket model.Ticket
	if err = json.Unmarshal(ticketBytes, &ticket); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal ticket: %v", err)
	}

	clientOrgID, err := utils.GetClientOrgMSPID(ctx, false)
	if err != nil {
		return nil, fmt.Errorf("Failed to get client Org MSPID: %v", err)
	}
	if clientOrgID != ticket.OwnerOrgMSPID {
		return nil, fmt.Errorf("Only owner can change description")
	}

	// process changes
	var ticketList []*model.Ticket
	ticket.Description = newDescription
	updateTicketJSON, err := json.Marshal(ticket)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal ticket: %v", err)
	}
	ticketList = append(ticketList, &ticket)

	err = ctx.GetStub().PutState(ticketID, updateTicketJSON)
	if err != nil {
		return nil, fmt.Errorf("When change tickets' public description, falied to write new ticket into ledger: %v", err)
	}

	return ticketList, nil
}
