package customer

import (
	"encoding/json"
	"fmt"
	"go-starter-api/domain/model"
	"go-starter-api/pkg/env"
	"go-starter-api/pkg/utils"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/service_bus/topic"

	"github.com/gofiber/fiber/v2"
)

type reciveHandler struct {
	reciveService ReciveService
	validate      *utils.Validator
	ai            appinsightsx.Appinsightsx
}

type ReciveHandler interface {
	GetCustomerAll(c *fiber.Ctx) error
	InsertRecive(c *fiber.Ctx) error
}

func NewReciveHandler(reciveService ReciveService, validate *utils.Validator, ai appinsightsx.Appinsightsx) ReciveHandler {
	return reciveHandler{reciveService, validate, ai}
}

// GetCustomerAll godoc
// @Summary Get all customers
// @Security BearerAuth
// @Description Get all customers
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {object} model.CustomerResponse
// @Failure 400 {object} utils.ErrorDTO
// @Router /recive [get]
func (cus reciveHandler) GetCustomerAll(c *fiber.Ctx) error {
	// data, err := cus.reciveService.GetCustomerAll()
	// if err != nil {
	// 	cus.ai.Error(appinsightsx.LoggerRequest{
	// 		Error: err.Error(),
	// 	})
	// 	return c.Status(fiber.StatusBadRequest).JSON(err)
	// }

	// res := model.CustomerResponse{
	// 	Data:    data,
	// 	Message: "success",
	// 	Errors:  nil,
	// }

	// cus.ai.Info(appinsightsx.LoggerRequest{
	// 	Process: fmt.Sprintf("GetCustomerAll: %v", res),
	// })
	return nil
}

// InsertRecive handles the insertion of a new customer.
// @Summary Insert a new customer
// @Security BearerAuth
// @Description Insert a new customer into the system
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body model.ReciveInsertRequest true "Customer Insert Request"
// @Success 200 {object} model.CustomerResponse
// @Failure 400 {object} utils.ErrorDTO
// @Router /recive [post]
func (cus reciveHandler) InsertRecive(c *fiber.Ctx) error {
	req := model.ReciveInsertRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		cus.ai.Error(appinsightsx.LoggerRequest{
			Error: err.Error(),
		})
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
	}

	// v := cus.validate.Validate(req)
	// if len(v) > 0 {
	// 	jsonString, err := json.Marshal(v)
	// 	if err != nil {
	// 		cus.ai.Error(appinsightsx.LoggerRequest{
	// 			Error: err.Error(),
	// 		})
	// 		return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
	// 	}
	// 	cus.ai.Error(appinsightsx.LoggerRequest{
	// 		Error: fmt.Errorf("validate len 0"),
	// 	})
	// 	return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", errors.New(string(jsonString))))
	// }

	// data, err := cus.reciveService.InsertRecive(req)
	// if err != nil {
	// 	cus.ai.Error(appinsightsx.LoggerRequest{
	// 		Error: err.Error(),
	// 	})
	// 	return c.Status(fiber.StatusBadRequest).JSON(err)
	// }

	res := model.CustomerResponse[map[string]interface{}]{
		TxID:      "123456",
		EventType: "recive",
		Payload:   map[string]interface{}{"test": "kazano"},
	}

	props := map[string]interface{}{
		"eventType": "recive",
		"txID":      "",
	}
	t := topic.New()

	// tpc := utils.IIf(eventType , env.Env().SB_TOPIC, env.Env().SB_ERROR_TOPIC)
	resBytes, err := json.Marshal(res)
	if err != nil {
		cus.ai.Error(appinsightsx.LoggerRequest{
			Error: err.Error(),
		})
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewErrorDTO(fiber.StatusInternalServerError, "", err))
	}
	t.SetTopic(env.Env().SB_TOPIC).Emit(resBytes, props)
	// t.SetTopic(env.Env().SB_TOPIC).Emit([]byte(res), props)

	cus.ai.Info(appinsightsx.LoggerRequest{
		Process: fmt.Sprintf("InsertRecive: %v", res),
	})
	return c.JSON(res)
}
