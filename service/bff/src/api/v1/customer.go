package v1

import (
	"net/http"

	"github.com/bosskrub9992/templatebe/corelib/errs"
	"github.com/bosskrub9992/templatebe/service/bff/src/model/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCustomer(c echo.Context) error {
	req := model.CreateCustomerRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errs.NewBind(err))
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errs.NewValidate(err))
	}
	ctx := c.Request().Context()
	resp, err := h.customerController.CreateCustomer(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errs.NewUnknown(err))
	}
	return c.JSON(http.StatusCreated, resp)
}
