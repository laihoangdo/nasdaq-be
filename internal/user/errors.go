package user

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserWrongPassword = errors.New("wrong password")
)
