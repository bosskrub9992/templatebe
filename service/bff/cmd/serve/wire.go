//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/bosskrub9992/templatebe/corelib/database"
	"github.com/bosskrub9992/templatebe/corelib/loggers"
	v1 "github.com/bosskrub9992/templatebe/service/bff/src/api/v1"
	"github.com/bosskrub9992/templatebe/service/bff/src/config"
	"github.com/bosskrub9992/templatebe/service/bff/src/controller"
	"github.com/bosskrub9992/templatebe/service/bff/src/repository/sqlcrepository"
	"github.com/bosskrub9992/templatebe/service/bff/src/repository/sqlcrepository/sqlc"
	"github.com/bosskrub9992/templatebe/service/bff/src/server"

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
		config.NewRESTServerConfig,
		v1.NewHandler,
		controllerSet,
		repositorySet,
		sqlc.New,
		database.NewPostgres,
		config.NewDBConfig,
		loggers.NewZerolog,
		config.NewLoggerConfig,

		wire.Bind(new(controller.CustomerRepository), new(*sqlcrepository.SQLCCustomerRepository)),
		wire.Bind(new(sqlc.DBTX), new(*sql.DB)),
	)
	return nil, nil, nil
}
