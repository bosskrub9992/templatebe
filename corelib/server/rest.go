package server

import (
	"context"

	"github.com/bosskrub9992/templatebe/corelib/loggers"
	"github.com/bosskrub9992/templatebe/corelib/validators"
	v1 "github.com/bosskrub9992/templatebe/service/bff/src/api/v1"
	"github.com/bosskrub9992/templatebe/service/bff/src/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type RESTServer struct {
	E       *echo.Echo
	Handler *v1.Handler
	config  *config.Server
}

func NewRESTServer(config *config.Server, handler *v1.Handler, logger *zerolog.Logger) *RESTServer {

	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(loggers.EchoMiddlewareZerolog(logger))

	e.Validator = validators.NewRequestValidator()

	return &RESTServer{
		E:       e,
		config:  config,
		Handler: handler,
	}
}

func (r *RESTServer) Serve() error {
	return r.E.Start(r.config.Port)
}

func (r *RESTServer) Shutdown(ctx context.Context) error {
	return r.E.Shutdown(ctx)
}
