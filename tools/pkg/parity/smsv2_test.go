// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package parity

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestBuildSMSv2MatrixFailsUnclassifiedGap(t *testing.T) {
	legacy := &LegacyManifest{Version: "0.12.2", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 1, Paths: []LegacyField{{Path: "missing", Type: "string", Optional: true}}}
	current := &CurrentManifest{Version: "2.1.225", PathCount: 1, Paths: []CurrentField{{Path: "spec.present", Type: "string", Cardinality: "single"}}, ChoiceGroups: map[string][]string{"choice": {"spec.present"}}}
	matrix, err := BuildSMSv2Matrix(legacy, current)
	if err == nil || len(matrix.Unclassified) != 1 || matrix.Unclassified[0] != "missing" {
		t.Fatalf("expected one unclassified gap, matrix=%+v err=%v", matrix, err)
	}
	found := false
	for _, entry := range matrix.Entries {
		if entry.LegacyPath == "missing" {
			found = entry.OwningRepository != "" && len(entry.RequiredTests) != 0
		}
	}
	if !found {
		t.Fatal("gap lacks ownership and validation requirements")
	}
}

func TestBuildSMSv2MatrixFromTerraformReportsGeneratorGap(t *testing.T) {
	legacy := &LegacyManifest{Version: "0.12.2", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 1, Paths: []LegacyField{{Path: "missing", Type: "string", Optional: true}}}
	current := &CurrentManifest{Version: "2.1.225", PathCount: 1, Paths: []CurrentField{{Path: "spec.missing", Type: "string", Cardinality: "single"}}, ChoiceGroups: map[string][]string{"choice": {"spec.missing"}}}
	matrix, err := BuildSMSv2MatrixFromTerraform(legacy, current, nil)
	if err == nil {
		t.Fatal("missing generated capability must block parity")
	}
	if matrix.Classification["generator_gap"] != 1 || matrix.Entries[0].Current.Generated {
		t.Fatalf("expected an explicit generator gap, got %+v", matrix)
	}
}

func TestBuildSMSv2MatrixFromTerraformUsesGeneratedRequiredness(t *testing.T) {
	legacy := &LegacyManifest{Version: "0.12.2", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 1, Paths: []LegacyField{{
		Path: "enable_upgrade_drain.drain_node_timeout", Type: "int64", Cardinality: "single", Required: true,
	}}}
	current := &CurrentManifest{Version: "2.1.225", PathCount: 1, Paths: []CurrentField{{
		Path: "spec.enable_upgrade_drain.drain_node_timeout", Type: "integer", Cardinality: "single", CreateRequired: true,
	}}, ChoiceGroups: map[string][]string{"choice": {"spec.enable_upgrade_drain.drain_node_timeout"}}}
	attrs := []openapi.TerraformAttribute{{
		TfsdkTag: "enable_upgrade_drain", JsonName: "enable_upgrade_drain", IsBlock: true, NestedBlockType: "single", Optional: true,
		NestedAttributes: []openapi.TerraformAttribute{{TfsdkTag: "drain_node_timeout", JsonName: "drain_node_timeout", Type: "int64", Optional: true}},
	}}
	matrix, err := BuildSMSv2MatrixFromTerraform(legacy, current, attrs)
	if err != nil {
		t.Fatal(err)
	}
	if len(matrix.Entries) != 1 || matrix.Entries[0].Current == nil {
		t.Fatalf("unexpected matrix: %+v", matrix)
	}
	if matrix.Entries[0].Current.Required || !matrix.Entries[0].Current.Optional {
		t.Fatalf("matrix reported manifest requiredness instead of generated schema: %+v", matrix.Entries[0].Current)
	}

	attrs[0].NestedAttributes[0].CreateRequired = true
	matrix, err = BuildSMSv2MatrixFromTerraform(legacy, current, attrs)
	if err != nil {
		t.Fatal(err)
	}
	if !matrix.Entries[0].Current.Required || matrix.Entries[0].Current.Optional {
		t.Fatalf("matrix did not report generated requiredness: %+v", matrix.Entries[0].Current)
	}
}

