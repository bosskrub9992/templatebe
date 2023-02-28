package v1

import "github.com/bosskrub9992/templatebe/service/bff/src/controller"

type Handler struct {
	customerController *controller.CustomerController
	healthController   *controller.HealthController
}

func NewHandler(
	customerController *controller.CustomerController,
	healthController *controller.HealthController,
) *Handler {
	return &Handler{
		customerController: customerController,
		healthController:   healthController,
	}
}
