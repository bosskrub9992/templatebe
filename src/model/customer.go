package model

type CreateCustomerRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name"`
}

type CreateCustomerResponse struct {
	ID int64 `json:"id"`
}
