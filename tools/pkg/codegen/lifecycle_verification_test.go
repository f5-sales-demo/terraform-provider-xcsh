package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestProtectedDomainGeneratorVerificationPersists(t *testing.T) {
	r := &openapi.ResourceTemplate{Name: "protected_domain", TitleCase: "ProtectedDomain", HasNamespaceInPath: true, ImportIDExtraFields: []string{"protected_domain"}, Attributes: []openapi.TerraformAttribute{
		{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string", Required: true},
		{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string", Required: true},
		{Name: "id", GoName: "ID", TfsdkTag: "id", Type: "string", Computed: true},
		{Name: "protected_domain", GoName: "ProtectedDomain", TfsdkTag: "protected_domain", Type: "string", Required: true},
	}}
	dir := t.TempDir()
	if err := GenerateResourceFile(r, dir); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(filepath.Join(dir, "protected_domain_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	text := string(body)
	for _, required := range []string{"VerifyProtectedDomain(ctx", "ValidateProtectedDomainCreateResponse", "client.HasHTTPStatus(err, 404)", "Unable to Verify Imported Registration", "data.ID = data.Name"} {
		if !strings.Contains(text, required) {
			t.Fatalf("missing %s", required)
		}
	}
	if strings.Contains(text, "keeping prior state") {
		t.Fatal("501 silently suppresses verification")
	}
	if err := GenerateDataSource(r, dir); err != nil {
		t.Fatal(err)
	}
	body, err = os.ReadFile(filepath.Join(dir, "protected_domain_data_source.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), "VerifyProtectedDomain(ctx, namespace, data.Name.ValueString(), data.ProtectedDomain.ValueString())") {
		t.Fatal("lookup does not use authoritative root")
	}
}

func TestRootBlockConflictsDeferValidationButBackstopPlanning(t *testing.T) {
	r := &openapi.ResourceTemplate{Name: "synthetic", TitleCase: "Synthetic", Attributes: []openapi.TerraformAttribute{
		{GoName: "Name", TfsdkTag: "name", Type: "string", Required: true},
		{GoName: "Left", TfsdkTag: "left", IsBlock: true, NestedBlockType: "single", OneOfGroup: "choice"},
		{GoName: "Right", TfsdkTag: "right", IsBlock: true, NestedBlockType: "single", OneOfGroup: "choice"},
	}}
	dir := t.TempDir()
	if err := GenerateResourceFile(r, dir); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(filepath.Join(dir, "synthetic_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	text := string(body)
	start := strings.Index(text, "func (r *SyntheticResource) ValidateConfig")
	end := strings.Index(text, "func (r *SyntheticResource) ModifyPlan")
	if !strings.Contains(text[start:end], "if req.Config.Raw.IsFullyKnown()") {
		t.Fatal("unresolved validation must defer ambiguous empty alternatives")
	}
	planned := text[end : strings.Index(text[end:], "func (r *SyntheticResource) Create")+end]
	if strings.Contains(planned, "IsFullyKnown()") || !strings.Contains(planned, `validators.ConfiguredObject(ctx, req.Config, path.Root("left"), &resp.Diagnostics)`) || !strings.Contains(planned, "Conflicting Configuration") {
		t.Fatal("planning must reject known conflicting selections despite unrelated unknowns")
	}
}

func TestWAFGeneratorRejectsExplicitBlockingPageConflict(t *testing.T) {
	r := &openapi.ResourceTemplate{Name: "app_firewall", TitleCase: "AppFirewall"}
	dir := t.TempDir()
	if err := GenerateResourceFile(r, dir); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(filepath.Join(dir, "app_firewall_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `validators.ConfiguredObject(ctx, req.Config, path.Root("blocking_page"), &resp.Diagnostics) && !data.UseDefaultBlockingPage.IsNull()`) {
		t.Fatal("mixed marker/block conflict missing")
	}
}
