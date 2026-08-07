package acctest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLintCoverageIncludesTools(t *testing.T) {
	// Find the project root
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	rootDir := filepath.Dir(filepath.Dir(cwd))

	testCases := []struct {
		file     string
		mustHave []string
	}{
		{
			file: ".github/workflows/_build-test.yml",
			mustHave: []string{
				"mapfile -t pkgs < <(go list -f '{{.Dir}}' ./internal/... . ./tools/... | grep -vE '/internal/(provider|client)$')",
			},
		},
		{
			file: "Makefile",
			mustHave: []string{
				"$(GOLINT) run --timeout=5m ./internal/... . ./tools/...",
			},
		},
		{
			file: "scripts/pre-commit-local.sh",
			mustHave: []string{
				"golangci-lint run --timeout=5m ./internal/... . ./tools/...",
			},
		},
		{
			file: "scripts/lint-generated-preview.sh",
			mustHave: []string{
				"golangci-lint run --timeout=5m ./internal/... . ./tools/...",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.file, func(t *testing.T) {
			path := filepath.Join(rootDir, tc.file)
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", tc.file, err)
			}

			strContent := string(content)
			for _, mustHave := range tc.mustHave {
				if !strings.Contains(strContent, mustHave) {
					t.Errorf("file %s is missing required snippet: %s", tc.file, mustHave)
				}
			}
		})
	}
}
