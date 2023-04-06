package server

import (
	"context"

	"github.com/bosskrub9992/templatebe/corelib/loggers"
	"github.com/bosskrub9992/templatebe/corelib/validators"
	v1 "github.com/bosskrub9992/templatebe/service/bff/internal/api/v1"
	"github.com/bosskrub9992/templatebe/service/bff/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type RESTServer struct {
	e       *echo.Echo
	config  *config.RESTServerConfig
	handler *v1.Handler
}

func NewRESTServer(config *config.RESTServerConfig, logger *zerolog.Logger, handler *v1.Handler) *RESTServer {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(loggers.EchoMiddlewareZerolog(logger))

	e.Validator = validators.NewRequestValidator()

	return &RESTServer{
		e:       e,
		config:  config,
		handler: handler,
	}
}

func (r *RESTServer) Serve() error {
	return r.e.Start(r.config.Port)
}

func (r *RESTServer) Shutdown(ctx context.Context) error {
	return r.e.Shutdown(ctx)
}
