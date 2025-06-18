package model

type CustomerInsertRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

type CustomerResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Errors  error       `json:"errors"`
}
