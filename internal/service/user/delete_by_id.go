package user

import (
	"context"
)

func (s service) DeleteByID(ctx context.Context, id string) error {
	return s.user.DeleteByID(ctx, id)
}
