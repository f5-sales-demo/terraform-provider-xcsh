// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConstitutionAllowsRegeneratedProviderFilesWithGeneratorChange(t *testing.T) {
	scriptPath := filepath.Join("..", "..", "scripts", "check-no-generated-files.sh")
	script, err := os.ReadFile(scriptPath) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	text := string(script)
	for _, pattern := range []string{
		`"^internal/provider/.*_resource\.go$"`,
		`"^internal/provider/.*_data_source\.go$"`,
	} {
		if strings.Count(text, pattern) < 2 {
			t.Fatalf("generated provider pattern %s must be both protected and allowed when its generator changes in check-no-generated-files.sh", pattern)
		}
	}
}
