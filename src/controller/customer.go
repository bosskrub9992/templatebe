package service

import (
	"context"
	"templatebe/src/domain"
	"templatebe/src/model"

	"github.com/rs/zerolog"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer domain.Customer) (int64, error)
}

type CustomerController struct {
	logger             *zerolog.Logger
	customerRepository CustomerRepository
}

func NewCustomerController(logger *zerolog.Logger, customerRepo CustomerRepository) *CustomerController {
	return &CustomerController{
		logger:             logger,
		customerRepository: customerRepo,
	}
}

func (s *CustomerController) CreateCustomer(ctx context.Context, req model.CreateCustomerRequest) (*model.CreateCustomerResponse, error) {
	s.logger.Info().Interface("req", req).Msg("CreateCustomer")

	customer := domain.Customer{
		ID:   req.ID,
		Name: req.Name,
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
