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
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.username"),
		SSLmode:  viper.GetString("database.sslmode"),
	}
}
