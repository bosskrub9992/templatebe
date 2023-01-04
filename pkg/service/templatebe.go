package service

import (
	"context"
	"templatebe/pkg/domain"
	"templatebe/pkg/model"

	"github.com/rs/zerolog"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer domain.Customer) (int64, error)
}

type CustomerService struct {
	logger             *zerolog.Logger
	customerRepository CustomerRepository
}

func NewCustomerService(logger *zerolog.Logger, customerRepo CustomerRepository) *CustomerService {
	return &CustomerService{
		logger:             logger,
		customerRepository: customerRepo,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, request model.CreateCustomerRequest) (*model.CreateCustomerResponse, error) {
	s.logger.Info().Msgf("CreateCustomer request: %+v", request)

	customer := domain.Customer{
		ID:   request.ID,
		Name: request.Name,
	}

	customerID, err := s.customerRepository.Create(ctx, customer)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, err
	}
	
	return &model.CreateCustomerResponse{
		ID: customerID,
	}, nil
}
