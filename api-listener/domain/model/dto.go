package model

type ResponseData struct {
	TxID      string `json:"txID"`
	Source    string `json:"source"`
	EventType string `json:"eventType"`
	Payload   PayloadData
}

type RequestData struct {
	TxID      string      `json:"txID"`
	Source    string      `json:"source"`
	EventType string      `json:"eventType"`
	Payload   PayloadData `json:"payload"`
}

type PayloadData struct {
	Message string `json:"message"`
	Email   string `json:"email"`
}
