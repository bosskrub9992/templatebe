package model

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateCustomerResponse struct {
	ID int64 `json:"id"`
}
