package controller

import (
	"context"
	"templatebe/src/model"
	"time"
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
