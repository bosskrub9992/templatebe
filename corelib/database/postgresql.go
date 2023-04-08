package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
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

func NewPostgres(cfg *PostgresConfig) (*sql.DB, error) {
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
		return nil, err
	}
	return sqlDB, nil
}

func NewGormDBPostgres(sqlDB *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
