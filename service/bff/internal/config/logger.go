package config

import (
	"github.com/spf13/viper"

	"github.com/bosskrub9992/templatebe/corelib/loggers"
)

func NewLoggerConfig() *loggers.ZerologConfig {
	return &loggers.ZerologConfig{
		JSON:           viper.GetBool("logger.json"),
		GlobalMinLevel: loggers.GlobalMinLevel(viper.GetString("logger.globalMinLevel")),
	}
}
