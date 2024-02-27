package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	ErrMsgForbidden           = "Forbidden"
	ErrMsgBadRequest          = "Bad request"
	ErrMsgInternalServerError = "Internal Server Error"
	ErrMsgNotFound            = "Not Found"
	ErrMsgValidationFailed    = "Validation failed"
	ErrMsgUnauthorized        = "Unauthorized"
)

var (
	BadRequest          = errors.New("Bad request")
	NotFound            = errors.New("Not Found")
	Unauthorized        = errors.New("Unauthorized")
	Forbidden           = errors.New("Forbidden")
	PermissionDenied    = errors.New("Permission Denied")
	NotRequiredFields   = errors.New("No such required fields")
	BadQueryParams      = errors.New("Invalid query params")
	InternalServerError = errors.New("Internal Server Error")
	RequestTimeoutError = errors.New("Request Timeout")
	InvalidJWTToken     = errors.New("Invalid JWT token")
	InvalidJWTClaims    = errors.New("Invalid JWT claims")
)

// Error struct
type Error struct {
	ErrCode     int    `json:"code,omitempty"`
	ErrMessage  string `json:"message,omitempty"`
	ErrCauses   any    `json:"-"`
	ErrValidate any    `json:"errors,omitempty"`
	StatusCode  int    `json:"-"`
}

// Error  Error() interface method
func (e Error) Error() string {
	return fmt.Sprintf("code: %d - message: %s - errors: %v", e.ErrCode, e.ErrMessage, e.ErrCauses)
}

// Error status
func (e Error) Status() int {
	if e.StatusCode == 0 {
		return http.StatusBadRequest
	}
	return e.StatusCode
}

// Error Causes
func (e Error) Causes() any {
	return e.ErrCauses
}

// New Error
func NewError(code int, message string, causes any) *Error {
	return &Error{
		ErrCode:    code,
		ErrMessage: message,
		ErrCauses:  causes,
		StatusCode: http.StatusBadRequest,
	}
}

// New Error With Message
func NewErrorWithMessage(code int, message string, causes any) *Error {
	return &Error{
		ErrCode:    code,
		ErrMessage: message,
		ErrCauses:  causes,
		StatusCode: http.StatusBadRequest,
	}
}

// New Error From Bytes
func NewErrorFromBytes(bytes []byte) (*Error, error) {
	apiErr := &Error{}
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// New Bad Request Error
func NewBadRequestError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusBadRequest,
		ErrMessage: ErrMsgBadRequest,
		ErrCauses:  causes,
		StatusCode: http.StatusBadRequest,
	}
}

// New Validate Error
func NewValidateError(validateErr any) *Error {
	return &Error{
		ErrCode:     http.StatusBadRequest,
		ErrMessage:  ErrMsgValidationFailed,
		ErrValidate: validateErr,
		StatusCode:  http.StatusBadRequest,
	}
}

// New Not Found Error
func NewNotFoundError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusNotFound,
		ErrMessage: ErrMsgNotFound,
		ErrCauses:  causes,
		StatusCode: http.StatusNotFound,
	}
}

// New Unauthorized Error
func NewUnauthorizedError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusUnauthorized,
		ErrMessage: ErrMsgUnauthorized,
		ErrCauses:  causes,
		StatusCode: http.StatusUnauthorized,
	}
}

// New Forbidden Error
func NewForbiddenError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusForbidden,
		ErrMessage: ErrMsgForbidden,
		ErrCauses:  causes,
		StatusCode: http.StatusForbidden,
	}
}

// New Internal Server Error
func NewInternalServerError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusInternalServerError,
		ErrMessage: ErrMsgInternalServerError,
		ErrCauses:  causes,
		StatusCode: http.StatusInternalServerError,
	}
}

type ErrorWithMessage struct {
	cause error
	msg   string
}

func (w *ErrorWithMessage) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *ErrorWithMessage) Cause() error  { return w.cause }

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &ErrorWithMessage{
		cause: err,
		msg:   message,
	}
}
