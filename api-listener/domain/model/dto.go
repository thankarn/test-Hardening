package model

type ResponseData struct {
	TxID      string `json:"txID"`
	EventType string `json:"eventType"`
	Payload   PayloadData
}

type RequestData[T any] struct {
	TxID      string `json:"txID"`
	EventType string `json:"eventType"`
	Payload   T      `json:"payload"`
}

type PayloadData struct {
	Message string `json:"message"`
}
