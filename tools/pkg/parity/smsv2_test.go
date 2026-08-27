// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package parity

import (
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestBuildSMSv2MatrixFailsUnclassifiedGap(t *testing.T) {
	legacy := &LegacyManifest{Version: "0.11.49", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 1, Paths: []LegacyField{{Path: "missing", Type: "string", Optional: true}}}
	current := &CurrentManifest{Version: "2.1.225", PathCount: 1, Paths: []CurrentField{{Path: "spec.present", Type: "string", Cardinality: "single"}}, ChoiceGroups: map[string][]string{"choice": {"spec.present"}}}
	matrix, err := BuildSMSv2Matrix(legacy, current)
	if err == nil || len(matrix.Unclassified) != 1 || matrix.Unclassified[0] != "missing" {
		t.Fatalf("expected one unclassified gap, matrix=%+v err=%v", matrix, err)
	}
}

func TestBuildSMSv2MatrixFromTerraformReportsGeneratorGap(t *testing.T) {
	legacy := &LegacyManifest{Version: "0.11.49", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 1, Paths: []LegacyField{{Path: "missing", Type: "string", Optional: true}}}
	current := &CurrentManifest{Version: "2.1.225", PathCount: 1, Paths: []CurrentField{{Path: "spec.missing", Type: "string", Cardinality: "single"}}, ChoiceGroups: map[string][]string{"choice": {"spec.missing"}}}
	matrix, err := BuildSMSv2MatrixFromTerraform(legacy, current, nil)
	if err != nil {
		t.Fatal(err)
	}
	if matrix.Classification["generator_gap"] != 1 || matrix.Entries[0].Current.Generated {
		t.Fatalf("expected an explicit generator gap, got %+v", matrix)
	}
}

func TestBuildSMSv2MatrixFromTerraformUsesGeneratedRequiredness(t *testing.T) {
	legacy := &LegacyManifest{Version: "0.11.49", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 1, Paths: []LegacyField{{
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
	legacy := &LegacyManifest{Version: "0.11.49", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 3, Paths: []LegacyField{
		{Path: "name", WireKey: "name", Type: "string", Cardinality: "single", Required: true, ForceNew: true},
		{Path: "private_adn", Type: "list", Optional: true, Deprecated: true},
		{Path: "segment_vrf[].segment_config.nameserver_v6", Type: "string", Optional: true},
	}}
	current := &CurrentManifest{Version: "2.1.225", PathCount: 1, Paths: []CurrentField{{Path: "metadata.name", WireKey: "name", Type: "string", Cardinality: "single", CreateRequired: true}}, ChoiceGroups: map[string][]string{"choice": {"metadata.name"}}}
	matrix, err := BuildSMSv2Matrix(legacy, current)
	if err != nil {
		t.Fatal(err)
	}
	for _, classification := range []string{"current_parity", "deprecated_exclusion", "current_platform_removal"} {
		if matrix.Classification[classification] != 1 {
			t.Fatalf("classification %s count=%d", classification, matrix.Classification[classification])
		}
	}
}
