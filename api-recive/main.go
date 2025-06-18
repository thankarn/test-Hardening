package main

import (
	"encoding/json"
	"fmt"
	"go-starter-api/external"
	"go-starter-api/pkg/db"
	"go-starter-api/pkg/env"

	_ "go-starter-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/fiber_middleware"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_base"
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_base/get_token"
)

// @title Go Starter API
// @version 1.0
// @description This is a sample API with Bearer Auth
// @host api2-dv.banpu.co.th
// @BasePath /starter-api/v2
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token (e.g., "Bearer <token>")
func main() {
	ai := appinsightsx.NewAppinsightsx(appinsightsx.InitProperties{
		InstrumentationKey: env.Env().APP_INSIGHTS_KEY,
		RoleName:           env.Env().APP_INSIGHTS_ROLE,
		EnableZlog: 	   true,
	})
	ai.Defer()

	banpuToken := get_token.NewBanpuAcquireToken()
	banpuHttp := http_base.NewHttpClient(banpuToken)

	app := fiber.New(fiber.Config{
		// BodyLimit: 100 * 1024 * 1024,
		AppName:     env.Env().MODULE_NAME,
		Immutable:   true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Application Insights Middleware: Logs each HTTP request.
	app.Use(fiber_middleware.RequestMiddleware(ai, banpuHttp))

	db, _ := db.GetDB()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	external.PublicRoutes(app, db, ai, banpuHttp)

	ai.Info(appinsightsx.LoggerRequest{
		Process: fmt.Sprintf("Start Project | Listen Port:%v", env.Env().PORT),
	})

	err := app.Listen(fmt.Sprintf(":%v", env.Env().PORT))
	if err != nil {
		ai.Error(appinsightsx.LoggerRequest{
			Process: fmt.Sprintf("Start Project | Listen Port:%v", env.Env().PORT),
			Error:   err.Error(),
		})
	}
}
