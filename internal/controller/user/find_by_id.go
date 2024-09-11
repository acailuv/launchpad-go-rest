package user

import (
	"launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

func (h controller) FindByID(c echo.Context) error {
	var req user.FindByIDRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.user.FindByID(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
