package user

type FindByIDRequest struct {
	ID string `param:"id"`
}

type CreateRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type UpdateByIDRequest struct {
	ID                   string `param:"id"`
	Email                string `json:"email"`
	OldPassword          string `json:"old_password"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type DeleteByIDRequest struct {
	ID string `param:"id"`
}
