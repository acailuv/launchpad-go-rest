package user

import (
	"context"
	"database/sql"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
	"net/http"
)

func (s service) FindByID(ctx context.Context, id string) (user.FindByIDResponse, error) {
	res, err := s.user.FindByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return res.ToFindByIDResponse(), errors.NewWithCode(http.StatusBadRequest, errors.USER_NOT_FOUND, "Invalid user id")
	} else if err != nil {
		return res.ToFindByIDResponse(), err
	}

	return res.ToFindByIDResponse(), nil
}
