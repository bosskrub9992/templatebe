package main

import (
	"templatebe/lib/database"
	"templatebe/lib/loggers"
	"templatebe/lib/validators"
	v1 "templatebe/src/api/v1"
	"templatebe/src/config"
	service "templatebe/src/controller"
	"templatebe/src/repository/sqlcrepository"
	"templatebe/src/repository/sqlcrepository/sqlc"
	"templatebe/src/router"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.New()
	sqlDB, err := database.NewPostgreSQLDB(cfg)
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	db := sqlc.New(sqlDB)
	customerRepo := sqlcrepository.NewSQLCCustomerRepository(db)
	logger := loggers.NewZerolog(cfg)
	CustomerController := service.NewCustomerController(logger, customerRepo)
	customerHandler := v1.NewCustomerHandler(CustomerController)

	e := echo.New()
	router.RegisterRoute(e, customerHandler)
	e.Validator = validators.NewRequestValidator()

	if err := e.Start(cfg.Server.Port); err != nil {
		logger.Fatal().Err(err).Send()
	}
}
