package sqlcrepository

import (
	"context"
	"templatebe/pkg/domain"
	"templatebe/pkg/infrastructure/sqlcrepository/sqlc"
)

type SQLCCustomerRepository struct {
	sqlcQueries *sqlc.Queries
}

func NewSQLCCustomerRepository(sqlcQueries *sqlc.Queries) *SQLCCustomerRepository {
	return &SQLCCustomerRepository{
		sqlcQueries: sqlcQueries,
	}
}

func (s *SQLCCustomerRepository) Create(ctx context.Context, customer domain.Customer) (int64, error) {
	newCustomer, err := s.sqlcQueries.CreateCustomer(ctx, customer.Name)
	if err != nil {
		return 0, err
	}
	return newCustomer.ID, nil
}
