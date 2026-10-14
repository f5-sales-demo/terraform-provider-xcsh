// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Evidence categories are fixed public values even when the free-text skip
// message does not happen to contain a historical reporter keyword.
func TestPrerequisiteEvidenceClassifiesSkipsWithoutPayloads(t *testing.T) {
	for _, tc := range []struct{ outcome, typ, want string }{
		{"quota_exhausted", "xcsh_container_registry", "QUOTA_EXHAUSTED"},
		{"quota_unconfigured", "xcsh_forwarding_class", "QUOTA_UNCONFIGURED"},
		{"operation_denied", "xcsh_bgp", "OPERATION_DENIED"},
		{"infrastructure_absent", "virtual_site", "INFRASTRUCTURE"},
		{"external_credential_missing", "vcs", "MISSING_CREDENTIALS"},
		{"non_disposable_scope", "tenant_configuration", "ENVIRONMENT"},
		{"cloud_appliance_out_of_scope", "xcsh_nfv_service", "OUT_OF_SCOPE"},
	} {
		t.Run(tc.outcome, func(t *testing.T) {
			stream := fmt.Sprintf(`{"Action":"start","Package":"fixture"}
{"Action":"run","Package":"fixture","Test":"TestAccFixture"}
{"Action":"output","Package":"fixture","Test":"TestAccFixture","Output":"[PREREQUISITE:%s method=GET type=%s status=403]\nSECRET_SENTINEL\n"}
{"Action":"skip","Package":"fixture","Test":"TestAccFixture"}
{"Action":"pass","Package":"fixture"}
`, tc.outcome, tc.typ)
			report := parseTestOutputReader(strings.NewReader(stream), 10, true)

			private := t.TempDir()
			for _, format := range []string{"json", "markdown"} {
				file, err := os.Create(filepath.Join(private, format))
				if err != nil {
					t.Fatal(err)
				}
				if format == "json" {
					outputJSON(file, report)
				} else {
					outputMarkdown(file, report, true)
				}
				if err = file.Close(); err != nil {
					t.Fatal(err)
				}
				data, err := os.ReadFile(filepath.Join(private, format))
				if err != nil {
					t.Fatal(err)
				}
				if strings.Contains(string(data), "SECRET_SENTINEL") {
					t.Fatal("skip output retained private payload")
				}
				if !strings.Contains(string(data), tc.want) {
					t.Fatal("skip classification missing from report")
				}
			}
			if len(report.Tests) != 1 || string(report.Tests[0].SkipReasonType) != tc.want {
				t.Fatalf("evidence did not classify skip: %+v", report.Tests)
			}
		})
	}
}
