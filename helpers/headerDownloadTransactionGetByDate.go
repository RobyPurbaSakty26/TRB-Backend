package helpers

import "time"

type HeaderDownloadTransactionVaByDate struct {
	Id               int       `xlsx:"id"`
	AccountNo        string    `xlsx:"account_no"`
	VirtualAccountNo string    `xlsx:"virtual_account_no"`
	Currency         string    `xlsx:"currency"`
	TransactionDate  time.Time `xlsx:"transaction_date"`
	TransactionTime  []uint8   `xlsx:"transaction_time"`
	Remark           string    `xlsx:"remark"`
	TellerId         int       `xlsx:"teller_id"`
	Category         string    `xlsx:"category"`
	Credit           string    `xlsx:"credit"`
}
