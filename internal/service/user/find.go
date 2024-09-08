package user

import (
	"context"
	"database/sql"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
	"net/http"
)

func (s service) Find(ctx context.Context) ([]user.FindResponse, error) {
	var response []user.FindResponse = make([]user.FindResponse, 0)

	res, err := s.user.Find(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return response, errors.NewWithCode(http.StatusBadRequest, errors.USER_NOT_FOUND, "Invalid user id")
	} else if err != nil {
		return response, err
	}

	for _, user := range res {
		response = append(response, user.ToFindResponse())
	}

	return response, nil
}
