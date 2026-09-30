// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParallelBuildAggregateFailClosed(t *testing.T) {
	script := extractWorkflowRunStep(t, "_build-test.yml", "aggregate", "Preserve aggregate required-check semantics")
	cases := []struct {
		name, build, vet, race, lint, runLint string
		success                               bool
	}{
		{"all pass", "success", "success", "success", "success", "true", true},
		{"optional lint skipped", "success", "success", "success", "skipped", "false", true},
		{"required lint skipped", "success", "success", "success", "skipped", "true", false},
		{"required lint failed", "success", "success", "success", "failure", "true", false},
		{"build failed", "failure", "success", "success", "success", "true", false},
		{"vet cancelled", "success", "cancelled", "success", "success", "true", false},
		{"race skipped", "success", "success", "skipped", "success", "true", false},
		{"missing race", "success", "success", "", "success", "true", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tmp := t.TempDir()
			outputFile := filepath.Join(tmp, "output")
			output, err := runWorkflowScript(tmp, script, []string{"GITHUB_OUTPUT=" + outputFile, "BUILD_RESULT=" + tc.build, "VET_RESULT=" + tc.vet, "RACE_RESULT=" + tc.race, "LINT_RESULT=" + tc.lint, "RUN_LINT=" + tc.runLint})
			if (err == nil) != tc.success {
				t.Fatalf("success=%v want=%v: %s", err == nil, tc.success, output)
			}
			body, err := os.ReadFile(outputFile)
			if err != nil {
				t.Fatal(err)
			}
			expected := "success=false"
			if tc.success {
				expected = "success=true"
			}
			if strings.TrimSpace(string(body)) != expected {
				t.Fatalf("aggregate output=%s want=%s", body, expected)
			}
		})
	}
}
