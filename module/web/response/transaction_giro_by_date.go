package response

type ResponseTransactionGiro struct {
	Status    string                                  `json:"status" `
	Limit     int                                     `json:"limit"`
	Total     int                                     `json:"total"`
	TotalPage int                                     `json:"total_page"`
	Page      int                                     `json:"page"`
	Data      []ResponseTransactionItemsGiroGetByDate `json:"data"`
}

type ResponseTransactionItemsGiroGetByDate struct {
	ID                uint   `json:"id"`
	NomorRekeningGiro string `json:"nomor_virtual_giro"`
	Currency          string `json:"currency"`
	TanggalTransaksi  string `json:"tanggal_transaksi"`
	Jam               string `json:"jam"`
	Remark            string `json:"remark"`
	Teller            int    `json:"teller"`
	Category          string `json:"category"`
	Amount            string `json:"amount"`
}
