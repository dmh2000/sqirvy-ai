package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// inputIsFromPipe determines if the program is receiving piped input on stdin.
// Returns true if stdin is a pipe, false if it's a terminal or other device.
func InputIsFromPipe() (bool, error) {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return (fileInfo.Mode() & os.ModeCharDevice) == 0, err
}

// ReadStdin reads and concatenates the contents of stdin,
func ReadStdin(maxTotalBytes int64) (data string, size int64, err error) {
	pipe, err := InputIsFromPipe()

	if err != nil {
		return "", 0, err
	}

	// not a pipe/stdin, so just return
	if !pipe {
		return "", 0, nil
	}

	stdinBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", 0, fmt.Errorf("error reading from stdin: %w", err)
	}
	size = int64(len(stdinBytes))
	if size > maxTotalBytes {
		return "", 0, fmt.Errorf("total size would exceed limit of %d bytes", maxTotalBytes)
	}
	// wrap thje input in backticks
	s := "```stdin\n" + string(stdinBytes) + "```"
	return s, size, nil
}

// validateFilePath checks if the given file path is safe and returns a cleaned version
func validateFilePath(fname string) (string, error) {
	// Sanitize path
	cleanPath := filepath.Clean(fname)

	cleanPath, err := filepath.EvalSymlinks(cleanPath)
	if err != nil {
		return "", fmt.Errorf("unsafe or invalid path specification %s: %w", fname, err)
	}
	return cleanPath, nil
}

// readFile reads and concatenates the contents of the given files,
// returning an error if any file doesn't exist, is suspicious or if total size exceeds maxTotalBytes
func ReadFile(fname string, maxTotalBytes int64) ([]byte, int64, error) {
	// Sanitize path
	cleanPath, err := validateFilePath(fname)
	if err != nil {
		return nil, 0, err
	}

	// Check if file exists
	info, err := os.Stat(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, 0, fmt.Errorf("file does not exist: %s", fname)
		}
		return nil, 0, fmt.Errorf("error accessing file %s: %v", fname, err)
	}

	// Check if new file would exceed size limit
	if info.Size() > maxTotalBytes {
		return nil, 0, fmt.Errorf("total size would exceed limit of %d bytes", maxTotalBytes)
	}

	// path is valid, open the file
	file, err := os.Open(cleanPath)
	if err != nil {
		return nil, 0, fmt.Errorf("error opening file %s: %w", fname, err)
	}
	defer file.Close()

	// Read file in chunks
	rdr := bufio.NewReader(file)
	buf := make([]byte, 1024)
	content := make([]byte, 0, maxTotalBytes)
	var size int64 = 0
	for {
		// Read chunk from file
		n, err := rdr.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, fmt.Errorf("error reading file %s: %w", fname, err)
		}
		size += int64(n)
		if size > maxTotalBytes {
			return nil, 0, fmt.Errorf("total size would exceed limit of %d bytes", maxTotalBytes)
		}
		// Append chunk to buffer
		content = append(content, buf[:n]...)
	}

	return content, int64(len(content)), nil
}

// readFiles reads and concatenates the contents of the given files,
// returning an error if any file doesn't exist, is suspicious or if total size exceeds MaxTotalBytes
func ReadFiles(filenames []string, maxTotalBytes int64) (string, int64, error) {
	// Check if we have any files
	if len(filenames) == 0 {
		return "", 0, nil
	}

	var totalSize int64
	builder := strings.Builder{}
	for _, fname := range filenames {
		start := []byte(fmt.Sprintf("```%s", fname))
		builder.Write(start)
		// Read file
		s, size, err := ReadFile(fname, maxTotalBytes)
		if err != nil {
			return "", size, fmt.Errorf("error reading file %s: %w", fname, err)
		}
		// check size
		totalSize += size
		if totalSize > maxTotalBytes {
			return "", totalSize, fmt.Errorf("total size would exceed limit of %d bytes", maxTotalBytes)
		}
		// add to builder
		builder.Write(s)

		end := []byte("```")
		builder.Write(end)
	}

	return builder.String(), totalSize, nil
}
