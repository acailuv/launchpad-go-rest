package user

import (
	"context"
	"database/sql"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
	"net/http"
)

func (s service) UpdateByID(ctx context.Context, p user.UpdateByIDRequest) error {
	if err := p.Validate(ctx); err != nil {
		return errors.NewWithCode(http.StatusBadRequest, errors.VALIDATION_ERROR, err.Error())
	}

	if p.Password != p.PasswordConfirmation {
		return errors.NewWithCode(http.StatusBadRequest, errors.PASSWORD_CONFIRMATION_MISMATCH, "Password confirmation does not match")
	}

	existingUser, err := s.user.FindByID(ctx, p.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.NewWithCode(http.StatusBadRequest, errors.USER_NOT_FOUND, "Invalid user id")
	} else if err != nil {
		return err
	}

	if !s.utils.ComparePassword(existingUser.Password, p.OldPassword) {
		return errors.NewWithCode(http.StatusBadRequest, errors.INVALID_OLD_PASSWORD, "Invalid old password")
	}

	newPasswordHash, err := s.utils.HashPassword(p.Password)
	if err != nil {
		return err
	}

	return s.user.UpdateByID(ctx, user.User{
		ID:       p.ID,
		Email:    p.Email,
		Password: newPasswordHash,
	})
}
