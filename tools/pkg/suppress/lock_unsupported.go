// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build !darwin && !linux && !windows

package suppress

import (
	"fmt"
	"os"
	"runtime"
)

func lockFile(_ *os.File) error {
	return fmt.Errorf("cross-process suppression locking is unsupported on %s", runtime.GOOS)
}

func unlockFile(_ *os.File) error {
	return nil
}
