package config

import (
	"github.com/spf13/viper"
)

type RESTServer struct {
	Port string
}

func NewRESTServerConfig() *RESTServer {
	return &RESTServer{
		Port: viper.GetString("server.port"),
	}
}
