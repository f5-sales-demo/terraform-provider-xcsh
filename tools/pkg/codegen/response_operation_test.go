// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestResponseOperationTypedMapConversion(t *testing.T) {
	fieldTypes := map[string]attr.Type{"url": types.StringType}
	entry, diags := types.ObjectValue(fieldTypes, map[string]attr.Value{"url": types.StringValue("https://example.com/image")})
	if diags.HasError() {
		t.Fatal(diags)
	}
	value, diags := types.MapValue(types.ObjectType{AttrTypes: fieldTypes}, map[string]attr.Value{"example": entry})
	if diags.HasError() || value.IsNull() || len(value.Elements()) != 1 {
		t.Fatalf("typed map conversion: %v", diags)
	}
	operation := responseOperationTemplate("query")
	operation.ResponseAttributes = []openapi.TerraformAttribute{{Name: "images", JsonName: "images", TfsdkTag: "images", GoName: "Images", Type: "map", IsBlock: true, NestedBlockType: "map", Computed: true, NestedAttributes: []openapi.TerraformAttribute{{Name: "url", JsonName: "url", TfsdkTag: "url", GoName: "URL", Type: "string", Computed: true}}}}
	dir := t.TempDir()
	if err := GenerateResponseOperation(operation, dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, operation.Name+"_data_source.go"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"MapNestedAttribute", "types.Map", "types.MapValue"} {
		if !strings.Contains(string(data), want) {
			t.Errorf("missing typed response map code %s", want)
		}
	}
}

func TestGeneratedResponseOperationsCompile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping generated response-operation compile check in short mode")
	}
	root := repoRootFromTest(t)
	providerDir, err := os.MkdirTemp(filepath.Join(root, "internal"), "zz_responsecompile_")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(providerDir) })

	roles := []string{"query", "collection", "issuance", "action"}
	for index, role := range roles {
		operation := responseOperationTemplate(role)
		operation.ResponseFields = []openapi.ResponseField{{Name: "secret", Type: "string", Required: true, MinLength: 1}}
		operation.ResponseAttributes = append(operation.ResponseAttributes, openapi.TerraformAttribute{Name: "images", JsonName: "images", TfsdkTag: "images", GoName: "Images", Type: "map", IsBlock: true, NestedBlockType: "map", Computed: true, NestedAttributes: []openapi.TerraformAttribute{{Name: "url", JsonName: "url", TfsdkTag: "url", GoName: "URL", Type: "string", Computed: true}}})
		operation.ResponseFields = append(operation.ResponseFields, openapi.ResponseField{Name: "images", Type: "object", MapFields: []openapi.ResponseField{{Name: "url", Type: "string", Required: true}}})
		operation.Name = "zz_response_" + role
		operation.TitleCase = "ZzResponse" + strings.ToUpper(role[:1]) + role[1:]
		if role == "query" || role == "issuance" {
			operation.Method = "GET"
			for inputIndex := range operation.Inputs {
				operation.Inputs[inputIndex].Bindings = []openapi.OperationBinding{{Location: "query", Name: operation.Inputs[inputIndex].Attribute.JsonName}}
			}
		}
		if err := GenerateResponseOperation(operation, providerDir); err != nil {
			t.Fatalf("generate role %s at index %d: %v", role, index, err)
		}
	}

	var image map[string]any
	if err := json.Unmarshal([]byte(syntheticKVMImageContract), &image); err != nil {
		t.Fatal(err)
	}
	if err := generateKVMImageDataSource(image, providerDir); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "build", "./internal/"+filepath.Base(providerDir)+"/")
	cmd.Dir = root
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated response operations failed to compile (%v):\n%s", err, output)
	}
}

func TestResponseOperationRetryPolicyUsesOperationRole(t *testing.T) {
	for _, role := range []string{"query", "issuance"} {
		operation := responseOperationTemplate(role)
		operation.Method = "GET"
		code := renderResponseOperationInvoke(operation, "r", true)
		want := ".Get("
		if role == "issuance" {
			want = ".GetOnce("
		}
		if !strings.Contains(code, want) {
			t.Fatalf("role %s omitted %s", role, want)
		}
		operation.Method = "POST"
		if code = renderResponseOperationInvoke(operation, "r", true); !strings.Contains(code, ".Post(") {
			t.Fatalf("role %s did not use non-retrying POST", role)
		}
	}
}

func TestResponseOperationDiagnosticsDoNotRenderRawBackendErrors(t *testing.T) {
	code := renderResponseOperationDiagnostic(responseOperationTemplate("query"))
	if strings.Contains(code, "fmt.Sprintf") || !strings.Contains(code, "Raw API diagnostics are suppressed") {
		t.Fatal("generic response operation would expose raw backend error data")
	}
}

