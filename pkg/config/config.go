package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct{}

func New() *Config {
	config := new(Config)

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}
