package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type ValidationError struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

// NewValidationError creates a new validation error.
func NewValidationError(field string, messages ...string) *ValidationError {
	return &ValidationError{
		Field:    field,
		Messages: messages,
	}
}

// Error implements the error interface.
func (ve ValidationError) Error() string {
	return fmt.Sprintf("validation messages: %s", strings.Join(ve.Messages, ", "))
}

// ErrorsCollector is a collector of errors.
type errorsCollector[T error] struct {
	ErrCode    int    `json:"code,omitempty"`
	ErrMessage string `json:"message,omitempty"`
	Errors     []T    `json:"errors,omitempty"`
	StatusCode int    `json:"-"`
}

func NewErrorsCollector(errCode int, msg string) *errorsCollector[error] {
	return &errorsCollector[error]{
		ErrCode:    errCode,
		ErrMessage: msg,
	}
}

func NewValidationErrorsCollector() *errorsCollector[*ValidationError] {
	return &errorsCollector[*ValidationError]{
		ErrCode:    http.StatusBadRequest,
		ErrMessage: ErrMsgValidationFailed,
		StatusCode: http.StatusOK,
	}
}

// Error implements the error interface.
func (ec errorsCollector[T]) Error() string {
	return fmt.Sprintf("errors: %s", ec.Errors)
}

// Add adds an error to the collector.
func (ec *errorsCollector[T]) Add(err T) {
	ec.Errors = append(ec.Errors, err)
}

// Status returns the status code
func (ec errorsCollector[T]) Status() int {
	if ec.StatusCode == 0 {
		return http.StatusBadRequest
	}
	return ec.StatusCode
}

// HasError returns true if the collector has errors.
func (ec errorsCollector[T]) HasError() bool {
	return len(ec.Errors) > 0
}
