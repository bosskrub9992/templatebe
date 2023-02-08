// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"templatebe/lib/database"
	"templatebe/lib/loggers"
	"templatebe/lib/server"
	"templatebe/src/api/v1"
	"templatebe/src/config"
	"templatebe/src/controller"
	"templatebe/src/repository/sqlcrepository"
	"templatebe/src/repository/sqlcrepository/sqlc"
)

// Injectors from wire.go:

func InitializeRestServer() (*server.RESTServer, func(), error) {
	configConfig := config.New()
	logger := loggers.NewZerolog(configConfig)
	db, cleanup, err := database.NewPostgreSQLDB(configConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	queries := sqlc.New(db)
	sqlcCustomerRepository := sqlcrepository.NewSQLCCustomerRepository(queries)
	customerController := controller.NewCustomerController(logger, sqlcCustomerRepository)
	healthController := controller.NewHealthController()
	handler := v1.NewHandler(customerController, healthController)
	restServer := server.NewRESTServer(configConfig, handler, logger)
	return restServer, func() {
		cleanup()
	}, nil
}

// wire.go:

var controllerSet = wire.NewSet(controller.NewCustomerController, controller.NewHealthController)

var repositorySet = wire.NewSet(sqlcrepository.NewSQLCCustomerRepository)
