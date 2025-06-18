package model

type UserprofileResponse struct {
	Error   []any  `json:"errors"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
