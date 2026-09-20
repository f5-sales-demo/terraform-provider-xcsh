package provider

import (
	"errors"
	"strings"
	"testing"
)

func TestImageLookupDiagnosticDoesNotInventOwnershipOrEchoResponse(t *testing.T) {
	contract := verifiedImageLookupPrerequisite()
	_, detail, matched := responseOperationPrerequisiteDiagnostic(errors.New(mauriceConfigCardinalityError+" synthetic-private-value"), []responseOperationPrerequisite{contract})
	if !matched {
		t.Fatal("expected the unresolved lookup diagnostic")
	}
	for _, forbidden := range []string{"synthetic-private-value", "tenant-owned", "external tenant prerequisite"} {
		if strings.Contains(strings.ToLower(detail), forbidden) {
			t.Fatalf("diagnostic contains forbidden claim or response data: %q", forbidden)
		}
	}
	if !strings.Contains(detail, "unknown") {
		t.Fatal("diagnostic must preserve uncertainty")
	}
}

func verifiedImageLookupPrerequisite() responseOperationPrerequisite {
	return responseOperationPrerequisite{
		ID: "maurice_config_cardinality_exactly_one", Resource: "maurice_config", Exactly: 1,
		Enforcement: "server", Availability: "unresolved_server_lookup",
		Reason:     "The server reported a lookup cardinality mismatch.",
		SourceKind: "runtime_api_error", SourceOperation: "ves.io.schema.registration.CustomAPI.GetImageDownloadUrl", SourceImmutable: true,
		LookupScope: "unknown", LookupCount: "unknown",
		SourceCommit: strings.Repeat("a", 40), SpecSHA256: strings.Repeat("b", 64),
		ReceiptPath: "config/evidence/test-receipt.json", ReceiptSHA256: strings.Repeat("c", 64),
	}
}

func TestImageLookupDiagnosticRejectsObsoleteTenantAssumption(t *testing.T) {
	contract := responseOperationPrerequisite{
		ID: "maurice_config_cardinality_exactly_one", Resource: "maurice_config", Exactly: 1,
		Enforcement: "server", Availability: "external_tenant_prerequisite",
		SourceKind: "runtime_api_error", SourceOperation: "ves.io.schema.registration.CustomAPI.GetImageDownloadUrl", SourceImmutable: true,
	}
	if _, _, matched := responseOperationPrerequisiteDiagnostic(errors.New(mauriceConfigCardinalityError), []responseOperationPrerequisite{contract}); matched {
		t.Fatal("an inferred tenant prerequisite must not be accepted as evidence")
	}
}
