// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAcceptanceGateRejectsUntestedSkipEvidence(t *testing.T) {
	script := extractWorkflowRunStep(t, "acc-tests.yml", "real-api-tests", "Run real API tests")
	start := strings.Index(script, "          jq -e '")
	if start < 0 {
		start = strings.Index(script, "jq -e '")
	}
	if start < 0 {
		t.Fatal("report gate missing")
	}
	expression := script[start+strings.Index(script[start:], "'")+1:]
	expression = expression[:strings.Index(expression, "' test-reports/real-tests.json")]
	report := `{"stream_completeness":{"complete":true},"total_tests":1,"total_passed":0,"total_failed":0,"failed_tests":[],"tests":[{"status":"skip","prerequisites":[{"outcome":"tested","method":"POST","resource_type":"xcsh_namespace"}]}]}`
	cmd := exec.Command("jq", "-e", expression)
	cmd.Stdin = strings.NewReader(report)
	if err := cmd.Run(); err == nil {
		t.Fatal("a probe that did not prove denial justified a skip")
	}
}
