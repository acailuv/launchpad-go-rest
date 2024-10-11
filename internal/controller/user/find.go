package user

import (
	"github.com/labstack/echo/v4"
)

// Users godoc
//
//	@Summary		List users
//	@Description	List all users without any filter and pagination support
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		user.FindResponse
//	@Failure		400	{object}	base.ErrorResponse
//	@Failure		401	{object}	base.ErrorResponse
//	@Failure		500	{object}	base.ErrorResponse
//	@Router			/v1/users [get]
func (h controller) Find(c echo.Context) error {
	res, err := h.user.Find(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
