package common

import "os"

// IsFileExists check if file exists
func IsFileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
