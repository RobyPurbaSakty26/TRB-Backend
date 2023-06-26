package entity

import "time"

type MasterAccount struct {
	AccountNo                     string    `json:"account_no"`
	Currency                      string    `json:"currency"`
	LastUpdate                    time.Time `json:"last_update"`
	AccountBalancePosition        string    `json:"account_balance_position"`
	TotalVirtualAccount           int       `json:"total_virtual_account"`
	VirtualAccountBalancePosition string    `json:"virtual_account_balance_position"`
}
