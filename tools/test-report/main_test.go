// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestClassifyFailureFingerprint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		output     string
		wantClass  FailureClass
		wantSum    string
		wantStatus int
		wantSource string
	}{
		{
			name:       "missing required argument",
			output:     "Error: Missing required argument\n  on certificate_resource_test.go line 417\nsecret detail",
			wantClass:  FailureClassFrameworkDiagnostic,
			wantSum:    "Missing required argument",
			wantSource: "certificate_resource_test.go:417",
		},
		{
			name:       "unsupported argument",
			output:     "Error: Unsupported argument\n  on fast_acl_resource_test.go line 93",
			wantClass:  FailureClassFrameworkDiagnostic,
			wantSum:    "Unsupported argument",
			wantSource: "fast_acl_resource_test.go:93",
		},
		{
			name:      "inconsistent result",
			output:    "Error: Provider produced inconsistent result after apply",
			wantClass: FailureClassFrameworkDiagnostic,
			wantSum:   "Provider produced inconsistent result after apply",
		},
		{
			name:      "invalid result object",
			output:    "Error: Provider returned invalid result object after apply",
			wantClass: FailureClassFrameworkDiagnostic,
			wantSum:   "Provider returned invalid result object",
		},
		{
			name:       "rate limit",
			output:     "request failed: HTTP response error StatusCode: 429, body secret\nresource_test.go:51",
			wantClass:  FailureClassRateLimit,
			wantStatus: 429,
			wantSource: "resource_test.go:51",
		},
		{
			name:       "indented go test location",
			output:     "    origin_pool_resource_test.go:144: Provider produced inconsistent result after apply",
			wantClass:  FailureClassFrameworkDiagnostic,
			wantSum:    "Provider produced inconsistent result after apply",
			wantSource: "origin_pool_resource_test.go:144",
		},
		{
			name:       "server error",
			output:     "request failed with HTTP 503 Service Unavailable (resource: https://private.invalid/id) (operation: DELETE)",
			wantClass:  FailureClassServerError,
			wantSum:    "Delete request failed",
			wantStatus: 503,
		},
		{
			name:       "create operation",
			output:     "request failed with status: 500 (resource: private-id) (operation: POST)",
			wantClass:  FailureClassServerError,
			wantSum:    "Create request failed",
			wantStatus: 500,
		},
		{
			name:       "http error",
			output:     "request failed: status code: 403 body={private}",
			wantClass:  FailureClassHTTPError,
			wantStatus: 403,
		},
		{
			name:      "assertion",
			output:    "expected value to equal wanted value, got something else",
			wantClass: FailureClassAssertion,
		},
		{
			name:      "timeout",
			output:    "context deadline exceeded while awaiting response",
			wantClass: FailureClassTimeout,
		},
		{
			name:      "connection",
			output:    "dial tcp: connection refused",
			wantClass: FailureClassConnection,
		},
		{
			name:      "other framework diagnostic",
			output:    "Error: Invalid Configuration for Read-Only Attribute",
			wantClass: FailureClassFrameworkDiagnostic,
		},
		{
			name:      "unclassified",
			output:    "opaque failure",
			wantClass: FailureClassUnclassified,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := classifyFailureFingerprint(tc.output)
			if got.Class != tc.wantClass || got.Summary != tc.wantSum ||
				got.HTTPStatus != tc.wantStatus || got.SourceLocation != tc.wantSource {
				t.Fatalf("classifyFailureFingerprint() = %#v", got)
			}
		})
	}
}

func TestTenantSafeReportOmitsSensitiveFailureAndSkipOutput(t *testing.T) {
	secretValues := []string{
		"token-secret-123",
		"https://private-tenant.example.invalid/api/config/namespaces/private-ns/resource/private-id",
		"private-tenant",
		"private-id",
		`{"state":{"password":"private-material"}}`,
	}
	input := strings.Join([]string{
		goTestEvent("run", "TestAccPrivateFailure", ""),
		goTestEvent("output", "TestAccPrivateFailure", "    certificate_resource_test.go:417: request failed with status: 500 (operation: POST) "+strings.Join(secretValues, " ")+"\n"),
		goTestEvent("fail", "TestAccPrivateFailure", ""),
		goTestEvent("run", "TestAccPrivateSkip", ""),
		goTestEvent("output", "TestAccPrivateSkip", "Skipping tenant "+strings.Join(secretValues, " ")+"\n"),
		goTestEvent("skip", "TestAccPrivateSkip", ""),
	}, "\n")

	report := parseTestOutputReader(strings.NewReader(input), 10, true)
	if report.TotalFailed != 1 || len(report.FailedTests) != 1 {
		t.Fatalf("unexpected report totals: %#v", report)
	}
	if report.FailedTests[0].FailureFingerprint == nil {
		t.Fatal("failed test has no fingerprint")
	}
	if got := report.FailedTests[0].FailureFingerprint.Summary; got != "Create request failed" {
		t.Fatalf("tenant-safe operation summary = %q", got)
	}

	jsonOutput, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	assertTenantSafe(t, string(jsonOutput), secretValues)

	markdownFile, err := os.CreateTemp(t.TempDir(), "report-*.md")
	if err != nil {
		t.Fatal(err)
	}
	outputMarkdown(markdownFile, report, true)
	if err := markdownFile.Close(); err != nil {
		t.Fatal(err)
	}
	markdownOutput, err := os.ReadFile(markdownFile.Name()) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	assertTenantSafe(t, string(markdownOutput), secretValues)
}

func TestFingerprintAggregatesAreDeterministic(t *testing.T) {
	failed := []TestResult{
		{Name: "TestAccZulu", FailureFingerprint: &FailureFingerprint{Class: FailureClassHTTPError, HTTPStatus: 400}},
		{Name: "TestAccAlpha", FailureFingerprint: &FailureFingerprint{Class: FailureClassFrameworkDiagnostic, Summary: "Missing required argument"}},
		{Name: "TestAccBravo", FailureFingerprint: &FailureFingerprint{Class: FailureClassHTTPError, HTTPStatus: 400}},
	}

	first := aggregateFailureFingerprints(failed)
	second := aggregateFailureFingerprints([]TestResult{failed[2], failed[0], failed[1]})
	firstJSON, _ := json.Marshal(first)
	secondJSON, _ := json.Marshal(second)
	if !bytes.Equal(firstJSON, secondJSON) {
		t.Fatalf("aggregate order depends on input:\n%s\n%s", firstJSON, secondJSON)
	}
	if got := first[1].Tests; len(got) != 2 || got[0] != "TestAccBravo" || got[1] != "TestAccZulu" {
		t.Fatalf("test names are not sorted: %v", got)
	}
}

func goTestEvent(action, testName, output string) string {
	event := TestEvent{Action: action, Package: "example.invalid/provider", Test: testName, Output: output}
	encoded, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	return string(encoded)
}

func assertTenantSafe(t *testing.T, output string, secrets []string) {
	t.Helper()
	if strings.Contains(output, "failure_output") || strings.Contains(output, "skip_reason\"") {
		t.Fatalf("tenant-safe output retains raw-output fields:\n%s", output)
	}
	for _, secret := range secrets {
		if strings.Contains(output, secret) {
			t.Fatalf("tenant-safe output retains sensitive fixture %q:\n%s", secret, output)
		}
	}
}
