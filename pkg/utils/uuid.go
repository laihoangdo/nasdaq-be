package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.New().String()
}

// IsValidUUID checks UUID validation
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
