package sqlcrepository

import (
	"database/sql"
	"templatebe/pkg/domain"
)

type SQLCCustomerRepository struct {
	sqlDB *sql.DB
}

func NewSQLCCustomerRepository(sqlDB *sql.DB) *SQLCCustomerRepository {
	return &SQLCCustomerRepository{
		sqlDB: sqlDB,
	}
}

func (s *SQLCCustomerRepository) Create(customer domain.Customer) error {
	return nil
}
