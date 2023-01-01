package main

import (
	"database/sql"
	v1 "templatebe/pkg/api/v1"
	"templatebe/pkg/infrastructure/sqlcrepository"
	"templatebe/pkg/infrastructure/sqlcrepository/sqlc"
	"templatebe/pkg/router"
	"templatebe/pkg/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	sqlDB, err := sql.Open("pgx", "localhost")
	if err != nil {
		panic(err)
	}
	queries := sqlc.New(sqlDB)
	customerRepo := sqlcrepository.NewSQLCCustomerRepository(queries)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := v1.NewCustomerHandler(customerService)

	router.RegisterRoute(e, customerHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
