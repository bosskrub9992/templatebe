// go:build wireinject
// +build wireinject

package main

import (
	"github.com/bosskrub9992/templatebe/corelib/database"
	"github.com/bosskrub9992/templatebe/corelib/loggers"
	v1 "github.com/bosskrub9992/templatebe/service/bff/internal/api/v1"
	"github.com/bosskrub9992/templatebe/service/bff/internal/config"
	"github.com/bosskrub9992/templatebe/service/bff/internal/controller"
	"github.com/bosskrub9992/templatebe/service/bff/internal/repository/gormrepo"
	"github.com/bosskrub9992/templatebe/service/bff/internal/server"

	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewCustomerController,
	controller.NewHealthController,
)

var repositorySet = wire.NewSet(
	gormrepo.NewCustomerRepo,
)

func InitializeRestServer() (*server.RESTServer, func(), error) {
	wire.Build(
		server.NewRESTServer,
		v1.NewHandler,
		controllerSet,
		repositorySet,
		database.NewGormDBPostgres,
		database.NewPostgres,
		config.NewPostgresConfig,
		loggers.NewZerolog,
		config.NewLoggerConfig,
		config.New,

		wire.Bind(new(controller.CustomerRepository), new(*gormrepo.CustomerRepo)),
	)
	return nil, nil, nil
}
