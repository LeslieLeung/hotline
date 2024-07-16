package misc

import (
	"os"
	"path/filepath"
)

// SafeWriteFile writes data to a file, creating any necessary directories.
func SafeWriteFile(path string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
