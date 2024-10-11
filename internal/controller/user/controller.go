package user

import (
	"launchpad-go-rest/internal/service/user"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	Find(c echo.Context) error
	FindByID(c echo.Context) error
	Create(c echo.Context) error
	UpdateByID(c echo.Context) error
	DeleteByID(c echo.Context) error
}

type controller struct {
	user user.Service
}

func New(user user.Service) Controller {
	return &controller{
		user: user,
	}
}
