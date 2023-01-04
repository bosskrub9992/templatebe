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
	var req model.CreateCustomerRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	resp, err := h.customerService.CreateCustomer(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resp)
}
