package main

import (
	"templatebe/pkg/api"
	v1 "templatebe/pkg/api/v1"
	"templatebe/pkg/config"
	"templatebe/pkg/infrastructure/log"
	"templatebe/pkg/infrastructure/postgresql"
	"templatebe/pkg/infrastructure/sqlcrepository"
	"templatebe/pkg/infrastructure/sqlcrepository/sqlc"
	"templatebe/pkg/router"
	"templatebe/pkg/service"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.New()
	sqlDB, err := postgresql.NewSQLDB(cfg)
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	db := sqlc.New(sqlDB)
	customerRepo := sqlcrepository.NewSQLCCustomerRepository(db)
	logger := log.NewZerolog(cfg)
	customerService := service.NewCustomerService(logger, customerRepo)
	customerHandler := v1.NewCustomerHandler(customerService)

	e := echo.New()
	router.RegisterRoute(e, customerHandler)
	e.Validator = api.NewRequestValidator()

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
