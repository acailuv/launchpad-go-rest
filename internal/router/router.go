package router

import (
	"launchpad-go-rest/internal/controller"
	"launchpad-go-rest/internal/middleware"
	"launchpad-go-rest/internal/router/user"
	"launchpad-go-rest/internal/router/utils"

	"github.com/labstack/echo/v4"
)

type routers struct {
	User  user.Router
	Utils utils.Router
}

func Init(e *echo.Echo, controllers *controller.Controller, middleware middleware.Middleware) {
	routers := routers{
		User:  user.New(e, controllers.User, middleware),
		Utils: utils.New(e, controllers.Utils, middleware),
	}

	routers.User.HandleRoutes()
	routers.Utils.HandleRoutes()
}
