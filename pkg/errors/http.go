package errors

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type errorWithStatus interface {
	error
	Status() int
}

// HTTP Parser of error string messages returns RestError
func HTTPParseErrors(err error) errorWithStatus {
	switch parsedErr := err.(type) {
	case *Error:
		return parsedErr
	case *errorsCollector[error]:
		return parsedErr
	case *errorsCollector[*ValidationError]:
		return parsedErr
	default:
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return NewError(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
		case strings.Contains(err.Error(), "Unmarshal"):
			return NewError(http.StatusBadRequest, BadRequest.Error(), err)
		case strings.Contains(err.Error(), "UUID"):
			return NewError(http.StatusBadRequest, err.Error(), err)
		case strings.Contains(strings.ToLower(err.Error()), "cookie"):
			return NewError(http.StatusUnauthorized, Unauthorized.Error(), err)
		case strings.Contains(strings.ToLower(err.Error()), "token"):
			return NewError(http.StatusUnauthorized, Unauthorized.Error(), err)
		case strings.Contains(strings.ToLower(err.Error()), "bcrypt"):
			return NewError(http.StatusBadRequest, BadRequest.Error(), err)
		case strings.Contains(strings.ToLower(err.Error()), "not found"):
			return NewError(http.StatusNotFound, NotFound.Error(), err)
		default:
			return NewInternalServerError(err)
		}
	}
}

// HTTP Error response
func HTTPErrorResponse(err error) (int, any) {
	return HTTPParseErrors(err).Status(), HTTPParseErrors(err)
}
