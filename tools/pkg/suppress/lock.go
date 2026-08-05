// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package suppress

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// outputLock is a cross-process advisory lock on a stable file adjacent to the
// generated output. Darwin, Linux, and Windows are supported; other targets
// fail closed. The lock file intentionally survives unlock: unlinking it would
// let a third process lock a new inode while a waiter still holds the old one.
// Its .tmp suffix keeps the coordination artifact out of version control.
type outputLock struct {
	file *os.File
}

func acquireOutputLock(outputPath string) (*outputLock, error) {
	directory := filepath.Dir(outputPath)
	lockPath := filepath.Join(directory, "."+filepath.Base(outputPath)+".emit-lock.tmp")
	file, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return nil, fmt.Errorf("open suppression lock for %s: %w", outputPath, err)
	}
	if err := lockFile(file); err != nil {
		_ = file.Close()
		return nil, fmt.Errorf("lock suppression output %s: %w", outputPath, err)
	}
	return &outputLock{file: file}, nil
}

func (lock *outputLock) Close() error {
	if lock == nil || lock.file == nil {
		return nil
	}
	unlockErr := unlockFile(lock.file)
	closeErr := lock.file.Close()
	lock.file = nil
	if err := errors.Join(unlockErr, closeErr); err != nil {
		return fmt.Errorf("release suppression output lock: %w", err)
	}
	return nil
}
