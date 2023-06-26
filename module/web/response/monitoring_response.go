package response

type MonitoringResponse struct {
	Status string           `json:"status"`
	Data   []ItemMonitoring `json:"data"`
}

type ItemMonitoring struct {
	NoRekeningGiro  string `json:"no_rekening_giro"`
	Currency        string `json:"currency"`
	Tanggal         string `json:"tanggal"`
	PosisiSaldoGiro int    `json:"posisi_saldo_giro"`
	JumlahNoVA      int64  `json:"jumlah_no_va"`
	PosisiSaldoVA   int    `json:"posisi_saldo_va"`
	Selisih         int    `json:"selisih"`
}
