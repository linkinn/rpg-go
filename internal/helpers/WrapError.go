package helpers

import (
	"fmt"
	"rpg-go/internal/dto"
)

type Error struct {
	err     error
	message string
	code    uint
}

const (
	ErrorCodeUnknown         = 400
	ErrorCodeNotFound        = 404
	ErrorCodeInvalidArgument = 402
)

func WrapError(err error, code uint, format string, a ...interface{}) *Error {
	return &Error{
		code:    code,
		err:     err,
		message: fmt.Sprintf(format, a...),
	}
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %v", e.message, e.err)
	}

	return e.message
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Code() uint {
	return e.code
}

func (e *Error) ErrorJSON() dto.ErrorOutputDTO {
	var out dto.ErrorOutputDTO

	out = dto.ErrorOutputDTO{Err: e.err.Error(), Message: e.message, Code: e.code}
	return out
}
