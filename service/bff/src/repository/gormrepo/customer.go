package gormrepo

import (
	"context"

	"github.com/bosskrub9992/templatebe/service/bff/src/domain"
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

func (repo *CustomerRepo) Create(ctx context.Context, customer domain.Customer) error {
	return repo.db.Create(&customer).Error
}
