package config

import (
	"github.com/spf13/viper"
)

type Server struct {
	Port string
}

func NewServerConfig() *Server {
	return &Server{
		Port: viper.GetString("server.port"),
	}
}
