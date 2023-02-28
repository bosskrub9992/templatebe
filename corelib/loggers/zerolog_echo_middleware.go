package loggers

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func EchoMiddlewareZerolog(logger *zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			req := c.Request()
			rid := req.Header.Get(echo.HeaderXRequestID)
			logger.Info().
				Str("rid", rid).
				Str("uri", req.RequestURI).
				Str("host", req.Host).
				Str("method", req.Method).
				Str("query", req.URL.RawQuery).
				Any("body", req.Body).
				Str("ip", c.RealIP()).
				Msgf("rid:%s request", rid)

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			res := c.Response()

			var (
				loggerWithLevel *zerolog.Event
				msg             string
			)

			switch {
			case res.Status >= 500:
				loggerWithLevel, msg = logger.Error(), "response server error"
			case res.Status >= 400:
				loggerWithLevel, msg = logger.Warn(), "response client error"
			case res.Status >= 300:
				loggerWithLevel, msg = logger.Info(), "response redirection"
			default:
				loggerWithLevel, msg = logger.Info(), "response success"
			}

			loggerWithLevel.
				Str("rid", rid).
				Int("code", res.Status).
				Dur("latency", time.Since(start)).
				AnErr("err", err).
				Msgf("rid:%s %s", rid, msg)

			return nil
		}
	}
}
