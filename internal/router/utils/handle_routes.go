package utils

func (r router) HandleRoutes() {
	r.echo.POST("/v1/utils/publish-task", r.controller.PublishTask)
}
