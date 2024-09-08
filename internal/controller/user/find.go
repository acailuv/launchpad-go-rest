package user

import (
	"github.com/labstack/echo/v4"
)

func (h controller) Find(c echo.Context) error {
	res, err := h.user.Find(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
