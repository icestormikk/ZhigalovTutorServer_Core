package structs

import (
	"errors"
	"net/http"
)

var (
	NotFound       = errors.New("not found")
	BadRequest     = errors.New("bad request")
	InternalServer = errors.New("internal server error")
)

type ApplicationError struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
	err     error
}

func (e *ApplicationError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return "application error"
}

func (e *ApplicationError) Unwrap() error {
	return e.err
}

func InternalServerError(message string) *ApplicationError {
	return &ApplicationError{
		Code:    http.StatusInternalServerError,
		Message: message,
		err:     InternalServer,
	}
}

func NotFoundError(message string) *ApplicationError {
	return &ApplicationError{
		Code:    http.StatusNotFound,
		Message: message,
		err:     NotFound,
	}
}

func BadRequestError(message string) *ApplicationError {
	return &ApplicationError{
		Code:    http.StatusBadRequest,
		Message: message,
		err:     BadRequest,
	}
}
