package errors

const (
	INTERNAL_SERVER_ERROR = iota + 1
	USER_NOT_FOUND
	INVALID_AUTH_TOKEN
	PASSWORD_CONFIRMATION_MISMATCH
	INVALID_OLD_PASSWORD
	VALIDATION_ERROR
)
