// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestExtractResponseOperationSchemaBindsPathQueryAndBody(t *testing.T) {
	spec := responseOperationProbeSpec()
	operation := openapi.ResolvedResponseOperation{
		Name: "site_registrations_by_state", Role: "collection", Method: "POST",
		Path: "/api/register/namespaces/{namespace}/probe", OperationID: "ves.io.schema.probe.CustomAPI.List",
		RequestSchema: "probeRequest", ResponseSchema: "probeResponse",
	}

	result, err := ExtractResponseOperationSchema(spec, operation)
	if err != nil {
		t.Fatalf("ExtractResponseOperationSchema() error = %v", err)
	}
	if result.Name != operation.Name || result.Role != operation.Role || result.TitleCase != "SiteRegistrationsByState" {
		t.Fatalf("identity fields = %+v", result)
	}
	inputs := operationInputsByTag(result.Inputs)
	if len(inputs) != 3 {
		t.Fatalf("len(inputs) = %d, want 3: %+v", len(inputs), inputs)
	}
	namespace := inputs["namespace"]
	if namespace == nil || namespace.Attribute.Required || !namespace.Attribute.Optional || namespace.Attribute.StringDefault != "system" {
		t.Fatalf("namespace default contract = %+v", namespace)
	}
	if len(namespace.Bindings) != 2 || namespace.Bindings[0].Location != "path" || namespace.Bindings[1].Location != "body" {
		t.Fatalf("namespace bindings = %+v, want path and body", namespace.Bindings)
	}
	if state := inputs["state"]; state == nil || !state.Attribute.Required || len(state.Bindings) != 1 || state.Bindings[0].Location != "body" {
		t.Fatalf("state input = %+v", state)
	}
	if limit := inputs["limit"]; limit == nil || !limit.Attribute.Optional || len(limit.Bindings) != 1 || limit.Bindings[0].Location != "query" {
		t.Fatalf("limit input = %+v", limit)
	}

	responses := terraformAttributesByTag(result.ResponseAttributes)
	if secret := responses["secret"]; secret == nil || !secret.Computed || !secret.Sensitive {
		t.Fatalf("secret response = %+v, want sensitive computed", secret)
	}
	items := responses["items"]
	if items == nil || !items.IsBlock || items.NestedBlockType != "list" || len(items.NestedAttributes) != 2 {
		t.Fatalf("items response = %+v, want typed referenced list", items)
	}
}

func TestExtractResponseOperationSchemaAppliesPublicInputPolicies(t *testing.T) {
	spec := responseOperationProbeSpec()
	imagePath := spec.Paths["/api/register/namespaces/{namespace}/probe"].(map[string]interface{})
	post := imagePath["post"].(map[string]interface{})
	post["parameters"] = []interface{}{
		map[string]interface{}{"name": "provider", "in": "query", "schema": map[string]interface{}{"type": "string"}},
	}
	post["x-f5xc-required-fields"] = []interface{}{"provider"}
	delete(post, "requestBody")

	result, err := ExtractResponseOperationSchema(spec, openapi.ResolvedResponseOperation{
		Name: "site_image", Role: "query", Method: "POST", Path: "/api/register/namespaces/{namespace}/probe",
		OperationID: "ves.io.schema.probe.CustomAPI.List", RequestSchema: "", ResponseSchema: "probeResponse",
	})
	if err == nil || !strings.Contains(err.Error(), "request schema") {
		t.Fatalf("POST operation without request schema error = %v", err)
	}

	post["requestBody"] = map[string]interface{}{"content": map[string]interface{}{
		"application/json": map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/probeRequest"}},
	}}
	result, err = ExtractResponseOperationSchema(spec, openapi.ResolvedResponseOperation{
		Name: "site_image", Role: "query", Method: "POST", Path: "/api/register/namespaces/{namespace}/probe",
		OperationID: "ves.io.schema.probe.CustomAPI.List", RequestSchema: "probeRequest", ResponseSchema: "probeResponse",
	})
	if err != nil {
		t.Fatal(err)
	}
	inputs := operationInputsByTag(result.Inputs)
	if input := inputs["provider_ref"]; input == nil || input.Attribute.JsonName != "provider" || !input.Attribute.Required {
		t.Fatalf("provider_ref input = %+v", input)
	}
}

