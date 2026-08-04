// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build windows

package suppress

import (
	"os"
)

func lockFile(file *os.File) error {
	return lockFileWindows(file)
}

func unlockFile(file *os.File) error {
	return unlockFileWindows(file)
}
