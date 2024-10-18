package utils

import (
	"context"

	"github.com/acailuv/numeric"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PublishTaskRequest struct {
	Queue  string          `json:"queue"`
	ID     string          `json:"id"`
	Amount numeric.Numeric `json:"amount"`
}

func (x PublishTaskRequest) Validate(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		&x,
		validation.Field(&x.Queue, validation.Required),
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.Amount, validation.Required),
	)
}
