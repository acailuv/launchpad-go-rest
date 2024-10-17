package errors

import (
	"fmt"
	"net/http"

	"github.com/go-errors/errors"
)

func New(e any) error {
	err := errors.Wrap(e, 1)

	return errors.New(&Error{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  INTERNAL_SERVER_ERROR,
		Err:        err,
	})
}

func NewWithCode(statusCode int, errorCode int, e any) error {
	var err error
	switch e := e.(type) {
	case error:
		err = e
	default:
		err = fmt.Errorf("%v", e)
	}

	return errors.New(&Error{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Err:        err,
	})
}

func Is(e error, original error) bool {
	var _e *Error
	if errors.As(e, &_e) {
		return errors.Is(_e.Err, original)
	}

	return errors.Is(e, original)
}

func As(e error, target any) bool {
	var _e *Error
	if errors.As(e, &_e) {
		return errors.As(_e.Err, target)
	}

	return errors.As(e, target)
}
