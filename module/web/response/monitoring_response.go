package response

import "time"

type MonitoringResponse struct {
	Status string           `json:"status"`
	Data   []ItemMonitoring `json:"data"`
}

type ItemMonitoring struct {
	NoRekeningGiro  string    `json:"no_rekening_giro"`
	Currency        string    `json:"currency"`
	Tanggal         time.Time `json:"tanggal"`
	PosisiSaldoGiro string    `json:"posisi_saldo_giro"`
	JumlahNoVA      int       `json:"jumlah_no_va"`
	PosisiSaldoVA   string    `json:"posisi_saldo_va"`
	Selisih         string    `json:"selisih"`
}
