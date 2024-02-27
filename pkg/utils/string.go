package utils

// Contains checks if arrStr contains any string s
func Contains(arrStr []string, s string) bool {
	for idx := range arrStr {
		if arrStr[idx] == s {
			return true
		}
	}
	return false
}
