package config

import (
	"github.com/bosskrub9992/templatebe/corelib/loggers"
)

func NewLoggerConfig(cfg *Config) *loggers.ZerologConfig {
	return &loggers.ZerologConfig{
		JSON:           cfg.Logger.JSON,
		GlobalMinLevel: loggers.GlobalMinLevel(cfg.Logger.GlobalMinLevel),
	}
}
