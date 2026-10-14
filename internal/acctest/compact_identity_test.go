// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"strings"
	"testing"
)

func TestCompactAcceptanceIdentityFitsAPILimitAndRejectsOtherRun(t *testing.T) {
	t.Setenv("XCSH_ACC_RUN_ID", "r38353711339a1")
	name := CompactRandomName(20)
	if len(name) > 20 || !strings.HasPrefix(name, TestResourcePrefix) {
		t.Fatal("compact identity violates API limit")
	}
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "r38353711339a1")
	if err != nil {
		t.Fatal(err)
	}
	if err = j.validate(OwnedResource{Type: "xcsh_origin_pool", Namespace: "shared", Name: name}); err != nil {
		t.Fatal("compact run identity rejected")
	}
	j.RunID = "r38353711340a1"
	if j.validate(OwnedResource{Type: "xcsh_origin_pool", Namespace: "shared", Name: name}) == nil {
		t.Fatal("another run compact identity accepted")
	}
}