func TestExtractResponseOperationSchemaSupportsScalarResponse(t *testing.T) {
	spec := responseOperationProbeSpec()
	spec.Components.Schemas["scalarResponse"] = openapi.Schema{Type: "string", Description: "Accepted version."}
	result, err := ExtractResponseOperationSchema(spec, openapi.ResolvedResponseOperation{
		Name: "probe_query", Role: "query", Method: "GET", Path: "/api/register/namespaces/{namespace}/probe",
		OperationID: "ves.io.schema.probe.CustomAPI.List", ResponseSchema: "scalarResponse",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.ResponseAttributes) != 1 || result.ResponseAttributes[0].TfsdkTag != "result" || result.ResponseAttributes[0].Type != "string" {
		t.Fatalf("scalar response attributes = %+v", result.ResponseAttributes)
	}
}

func TestExtractResponseOperationSchemaRejectsUnsupportedShape(t *testing.T) {
	spec := responseOperationProbeSpec()
	spec.Components.Schemas["badResponse"] = openapi.Schema{Type: "number"}
	_, err := ExtractResponseOperationSchema(spec, openapi.ResolvedResponseOperation{
		Name: "bad_query", Role: "query", Method: "GET", Path: "/api/register/namespaces/{namespace}/probe",
		OperationID: "ves.io.schema.probe.CustomAPI.List", ResponseSchema: "badResponse",
	})
	if err == nil || !strings.Contains(err.Error(), "unsupported response schema type") {
		t.Fatalf("ExtractResponseOperationSchema() error = %v", err)
	}
}

func TestExtractResponseOperationSchemaRejectsNestedRequestInput(t *testing.T) {
	spec := responseOperationProbeSpec()
	request := spec.Components.Schemas["probeRequest"]
	request.Properties["nested"] = openapi.Schema{Type: "object", Properties: map[string]openapi.Schema{"value": {Type: "string"}}}
	spec.Components.Schemas["probeRequest"] = request
	_, err := ExtractResponseOperationSchema(spec, openapi.ResolvedResponseOperation{
		Name: "probe_action", Role: "action", Method: "POST", Path: "/api/register/namespaces/{namespace}/probe",
		OperationID: "ves.io.schema.probe.CustomAPI.List", RequestSchema: "probeRequest", ResponseSchema: "probeResponse",
	})
	if err == nil || !strings.Contains(err.Error(), "unsupported body shape") {
		t.Fatalf("nested request error = %v", err)
	}
}

func responseOperationProbeSpec() *openapi.Spec {
	return &openapi.Spec{
		Paths: map[string]interface{}{
			"/api/register/namespaces/{namespace}/probe": map[string]interface{}{
				"post": map[string]interface{}{
					"operationId":            "ves.io.schema.probe.CustomAPI.List",
					"x-f5xc-required-fields": []interface{}{"namespace", "state"},
					"parameters": []interface{}{
						map[string]interface{}{"name": "namespace", "in": "path", "required": true, "schema": map[string]interface{}{"type": "string"}},
						map[string]interface{}{"name": "limit", "in": "query", "schema": map[string]interface{}{"type": "integer", "format": "int64"}},
					},
					"requestBody": map[string]interface{}{"content": map[string]interface{}{
						"application/json": map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/probeRequest"}},
					}},
				},
				"get": map[string]interface{}{
					"operationId": "ves.io.schema.probe.CustomAPI.List",
					"parameters":  []interface{}{},
				},
			},
		},
		Components: openapi.Components{Schemas: map[string]openapi.Schema{
			"probeRequest": {
				Type: "object", Required: []string{"namespace", "state"},
				Properties: map[string]openapi.Schema{
					"namespace": {Type: "string"},
					"state":     {Type: "string"},
				},
			},
			"probeResponse": {
				Type: "object", Description: "Probe response.",
				Properties: map[string]openapi.Schema{
					"secret": {Type: "string", XF5XCSensitive: true},
					"items":  {Type: "array", Items: &openapi.Schema{Ref: "#/components/schemas/probeItem"}},
				},
			},
			"probeItem": {
				Type: "object", Properties: map[string]openapi.Schema{
					"name":    {Type: "string"},
					"details": {AllOf: []openapi.Schema{{Ref: "#/components/schemas/probeDetails"}}},
				},
			},
			"probeDetails": {Type: "object", Properties: map[string]openapi.Schema{"enabled": {Type: "boolean"}}},
		}},
	}
}

func operationInputsByTag(inputs []openapi.ResponseOperationInput) map[string]*openapi.ResponseOperationInput {
	result := make(map[string]*openapi.ResponseOperationInput, len(inputs))
	for index := range inputs {
		result[inputs[index].Attribute.TfsdkTag] = &inputs[index]
	}
	return result
}

func terraformAttributesByTag(attributes []openapi.TerraformAttribute) map[string]*openapi.TerraformAttribute {
	result := make(map[string]*openapi.TerraformAttribute, len(attributes))
	for index := range attributes {
		result[attributes[index].TfsdkTag] = &attributes[index]
	}
	return result
}
