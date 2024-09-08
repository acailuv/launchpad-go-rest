package user

type FindByIDResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type FindResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
