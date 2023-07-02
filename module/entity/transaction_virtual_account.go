package entity

import "time"

type TransactionVirtualAccount struct {
	Id               int       `json:"id"`
	AccountNo        string    `json:"account_no"`
	VirtualAccountNo string    `json:"virtual_account_no"`
	Currency         string    `json:"currency"`
	TransactionDate  time.Time `json:"transaction_date"`
	TransactionTime  []uint8   `json:"transaction_time"`
	Remark           string    `json:"remark"`
	TellerId         int       `json:"teller_id"`
	Category         string    `json:"category"`
	Credit           string    `json:"credit"`
}

func (TransactionVirtualAccount) TableName() string {
	return "transaction_virtual_account"
}
