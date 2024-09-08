package errors

import (
	"fmt"
	"net/http"

	"github.com/go-errors/errors"
)

func New(e any) error {
	var err error
	switch e := e.(type) {
	case error:
		err = e
	default:
		err = fmt.Errorf("%v", e)
	}

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
	return errors.Is(e, original)
}

func As(e error, target any) bool {
	return errors.As(e, target)
}
