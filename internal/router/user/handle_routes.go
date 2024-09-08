package user

func (r router) HandleRoutes() {
	r.echo.GET("/v1/users", r.controller.FindByID, r.middleware.VerifyToken())
	r.echo.GET("/v1/users/:id", r.controller.FindByID, r.middleware.VerifyToken())
	r.echo.POST("/v1/users", r.controller.FindByID, r.middleware.VerifyToken())
	r.echo.PUT("/v1/users/:id", r.controller.FindByID, r.middleware.VerifyToken())
	r.echo.DELETE("/v1/users/:id", r.controller.FindByID, r.middleware.VerifyToken())
}
