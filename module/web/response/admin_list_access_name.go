package response

type ResponseAccessName struct {
	Status string           `json:"status"`
	Data   []ItemAccessName `json:"data"`
}

type ItemAccessName struct {
	Name string `json:"name"`
}
