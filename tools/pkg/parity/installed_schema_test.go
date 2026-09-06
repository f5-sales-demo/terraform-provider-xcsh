// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package parity

import (
	"reflect"
	"testing"
)

func TestInstalledPaths(t *testing.T) {
	data := []byte(`{"provider_schemas":{"example/provider":{"resource_schemas":{"site":{"block":{"attributes":{"name":{"type":"string"},"servers":{"type":["list","string"]}},"block_types":{"aws":{"max_items":1,"block":{"block_types":{"nodes":{"block":{"attributes":{"hostname":{"type":"string"}}}}}}}}}}}}}}`)
	got, err := InstalledPaths(data, "example/provider", "site")
	want := []string{"aws", "aws.nodes[]", "aws.nodes[].hostname", "name", "servers[]"}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("paths=%v err=%v", got, err)
	}
}

func TestInstalledPathsRejectsMissingAndMalformedSchemas(t *testing.T) {
	for _, data := range []string{`{}`, `null`, `invalid`, `{"provider_schemas":{"example/provider":{"resource_schemas":{"site":{"block":{}}}}}}`} {
		if _, err := InstalledPaths([]byte(data), "example/provider", "site"); err == nil {
			t.Fatalf("accepted invalid schema %s", data)
		}
	}
}

func TestInstalledPathsPreservesEmptyChoiceBlocks(t *testing.T) {
	data := []byte(`{"provider_schemas":{"example/provider":{"resource_schemas":{"site":{"block":{"block_types":{"disable":{"max_items":1,"block":{}}}}}}}}}`)
	got, err := InstalledPaths(data, "example/provider", "site")
	if err != nil || !reflect.DeepEqual(got, []string{"disable"}) {
		t.Fatalf("empty choice paths=%v err=%v", got, err)
	}
}
