package main

import (
	"database/sql"
	"fmt"
	"templatebe/pkg/api"
	v1 "templatebe/pkg/api/v1"
	"templatebe/pkg/config"
	"templatebe/pkg/infrastructure/sqlcrepository"
	"templatebe/pkg/infrastructure/sqlcrepository/sqlc"
	"templatebe/pkg/log"
	"templatebe/pkg/router"
	"templatebe/pkg/service"

	"github.com/labstack/echo/v4"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.New()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLmode,
	)
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	queries := sqlc.New(sqlDB)
	customerRepo := sqlcrepository.NewSQLCCustomerRepository(queries)
	logger := log.NewZerolog(cfg)
	customerService := service.NewCustomerService(logger, customerRepo)
	customerHandler := v1.NewCustomerHandler(customerService)

	e := echo.New()
	router.RegisterRoute(e, customerHandler)
	e.Validator = api.NewRequestValidator()

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
