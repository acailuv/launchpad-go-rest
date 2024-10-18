package controller

import (
	"launchpad-go-rest/internal/controller/user"
	"launchpad-go-rest/internal/controller/utils"
	"launchpad-go-rest/internal/service"
)

type Controller struct {
	User  user.Controller
	Utils utils.Controller
}

func Init(services *service.Service) *Controller {
	return &Controller{
		User:  user.New(services.User),
		Utils: utils.New(services.Utils),
	}
}
