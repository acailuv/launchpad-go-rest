package user

import (
	"launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

// Users godoc
//
//	@Summary		Create user
//	@Description	Create user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		user.CreateRequest	true	"body"
//	@Failure		400		{object}	base.ErrorResponse
//	@Failure		401		{object}	base.ErrorResponse
//	@Failure		500		{object}	base.ErrorResponse
//	@Router			/v1/users [post]
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
