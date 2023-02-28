package router

import "github.com/bosskrub9992/templatebe/corelib/server"

func RegisterRoute(restServer *server.RESTServer) {
	v1Group := restServer.E.Group("/api/v1")

	v1Group.GET("/health", restServer.Handler.GetHealth)
	v1Group.POST("/customers", restServer.Handler.CreateCustomer)
}
