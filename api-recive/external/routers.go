package external

import (
	"fmt"
	"go-starter-api/internal/customer"
	"go-starter-api/internal/userprofile"
	"go-starter-api/pkg/env"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_base"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// HealtCheck handles the health check endpoint.
// @Summary Health Check
// @Security BearerAuth
// @Description Returns the health status of the application along with some environment details.
// @Tags Common
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Router /health-check [get]
func HealtCheck(c *fiber.Ctx) error {
	return c.JSON(map[string]any{
		"ProductCode": env.Env().PRODUCT_CODE,
		"ModuleName":  env.Env().MODULE_NAME,
		"Build":       env.Env().BUILD,
		"Release":     env.Env().RELEASE,
		"Port":        env.Env().PORT,
		"PathPrefix":  env.Env().PATH_PREFIX,
	})
}

func PublicRoutes(app *fiber.App, db *gorm.DB, ai appinsightsx.Appinsightsx, http_base http_base.HttpBase) {
	// app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: fmt.Sprintf("%s/swagger/doc.json", env.Env().PATH_PREFIX),
	}))

	routerV1 := app.Group("/")
	routerV1.Get("/health-check", HealtCheck)

	customer.Router(routerV1, db, ai)
	userprofile.Router(routerV1, db, ai, http_base)
}
