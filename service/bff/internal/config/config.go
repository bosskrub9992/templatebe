package config

import (
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")     // local
	viper.AddConfigPath("../../config") // unit test
	viper.AddConfigPath("/app/config")  // docker
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
}

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string
	}
	Logger struct {
		GlobalMinLevel string
		JSON           bool
	}
}

func New() *Config {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
