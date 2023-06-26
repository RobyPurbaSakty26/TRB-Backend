package entity

import "time"

type TransactionAccount struct {
	Id              int       `json:"id"`
	AccountNo       string    `json:"account_no"`
	Currency        string    `json:"currency"`
	TransactionDate time.Time `json:"transaction_date"`
	TransactionTime time.Time `json:"transaction_time"`
	Remark          string    `json:"remark"`
	TellerId        int       `json:"teller_id"`
	Category        string    `json:"category"`
	Amount          string    `json:"amount"`
}
