package utils

import (
	"errors"
	"regexp"
)

// Errors
var (
	// ErrPhoneInvalid the invalid phone number format error
	ErrPhoneInvalid = errors.New("invalid phone number format")
)

var regexpPhone = regexp.MustCompile(`(84|0[3|5|7|8|9])+([0-9]{8})\b`)

// Phone validates the phone number format
func Phone(s string) error {
	if !regexpPhone.MatchString(s) {
		return ErrPhoneInvalid
	}
	return nil
}
