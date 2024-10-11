package router

import (
	"launchpad-go-rest/internal/controller"
	"launchpad-go-rest/internal/middleware"
	"launchpad-go-rest/internal/router/user"

	"github.com/labstack/echo/v4"
)

type routers struct {
	User user.Router
}

func Init(e *echo.Echo, controllers *controller.Controller, middleware middleware.Middleware) {
	routers := routers{
		User: user.New(e, controllers.User, middleware),
	}

	routers.User.HandleRoutes()
}
