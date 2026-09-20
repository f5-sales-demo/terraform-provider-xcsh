// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"errors"
	"strings"
	"testing"
)

func TestResponseOperationPrerequisiteDiagnosticRecognizesImmutableMauriceContract(t *testing.T) {
	prerequisites := []responseOperationPrerequisite{verifiedImageLookupPrerequisite()}
	title, detail, matched := responseOperationPrerequisiteDiagnostic(errors.New(mauriceConfigCardinalityError), prerequisites)
	if !matched || title != "Image lookup unresolved" {
		t.Fatalf("diagnostic = (%q, %q, %t), want unresolved lookup", title, detail, matched)
	}
	for _, want := range []string{"exactly one", "unknown", "test-receipt.json", strings.Repeat("c", 64)} {
		if !strings.Contains(detail, want) {
			t.Fatalf("detail is missing %q", want)
		}
	}
}

func TestResponseOperationPrerequisiteDiagnosticLeavesUnrelatedErrorsUntouched(t *testing.T) {
	for _, err := range []error{nil, errors.New("connection refused")} {
		if title, detail, matched := responseOperationPrerequisiteDiagnostic(err, []responseOperationPrerequisite{verifiedImageLookupPrerequisite()}); matched || title != "" || detail != "" {
			t.Fatal("unrelated errors must not become lookup diagnostics")
		}
	}
}

func TestResponseOperationPrerequisiteDiagnosticRejectsMutableOrUnprovenContract(t *testing.T) {
	for _, mutate := range []func(*responseOperationPrerequisite){
		func(p *responseOperationPrerequisite) { p.Enforcement = "provider" },
		func(p *responseOperationPrerequisite) { p.LookupScope = "tenant" },
		func(p *responseOperationPrerequisite) { p.LookupCount = "0" },
		func(p *responseOperationPrerequisite) { p.SourceImmutable = false },
		func(p *responseOperationPrerequisite) { p.SourceCommit = "" },
		func(p *responseOperationPrerequisite) { p.SpecSHA256 = strings.Repeat("x", 64) },
		func(p *responseOperationPrerequisite) { p.ReceiptSHA256 = "" },
		func(p *responseOperationPrerequisite) { p.ReceiptPath = "config/evidence/../../unverified.json" },
	} {
		prerequisite := verifiedImageLookupPrerequisite()
		mutate(&prerequisite)
		if _, _, matched := responseOperationPrerequisiteDiagnostic(errors.New(mauriceConfigCardinalityError), []responseOperationPrerequisite{prerequisite}); matched {
			t.Fatal("mutable or unproven contract must not produce a lookup diagnostic")
		}
	}
}
