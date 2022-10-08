package dto

type ResJSON struct {
	Error       bool        `json:"error"`
	Code        int         `json:"code"`
	Transaction string      `json:"transaction_id"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}
