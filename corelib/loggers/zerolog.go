package loggers

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type GlobalMinLevel string

const (
	GlobalMinLevelInfo  GlobalMinLevel = "info"
	GlobalMinLevelWarn  GlobalMinLevel = "warn"
	GlobalMinLevelError GlobalMinLevel = "error"
	GlobalMinLevelFatal GlobalMinLevel = "fatal"
	GlobalMinLevelPanic GlobalMinLevel = "panic"
	GlobalMinLevelNo    GlobalMinLevel = "no"
	GlobalMinLevelDis   GlobalMinLevel = "dis"
	GlobalMinLevelTrace GlobalMinLevel = "trace"
)

type ZerologConfig struct {
	JSON           bool
	GlobalMinLevel GlobalMinLevel
}

func NewZerolog(cfg *ZerologConfig) *zerolog.Logger {
	logger := zerolog.New(os.Stderr)
	if !cfg.JSON {
		logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		})
	}

	switch cfg.GlobalMinLevel {
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
