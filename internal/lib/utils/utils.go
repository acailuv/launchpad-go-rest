package utils

type Utils interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) bool
}

type utils struct{}

func New() Utils {
	return &utils{}
}
