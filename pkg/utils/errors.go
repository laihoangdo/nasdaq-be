package utils

import "errors"

var (
	ErrInvalidHourFormat = errors.New("invalid hour format")
	ErrInvalidUUID       = errors.New("invalid uuid format")
)
