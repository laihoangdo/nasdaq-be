package utils

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// ValidatePassword validates password string
func ValidatePassword(password string) bool {
	check, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, password)
	return check
}
