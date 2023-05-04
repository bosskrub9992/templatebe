package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
	SSLmode  string
}

func NewPostgres(cfg *PostgresConfig, logger *zerolog.Logger) (*sql.DB, func(), error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.SSLmode,
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Err(err).Send()
		return nil, nil, err
	}

	cleanUp := func() {
		if err := sqlDB.Close(); err != nil {
			logger.Err(err).Send()
		}
	}

	return sqlDB, cleanUp, nil
}

func NewGormDBPostgres(sqlDB *sql.DB, logger *zerolog.Logger) (*gorm.DB, error) {
	dialector := postgres.New(postgres.Config{
		Conn: sqlDB,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logger.Err(err).Send()
		return nil, err
	}
	
	return gormDB, nil
}
