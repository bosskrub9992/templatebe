package router

import (
	v1 "templatebe/src/api/v1"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo, customerHandler *v1.CustomerHandler) {
	v1Group := e.Group("/v1")

	v1Group.POST("/create-customer", customerHandler.CreateCustomer)
}
