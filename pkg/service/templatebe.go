package service

import (
	"templatebe/pkg/domain"
	"templatebe/pkg/model"
)

type CustomerRepository interface {
	Create(customer domain.Customer) error
}

type CustomerService struct {
	customerRepository CustomerRepository
}

func NewCustomerService(customerRepo CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepo,
	}
}

func (c *CustomerService) CreateCustomer(customer model.Customer) error {
	domainCustomer := domain.Customer{
		ID:   customer.ID,
		Name: customer.Name,
	}
	return c.customerRepository.Create(domainCustomer)
}
