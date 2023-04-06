package config

import (
	"github.com/bosskrub9992/templatebe/corelib/database"
	"github.com/spf13/viper"
)

func NewPostgresConfig() *database.PostgresConfig {
	return &database.PostgresConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.dbname"),
		Password: viper.GetString("database.username"),
		DBName:   viper.GetString("database.password"),
		SSLmode:  viper.GetString("database.sslmode"),
	}
}
