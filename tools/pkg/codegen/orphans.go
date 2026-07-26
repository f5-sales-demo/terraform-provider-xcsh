// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// PruneOrphanFiles deletes every <dir>/<name><ext> whose name is not in keep,
// and returns the paths it removed, sorted.
//
// Generators that only write and overwrite cannot express a DELETION. When F5
// removes a schema family upstream, codegen stops emitting that resource — but
// every page already on disk stays there, indefinitely, describing a resource
// the provider no longer ships. docs/_llms-txt/resources/apm.txt outlived the
// apm* schemas by several releases exactly that way (#1351).
//
// keep is the set of names the generator just emitted, so pruning is derived
// from the same source of truth as generation and there is no second list to
// maintain. An empty keep set is refused: that is what an upstream fetch failure
// or a bad glob looks like, and wiping the entire directory in response would
// turn a transient error into deleted documentation.
func PruneOrphanFiles(dir, ext string, keep map[string]bool) ([]string, error) {
	if ext == "" || !strings.HasPrefix(ext, ".") {
		return nil, fmt.Errorf("extension %q must be non-empty and start with a dot", ext)
	}
	if len(keep) == 0 {
		return nil, fmt.Errorf("refusing to prune %s: the keep set is empty, which would delete every %s file", dir, ext)
	}

	matches, err := filepath.Glob(filepath.Join(dir, "*"+ext))
	if err != nil {
		return nil, fmt.Errorf("scanning %s for orphans: %w", dir, err)
	}

	var removed []string
	for _, path := range matches {
		name := strings.TrimSuffix(filepath.Base(path), ext)
		if keep[name] {
			continue
		}
		if err := os.Remove(path); err != nil {
			return removed, fmt.Errorf("removing orphan %s: %w", path, err)
		}
		removed = append(removed, path)
	}

	sort.Strings(removed)
	return removed, nil
}
