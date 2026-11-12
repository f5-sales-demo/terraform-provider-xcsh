//go:build !linux

package provider

import "os"

// Opt-in provenance currently requires independently observable Linux /proc.
// Disabled capture does not alter any supported provider platform.
func captureOwned(os.FileInfo) bool { return false }
func originalCreationProcess() (OriginalCreationProcess, error) {
	return OriginalCreationProcess{}, captureError()
}
