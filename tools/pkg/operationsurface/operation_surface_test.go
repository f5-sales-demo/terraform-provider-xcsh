package operationsurface

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/releasesurface"
)

func TestV9OperationSurfaceIsExact(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	specDir := filepath.Join(root, "docs", "specifications", "api")
	catalog, err := openapi.ParseOperationCatalogFromDir(specDir)
	if err != nil {
		t.Fatal(err)
	}
	manifest, err := LoadAndApply(
		filepath.Join(root, "operation-surface.json"),
		filepath.Join(specDir, "upstream-contract-changes.json"),
		"v9.0.0",
		catalog,
	)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(manifest.Operations), 44; got != want {
		t.Fatalf("operation mappings = %d, want %d", got, want)
	}
	counts := map[string]int{}
	definitions := map[string]bool{}
	for _, operation := range manifest.Operations {
		counts[operation.Kind]++
		definitions[operation.Kind+" "+operation.Name] = true
	}
	if counts["resource"] != 3 || counts["data_source"] != 30 || counts["action"] != 9 || counts["ephemeral_resource"] != 2 {
		t.Fatalf("operation mapping counts = %#v", counts)
	}
	if got, want := len(definitions), 41; got != want {
		t.Fatalf("Terraform definitions = %d, want %d", got, want)
	}
	surface, err := releasesurface.Load(filepath.Join(root, "provider-release-surface.json"))
	if err != nil {
		t.Fatal(err)
	}
	for _, operation := range manifest.Operations {
		allowed := map[string]bool{
			"resource":           surface.AllowsResource(operation.Name),
			"data_source":        surface.AllowsDataSource(operation.Name),
			"action":             surface.AllowsAction(operation.Name),
			"ephemeral_resource": surface.AllowsEphemeralResource(operation.Name),
		}[operation.Kind]
		if !allowed {
			t.Errorf("%s %s is absent from the provider release surface", operation.Kind, operation.Name)
		}
	}
}

func TestV9OperationSurfaceRejectsIncompleteDuplicateAndMismatchedMappings(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	specDir := filepath.Join(root, "docs", "specifications", "api")
	data, err := os.ReadFile(filepath.Join(root, "operation-surface.json"))
	if err != nil {
		t.Fatal(err)
	}
	var baseline Manifest
	if err := json.Unmarshal(data, &baseline); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name   string
		mutate func(*Manifest)
	}{
		{name: "missing", mutate: func(manifest *Manifest) { manifest.Operations = manifest.Operations[1:] }},
		{name: "duplicate", mutate: func(manifest *Manifest) { manifest.Operations = append(manifest.Operations, manifest.Operations[0]) }},
		{name: "operation id", mutate: func(manifest *Manifest) { manifest.Operations[0].OperationID = "wrong.operation" }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			manifest := baseline
			manifest.Operations = append([]Operation(nil), baseline.Operations...)
			test.mutate(&manifest)
			encoded, err := json.Marshal(manifest)
			if err != nil {
				t.Fatal(err)
			}
			path := filepath.Join(t.TempDir(), "operation-surface.json")
			if err := os.WriteFile(path, encoded, 0o600); err != nil {
				t.Fatal(err)
			}
			catalog, err := openapi.ParseOperationCatalogFromDir(specDir)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := LoadAndApply(path, filepath.Join(specDir, "upstream-contract-changes.json"), "v9.0.0", catalog); err == nil {
				t.Fatal("invalid operation surface was accepted")
			}
		})
	}
}
