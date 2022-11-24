package utils

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/statebased"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func VerifyClientOrgMatchesPeerOrg(clientOrgID string) error {
	peerOrgID, err := shim.GetMSPID()
	if err != nil {
		return fmt.Errorf("failed getting peer's orgID: %v", err)
	}

	if clientOrgID != peerOrgID {
		return fmt.Errorf("client from org %s is not authorized to read or write private from an org %s peer",
			clientOrgID,
			peerOrgID,
		)
	}
	return nil
}

// func GetClientImplicitCollectionName(ctx contractapi.TransactionContextInterface) (string, error) {
// 	clientOrgID, err := GetClientOrgID(ctx, true)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to get verified OrgID: %v", err)
// 	}

// 	err = VerifyClientOrgMatchesPeerOrg(clientOrgID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return BuildCollectionName(clientOrgID), nil
// }

func SetTicketStateBasedEndorsement(ctx contractapi.TransactionContextInterface, ticketID string, orgToEndorse string) error {
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
		return fmt.Errorf("failed to create endorsement policy bytes for org: %v", err)
	}
	err = ctx.GetStub().SetStateValidationParameter(ticketID, policy)
	if err != nil {
		return fmt.Errorf("failed to set validation parameter on ticket: %v", err)
	}

	return nil
}
