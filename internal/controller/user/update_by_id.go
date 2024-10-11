package user

import (
	"launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

// Users godoc
//
//	@Summary		Update user by ID
//	@Description	Update user by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"User ID"
//	@Param			request	body		user.UpdateByIDRequest	true	"body"
//	@Failure		400		{object}	base.ErrorResponse
//	@Failure		401		{object}	base.ErrorResponse
//	@Failure		500		{object}	base.ErrorResponse
//	@Router			/v1/users/{id} [put]
func (h controller) UpdateByID(c echo.Context) error {
	var req user.UpdateByIDRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.user.UpdateByID(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(200, nil)
}
