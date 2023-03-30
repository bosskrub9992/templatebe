package v1

import (
	"net/http"

	"github.com/bosskrub9992/templatebe/corelib/errs"
	"github.com/bosskrub9992/templatebe/service/bff/src/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCustomer(c echo.Context) error {
	request := model.CreateCustomerRequest{}

	if err := c.Bind(&request); err != nil {
		errorResp := errs.NewBind(err)
		return c.JSON(errorResp.Status, errorResp)
	}

	if err := c.Validate(&request); err != nil {
		errorResp := errs.NewValidate(err)
		return c.JSON(errorResp.Status, errorResp)
	}

	ctx := c.Request().Context()
	response, err := h.customerController.CreateCustomer(ctx, request)
	if err != nil {
		errorResp := errs.NewUnknown(err)
		return c.JSON(errorResp.Status, errorResp)
	}

	return c.JSON(http.StatusCreated, response)
}
