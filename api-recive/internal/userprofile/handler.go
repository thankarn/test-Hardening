package userprofile

import (
	"encoding/json"
	"fmt"

	"go-starter-api/pkg/utils"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"

	"github.com/gofiber/fiber/v2"
)

type userprofileHandler struct {
	userprofileService UserprofileService
	validate           *utils.Validator
	ai                 appinsightsx.Appinsightsx
}

type UserprofileHandler interface {
	GetEmail(c *fiber.Ctx) error
}

func NewUserprofileHandler(userprofileService UserprofileService, validate *utils.Validator, ai appinsightsx.Appinsightsx) UserprofileHandler {
	return userprofileHandler{userprofileService, validate, ai}
}

// GetEmail godoc
// @Summary Get user email
// @Security BearerAuth
// @Description Get user email by email parameter
// @Tags Userprofile
// @Accept json
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} model.UserprofileResponse
// @Failure 400 {object} utils.ErrorDTO
// @Router /user-profile/email/{email} [get]
func (u userprofileHandler) GetEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	res, err := u.userprofileService.GetEmail(email)
	if err != nil {
		u.ai.Error(appinsightsx.LoggerRequest{
			Error: err.Error(),
		})
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
	}

	var data []byte
	data, err = json.Marshal(res.Data)
	if err != nil {
		u.ai.Error(appinsightsx.LoggerRequest{
			Error: err.Error(),
		})
		return c.Status(fiber.StatusInternalServerError).JSON(utils.NewErrorDTO(fiber.StatusInternalServerError, "", err))
	}

	u.ai.Info(appinsightsx.LoggerRequest{
		Process: fmt.Sprintf("GetEmail: %s", string(data)),
	})
	return c.JSON(res)
}
