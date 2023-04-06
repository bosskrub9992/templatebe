package config

import (
	"github.com/spf13/viper"
)

type RESTServerConfig struct {
	Port string
}

func NewRESTServerConfig() *RESTServerConfig {
	return &RESTServerConfig{
		Port: viper.GetString("server.port"),
	}
}
