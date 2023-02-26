package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Port     string
		Host     string
		DBName   string
		Username string
		Password string
		SSLmode  string
	}
	Logger struct {
		GlobalMinLevel string
		JSON           bool
	}
}

func New() *Config {
	config := new(Config)

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")     // local
	viper.AddConfigPath("../../config") // testing
	viper.AddConfigPath("/app/config")  // docker
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
