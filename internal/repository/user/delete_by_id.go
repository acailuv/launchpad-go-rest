package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
)

func (r repository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, r.db.Rebind(`
		DELETE FROM users WHERE id = ?
	`), id)
	if err != nil {
		return errors.New(err)
	}

	return nil
}
