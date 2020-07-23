package utils

import "os"

// CreateFolder is
func CreateFolder(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// RemoveFolder is
func RemoveFolder(path string) error {
	return os.RemoveAll(path)
}
