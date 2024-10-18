package utils

import (
	"launchpad-go-rest/internal/service/utils"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	PublishTask(c echo.Context) error
}

type controller struct {
	utils utils.Service
}

func New(utils utils.Service) Controller {
	return &controller{
		utils: utils,
	}
}
