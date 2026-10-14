// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"os"
	"testing"
)

func TestImportCaptureUsesSchemaFieldsOnly(t *testing.T) {
	path := t.TempDir() + "/schema.json"
	data := `{"provider_schemas":{"xcsh":{"resource_schemas":{"xcsh_origin_pool":{"block":{"attributes":{"port":{"required":true}}}}}}}`
	if err := os.WriteFile(path, []byte(data), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("XCSH_ACC_SCHEMA", path)
	capture, check := CaptureImportFieldDiagnostics(t, "xcsh_origin_pool", "xcsh_origin_pool.test")
	state := &terraform.State{Modules: []*terraform.ModuleState{{Path: []string{"root"}, Resources: map[string]*terraform.ResourceState{"xcsh_origin_pool.test": {Type: "xcsh_origin_pool", Primary: &terraform.InstanceState{Attributes: map[string]string{"port": "80", "annotations.SECRET_SENTINEL": "private"}}}}}}}
	if err := capture(state); err != nil {
		t.Fatal(err)
	}
	if err := check([]*terraform.InstanceState{{Attributes: map[string]string{"port": "81", "annotations.SECRET_SENTINEL": "changed"}}}); err != nil {
		t.Fatal(err)
	}
}
