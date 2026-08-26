// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package requiredness

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFixture(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func resourceFixture(required bool) string {
	mode := "Optional: true"
	if required {
		mode = "Required: true"
	}
	return `package provider
import "github.com/hashicorp/terraform-plugin-framework/resource/schema"
func (r *ProbeResource) Schema() {
  _ = schema.Schema{Attributes: map[string]schema.Attribute{
    "guided": schema.StringAttribute{` + mode + `},
    "stable": schema.StringAttribute{Required: true},
  }}
}`
}

func specFixture(minimum bool, explicit bool) string {
	minimumConfig := "null"
	if minimum {
		minimumConfig = `{"required_fields":["spec.guided"]}`
	}
	requiredFor := "false"
	if explicit {
		requiredFor = "true"
	}
	return `{
  "openapi":"3.0.0",
  "info":{"title":"fixture","version":"1"},
  "paths":{},
  "components":{"schemas":{"probeCreateSpecType":{
    "type":"object",
    "x-f5xc-minimum-configuration":` + minimumConfig + `,
    "properties":{"guided":{"type":"string","x-f5xc-required-for":{"create":` + requiredFor + `}}}
  }}}
}`
}

func TestCompareTracesMinimumConfigurationPromotionRemoval(t *testing.T) {
	baselineProvider := filepath.Join(t.TempDir(), "provider")
	candidateProvider := filepath.Join(t.TempDir(), "provider")
	baselineSpecs := t.TempDir()
	candidateSpecs := t.TempDir()
	writeFixture(t, filepath.Join(baselineProvider, "probe_resource.go"), resourceFixture(true))
	writeFixture(t, filepath.Join(candidateProvider, "probe_resource.go"), resourceFixture(false))
	writeFixture(t, filepath.Join(baselineSpecs, "domains", "probe.json"), specFixture(true, false))
	writeFixture(t, filepath.Join(candidateSpecs, "domains", "probe.json"), specFixture(false, false))

	report, err := Compare(baselineProvider, candidateProvider, baselineSpecs, candidateSpecs, Report{
		BaselineProvider: "main@abc",
		BaselineSpec:     "v1",
		CandidateSpec:    "v2",
	})
	if err != nil {
		t.Fatalf("Compare: %v", err)
	}
	if len(report.Transitions) != 1 {
		t.Fatalf("transitions = %v, want one", report.Transitions)
	}
	got := report.Transitions[0]
	if got.Resource != "probe" || got.Attribute != "guided" || got.Reason != RemovedMinimumConfigurationPromotion {
		t.Fatalf("unexpected transition: %+v", got)
	}
}

func TestCompareRejectsUntracedOrStillRequiredTransition(t *testing.T) {
	for name, tc := range map[string]struct {
		baselineMinimum   bool
		baselineExplicit  bool
		candidateExplicit bool
		want              string
	}{
		"not minimum config":        {false, false, false, "absent from baseline minimum configuration"},
		"genuinely required before": {true, true, false, "baseline contract genuinely required it for create"},
		"still required":            {true, false, true, "still requires it for create"},
	} {
		t.Run(name, func(t *testing.T) {
			baselineProvider := filepath.Join(t.TempDir(), "provider")
			candidateProvider := filepath.Join(t.TempDir(), "provider")
			baselineSpecs := t.TempDir()
			candidateSpecs := t.TempDir()
			writeFixture(t, filepath.Join(baselineProvider, "probe_resource.go"), resourceFixture(true))
			writeFixture(t, filepath.Join(candidateProvider, "probe_resource.go"), resourceFixture(false))
			writeFixture(t, filepath.Join(baselineSpecs, "domains", "probe.json"), specFixture(tc.baselineMinimum, tc.baselineExplicit))
			writeFixture(t, filepath.Join(candidateSpecs, "domains", "probe.json"), specFixture(false, tc.candidateExplicit))
			_, err := Compare(baselineProvider, candidateProvider, baselineSpecs, candidateSpecs, Report{})
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error = %v, want substring %q", err, tc.want)
			}
		})
	}
}

func TestCompareClassifiesVerifiedSingleNamespaceDefault(t *testing.T) {
	baselineProvider := filepath.Join(t.TempDir(), "provider")
	candidateProvider := filepath.Join(t.TempDir(), "provider")
	baselineSpecs := t.TempDir()
	candidateSpecs := t.TempDir()
	provider := func(required bool) string {
		mode := "Optional: true"
		if required {
			mode = "Required: true"
		}
		return `package provider
import "github.com/hashicorp/terraform-plugin-framework/resource/schema"
func (r *ProbeResource) Schema() {
  _ = schema.Schema{Attributes: map[string]schema.Attribute{
    "namespace": schema.StringAttribute{` + mode + `},
  }}
}`
	}
	baselineSpec := `{"openapi":"3.0.0","info":{"version":"1"},"paths":{},"components":{"schemas":{"probeCreateSpecType":{"type":"object","properties":{}}}}}`
	candidateSpec := `{"openapi":"3.0.0","info":{"version":"2"},"paths":{},"components":{"schemas":{"viewsprobeCreateSpecType":{"type":"object","x-f5xc-namespace-profile":{"constraint":{"allowed":["system"],"enforced":true}},"properties":{}}}}}`
	writeFixture(t, filepath.Join(baselineProvider, "probe_resource.go"), provider(true))
	writeFixture(t, filepath.Join(candidateProvider, "probe_resource.go"), provider(false))
	writeFixture(t, filepath.Join(baselineSpecs, "domains", "probe.json"), baselineSpec)
	writeFixture(t, filepath.Join(candidateSpecs, "domains", "probe.json"), candidateSpec)

	report, err := Compare(baselineProvider, candidateProvider, baselineSpecs, candidateSpecs, Report{})
	if err != nil {
		t.Fatalf("Compare: %v", err)
	}
	if len(report.Transitions) != 1 || report.Transitions[0].Reason != VerifiedSingleNamespaceDefault {
		t.Fatalf("unexpected transitions: %+v", report.Transitions)
	}
}
