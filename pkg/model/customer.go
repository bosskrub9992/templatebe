package model

type CreateCustomerRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateCustomerResponse struct {
	ID int64 `json:"id"`
}
