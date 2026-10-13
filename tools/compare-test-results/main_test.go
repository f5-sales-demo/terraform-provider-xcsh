// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"os"
	"testing"
)

func TestComparisonRejectsIncompleteAndUnaccountedReports(t *testing.T) {
	for _, report := range []string{
		`{"total_tests":1,"total_passed":1,"tests":[{"name":"TestAccFixture","status":"pass"}]}`,
		`{"stream_completeness":{"complete":true},"total_tests":2,"total_passed":2,"tests":[{"name":"TestAccFixture","status":"pass"}]}`,
		`{"stream_completeness":{"complete":true},"total_tests":2,"total_passed":2,"tests":[{"name":"TestAccFixture","status":"pass"},{"name":"TestAccFixture","status":"pass"}]}`,
		`{"stream_completeness":{"complete":true},"total_tests":1,"total_skipped":1,"tests":[{"name":"TestAccFixture","status":"skip"}]}`,
	} {
		path := t.TempDir() + "/report.json"
		if os.WriteFile(path, []byte(report), 0600) != nil {
			t.Fatal("report write failed")
		}
		if _, err := loadReport(path); err == nil {
			t.Error("incomplete or unaccounted comparison input accepted")
		}
	}
}

func TestComparisonDoesNotCallMissingFailedTestSkipped(t *testing.T) {
	mock := &Report{Tests: []TestResult{{Name: "TestAccFixture", Status: "fail"}}}
	report := compare(mock, &Report{})
	if report.Summary.TestNeedsFix != 1 || report.Summary.BothSkipped != 0 {
		t.Fatal("missing failure lost")
	}
}
