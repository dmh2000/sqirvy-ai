package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// MaxTotalBytes is the maximum allowed size for all input files combined
const MaxTotalBytes = 100 * 1024 // 100KB limit

// ReadFiles reads and concatenates the contents of the given files,
// returning an error if any file doesn't exist or if total size exceeds MaxTotalBytes
func ReadFiles(filenames []string) (string, error) {
	if len(filenames) == 0 {
		return "", nil
	}

	var totalSize int64
	var result string

	for _, fname := range filenames {
		// Sanitize path
		cleanPath := filepath.Clean(fname)

		// Check if file exists
		info, err := os.Stat(cleanPath)
		if err != nil {
			if os.IsNotExist(err) {
				return "", fmt.Errorf("file does not exist: %s", fname)
			}
			return "", fmt.Errorf("error accessing file %s: %v", fname, err)
		}

		// Check if new file would exceed size limit
		newSize := totalSize + info.Size()
		if newSize > MaxTotalBytes {
			return "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
		}

		// Read file
		content, err := os.ReadFile(cleanPath)
		if err != nil {
			return "", fmt.Errorf("error reading file %s: %v", fname, err)
		}

		result += string(content)
		totalSize = newSize
	}

	return result, nil
}
