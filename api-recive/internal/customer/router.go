package customer

import (
	"go-starter-api/pkg/utils"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/guard"
)

func Router(r fiber.Router, db *gorm.DB, ai appinsightsx.Appinsightsx) {
	validate := utils.NewValidator()

	repository := NewCustomerRepository(db)
	service := NewCustomerService(repository)
	handler := NewCustomerHandler(service, validate, ai)

	groupRoute := r.Group("/customer", guard.ADVerifyToken())
	groupRoute.Get("", utils.CheckRolesStarterHandler, handler.GetCustomerAll)
	groupRoute.Post("", handler.InsertCustomer)
}
