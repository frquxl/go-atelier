package fs

import (
	"fmt"
	"os"
)

// WriteFile creates a file with the given content.
func WriteFile(path string, content []byte) error {
	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	return nil
}

// CreateDir creates a directory if it doesn't exist.
func CreateDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}