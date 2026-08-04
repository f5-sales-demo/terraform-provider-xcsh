// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build windows

package suppress

// Windows does not expose the Unix directory-fsync operation. WRITE_THROUGH is
// the platform durability equivalent: MoveFileEx does not return until the
// replacement has been flushed to disk.
func replaceFileAndSyncDirectory(temporaryPath, path, _ string) error {
	return replaceFileWindows(temporaryPath, path)
}
