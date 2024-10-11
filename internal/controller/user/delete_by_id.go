package user

import (
	"launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

// Users godoc
//
//	@Summary		Delete user by ID
//	@Description	Delete user by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Failure		400	{object}	base.ErrorResponse
//	@Failure		401	{object}	base.ErrorResponse
//	@Failure		500	{object}	base.ErrorResponse
//	@Router			/v1/users/{id} [delete]
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
