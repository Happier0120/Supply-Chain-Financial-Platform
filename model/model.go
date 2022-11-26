package model

type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名称
	Balance   float64 `json:"balance"`   //余额
}

type Ticket struct {
	Type          string `json:"objectType"`
	ID            string `json:"ticketID"`    //票据ID
	OwnerOrgMSPID string `json:"ownerOrg"`    //票据所有者ID
	Description   string `json:"description"` //票据备注
}

type TicketPrivate struct {
	Price     string `json:"price"`
	FromOrder string `json:"fromOrder"`
}

const (
	AccountKey = "A"
	TicketKey  = "T"
	TestKey    = "Test"
)
