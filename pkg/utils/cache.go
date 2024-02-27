package utils

import (
	"fmt"
	"strconv"
)

// Generate key
func GenKeyForCache(data interface{}) string {
	switch v := data.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}
