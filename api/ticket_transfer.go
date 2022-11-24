package api

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/statebased"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const ticketCollection = "ticketCollection"
const (
	AccountKey = "A"
	TicketKey  = "T"
)

// 默认设定核心企业为初始化环境的组织
var coreOrgMSPID string

type SmartContract struct {
	contractapi.Contract
}

type Ticket struct {
	Type          string `json:"objectType"`
	ID            string `json:"ticketID"`    //票据ID
	OwnerOrgMSPID string `json:"ownerOrg"`    //票据所有者ID
	Description   string `json:"description"` //票据备注
}

type TicketPrivateDetails struct {
	ID    string `json:"assetID"`
	Value int    `json:"value"`
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
	var userNames = [4]string{"管理员", "①号物流", "②号物流", "③号物流"}
	var balances = [4]float64{10000, 0, 0, 0}
	//初始化账号数据
	for i, val := range accountIds {
		account := Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		// 写入账本
		accountBytes, err := json.Marshal(account)
		if err != nil {
			return "", fmt.Errorf("failed to marshal account")
		}
		if err = ctx.GetStub().PutState(val, accountBytes); err != nil {
			return "", fmt.Errorf("Failed to write accountlists into ledger: %v", err)
		}
	}

	result := coreOrgMSPID
	return result, nil
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
	clientOrgMSPID, err := getClientOrgMSPID(ctx, false)
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
	ticket := Ticket{
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
	if err = setAssetStateBasedEndorsement(ctx, ticket.ID, coreOrgMSPID); err != nil {
		return "", fmt.Errorf("failed to set ticket endorsement: %v", err)
	}

	collection := buildCollectionName(coreOrgMSPID)
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
	var ticket Ticket
	if err = json.Unmarshal(ticketBytes, &ticket); err != nil {
		return "", fmt.Errorf("Failed to unmarshal ticket: %v", err)
	}

	clientOrgID, err := getClientOrgMSPID(ctx, false)
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
	if err = setAssetStateBasedEndorsement(ctx, ticketID, toOrgMSPID); err != nil {
		return "", fmt.Errorf("failed to set ticket endorsement: %v", err)
	}
	// collectionSender := buildCollectionName(clientOrgID)

	// immutablePropertiesJSON, err := ctx.GetStub().GetPrivateData(collectionSender, ticketID)
	// if err != nil {
	// 	return "", fmt.Errorf("Failed to get ticket private data: %v", err)
	// }

	// err = ctx.GetStub().DelPrivateData(collectionSender, ticket.ID)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to delete Asset private details from seller: %v", err)
	// }

	// collectionBuyer := buildCollectionName(toOrgMSPID)
	// err = ctx.GetStub().PutPrivateData(collectionBuyer, ticket.ID, immutablePropertiesJSON)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to put Asset private properties for buyer: %v", err)
	// }

	// if err = s.TransferTicketPrivate(ctx, updateTicketJSON, ticket.ID, clientOrgID, toOrgMSPID); err != nil {
	// 	return "", err
	// }

	return fmt.Sprintf("Succeed! Transfer %v from %v to %v \n", ticket.ID, clientOrgID, toOrgMSPID), nil
}

func (s *SmartContract) TransferTicketPrivate(ctx contractapi.TransactionContextInterface, updateTicketJSON []byte, ticketID, clientOrgID, toOrgMSPID string) error {
	// process private data
	collectionSender := buildCollectionName(clientOrgID) //_implicit_org_Org1MSP
	privatedata, err := ctx.GetStub().GetPrivateData(collectionSender, ticketID)
	if err != nil {
		return fmt.Errorf("Failed to get ticket private data: %v", err)
	}
	if err = ctx.GetStub().DelPrivateData(collectionSender, ticketID); err != nil {
		return fmt.Errorf("Failed to delete ticket private data: %v", err)
	}

	collectionReciever := buildCollectionName(toOrgMSPID)
	if err = ctx.GetStub().PutPrivateData(collectionReciever, ticketID, privatedata); err != nil {
		return fmt.Errorf("Failed to put update ticket's private data: %v", err)
	}

	return nil
}

// 更改票据公开信息备注
func (s *SmartContract) ChangeDescription(ctx contractapi.TransactionContextInterface, ticketID, newDescription string) (string, error) {
	// check that only owner can change description
	ticketBytes, err := ctx.GetStub().GetState(ticketID)
	if err != nil {
		return "", fmt.Errorf("Failed to get ticket state: %v", err)
	}
	var ticket Ticket
	if err = json.Unmarshal(ticketBytes, &ticket); err != nil {
		return "", fmt.Errorf("Failed to unmarshal ticket: %v", err)
	}

	clientOrgID, err := getClientOrgMSPID(ctx, false)
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

func buildCollectionName(clientOrgID string) string {
	return fmt.Sprintf("_implicit_org_%s", clientOrgID)
}

func getClientOrgMSPID(ctx contractapi.TransactionContextInterface, verifyOrg bool) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("failed getting client's orgID: %v", err)
	}

	return clientOrgID, nil
}

// setAssetStateBasedEndorsement adds an endorsement policy to a asset so that only a peer from an owning org
// can update or transfer the asset.
func setAssetStateBasedEndorsement(ctx contractapi.TransactionContextInterface, assetID string, orgToEndorse string) error {
	endorsementPolicy, err := statebased.NewStateEP(nil)
	if err != nil {
		return err
	}
	err = endorsementPolicy.AddOrgs(statebased.RoleTypePeer, orgToEndorse)
	if err != nil {
		return fmt.Errorf("failed to add org to endorsement policy: %v", err)
	}
	policy, err := endorsementPolicy.Policy()
	if err != nil {
		return fmt.Errorf("failed to create endorsement policy bytes from org: %v", err)
	}
	err = ctx.GetStub().SetStateValidationParameter(assetID, policy)
	if err != nil {
		return fmt.Errorf("failed to set validation parameter on asset: %v", err)
	}

	return nil
}
