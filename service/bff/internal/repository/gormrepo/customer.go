package gormrepo

import (
	"context"

	"github.com/bosskrub9992/templatebe/service/bff/internal/domain"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

func (repo *CustomerRepo) Create(ctx context.Context, customer domain.Customer) (int64, error) {
	tx := repo.db.Create(&customer)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return customer.ID, nil
}
