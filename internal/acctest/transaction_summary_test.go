// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOnMergeTransactionSummary(t *testing.T) {
	script := extractWorkflowRunStep(t, "on-merge.yml", "summary", "Print summary")
	type summaryCase struct {
		name, dryRun, createPR, release, state, pr, tag, receipt string
		wantSuccess                                              bool
	}
	cases := []summaryCase{
		{"dry run no publication", "true", "false", "false", "success", "skipped", "skipped", "skipped", true},
		{"dry run regeneration", "true", "true", "false", "success", "skipped", "skipped", "skipped", true},
		{"dry run release", "true", "false", "true", "success", "skipped", "skipped", "skipped", true},
		{"normal no publication", "false", "false", "false", "success", "skipped", "skipped", "skipped", true},
		{"normal regeneration", "false", "true", "false", "success", "success", "skipped", "skipped", true},
		{"normal release", "false", "false", "true", "success", "skipped", "success", "success", true},
		{"push release", "", "false", "true", "success", "skipped", "success", "skipped", true},
	}
	for _, result := range []string{"", "failure", "cancelled", "skipped"} {
		cases = append(cases,
			summaryCase{"normal required PR " + result, "false", "true", "false", "success", result, "skipped", "skipped", false},
			summaryCase{"normal required release " + result, "false", "false", "true", "success", "skipped", result, "skipped", false})
		for _, dryRun := range []string{"false", "true"} {
			cases = append(cases, summaryCase{"classifier " + dryRun + " " + result, dryRun, "false", "false", result, "skipped", "skipped", "skipped", false})
		}
	}
	for _, result := range []string{"", "failure", "cancelled", "success"} {
		cases = append(cases,
			summaryCase{"dry run PR " + result, "true", "true", "false", "success", result, "skipped", "skipped", false},
			summaryCase{"dry run tag " + result, "true", "false", "true", "success", "skipped", result, "skipped", false},
			summaryCase{"dry run receipt " + result, "true", "false", "false", "success", "skipped", "skipped", result, false})
	}
	for _, result := range []string{"failure", "cancelled"} {
		cases = append(cases, summaryCase{"normal receipt " + result, "false", "false", "true", "success", "skipped", "success", result, false})
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tmp := t.TempDir()
			summary := filepath.Join(tmp, "summary.md")
			output, err := runWorkflowScript(tmp, script, []string{
				"GITHUB_STEP_SUMMARY=" + summary, "DRY_RUN=" + tc.dryRun,
				"SHOULD_PROCESS=true", "IS_REGENERATION_COMMIT=false", "NEEDS_REGENERATION=true",
				"SPECS_CHANGED=false", "CODE_CHANGED=true", "TOOLS_CHANGED=false", "DOCS_CHANGED=false",
				"BUILD_TEST_RESULT=success", "PROVIDER_RESULT=success", "DOCS_RESULT=success",
				"GENERATION_ARTIFACT=combined-generation-1-1", "GENERATION_SOURCE_SHA=" + strings.Repeat("a", 40),
				"GENERATION_DIGEST=sha256:" + strings.Repeat("b", 64),
				"GENERATION_STATE_RESULT=" + tc.state, "CREATE_PR_REQUIRED=" + tc.createPR,
				"RELEASE_REQUIRED=" + tc.release, "PR_RESULT=" + tc.pr,
				"TAG_RELEASE_RESULT=" + tc.tag, "DELIVERY_RECEIPT_RESULT=" + tc.receipt,
			})
			if (err == nil) != tc.wantSuccess {
				t.Fatalf("success=%v want=%v: %s", err == nil, tc.wantSuccess, output)
			}
			if tc.wantSuccess {
				contents, readErr := os.ReadFile(summary)
				if readErr != nil {
					t.Fatal(readErr)
				}
				for _, row := range []string{"| Regeneration PR Required | " + tc.createPR + " |", "| Release Required | " + tc.release + " |"} {
					if !strings.Contains(string(contents), row) {
						t.Errorf("summary omits classifier decision %q", row)
					}
				}
				if tc.dryRun == "true" && !strings.Contains(string(contents), "| Dry Run | true |") {
					t.Error("summary omits dry-run mode")
				}
			}
		})
	}
}
