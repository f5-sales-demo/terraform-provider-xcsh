// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"crypto/sha256"
	"encoding/base32"
	helper "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"os"
	"strings"
	"sync"
)

var compactNames = struct {
	sync.Mutex
	names map[string]bool
}{names: map[string]bool{}}

func compactRunPrefix(run string) string {
	digest := sha256.Sum256([]byte(run))
	return TestResourcePrefix + strings.ToLower(base32.StdEncoding.EncodeToString(digest[:])[:5]) + "-"
}

// CompactRandomName preserves the private full run identity in the ownership
// journal and annotation while satisfying stricter API-specific name limits.
func CompactRandomName(limit int) string {
	run := os.Getenv("XCSH_ACC_RUN_ID")
	if run == "" {
		run = "mock"
	}
	if !ownershipRunPattern.MatchString(run) {
		panic("invalid acceptance run identifier")
	}
	prefix := compactRunPrefix(run)
	remaining := limit - len(prefix)
	if remaining < 2 {
		panic("API name limit cannot carry a run identity")
	}
	compactNames.Lock()
	defer compactNames.Unlock()
	for attempt := 0; attempt < 4096; attempt++ {
		name := prefix + helper.RandStringFromCharSet(remaining, helper.CharSetAlphaNum)
		if !compactNames.names[name] {
			compactNames.names[name] = true
			return name
		}
	}
	panic("compact acceptance identity capacity exhausted")
}
func matchesOwnershipRun(name, run string) bool {
	return strings.HasPrefix(name, TestResourcePrefix+run+"-") || strings.HasPrefix(name, compactRunPrefix(run))
}
