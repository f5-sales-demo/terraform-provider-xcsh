package codegen

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBlindfoldAdaptersRegenerate(t *testing.T) {
	attrs := []openapi.TerraformAttribute{{Name: "certificate_url", TfsdkTag: "certificate_url", GoName: "CertificateURL", Type: "string", Required: true, IsSpecField: true}, {Name: "private_key", TfsdkTag: "private_key", GoName: "PrivateKey", IsBlock: true, NestedBlockType: "single", Optional: true, IsSpecField: true}}
	r := &openapi.ResourceTemplate{Name: "certificate", TitleCase: "Certificate", APIPath: "/api/config/namespaces/%s/certificates", APIPathItem: "/api/config/namespaces/%s/certificates/%s", HasNamespaceInPath: true, HasBlocks: true, Attributes: attrs}
	dir := t.TempDir()
	if err := GenerateResourceFile(r, dir); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(dir, "certificate_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), "newBlindfoldResource(") {
		t.Fatal("native adapter missing")
	}
	if err := GenerateResourceFile(r, dir); err != nil {
		t.Fatal(err)
	}
	again, _ := os.ReadFile(filepath.Join(dir, "certificate_resource.go"))
	if string(b) != string(again) {
		t.Fatal("regeneration differs")
	}
	if err := GenerateDataSource(r, dir); err != nil {
		t.Fatal(err)
	}
	d, _ := os.ReadFile(filepath.Join(dir, "certificate_data_source.go"))
	if strings.Contains(string(d), "newBlindfoldResource") || strings.Contains(string(d), `"blindfold"`) {
		t.Fatal("provider inputs exposed in data source")
	}
}