func TestBuildSMSv2MatrixClassifiesSupportedCases(t *testing.T) {
	legacy := &LegacyManifest{Version: "0.12.2", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 1, Paths: []LegacyField{
		{Path: "name", WireKey: "name", Type: "string", Cardinality: "single", Required: true, ForceNew: true},
	}}
	current := &CurrentManifest{Version: "2.1.225", PathCount: 1, Paths: []CurrentField{{Path: "metadata.name", WireKey: "name", Type: "string", Cardinality: "single", CreateRequired: true}}, ChoiceGroups: map[string][]string{"choice": {"metadata.name"}}}
	matrix, err := BuildSMSv2Matrix(legacy, current)
	if err != nil {
		t.Fatal(err)
	}
	for _, classification := range []string{"current_parity"} {
		if matrix.Classification[classification] != 1 {
			t.Fatalf("classification %s count=%d", classification, matrix.Classification[classification])
		}
	}
}

func TestFlattenTerraformAttributesUsesWireNamesForReservedNames(t *testing.T) {
	t.Parallel()
	attributes := []openapi.TerraformAttribute{
		{Name: "description", TfsdkTag: "description_spec", JsonName: "description", IsSpecField: true},
		{Name: "provider", TfsdkTag: "provider_ref", JsonName: "provider", IsSpecField: true},
	}
	got := flattenTerraformAttributes(attributes)
	for _, path := range []string{"spec.description", "spec.provider"} {
		if _, ok := got[path]; !ok {
			t.Errorf("missing wire path %q", path)
		}
	}
	for _, path := range []string{"spec.description_spec", "spec.provider_ref"} {
		if _, ok := got[path]; ok {
			t.Errorf("Terraform-only alias leaked into parity path %q", path)
		}
	}
}

func TestDeprecatedCapabilityDoesNotProveRemoval(t *testing.T) {
	legacy := &LegacyManifest{Paths: []LegacyField{{Path: "log_receiver", Deprecated: true, Type: "list", Optional: true}}}
	current := &CurrentManifest{}
	matrix, err := BuildSMSv2Matrix(legacy, current)
	if err == nil || len(matrix.Unclassified) != 1 {
		t.Fatalf("deprecation cannot establish removal: matrix=%+v err=%v", matrix, err)
	}
}

func TestMissingManagedPlatformDoesNotProveEquivalentLifecycle(t *testing.T) {
	for _, platform := range []string{"aws", "azure", "gcp"} {
		t.Run(platform, func(t *testing.T) {
			legacy := &LegacyManifest{Paths: []LegacyField{{Path: platform + ".managed", Type: "list", Optional: true}}}
			matrix, err := BuildSMSv2Matrix(legacy, &CurrentManifest{})
			if err == nil || len(matrix.Unclassified) != 1 {
				t.Fatalf("another resource name does not prove equivalence: %+v %v", matrix, err)
			}
		})
	}
}

func TestPlatformRemovalRequiresIndependentEvidence(t *testing.T) {
	legacy := &LegacyManifest{Paths: []LegacyField{{Path: "segment_vrf[].segment_config.nameserver_v6", Type: "string", Optional: true}}}
	matrix, err := BuildSMSv2Matrix(legacy, &CurrentManifest{})
	if err == nil || len(matrix.Unclassified) != 1 {
		t.Fatalf("a hard-coded field name is not removal evidence: %+v %v", matrix, err)
	}
}

