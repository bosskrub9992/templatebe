package model

import "time"

type GetHealthResponse struct {
	ServerStartTime time.Time `json:"serverStartTime"`
}
