package v1

import (
	"net/http"

	"github.com/bosskrub9992/templatebe/corelib/errs"
	"github.com/bosskrub9992/templatebe/service/bff/internal/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCustomer(c echo.Context) error {
	request := model.CreateCustomerRequest{}

	if err := c.Bind(&request); err != nil {
		bindErr := errs.NewBind(err)
		return c.JSON(bindErr.Status, bindErr)
	}

	if err := c.Validate(&request); err != nil {
		validateErr := errs.NewValidate(err)
		return c.JSON(validateErr.Status, validateErr)
	}

	ctx := c.Request().Context()
	response, err := h.customerController.CreateCustomer(ctx, request)
	if err != nil {
		unknownErr := errs.NewUnknown(err)
		return c.JSON(unknownErr.Status, unknownErr)
	}

	return c.JSON(http.StatusCreated, response)
}
