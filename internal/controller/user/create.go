package user

import (
	"launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

func (h controller) Create(c echo.Context) error {
	var req user.CreateRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.user.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(201, nil)
}
