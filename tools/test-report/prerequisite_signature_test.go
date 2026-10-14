// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import "testing"

func TestReproducedConnectorPrerequisiteSignature(t *testing.T) {
	f := classifyFailureFingerprint("connector node and interface prerequisite unverified SECRET_SENTINEL")
	if f.Class != FailureClassFrameworkDiagnostic || f.Summary != "connector node and interface prerequisite unverified" {
		t.Fatal("reproduced prerequisite failure unclassified")
	}
}
