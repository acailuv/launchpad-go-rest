package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
)

func (r repository) Find(ctx context.Context) ([]user.User, error) {
	var res []user.User

	err := r.db.SelectContext(ctx, &res, r.db.Rebind(`
		SELECT id, email, password FROM users
	`))
	if err != nil {
		return res, errors.New(err)
	}

	return res, nil
}
