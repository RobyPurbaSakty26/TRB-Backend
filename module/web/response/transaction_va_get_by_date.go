package response

type ResponseTransactionVitualAccount struct {
	Status    string                                `json:"status" `
	Limit     int                                   `json:"limit"`
	Total     int                                   `json:"total"`
	TotalPage int                                   `json:"total_page"`
	Page      int                                   `json:"page"`
	Data      []ResponseTransactionItemsVaGetByDate `json:"data"`
}

type ResponseTransactionItemsVaGetByDate struct {
	ID                          uint   `json:"id"`
	NomorRekeningGiro           string `json:"nomor_virtual_giro"`
	NomorRekeningVirtualAccount string `json:"nomor_virtual_account"`
	Currency                    string `json:"currency"`
	TanggalTransaksi            string `json:"tanggal_transaksi"`
	Jam                         string `json:"jam"`
	Remark                      string `json:"remark"`
	Teller                      int    `json:"teller"`
	Category                    string `json:"category"`
	Credit                      string `json:"credit"`
}
