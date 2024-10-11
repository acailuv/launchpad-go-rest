package user

import (
	"context"
	"database/sql"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/user"
)

func (r repository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	var res user.User

	err := r.db.GetContext(ctx, &res, r.db.Rebind(`
		SELECT id, email, password FROM users WHERE email = ?
	`), email)
	if errors.Is(err, sql.ErrNoRows) {
		return res, err
	} else if err != nil {
		return res, errors.New(err)
	}

	return res, nil
}
