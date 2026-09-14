// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"errors"
	"strings"
	"testing"
)

func TestResponseOperationPrerequisiteDiagnosticRecognizesImmutableMauriceContract(t *testing.T) {
	prerequisites := []responseOperationPrerequisite{{
		ID:              "maurice_config_cardinality_exactly_one",
		Resource:        "maurice_config",
		Exactly:         1,
		Enforcement:     "server",
		Availability:    "external_tenant_prerequisite",
		Reason:          "The tenant must contain exactly one maurice_config object before the platform can issue a Customer Edge image download URL.",
		SourceKind:      "runtime_api_error",
		SourceOperation: "ves.io.schema.registration.CustomAPI.GetImageDownloadUrl",
		SourceImmutable: true,
	}}
	title, detail, matched := responseOperationPrerequisiteDiagnostic(errors.New("cannot create dow"+"load url, err: number of maurice_config object is not one"), prerequisites)
	if !matched || title != "External tenant prerequisite unavailable" {
		t.Fatalf("diagnostic = (%q, %q, %t), want matched external prerequisite", title, detail, matched)
	}
	for _, want := range []string{"maurice_config_cardinality_exactly_one", "exactly one maurice_config", "cannot create or repair"} {
		if !strings.Contains(detail, want) {
			t.Fatalf("detail = %q, want %q", detail, want)
		}
	}
}

func TestResponseOperationPrerequisiteDiagnosticLeavesUnrelatedErrorsUntouched(t *testing.T) {
	prerequisites := []responseOperationPrerequisite{{
		ID:              "maurice_config_cardinality_exactly_one",
		Resource:        "maurice_config",
		Exactly:         1,
		Enforcement:     "server",
		Availability:    "external_tenant_prerequisite",
		Reason:          "tenant prerequisite",
		SourceKind:      "runtime_api_error",
		SourceOperation: "ves.io.schema.registration.CustomAPI.GetImageDownloadUrl",
		SourceImmutable: true,
	}}
	if title, detail, matched := responseOperationPrerequisiteDiagnostic(errors.New("connection refused"), prerequisites); matched || title != "" || detail != "" {
		t.Fatalf("diagnostic = (%q, %q, %t), want no match", title, detail, matched)
	}
}

func TestResponseOperationPrerequisiteDiagnosticRejectsMutableOrUnprovenContract(t *testing.T) {
	prerequisite := responseOperationPrerequisite{
		ID:              "maurice_config_cardinality_exactly_one",
		Resource:        "maurice_config",
		Exactly:         1,
		Enforcement:     "provider",
		Availability:    "external_tenant_prerequisite",
		Reason:          "tenant prerequisite",
		SourceKind:      "runtime_api_error",
		SourceOperation: "ves.io.schema.registration.CustomAPI.GetImageDownloadUrl",
		SourceImmutable: true,
	}
	if _, _, matched := responseOperationPrerequisiteDiagnostic(errors.New(mauriceConfigCardinalityError), []responseOperationPrerequisite{prerequisite}); matched {
		t.Fatal("mutable contract must not produce a prerequisite diagnostic")
	}
}
