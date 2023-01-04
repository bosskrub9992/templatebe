package log

import (
	"os"
	"templatebe/pkg/config"
	"time"

	"github.com/rs/zerolog"
)

func NewZerolog(cfg *config.Config) *zerolog.Logger {

	logger := zerolog.New(os.Stderr)
	if !cfg.Logger.JSON {
		logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			NoColor:    false,
			TimeFormat: time.RFC3339,
		})
	}

	switch cfg.Logger.GlobalMinLevel {
	case "info":
		logger = logger.Level(zerolog.InfoLevel)
	case "warn":
		logger = logger.Level(zerolog.WarnLevel)
	case "error":
		logger = logger.Level(zerolog.ErrorLevel)
	case "fatal":
		logger = logger.Level(zerolog.FatalLevel)
	case "panic":
		logger = logger.Level(zerolog.PanicLevel)
	case "no":
		logger = logger.Level(zerolog.NoLevel)
	case "dis":
		logger = logger.Level(zerolog.Disabled)
	case "trace":
		logger = logger.Level(zerolog.TraceLevel)
	default:
		logger = logger.Level(zerolog.DebugLevel)
	}

	logger = logger.With().Timestamp().Caller().Logger()

	return &logger
}
