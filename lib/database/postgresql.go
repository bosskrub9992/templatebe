package database

import (
	"database/sql"
	"fmt"
	"templatebe/src/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
)

func NewPostgreSQLDB(cfg *config.Config, logger *zerolog.Logger) (*sql.DB, func(), error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLmode,
	)
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, nil, err
	}
	cleanUp := func() {
		if err := sqlDB.Close(); err != nil {
			logger.Err(err).Send()
		}
	}
	return sqlDB, cleanUp, nil
}
