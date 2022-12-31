package v1

import (
	"net/http"
	"templatebe/pkg/model"
	"templatebe/pkg/service"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(customerService *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var customer model.Customer
	if err := c.Bind(&customer); err != nil {
		return err
	}
	if err := h.customerService.CreateCustomer(customer); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, nil)
}