func TestResponseOperationArrayValidatorsPreserveUniqueness(t *testing.T) {
	var code strings.Builder
	renderResponseOperationValidators(&code, openapi.TerraformAttribute{Type: "list", MinItems: 1, MaxItems: 100, UniqueItems: true}, "")
	if !strings.Contains(code.String(), "listvalidator.SizeBetween(1, 100)") || !strings.Contains(code.String(), "listvalidator.UniqueValues()") {
		t.Fatal("missing image query array validators")
	}
}

func TestResponseOperationValidationPrecedesStateAssignment(t *testing.T) {
	operation := responseOperationTemplate("query")
	operation.ResponseFields = []openapi.ResponseField{{Name: "secret", Type: "string", Required: true, MinLength: 1, Format: "uri"}}
	dir := t.TempDir()
	if err := GenerateResponseOperation(operation, dir); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(filepath.Join(dir, operation.Name+"_data_source.go"))
	if err != nil {
		t.Fatal(err)
	}
	source := string(content)
	validation := strings.Index(source, "client.ValidateResponseFields")
	state := strings.Index(source, "resp.State.Set")
	if validation < 0 || state < validation {
		t.Fatal("response validation must precede writing state")
	}
	for _, want := range []string{`Required: true`, `MinLength: 1`, `Format: "uri"`, `AddError("Incomplete API Response", err.Error())`} {
		if !strings.Contains(source, want) {
			t.Errorf("missing %s", want)
		}
	}
}

func TestGenerateResponseOperationEveryRole(t *testing.T) {
	for _, test := range []struct {
		role, suffix, marker string
	}{
		{role: "query", suffix: "_data_source.go", marker: "datasource.DataSource"},
		{role: "collection", suffix: "_data_source.go", marker: "datasource.DataSource"},
		{role: "issuance", suffix: "_resource.go", marker: "resource.Resource"},
		{role: "action", suffix: "_action.go", marker: "action.Action"},
	} {
		t.Run(test.role, func(t *testing.T) {
			dir := t.TempDir()
			template := responseOperationTemplate(test.role)
			if err := GenerateResponseOperation(template, dir); err != nil {
				t.Fatalf("GenerateResponseOperation() error = %v", err)
			}
			path := filepath.Join(dir, template.Name+test.suffix)
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			text := string(content)
			for _, want := range []string{"Code generated", test.marker, `"/api/probe/{namespace}/{name}"`, "url.PathEscape", `stringvalidator.OneOf("v1", "v2")`} {
				if !strings.Contains(text, want) {
					t.Errorf("generated %s missing %q", test.role, want)
				}
			}
			if test.role == "action" {
				for _, want := range []string{"func (a *ProbeAction) Invoke", "types.BoolValue(false)", `body["force"]`} {
					if !strings.Contains(text, want) {
						t.Errorf("generated action missing %q", want)
					}
				}
				for _, forbidden := range []string{"time.Sleep", "CurrentState", "poll"} {
					if strings.Contains(text, forbidden) {
						t.Errorf("generated action unexpectedly contains %q", forbidden)
					}
				}
				if strings.Contains(text, "apiResult :=") || strings.Contains(text, "scalarResult") {
					t.Error("generated action declares an unused response value")
				}
			}
			if test.role == "issuance" {
				if !strings.Contains(text, `"id": schema.StringAttribute{`) {
					t.Error("generated issuance schema is missing its computed id")
				}
				if !strings.Contains(text, `AddError("Update Not Supported"`) {
					t.Error("generated issuance is missing the canonical fail-closed update stub")
				}
			}
			if test.role != "action" && (!strings.Contains(text, "Sensitive:") || !strings.Contains(text, "true")) {
				t.Errorf("generated %s lost sensitive response propagation", test.role)
			}
		})
	}
}

