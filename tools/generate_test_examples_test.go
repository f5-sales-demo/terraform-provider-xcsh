// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCleanConfigUsesValidConstrainedValues(t *testing.T) {
	input := `resource "xcsh_healthcheck" "test" {
  healthy_threshold   = %[2]d
  unhealthy_threshold = %[3]d
  jitter_percent      = %[4]d
}

resource "xcsh_rate_limiter" "test" {
  limits {
    unit = %[2]q
  }
}`

	got := cleanConfig("", input)
	for _, want := range []string{
		"healthy_threshold   = 3",
		"unhealthy_threshold = 2",
		"jitter_percent      = 30",
		`unit = "MINUTE"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("cleanConfig() did not contain %q:\n%s", want, got)
		}
	}
}

func TestValidateExampleContentRejectsMeasuredMutations(t *testing.T) {
	valid := addProviderRequirements(`resource "xcsh_healthcheck" "test" {
  healthy_threshold   = 3
  unhealthy_threshold = 2
  jitter_percent      = 30
}`)
	if err := validateExampleContent(valid); err != nil {
		t.Fatalf("valid example rejected: %v", err)
	}

	tests := map[string]string{
		"health threshold outside schema range": strings.Replace(valid, "healthy_threshold   = 3", "healthy_threshold   = 443", 1),
		"jitter outside percentage range":       strings.Replace(valid, "jitter_percent      = 30", "jitter_percent      = 443", 1),
		"wrong provider source":                 strings.Replace(valid, xcshProviderSource, "hashicorp/xcsh", 1),
		"wrong provider version constraint":     strings.Replace(valid, xcshVersionConstraint, ">= 0.0.1", 1),
	}
	for name, mutated := range tests {
		t.Run(name, func(t *testing.T) {
			if err := validateExampleContent(mutated); err == nil {
				t.Fatal("mutated example was accepted")
			}
		})
	}

	rateLimiter := addProviderRequirements(`resource "xcsh_rate_limiter" "test" {
  limits {
    unit = "MINUTE"
  }
}`)
	mutatedUnit := strings.Replace(rateLimiter, `unit = "MINUTE"`, `unit = "example-value"`, 1)
	if err := validateExampleContent(mutatedUnit); err == nil {
		t.Fatal("invalid rate limiter unit was accepted")
	}
}

func TestTimeProviderIsPinnedWhenRequired(t *testing.T) {
	content := addProviderRequirements(`resource "time_sleep" "wait" {
  create_duration = "5s"
}`)
	if err := validateExampleContent(content); err != nil {
		t.Fatalf("time_sleep provider requirements rejected: %v", err)
	}

	mutated := strings.Replace(content, `version = "= 0.13.1"`, `version = ">= 0.9.0"`, 1)
	if err := validateExampleContent(mutated); err == nil {
		t.Fatal("unpinned time provider was accepted")
	}
}

func TestRenderExamplesHasExpectedSelection(t *testing.T) {
	testDir := providerTestDirectory(t)
	examples, err := renderExamples(testDir, filepath.Join(t.TempDir(), "examples"))
	if err != nil {
		t.Fatalf("renderExamples() error: %v", err)
	}
	if err := validateRenderedExamples(examples); err != nil {
		t.Fatalf("rendered examples failed contract validation: %v", err)
	}

	wantCounts := map[string]int{
		"xcsh_http_loadbalancer":         14,
		"xcsh_tcp_loadbalancer":          6,
		"xcsh_healthcheck":               13,
		"xcsh_app_firewall":              11,
		"xcsh_origin_pool":               7,
		"xcsh_rate_limiter":              7,
		"xcsh_service_policy":            4,
		"xcsh_user_identification":       8,
		"xcsh_malicious_user_mitigation": 8,
	}
	gotCounts := make(map[string]int)
	gotPaths := make(map[string]struct{})
	for _, example := range examples {
		resource := filepath.Base(filepath.Dir(example.Path))
		gotCounts[resource]++
		if _, exists := gotPaths[example.Path]; exists {
			t.Fatalf("duplicate generated path %s", example.Path)
		}
		gotPaths[example.Path] = struct{}{}
		if filepath.Base(example.Path) == "basic.tf" || filepath.Base(example.Path) == "basic-system.tf" {
			t.Fatalf("basic example must remain canonical, got named path %s", example.Path)
		}
	}
	for path := range gotPaths {
		if strings.HasSuffix(path, "xcsh_http_loadbalancer/conflict-protocol.tf") {
			t.Errorf("negative acceptance fixture was published as a verified example: %s", path)
		}
	}
	for resource, want := range wantCounts {
		if got := gotCounts[resource]; got != want {
			t.Errorf("%s generated count = %d, want %d", resource, got, want)
		}
	}
	if len(gotCounts) != len(wantCounts) {
		t.Errorf("generated resource count = %d, want %d: %#v", len(gotCounts), len(wantCounts), gotCounts)
	}

	for _, suffix := range []string{
		"xcsh_healthcheck/thresholds.tf",
		"xcsh_healthcheck/with-jitter.tf",
		"xcsh_rate_limiter/unit.tf",
		"xcsh_malicious_user_mitigation/with-labels.tf",
	} {
		found := false
		for path := range gotPaths {
			if strings.HasSuffix(path, suffix) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected generated selection %s was missing", suffix)
		}
	}
}

func TestNegativeConfigHelpersAndStaleExamplePruning(t *testing.T) {
	source := []byte(`package provider
func TestAccExample(t *testing.T) {
  resource.Test(t, resource.TestCase{Steps: []resource.TestStep{{
    Config: testAccWidgetConfig_valid("ok"),
  }, {
    Config: testAccWidgetConfig_conflict("bad"),
    ExpectError: regexp.MustCompile("conflict"),
  }}})
}
func testAccWidgetConfig_valid(name string) string { return fmt.Sprintf(` + "`resource \"xcsh_widget\" \"test\" { name = %[1]q }`" + `, name) }
func testAccWidgetConfig_conflict(name string) string { return fmt.Sprintf(` + "`resource \"xcsh_widget\" \"test\" { name = %[1]q }`" + `, name) }
`)
	excluded, err := negativeConfigHelpers(source)
	if err != nil {
		t.Fatalf("negativeConfigHelpers() error: %v", err)
	}
	if _, ok := excluded["testAccWidgetConfig_conflict"]; !ok {
		t.Fatal("negative helper was not excluded")
	}
	if _, ok := excluded["testAccWidgetConfig_valid"]; ok {
		t.Fatal("positive helper was excluded")
	}

	root := t.TempDir()
	directory := filepath.Join(root, "xcsh_http_loadbalancer")
	if err := os.MkdirAll(directory, 0755); err != nil {
		t.Fatal(err)
	}
	stale := filepath.Join(directory, "conflict-protocol.tf")
	manual := filepath.Join(directory, "manual.tf")
	if err := os.WriteFile(stale, []byte("# heading\n"+generatedExampleMarker+"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(manual, []byte("# maintained manually\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := pruneStaleGeneratedExamples(root, nil); err != nil {
		t.Fatalf("pruneStaleGeneratedExamples() error: %v", err)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("stale generated example still exists: %v", err)
	}
	if _, err := os.Stat(manual); err != nil {
		t.Fatalf("manual example was removed: %v", err)
	}
}

func providerTestDirectory(t *testing.T) string {
	t.Helper()
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}
	for _, candidate := range []string{
		filepath.Join(workingDirectory, "internal", "provider"),
		filepath.Join(workingDirectory, "..", "internal", "provider"),
	} {
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
	}
	t.Fatalf("cannot locate internal/provider from %s", workingDirectory)
	return ""
}
