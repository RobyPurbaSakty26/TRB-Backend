package response

type ErrorResponseLogin struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	InputFalse int    `json:"input_false"`
}
