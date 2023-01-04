package postgresql

import (
	"database/sql"
	"fmt"
	"templatebe/pkg/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewSQLDB(cfg *config.Config) (*sql.DB, error) {
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
		return nil, err
	}
	return sqlDB, nil
}
