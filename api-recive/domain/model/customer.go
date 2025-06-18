package model

type CustomerInsertRequest struct {
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

type CustomerResponse[T any] struct {
	TxID      string `json:"txID"`
	EventType string `json:"eventType"`
	Payload   T      `json:"payload"`
}
