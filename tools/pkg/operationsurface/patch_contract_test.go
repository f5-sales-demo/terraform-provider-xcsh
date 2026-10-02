package operationsurface

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"os"
	"path/filepath"
	"testing"
)

func TestOperationBaselineSurvivesSpecPatchDelta(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	catalog, err := openapi.ParseOperationCatalogFromDir(filepath.Join(root, "docs/specifications/api"))
	if err != nil {
		t.Fatal(err)
	}
	report := filepath.Join(t.TempDir(), "patch.json")
	if err = os.WriteFile(report, []byte(`{"operations":{"additions":[],"removals":[]}}`), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err = LoadAndApply(filepath.Join(root, "operation-surface.json"), report, "v9.0.1", catalog); err != nil {
		t.Fatal(err)
	}
	for _, release := range []string{"v8.0.9", "v10.0.0", "invalid"} {
		if _, err = LoadAndApply(filepath.Join(root, "operation-surface.json"), report, release, catalog); err == nil {
			t.Fatalf("accepted %s", release)
		}
	}
	for _, content := range []string{
		`{"operations":{"additions":[{"method":"POST","path":"/unmapped"}],"removals":[]}}`,
		`{"operations":{"additions":[],"removals":[{"method":"POST","path":"/api/config/dns/namespaces/system/dns_zone/add_cryptokey"}]}}`,
	} {
		if err = os.WriteFile(report, []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
		if _, err = LoadAndApply(filepath.Join(root, "operation-surface.json"), report, "v9.0.1", catalog); err == nil {
			t.Fatal("accepted invalid patch contract")
		}
	}
}
