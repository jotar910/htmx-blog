package cerrors

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (ce CustomError) Error() string {
	return fmt.Sprintf("%s (%d)", ce.Message, ce.Code)
}

var (
	BadRequest          = CustomError{Code: 400, Message: "Bad Request"}
	NotFound            = CustomError{Code: 404, Message: "Not Found"}
	InternalServerError = CustomError{Code: 500, Message: "Internal Server Error"}
)

func Wrap(cause error, customError CustomError) error {
	return fmt.Errorf("%w: %w", cause, customError)
}

func Unwrap(err error) CustomError {
	var cerr CustomError
	if !errors.As(err, &cerr) {
		cerr = InternalServerError
	}
	return cerr
}
