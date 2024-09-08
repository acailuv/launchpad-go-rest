package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
)

func (r repository) Create(ctx context.Context, u user.User) error {
	_, err := r.db.ExecContext(ctx, r.db.Rebind(`
		INSERT INTO users (id, email, password) VALUES (?, ?, ?)
	`), u.ID, u.Email, u.Password)
	if err != nil {
		return errors.New(err)
	}

	return nil
}
