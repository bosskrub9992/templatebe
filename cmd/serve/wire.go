// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"templatebe/lib/database"
	"templatebe/lib/loggers"
	"templatebe/lib/server"
	v1 "templatebe/src/api/v1"
	"templatebe/src/config"
	"templatebe/src/controller"
	"templatebe/src/repository/sqlcrepository"
	"templatebe/src/repository/sqlcrepository/sqlc"

	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewCustomerController,
	controller.NewHealthController,
)

var repositorySet = wire.NewSet(
	sqlcrepository.NewSQLCCustomerRepository,
)

func InitializeRestServer() (*server.RESTServer, func(), error) {
	wire.Build(
		server.NewRESTServer,
		v1.NewHandler,
		controllerSet,
		repositorySet,
		sqlc.New,
		database.NewPostgreSQLDB,
		config.New,
		loggers.NewZerolog,

		wire.Bind(new(controller.CustomerRepository), new(*sqlcrepository.SQLCCustomerRepository)),
		wire.Bind(new(sqlc.DBTX), new(*sql.DB)),
	)
	return nil, nil, nil
}
