package errors

type Error struct {
	StatusCode int
	ErrorCode  int
	Err        error
}

func (e *Error) Error() string {
	return e.Err.Error()
}
