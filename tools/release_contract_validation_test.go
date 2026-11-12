package main

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestReleaseContractValidation(t *testing.T) {
	for _, test := range []struct{ name, command, script string }{
		{"smsv2_reviewed_contract", "bash", "test-smsv2-release-validator.sh"},
	} {
		t.Run(test.name, func(t *testing.T) {
			output, err := exec.Command(test.command, filepath.Join("..", "tests", test.script)).CombinedOutput()
			if err != nil {
				t.Fatalf("release validation failed: %v\n%s", err, output)
			}
		})
	}
}
