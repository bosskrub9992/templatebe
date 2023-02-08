package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetHealth(c echo.Context) error {
	ctx := c.Request().Context()
	response := h.healthController.GetHealth(ctx)
	return c.JSON(http.StatusOK, response)
}
