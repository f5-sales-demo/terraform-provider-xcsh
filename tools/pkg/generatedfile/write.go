// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package generatedfile avoids rewriting identical generator output.
package generatedfile

import (
	"bytes"
	"errors"
	"os"
)

// Matches reports whether an existing output has exactly the requested bytes.
// Missing output is a normal first generation; other read errors are failures.
func Matches(path string, data []byte) (bool, error) {
	existing, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return bytes.Equal(existing, data), nil
}

// WriteFile preserves the inode, timestamp and permissions of identical output.
// New or changed output retains os.WriteFile semantics, including its errors.
func WriteFile(path string, data []byte, mode os.FileMode) error {
	same, err := Matches(path, data)
	if err != nil {
		return err
	}
	if same {
		return nil
	}
	return os.WriteFile(path, data, mode)
}
