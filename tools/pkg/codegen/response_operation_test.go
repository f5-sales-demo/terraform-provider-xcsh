// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

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
		operation.Name = "zz_response_" + role
		operation.TitleCase = "ZzResponse" + strings.ToUpper(role[:1]) + role[1:]
		if role == "query" {
			operation.Method = "GET"
			for inputIndex := range operation.Inputs {
				operation.Inputs[inputIndex].Bindings = []openapi.OperationBinding{{Location: "query", Name: operation.Inputs[inputIndex].Attribute.JsonName}}
			}
		}
		if err := GenerateResponseOperation(operation, providerDir); err != nil {
			t.Fatalf("generate role %s at index %d: %v", role, index, err)
		}
	}

	cmd := exec.Command("go", "build", "./internal/"+filepath.Base(providerDir)+"/")
	cmd.Dir = root
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated response operations failed to compile (%v):\n%s", err, output)
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
			if test.role == "issuance" && !strings.Contains(text, `"id": schema.StringAttribute{`) {
				t.Error("generated issuance schema is missing its computed id")
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
		ResponseAttributes: []openapi.TerraformAttribute{{Name: "secret", GoName: "Secret", TfsdkTag: "secret", JsonName: "secret", Type: "string", Computed: true, Sensitive: true, IsSpecField: true}},
	}
}
