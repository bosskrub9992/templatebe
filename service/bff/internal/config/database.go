package config

import (
	"github.com/bosskrub9992/templatebe/corelib/database"
)

func NewPostgresConfig(cfg *Config) *database.PostgresConfig {
	return &database.PostgresConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLmode:  cfg.Database.SSLMode,
	}
}
