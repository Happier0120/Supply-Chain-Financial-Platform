package model

import "time"

type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名称
	Balance   float64 `json:"balance"`   //余额
}

type Ticket struct {
	Type          string `json:"objectType"`
	ID            string `json:"ticketID"`    // 票据ID
	OwnerOrgMSPID string `json:"ownerOrg"`    // 票据所有者ID
	ChannelID     string `json:"channelID"`   // 通道ID
	TxID          string `json:"txID"`        // 交易ID
	TxTimestamp   string `json:"txTimestamp"` // 交易时间戳
	CreateTime    string `json:"createTime"`  // 创建日期
	DueDate       string `json:"duedate"`     // 到期日期
	Guarantor     string `json:"guarantor"`   // 担保人
	Description   string `json:"description"` // 票据备注
}

type TicketPrivate struct {
	Value     string `json:"value"`
	FromOrder string `json:"fromOrder"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Record    Ticket
	TxId      string    `json:"txId"`
	Timestamp time.Time `json:"timestamp"`
}

const (
	AccountKey = "A"
	TicketKey  = "T"
	TestKey    = "Test"
)
