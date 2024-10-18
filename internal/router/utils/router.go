package utils

import (
	"launchpad-go-rest/internal/controller/utils"
	"launchpad-go-rest/internal/middleware"

	"github.com/labstack/echo/v4"
)

type Router interface {
	HandleRoutes()
}

type router struct {
	echo       *echo.Echo
	controller utils.Controller
	middleware middleware.Middleware
}

func New(e *echo.Echo, controller utils.Controller, middleware middleware.Middleware) Router {
	return &router{
		echo:       e,
		controller: controller,
		middleware: middleware,
	}
}
