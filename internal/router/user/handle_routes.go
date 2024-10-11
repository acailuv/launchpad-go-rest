package user

func (r router) HandleRoutes() {
	r.echo.GET("/v1/users", r.controller.Find)
	r.echo.GET("/v1/users/:id", r.controller.FindByID)
	r.echo.POST("/v1/users", r.controller.Create)
	r.echo.PUT("/v1/users/:id", r.controller.UpdateByID)
	r.echo.DELETE("/v1/users/:id", r.controller.DeleteByID)

	// Verify Token Middleware Usage Example
	// r.echo.DELETE("/v1/users/:id", r.controller.DeleteByID)
}
