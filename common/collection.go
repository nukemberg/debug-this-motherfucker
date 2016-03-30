package common

// StringInSlice check if string in slice
func StringInSlice(slice []string, s string) bool {
	for _, x := range slice {
		if x == s {
			return true
		}
	}
	return false
}
