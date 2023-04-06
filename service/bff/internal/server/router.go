package server

func (r *RESTServer) RegisterRoute() {
	apiGroup := r.e.Group("/api")

	v1Group := apiGroup.Group("/v1")
	v1Group.GET("/health", r.handler.GetHealth)
	v1Group.POST("/customers", r.handler.CreateCustomer)
}
