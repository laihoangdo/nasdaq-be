package utils

// ToPointer convert a value to pointer
func ToPointer[T any](v T) *T {
	return &v
}
