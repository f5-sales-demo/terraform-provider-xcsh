// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package parity

import "testing"

func TestBuildSMSv2MatrixFailsUnclassifiedGap(t *testing.T) {
	legacy := &LegacyManifest{Version: "0.11.49", SourceURL: "source", SourceSHA256: "sha256:test", PathCount: 1, Paths: []LegacyField{{Path: "missing", Type: "string", Optional: true}}}
	current := &CurrentManifest{Version: "2.1.225", PathCount: 1, Paths: []CurrentField{{Path: "spec.present", Type: "string", Cardinality: "single"}}, ChoiceGroups: map[string][]string{"choice": {"spec.present"}}}
	matrix, err := BuildSMSv2Matrix(legacy, current)
	if err == nil || len(matrix.Unclassified) != 1 || matrix.Unclassified[0] != "missing" {
		t.Fatalf("expected one unclassified gap, matrix=%+v err=%v", matrix, err)
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
