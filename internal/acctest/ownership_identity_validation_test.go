// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"encoding/json"
	"os"
	"testing"
)

func TestLoadedJournalRejectsUnverifiableIdentities(t *testing.T) {
	for _, tc := range []struct {
		name     string
		resource OwnedResource
	}{
		{"raw uid", OwnedResource{Type: "xcsh_origin_pool", Namespace: "system", Name: "tf-acc-test-runone-pool", UID: "SECRET_SENTINEL", Status: "created"}},
		{"unknown handler", OwnedResource{Type: "xcsh_" + "unknown", Namespace: "system", Name: "tf-acc-test-runone-pool", Status: "intent"}},
		{"created without uid", OwnedResource{Type: "xcsh_origin_pool", Namespace: "system", Name: "tf-acc-test-runone-pool", Status: "created"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			path := t.TempDir() + "/journal.json"
			data, err := json.Marshal(OwnershipJournal{Version: 1, RunID: "runone", Resources: []OwnedResource{tc.resource}})
			if err != nil {
				t.Fatal(err)
			}
			if os.WriteFile(path, data, 0600) != nil {
				t.Fatal("journal write failed")
			}
			if _, err := LoadOwnershipJournal(path, "runone"); err == nil {
				t.Fatal("unverifiable identity accepted")
			}
		})
	}
}
