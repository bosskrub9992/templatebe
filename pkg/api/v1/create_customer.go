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
	var (
		req model.CreateCustomerRequest
		ctx = c.Request().Context()
	)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	resp, err := h.customerService.CreateCustomer(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, resp)
}
