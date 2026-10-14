// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import "testing"

func TestCompactNamesRemainUniqueAcrossSerialSuite(t *testing.T) {
	t.Setenv("XCSH_ACC_RUN_ID", "compactunique")
	seen := map[string]bool{}
	for i := 0; i < 400; i++ {
		name := CompactRandomName(20)
		if seen[name] {
			t.Fatal("compact acceptance identities collided")
		}
		seen[name] = true
	}
}
