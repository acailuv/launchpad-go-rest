package user

import (
	"launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

// Users godoc
//
//	@Summary		Find user by ID
//	@Description	Find a user by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	user.FindByIDResponse
//	@Failure		400	{object}	base.ErrorResponse
//	@Failure		401	{object}	base.ErrorResponse
//	@Failure		500	{object}	base.ErrorResponse
//	@Router			/v1/users/{id} [get]
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
