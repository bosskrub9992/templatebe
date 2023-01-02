package service

import (
	"context"
	"templatebe/pkg/domain"
	"templatebe/pkg/model"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer domain.Customer) (int64, error)
}

type CustomerService struct {
	customerRepository CustomerRepository
}

func NewCustomerService(customerRepo CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepo,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, customer model.CreateCustomerRequest) (*model.CreateCustomerResponse, error) {
	domainCustomer := domain.Customer{
		ID:   customer.ID,
		Name: customer.Name,
	}
	customerID, err := s.customerRepository.Create(ctx, domainCustomer)
	if err != nil {
		return nil, err
	}
	return &model.CreateCustomerResponse{
		ID: customerID,
	}, nil
}
