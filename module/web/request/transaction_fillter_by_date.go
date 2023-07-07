package request

type FillterTransactionByDate struct {
	AccNo     string
	StartDate string
	EndDate   string
	Page      int
	Limit     int
}
