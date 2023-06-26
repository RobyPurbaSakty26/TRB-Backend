package entity

import "time"

type MasterAccount struct {
	AccountNo                     string    `json:"account_no"`
	Currency                      string    `json:"currency"`
	LastUpdate                    time.Time `json:"last_update"`
	AccountBalancePosition        int       `json:"account_balance_position"`
	TotalVirtualAccount           int       `json:"total_virtual_account"`
	VirtualAccountBalancePosition int       `json:"virtual_account_balance_position"`
}

func (MasterAccount) TableName() string {
	return "master_account"
}
