package controller

import (
	"context"
	"time"

	"github.com/bosskrub9992/templatebe/service/bff/internal/model"
)

type HealthController struct {
	ServerStartTime time.Time
}

func NewHealthController() *HealthController {
	return &HealthController{
		ServerStartTime: time.Now(),
	}
}

func (con *HealthController) GetHealth(ctx context.Context) *model.GetHealthResponse {
	return &model.GetHealthResponse{
		ServerStartTime: con.ServerStartTime,
	}
}
