package utils

import (
	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/guard"
	"github.com/gofiber/fiber/v2"
)

func CheckRolesStarterHandler(c *fiber.Ctx) error {
	return guard.CheckRols(c, []string{"BTH-STARTER-USER"})
}