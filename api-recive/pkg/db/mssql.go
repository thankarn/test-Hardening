package db

import (
	"fmt"
	"go-starter-api/pkg/env"

	_ "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/azuread"

	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Querier interface {
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func GetDB() (*gorm.DB, error) {
	server := env.Env().DB_HOST
	port := env.Env().DB_PORT
	database := env.Env().DB_DATABASE
	tenantID := env.Env().AAD_SP_TENANT_ID
	clientID := env.Env().AAD_SP_CLIENT_ID
	clientSecret := env.Env().AAD_SP_CLIENT_SECRET
	fedauth := env.Env().DB_FEDAUTH

	dsn := fmt.Sprintf("server=%s;user id=%s@%s;password=%s;port=%s;database=%s;fedauth=%s;", server, clientID, tenantID, clientSecret, port, database, fedauth)
	dial := sqlserver.New(sqlserver.Config{DriverName: "azuresql", DSN: dsn})
	db, _ := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		DryRun: false,
	})

	return db, nil
}
