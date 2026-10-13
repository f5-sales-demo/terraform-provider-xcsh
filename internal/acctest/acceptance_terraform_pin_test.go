// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"os"
	"strings"
	"testing"
)

func TestAcceptanceJobsUseInstalledPinnedTerraform(t *testing.T) {
	data, err := os.ReadFile("../../.github/workflows/acc-tests.yml")
	if err != nil {
		t.Fatal(err)
	}
	for _, job := range []string{"mock-tests", "real-api-tests"} {
		source := string(data)
		start := strings.Index(source, "  "+job+":")
		if start < 0 {
			t.Fatal("job missing")
		}
		body := source[start:]
		if next := strings.Index(body[3:], "\n  # ═"); next >= 0 {
			body = body[:next+3]
		}
		for _, required := range []string{"terraform_version: 1.16.3", "TF_ACC_TERRAFORM_PATH="} {
			if !strings.Contains(body, required) {
				t.Errorf("%s missing %s", job, required)
			}
		}
	}
}