func TestGenerateResponseOperationGETDoesNotDeclareBody(t *testing.T) {
	dir := t.TempDir()
	template := responseOperationTemplate("query")
	template.Method = "GET"
	for index := range template.Inputs {
		template.Inputs[index].Bindings = []openapi.OperationBinding{{Location: "query", Name: template.Inputs[index].Attribute.JsonName}}
	}
	if err := GenerateResponseOperation(template, dir); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(filepath.Join(dir, "probe_data_source.go"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(content), "body :=") {
		t.Fatal("generated GET declares an unused request body")
	}
}

func TestGenerateResponseOperationRendersImmutablePrerequisiteDiagnostic(t *testing.T) {
	dir := t.TempDir()
	template := responseOperationTemplate("query")
	template.Prerequisites = []openapi.ResponseOperationPrerequisite{{
		ID:              "maurice_config_cardinality_exactly_one",
		Resource:        "maurice_config",
		Exactly:         1,
		Enforcement:     "server",
		Availability:    "unresolved_server_lookup",
		Reason:          "Server lookup scope and count are unknown.",
		SourceKind:      "runtime_api_error",
		SourceOperation: "ves.io.schema.registration.CustomAPI.GetImageDownloadUrl",
		SourceImmutable: true,
		LookupScope:     "unknown", LookupCount: "unknown",
		SourceCommit: strings.Repeat("a", 40), SpecSHA256: strings.Repeat("b", 64),
		ReceiptPath: "config/evidence/test-receipt.json", ReceiptSHA256: strings.Repeat("c", 64),
	}}
	if err := GenerateResponseOperation(template, dir); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(filepath.Join(dir, "probe_data_source.go"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"responseOperationPrerequisiteDiagnostic", "maurice_config_cardinality_exactly_one", "unresolved_server_lookup", "LookupScope", "LookupCount", "SourceCommit", "SpecSHA256", "ReceiptPath", "ReceiptSHA256", strings.Repeat("c", 64)} {
		if !strings.Contains(string(content), want) {
			t.Fatalf("generated response operation missing %q:\n%s", want, content)
		}
	}
}

func TestResponseOperationEmptyChoiceMarkerUsesObjectAttribute(t *testing.T) {
	marker := openapi.TerraformAttribute{
		Name: "marker", GoName: "Marker", TfsdkTag: "marker", Type: "object",
		Computed: true, EmptyObjectMarker: true,
	}
	got := renderResponseOperationSchemaMap([]openapi.TerraformAttribute{marker}, "", false)
	for _, want := range []string{`"marker": schema.ObjectAttribute{`, `AttributeTypes: map[string]attr.Type{},`} {
		if !strings.Contains(got, want) {
			t.Fatalf("response-operation marker schema missing %q: %s", want, got)
		}
	}
}

func TestGenerateResponseOperationRejectsUnsupportedRoleAndHandwrittenTarget(t *testing.T) {
	dir := t.TempDir()
	template := responseOperationTemplate("unknown")
	if err := GenerateResponseOperation(template, dir); err == nil || !strings.Contains(err.Error(), "unsupported response-operation role") {
		t.Fatalf("unsupported role error = %v", err)
	}

	template = responseOperationTemplate("query")
	path := filepath.Join(dir, template.Name+"_data_source.go")
	if err := os.WriteFile(path, []byte("package provider\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := GenerateResponseOperation(template, dir); err == nil || !strings.Contains(err.Error(), "refusing to overwrite handwritten file") {
		t.Fatalf("handwritten target error = %v", err)
	}
}

func responseOperationTemplate(role string) *openapi.ResponseOperationTemplate {
	falseValue := false
	return &openapi.ResponseOperationTemplate{
		Name: "probe", Role: role, TitleCase: "Probe", Method: "POST", APIPath: "/api/probe/{namespace}/{name}",
		OperationID: "ves.io.schema.probe.CustomAPI.Invoke", RequestSchema: "probeRequest", ResponseSchema: "probeResponse",
		Description: "Probe operation.",
		Inputs: []openapi.ResponseOperationInput{
			{Attribute: openapi.TerraformAttribute{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", JsonName: "namespace", Type: "string", Required: true}, Bindings: []openapi.OperationBinding{{Location: "path", Name: "namespace"}, {Location: "body", Name: "namespace"}}},
			{Attribute: openapi.TerraformAttribute{Name: "name", GoName: "Name", TfsdkTag: "name", JsonName: "name", Type: "string", Required: true, EnumValues: []string{"v1", "v2"}}, Bindings: []openapi.OperationBinding{{Location: "path", Name: "name"}, {Location: "body", Name: "name"}}},
			{Attribute: openapi.TerraformAttribute{Name: "force", GoName: "Force", TfsdkTag: "force", JsonName: "force", Type: "bool", Optional: true, Default: falseValue}, Bindings: []openapi.OperationBinding{{Location: "body", Name: "force"}}},
		},
		ResponseAttributes: []openapi.TerraformAttribute{
			{Name: "secret", GoName: "Secret", TfsdkTag: "secret", JsonName: "secret", Type: "string", Computed: true, Sensitive: true, IsSpecField: true},
			{Name: "details", GoName: "Details", TfsdkTag: "details", JsonName: "details", IsBlock: true, NestedBlockType: "single", Computed: true, IsSpecField: true, NestedAttributes: []openapi.TerraformAttribute{
				{Name: "value", GoName: "Value", TfsdkTag: "value", JsonName: "value", Type: "string", Computed: true, IsSpecField: true},
			}},
			{Name: "items", GoName: "Items", TfsdkTag: "items", JsonName: "items", IsBlock: true, NestedBlockType: "list", Computed: true, IsSpecField: true, NestedAttributes: []openapi.TerraformAttribute{
				{Name: "name", GoName: "Name", TfsdkTag: "name", JsonName: "name", Type: "string", Computed: true, IsSpecField: true},
			}},
		},
	}
}
