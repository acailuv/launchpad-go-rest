package user

type User struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (x User) ToFindByIDResponse() FindByIDResponse {
	return FindByIDResponse{
		ID:    x.ID,
		Email: x.Email,
	}
}

func (x User) ToFindResponse() FindResponse {
	return FindResponse{
		ID:    x.ID,
		Email: x.Email,
	}
}
