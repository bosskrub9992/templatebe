package server

import (
	"context"
	"templatebe/lib/loggers"
	"templatebe/lib/validators"
	v1 "templatebe/src/api/v1"
	"templatebe/src/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type RESTServer struct {
	E       *echo.Echo
	config  *config.Config
	Handler *v1.Handler
}

func NewRESTServer(
	config *config.Config,
	handler *v1.Handler,
	logger *zerolog.Logger,
) *RESTServer {

	e := echo.New()
	e.Validator = validators.NewRequestValidator()
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(loggers.EchoMiddleware(logger))

	return &RESTServer{
		E:       e,
		config:  config,
		Handler: handler,
	}
}

func (r *RESTServer) Serve() error {
	return r.E.Start(r.config.Server.Port)
}

func (r *RESTServer) Shutdown(ctx context.Context) error {
	return r.E.Shutdown(ctx)
}
