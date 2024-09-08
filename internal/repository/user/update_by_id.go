package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
)

func (r repository) UpdateByID(ctx context.Context, u user.User) error {
	_, err := r.db.ExecContext(ctx, r.db.Rebind(`
		UPDATE users SET email = ?, password = ? WHERE id = ?
	`), u.Email, u.Password, u.ID)
	if err != nil {
		return errors.New(err)
	}

	return nil
}
