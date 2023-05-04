package server

import (
	"context"
	"fmt"

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
	config  *config.Config
	handler *v1.Handler
}

func NewRESTServer(config *config.Config, logger *zerolog.Logger, handler *v1.Handler) *RESTServer {
	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.CORS(),
		middleware.RequestID(),
		loggers.EchoMiddlewareZerolog(logger),
	)

	e.Validator = validators.NewRequestValidator()

	return &RESTServer{
		e:       e,
		config:  config,
		handler: handler,
	}
}

func (r *RESTServer) Serve() error {
	address := fmt.Sprintf(":%s", r.config.Server.Port)
	return r.e.Start(address)
}

func (r *RESTServer) Shutdown(ctx context.Context) error {
	return r.e.Shutdown(ctx)
}
