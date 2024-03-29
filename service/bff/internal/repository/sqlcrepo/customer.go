package sqlcrepo

import (
	"context"

	"github.com/bosskrub9992/templatebe/service/bff/internal/domain"
	"github.com/bosskrub9992/templatebe/service/bff/internal/repository/sqlcrepo/sqlc"
)

type SQLCCustomerRepository struct {
	db *sqlc.Queries
}

func NewSQLCCustomerRepository(db *sqlc.Queries) *SQLCCustomerRepository {
	return &SQLCCustomerRepository{
		db: db,
	}
}

func (r *SQLCCustomerRepository) Create(ctx context.Context, customer domain.Customer) (int64, error) {
	newCustomer, err := r.db.CreateCustomer(ctx, customer.Name)
	if err != nil {
		return 0, err
	}
	return newCustomer.ID, nil
}
