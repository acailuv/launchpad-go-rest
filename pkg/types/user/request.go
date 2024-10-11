package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type FindByIDRequest struct {
	ID string `param:"id"`
}

func (x FindByIDRequest) Validate(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		&x,
		validation.Field(&x.ID, validation.Required),
	)
}

type CreateRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (x CreateRequest) Validate(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		&x,
		validation.Field(&x.Email, validation.Required),
		validation.Field(&x.Password, validation.Required),
		validation.Field(&x.PasswordConfirmation, validation.Required),
	)
}

type UpdateByIDRequest struct {
	ID                   string `param:"id" json:"-"`
	Email                string `json:"email"`
	OldPassword          string `json:"old_password"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (x UpdateByIDRequest) Validate(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		&x,
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.Email, validation.Required),
		validation.Field(&x.OldPassword, validation.Required),
		validation.Field(&x.Password, validation.Required),
		validation.Field(&x.PasswordConfirmation, validation.Required),
	)
}

type DeleteByIDRequest struct {
	ID string `param:"id"`
}

func (x DeleteByIDRequest) Validate(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		&x,
		validation.Field(&x.ID, validation.Required),
	)
}
