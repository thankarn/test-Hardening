package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ReqeustDTO struct {
	TxID  string   `json:"txID"`
	User  string   `json:"user"`
	Roles []string `json:"roles"`
}
type ReqeustWithDataDTO struct {
	TxID  string   `json:"txID"`
	User  string   `json:"user"`
	Roles []string `json:"roles"`
	Data  any      `json:"data"`
}
type ErrorDTO struct {
	Code    int                    `json:"code"`
	Message any                    `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}
type ResponseDTO struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
type ResponseOffsetDTO struct {
	Data    any    `json:"data"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
	Total   int    `json:"total"`
	Message string `json:"message"`
}

// add switch other variant
func NewErrorDTO(code int, message string, err error) ErrorDTO {
	e := ErrorDTO{}
	e.Code = code
	e.Message = message
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	default:
		e.Errors["code"] = v.Error()
	}
	return e
}

func NewErrorDTOFiber(err *fiber.Error) ErrorDTO {
	return ErrorDTO{
		Code:    err.Code,
		Message: err.Message,
		Errors: map[string]interface{}{
			fmt.Sprintf("%d", err.Code): err.Error(),
		},
	}
}

func NewResponseDTO(data interface{}, msg string) ResponseDTO {
	return ResponseDTO{
		Data:    data,
		Message: msg,
	}
}

func NewResponseOffsetDTO(data interface{}, msg string, total, limit, offset int) ResponseOffsetDTO {
	return ResponseOffsetDTO{
		Data:    data,
		Message: msg,
		Limit:   limit,
		Offset:  offset,
		Total:   total,
	}
}