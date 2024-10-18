package utils

import (
	"launchpad-go-rest/pkg/types/utils"

	"github.com/labstack/echo/v4"
)

// Utils godoc
//
//	@Summary		Publish a task to queue
//	@Description	Publish a task to queue
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		utils.PublishTaskRequest	true	"body"
//	@Failure		400		{object}	base.ErrorResponse
//	@Failure		401		{object}	base.ErrorResponse
//	@Failure		500		{object}	base.ErrorResponse
//	@Router			/v1/utils/publish-task [post]
func (h controller) PublishTask(c echo.Context) error {
	var req utils.PublishTaskRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.utils.PublishTask(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(200, nil)
}
