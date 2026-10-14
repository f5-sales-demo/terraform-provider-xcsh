// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"encoding/json"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
	"os"
	"sort"
	"testing"
)

func TestFixtureSchemaAudit(t *testing.T) {
	path := os.Getenv("XCSH_ACC_SCHEMA")
	if path == "" {
		t.Skip("installed-schema fixture audit requires XCSH_ACC_SCHEMA")
	}
	diagnostics, err := acctest.AuditAcceptanceFixtures(path)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(diagnostics, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if report := os.Getenv("XCSH_ACC_FIXTURE_REPORT"); report != "" {
		if err = os.WriteFile(report, append(data, '\n'), 0600); err != nil {
			t.Fatal(err)
		}
	}
	counts := map[string]int{}
	for _, d := range diagnostics {
		counts[d.Category]++
	}
	names := []string{}
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		t.Logf("%s: %d", name, counts[name])
	}
	if len(diagnostics) > 0 {
		t.Fatalf("%d schema diagnostics; inspect structured fixture audit", len(diagnostics))
	}
}
