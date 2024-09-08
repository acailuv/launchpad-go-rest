package controller

import (
	"launchpad-go-rest/internal/controller/user"
	"launchpad-go-rest/internal/service"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	User user.Controller
}

func Init(services *service.Service, logger echo.Logger) *Controller {
	return &Controller{
		User: user.New(services.User, logger),
	}
}
