// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConstitutionAllowsRegeneratedProviderFilesWithGeneratorChange(t *testing.T) {
	workflowPath := filepath.Join("..", "..", ".github", "workflows", "ci.yml")
	workflow, err := os.ReadFile(workflowPath) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	text := string(workflow)
	for _, pattern := range []string{
		`"^internal/provider/.*_resource\.go$"`,
		`"^internal/provider/.*_data_source\.go$"`,
	} {
		if strings.Count(text, pattern) < 2 {
			t.Fatalf("generated provider pattern %s must be both protected and allowed when its generator changes", pattern)
		}
	}
}
