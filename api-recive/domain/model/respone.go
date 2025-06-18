package model

type Response[T any] struct {
	Errors  any    `json:"errors"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
