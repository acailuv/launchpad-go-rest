package user

import (
	"launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

func (h controller) DeleteByID(c echo.Context) error {
	var req user.DeleteByIDRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.user.DeleteByID(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(200, nil)
}
