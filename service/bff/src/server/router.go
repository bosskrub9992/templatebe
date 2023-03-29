package server

func (r *RESTServer) RegisterRoute() {
	v1Group := r.e.Group("/api/v1")

	v1Group.GET("/health", r.handler.GetHealth)
	v1Group.POST("/customers", r.handler.CreateCustomer)
}
