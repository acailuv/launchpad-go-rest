package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
	"net/http"

	"github.com/oklog/ulid/v2"
)

func (s service) Create(ctx context.Context, p user.CreateRequest) error {
	if err := p.Validate(ctx); err != nil {
		return errors.NewWithCode(http.StatusBadRequest, errors.VALIDATION_ERROR, err.Error())
	}

	if p.Password != p.PasswordConfirmation {
		return errors.NewWithCode(http.StatusBadRequest, errors.PASSWORD_CONFIRMATION_MISMATCH, "Password confirmation does not match")
	}

	passwordHash, err := s.utils.HashPassword(p.Password)
	if err != nil {
		return err
	}

	return s.user.Create(ctx, user.User{
		ID:       ulid.Make().String(),
		Email:    p.Email,
		Password: passwordHash,
	})
}
