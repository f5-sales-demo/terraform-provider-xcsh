// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build !darwin && !linux && !windows

package suppress

import (
	"fmt"
	"runtime"
)

func replaceFileAndSyncDirectory(_, _, _ string) error {
	return fmt.Errorf("durable suppression replacement is unsupported on %s", runtime.GOOS)
}