func TestGeneratedParitySeparatesSchemaPathFromWireName(t *testing.T) {
	legacy := &LegacyManifest{Paths: []LegacyField{{Path: "blocked_services.blocked_sevice[]", WireKey: "blocked_sevice", Type: "list", Cardinality: "list", Optional: true}}}
	current := &CurrentManifest{Paths: []CurrentField{{Path: "spec.blocked_services.blocked_service[]", WireKey: "blocked_sevice", Type: "array", Cardinality: "list"}}}
	attrs := []openapi.TerraformAttribute{{Name: "blocked_services", IsBlock: true, NestedBlockType: "single", NestedAttributes: []openapi.TerraformAttribute{{Name: "blocked_service", JsonName: "blocked_sevice", TfsdkTag: "blocked_service", IsBlock: true, NestedBlockType: "list", Optional: true}}}}
	matrix, err := BuildSMSv2MatrixFromTerraform(legacy, current, attrs)
	if err != nil {
		t.Fatalf("logical rename must retain generated capability: %v", err)
	}
	if len(matrix.Entries) != 1 || !matrix.Entries[0].Current.Generated || matrix.Entries[0].Current.WireKey != "blocked_sevice" {
		t.Fatalf("wire identity lost: %+v", matrix)
	}
}

func TestVerifiedPlatformRemovalIsScopedAndRetainsEvidence(t *testing.T) {
	proof := `{"current_platform_removals":["spec.rseries"],"platform_removal_evidence":{"spec.rseries":{"proof_kind":"explicit_api_rejection","http_status":400,"server_message":"Rseries provider is not supported for SecureMeshSite","observed_date":"2026-09-06","legacy_fixture_sha256":"sha256:` + strings.Repeat("a", 64) + `","probe_receipt_sha256":"sha256:` + strings.Repeat("b", 64) + `"}}}`
	var current CurrentManifest
	if err := json.Unmarshal([]byte(proof), &current); err != nil {
		t.Fatal(err)
	}
	legacy := &LegacyManifest{Paths: []LegacyField{{Path: "rseries"}, {Path: "rseries.not_managed.node_list[].hostname"}, {Path: "rseries_other"}}}
	matrix, err := BuildSMSv2Matrix(legacy, &current)
	if err == nil || len(matrix.Unclassified) != 1 || matrix.Unclassified[0] != "rseries_other" || matrix.Classification["current_platform_removal"] != 2 {
		t.Fatalf("removal scope incorrect: %+v %v", matrix, err)
	}
	data, _ := json.Marshal(matrix)
	if !strings.Contains(string(data), "probe_receipt_sha256") {
		t.Fatal("matrix discarded evidence provenance")
	}
	var unsupported CurrentManifest
	if err := json.Unmarshal([]byte(strings.ReplaceAll(proof, "explicit_api_rejection", "deprecation_annotation")), &unsupported); err != nil {
		t.Fatal(err)
	}
	matrix, _ = BuildSMSv2Matrix(legacy, &unsupported)
	if len(matrix.Unclassified) != 3 {
		t.Fatal("unsupported removal must remain unresolved")
	}
}

func TestPlatformRemovalRejectsUnrelatedAPIRejections(t *testing.T) {
	for _, message := range []string{
		"A subscription to addon f5xc-ipv6-standard is required",
		"Invalid interface configuration",
		"Aws provider is not supported for SecureMeshSite",
	} {
		t.Run(message, func(t *testing.T) {
			current := &CurrentManifest{CurrentPlatformRemovals: []string{"spec.rseries"}, PlatformRemovalEvidence: map[string]RemovalEvidence{"spec.rseries": {
				ProofKind: "explicit_api_rejection", HTTPStatus: 400, ServerMessage: message, ObservedDate: "2026-09-06", LegacyFixtureSHA256: "sha256:" + strings.Repeat("a", 64), ProbeReceiptSHA256: "sha256:" + strings.Repeat("b", 64),
			}}}
			legacy := &LegacyManifest{Paths: []LegacyField{{Path: "rseries"}}}
			matrix, err := BuildSMSv2Matrix(legacy, current)
			if err == nil || len(matrix.Unclassified) != 1 || matrix.Classification["current_platform_removal"] != 0 {
				t.Fatalf("unrelated rejection incorrectly waived parity: %+v err=%v", matrix, err)
			}
		})
	}
}
