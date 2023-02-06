package v1

import (
	"net/http"
	service "templatebe/src/controller"
	"templatebe/src/model"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	CustomerController *service.CustomerController
}

func NewCustomerHandler(CustomerController *service.CustomerController) *CustomerHandler {
	return &CustomerHandler{
		CustomerController: CustomerController,
	}
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	req := model.CreateCustomerRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	ctx := c.Request().Context()
	resp, err := h.CustomerController.CreateCustomer(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, resp)
}
