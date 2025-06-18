package userprofile

import (
	"go-starter-api/api"
	"go-starter-api/pkg/utils"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_base"
	"gorm.io/gorm"
)

func Router(r fiber.Router, db *gorm.DB, ai appinsightsx.Appinsightsx, http_base http_base.HttpBase) {
	validate := utils.NewValidator()

	userprofileApi := api.NewUserprofileApi(http_base, ai)
	service := NewCustomerService(userprofileApi)
	handler := NewUserprofileHandler(service, validate, ai)

	groupRoute := r.Group("/user-profile")
	groupRoute.Get("/email/:email", handler.GetEmail)
}
