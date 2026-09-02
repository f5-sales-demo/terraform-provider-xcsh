// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateSMSv2ContractConstantsIsDeterministic(t *testing.T) {
	t.Parallel()
	specDir, outputDir := t.TempDir(), t.TempDir()
	contract := strings.NewReplacer(
		`"availability":"evidence_backed"`, `"availability":"schema_only"`,
		`"aws_ce_create":"available","runtime_status":"available","tgw_connect":"available"`, `"aws_ce_create":"unavailable","runtime_status":"unavailable","tgw_connect":"unavailable"`,
		`"unavailable_capabilities":[]`, `"unavailable_capabilities":["aws_ce_create","runtime_status","tgw_connect"]`,
		`"availability":"available","complete":true`, `"availability":"unavailable","complete":false`,
	).Replace(syntheticSMSv2V5Contract)
	manifest := `{"contract_id":"f5xc-ce-automation/v2","contract_version":"5.0.0","release":{"tag":"v5.0.1","commit":"3a647f1bf0c2447a71750c69136fab96fb073902"}}`
	if err := os.WriteFile(filepath.Join(specDir, "smsv2-contract.json"), []byte(contract), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(specDir, "smsv2-contract-manifest.json"), []byte(manifest), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := GenerateSMSv2ContractConstants(specDir, outputDir); err != nil {
		t.Fatal(err)
	}
	first, err := os.ReadFile(filepath.Join(outputDir, "smsv2_contract_generated.go"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := GenerateSMSv2ContractConstants(specDir, outputDir); err != nil {
		t.Fatal(err)
	}
	second, err := os.ReadFile(filepath.Join(outputDir, "smsv2_contract_generated.go"))
	if err != nil {
		t.Fatal(err)
	}
	if string(first) != string(second) {
		t.Fatal("SMSv2 contract generation is not deterministic")
	}
	for _, want := range []string{"f5xc-ce-automation/v2", "v5.0.1", `"aws_ce_create": "unavailable"`, "3a647f1bf0c2447a71750c69136fab96fb073902", "f5xc-smsv2-aws-tgw-telemetry/v1"} {
		if !strings.Contains(string(first), want) {
			t.Errorf("generated constants missing %q", want)
		}
	}
}
