package errors

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type HTTPError struct {
	statusCode int
	message    string
	cause      error
}

var _ error = HTTPError{}

func (e HTTPError) StatusCode() int {
	return e.statusCode
}

func (e HTTPError) Error() string {
	return e.message
}

func (e HTTPError) WithMessage(message string) HTTPError {
	return HTTPError{
		statusCode: e.statusCode,
		message:    message,
		cause:      e.cause,
	}
}

func (e HTTPError) WithCause(cause error) HTTPError {
	return HTTPError{
		statusCode: e.statusCode,
		message:    e.message,
		cause:      cause,
	}
}

func (e HTTPError) Unwrap() error {
	return e.cause
}

var (
	ErrBadRequest   = HTTPError{statusCode: http.StatusBadRequest, message: "bad request"}
	ErrUnauthorized = HTTPError{statusCode: http.StatusUnauthorized, message: "unauthorized"}
	ErrForbidden    = HTTPError{statusCode: http.StatusForbidden, message: "forbidden"}
	ErrNotFound     = HTTPError{statusCode: http.StatusNotFound, message: "not found"}

	ErrInternal = HTTPError{statusCode: http.StatusInternalServerError, message: "bad request"}
)

func TransformGormErr(err error) error {
	if err == nil || errors.As(err, &HTTPError{}) {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound.WithCause(err)
	}
	return ErrInternal.WithCause(err)
}
