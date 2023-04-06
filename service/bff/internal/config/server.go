package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type RESTServerConfig struct {
	Port string
}

func NewRESTServerConfig() *RESTServerConfig {
	return &RESTServerConfig{
		Port: fmt.Sprintf(":%d", viper.GetInt("server.port")),
	}
}
