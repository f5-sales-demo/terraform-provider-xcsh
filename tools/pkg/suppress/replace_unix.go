// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build darwin || linux

package suppress

import (
	"fmt"
	"os"
)

func replaceFileAndSyncDirectory(temporaryPath, path, directoryPath string) error {
	if err := os.Rename(temporaryPath, path); err != nil {
		return err
	}
	directory, err := os.Open(directoryPath)
	if err != nil {
		return fmt.Errorf("open containing directory for sync: %w", err)
	}
	if err := directory.Sync(); err != nil {
		_ = directory.Close()
		return fmt.Errorf("sync containing directory: %w", err)
	}
	if err := directory.Close(); err != nil {
		return fmt.Errorf("close containing directory after sync: %w", err)
	}
	return nil
}
