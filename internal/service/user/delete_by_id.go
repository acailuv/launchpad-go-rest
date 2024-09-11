package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
	"net/http"
)

func (s service) DeleteByID(ctx context.Context, p user.DeleteByIDRequest) error {
	if err := p.Validate(ctx); err != nil {
		return errors.NewWithCode(http.StatusBadRequest, errors.VALIDATION_ERROR, err.Error())
	}

	return s.user.DeleteByID(ctx, p.ID)
}
