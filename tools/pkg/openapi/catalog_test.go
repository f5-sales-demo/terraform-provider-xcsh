// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const validOperationCatalog = `{
  "service": "f5xc",
  "displayName": "F5 Distributed Cloud API",
  "version": "2.1.214",
  "specSource": "api-specs-enriched",
  "auth": {},
  "defaults": {},
  "apiOperations": [
    {
      "apiIdentity": "ves.io.schema.probe",
      "operations": [
        {
          "method": "GET",
          "path": "/api/config/namespaces/{namespace}/irregular_policys",
          "operationId": "ves.io.schema.probe.CustomApi.List",
          "surface": "config"
        },
        {
          "method": "GET",
          "path": "/api/config/namespaces/{namespace}/irregular_policys/{name}",
          "operationId": "ves.io.schema.probe.CustomApi.Get",
          "surface": "config"
        },
        {
          "method": "POST",
          "path": "/api/config/namespaces/{metadata.namespace}/irregular_policys",
          "operationId": "ves.io.schema.probe.CustomApi.Create",
          "surface": "config",
          "requestSchema": "probeCreateRequest"
        },
        {
          "method": "PUT",
          "path": "/api/config/namespaces/{metadata.namespace}/irregular_policys/{metadata.name}",
          "operationId": "ves.io.schema.probe.CustomApi.Replace",
          "surface": "config",
          "requestSchema": "probeReplaceRequest"
        }
      ]
    }
  ],
  "apiExclusions": [],
  "categories": []
}`

func TestParseOperationCatalogPreservesExactOperationFacts(t *testing.T) {
	catalog, err := ParseOperationCatalog([]byte(validOperationCatalog))
	if err != nil {
		t.Fatalf("ParseOperationCatalog() error = %v", err)
	}
	if catalog.Version != "2.1.214" {
		t.Fatalf("Version = %q, want 2.1.214", catalog.Version)
	}
	identity, ok := catalog.Identity("ves.io.schema.probe")
	if !ok {
		t.Fatal("Identity(ves.io.schema.probe) was not found")
	}
	if len(identity.Operations) != 4 {
		t.Fatalf("len(Operations) = %d, want 4", len(identity.Operations))
	}
	create := identity.Operations[2]
	if create.Method != "POST" || create.Path != "/api/config/namespaces/{metadata.namespace}/irregular_policys" ||
		create.OperationID != "ves.io.schema.probe.CustomApi.Create" || create.Surface != "config" ||
		create.RequestSchema != "probeCreateRequest" {
		t.Fatalf("Create operation was not preserved exactly: %+v", create)
	}
}

func TestParseOperationCatalogPreservesOptionalResponseSchema(t *testing.T) {
	raw := strings.Replace(
		validOperationCatalog,
		`"surface": "config"`,
		`"surface": "config", "responseSchema": "probeResponse"`,
		1,
	)
	catalog, err := ParseOperationCatalog([]byte(raw))
	if err != nil {
		t.Fatalf("ParseOperationCatalog() error = %v", err)
	}
	identity, ok := catalog.Identity("ves.io.schema.probe")
	if !ok {
		t.Fatal("Identity(ves.io.schema.probe) was not found")
	}
	if got := identity.Operations[0].ResponseSchema; got != "probeResponse" {
		t.Fatalf("ResponseSchema = %q, want probeResponse", got)
	}
	if got := identity.Operations[1].ResponseSchema; got != "" {
		t.Fatalf("absent ResponseSchema = %q, want empty", got)
	}
}

func TestParseOperationCatalogRejectsInvalidResponseSchema(t *testing.T) {
	tests := []struct {
		name     string
		fragment string
		wantErr  string
	}{
		{name: "null", fragment: `"responseSchema": null`, wantErr: "responseSchema must be absent or a non-empty schema name"},
		{name: "empty", fragment: `"responseSchema": ""`, wantErr: "responseSchema must be absent or a non-empty schema name"},
		{name: "non-string", fragment: `"responseSchema": 42`, wantErr: "responseSchema must be absent or a non-empty schema name"},
		{name: "malformed", fragment: `"responseSchema": "probe-response"`, wantErr: "responseSchema must be absent or a non-empty schema name"},
		{name: "duplicate", fragment: `"responseSchema": "first", "responseSchema": "second"`, wantErr: "duplicate field"},
		{name: "unknown field", fragment: `"unexpectedResponseSchema": "probeResponse"`, wantErr: `unknown field "unexpectedResponseSchema"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := strings.Replace(
				validOperationCatalog,
				`"surface": "config"`,
				`"surface": "config", `+tt.fragment,
				1,
			)
			_, err := ParseOperationCatalog([]byte(raw))
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("ParseOperationCatalog() error = %v, want substring %q", err, tt.wantErr)
			}
		})
	}
}

func TestParseOperationCatalogPreservesOptionalOperationRole(t *testing.T) {
	raw := strings.Replace(
		validOperationCatalog,
		`"surface": "config"`,
		`"surface": "config", "role": "query", "terraformName": "probe_query", "responseSchema": "probeResponse"`,
		1,
	)
	catalog, err := ParseOperationCatalog([]byte(raw))
	if err != nil {
		t.Fatalf("ParseOperationCatalog() error = %v", err)
	}
	identity, ok := catalog.Identity("ves.io.schema.probe")
	if !ok {
		t.Fatal("Identity(ves.io.schema.probe) was not found")
	}
	if got := identity.Operations[0].Role; got != "query" {
		t.Fatalf("Role = %q, want query", got)
	}
	if got := identity.Operations[1].Role; got != "" {
		t.Fatalf("absent Role = %q, want empty", got)
	}
}

func TestParseOperationCatalogPreservesTerraformResponseOperationContract(t *testing.T) {
	raw := strings.Replace(
		validOperationCatalog,
		`"surface": "config"`,
		`"surface": "config", "role": "collection", "terraformName": "site_registrations", "responseSchema": "registrationListResponse"`,
		1,
	)
	catalog, err := ParseOperationCatalog([]byte(raw))
	if err != nil {
		t.Fatalf("ParseOperationCatalog() error = %v", err)
	}
	operation := catalog.APIOperations[0].Operations[0]
	if operation.Role != "collection" || operation.TerraformName != "site_registrations" || operation.ResponseSchema != "registrationListResponse" {
		t.Fatalf("response operation contract was not preserved exactly: %+v", operation)
	}
}

func TestParseOperationCatalogRejectsInvalidOperationRole(t *testing.T) {
	tests := []struct {
		name     string
		fragment string
		wantErr  string
	}{
		{name: "null", fragment: `"role": null`, wantErr: "role must be absent, action, collection, issuance, or query"},
		{name: "empty", fragment: `"role": ""`, wantErr: "role must be absent, action, collection, issuance, or query"},
		{name: "non-string", fragment: `"role": 42`, wantErr: "role must be absent, action, collection, issuance, or query"},
		{name: "unsupported", fragment: `"role": "lookup"`, wantErr: "role must be absent, action, collection, issuance, or query"},
		{name: "duplicate", fragment: `"role": "query", "role": "issuance"`, wantErr: "duplicate field"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := strings.Replace(
				validOperationCatalog,
				`"surface": "config"`,
				`"surface": "config", `+tt.fragment,
				1,
			)
			_, err := ParseOperationCatalog([]byte(raw))
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("ParseOperationCatalog() error = %v, want substring %q", err, tt.wantErr)
			}
		})
	}
}

func TestParseOperationCatalogRequiresCompleteResponseOperationContract(t *testing.T) {
	base := strings.Replace(
		validOperationCatalog,
		`"surface": "config"`,
		`"surface": "config", "role": "query", "terraformName": "site_image", "responseSchema": "probeResponse"`,
		1,
	)
	tests := []struct {
		name    string
		mutate  func(string) string
		wantErr string
	}{
		{name: "missing name", mutate: func(raw string) string { return strings.Replace(raw, `, "terraformName": "site_image"`, "", 1) }, wantErr: "terraformName is required"},
		{name: "invalid name", mutate: func(raw string) string { return strings.Replace(raw, `"site_image"`, `"xcsh-site-image"`, 1) }, wantErr: "terraformName"},
		{name: "missing response", mutate: func(raw string) string { return strings.Replace(raw, `, "responseSchema": "probeResponse"`, "", 1) }, wantErr: "responseSchema is required"},
		{name: "name without role", mutate: func(raw string) string { return strings.Replace(raw, `, "role": "query"`, "", 1) }, wantErr: "terraformName requires a response-operation role"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseOperationCatalog([]byte(tt.mutate(base)))
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("ParseOperationCatalog() error = %v, want substring %q", err, tt.wantErr)
			}
		})
	}
}

func TestParseOperationCatalogFromDir(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "api-catalog.json"), []byte(validOperationCatalog), 0o600); err != nil {
		t.Fatal(err)
	}
	catalog, err := ParseOperationCatalogFromDir(dir)
	if err != nil {
		t.Fatalf("ParseOperationCatalogFromDir() error = %v", err)
	}
	if len(catalog.APIOperations) != 1 {
		t.Fatalf("len(APIOperations) = %d, want 1", len(catalog.APIOperations))
	}
}

func TestParseOperationCatalogRejectsInvalidContracts(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(string) string
		wantErr string
	}{
		{
			name: "missing apiOperations",
			mutate: func(raw string) string {
				return mutateCatalogDocument(raw, func(document map[string]interface{}) {
					delete(document, "apiOperations")
				})
			},
			wantErr: "apiOperations",
		},
		{
			name: "missing apiExclusions",
			mutate: func(raw string) string {
				return mutateCatalogDocument(raw, func(document map[string]interface{}) {
					delete(document, "apiExclusions")
				})
			},
			wantErr: "apiExclusions",
		},
		{
			name: "unsupported method",
			mutate: func(raw string) string {
				return strings.Replace(raw, `"method": "GET"`, `"method": "TRACE"`, 1)
			},
			wantErr: "method",
		},
		{
			name: "relative path",
			mutate: func(raw string) string {
				return strings.Replace(raw, `"path": "/api/`, `"path": "api/`, 1)
			},
			wantErr: "path",
		},
		{
			name: "operation identity mismatch",
			mutate: func(raw string) string {
				return strings.Replace(raw, `ves.io.schema.probe.CustomApi.List`, `ves.io.schema.other.CustomApi.List`, 1)
			},
			wantErr: "operationId",
		},
		{
			name: "duplicate identity",
			mutate: func(raw string) string {
				return mutateCatalogDocument(raw, func(document map[string]interface{}) {
					operations := document["apiOperations"].([]interface{})
					document["apiOperations"] = append(operations, operations[0])
				})
			},
			wantErr: "duplicate apiIdentity",
		},
		{
			name: "duplicate JSON field",
			mutate: func(raw string) string {
				return strings.Replace(raw, `"service": "f5xc",`, "\"service\": \"f5xc\",\n  \"service\": \"duplicate\",", 1)
			},
			wantErr: "duplicate field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseOperationCatalog([]byte(tt.mutate(validOperationCatalog)))
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("ParseOperationCatalog() error = %v, want substring %q", err, tt.wantErr)
			}
		})
	}
}

func TestParseOperationCatalogRejectsDuplicateOperationAndExclusionOverlap(t *testing.T) {
	duplicateOperation := mutateCatalogDocument(validOperationCatalog, func(document map[string]interface{}) {
		identity := document["apiOperations"].([]interface{})[0].(map[string]interface{})
		operations := identity["operations"].([]interface{})
		identity["operations"] = append(operations, operations[0])
	})
	if _, err := ParseOperationCatalog([]byte(duplicateOperation)); err == nil || !strings.Contains(err.Error(), "duplicate method/path") {
		t.Fatalf("duplicate operation error = %v", err)
	}

	overlap := strings.Replace(
		validOperationCatalog,
		`"apiExclusions": []`,
		`"apiExclusions": [{"apiIdentity":"ves.io.schema.probe","classification":"path-collision","reason":"synthetic overlap"}]`,
		1,
	)
	if _, err := ParseOperationCatalog([]byte(overlap)); err == nil || !strings.Contains(err.Error(), "both apiOperations and apiExclusions") {
		t.Fatalf("exclusion overlap error = %v", err)
	}
}

func mutateCatalogDocument(raw string, mutate func(map[string]interface{})) string {
	var document map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &document); err != nil {
		panic(err)
	}
	mutate(document)
	result, err := json.Marshal(document)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func TestOperationCatalogResolveResourceUsesExactCatalogPaths(t *testing.T) {
	catalog, err := ParseOperationCatalog([]byte(validOperationCatalog))
	if err != nil {
		t.Fatal(err)
	}
	spec := catalogProbeSpec()

	resolved, err := catalog.ResolveResource(spec, "probe")
	if err != nil {
		t.Fatalf("ResolveResource() error = %v", err)
	}
	if resolved.APIIdentity != "ves.io.schema.probe" {
		t.Fatalf("APIIdentity = %q", resolved.APIIdentity)
	}
	if resolved.CollectionPath != "/api/config/namespaces/%s/irregular_policys" {
		t.Fatalf("CollectionPath = %q; path was reconstructed instead of consumed", resolved.CollectionPath)
	}
	if resolved.ItemPath != "/api/config/namespaces/%s/irregular_policys/%s" {
		t.Fatalf("ItemPath = %q; path was reconstructed instead of consumed", resolved.ItemPath)
	}
	if !resolved.HasNamespace || !resolved.HasCreate {
		t.Fatalf("resolved flags = namespace:%t create:%t", resolved.HasNamespace, resolved.HasCreate)
	}
	if resolved.Create == nil || resolved.Create.OperationID != "ves.io.schema.probe.CustomApi.Create" {
		t.Fatalf("resolved Create = %+v", resolved.Create)
	}
}

func TestOperationCatalogValidatesEveryOperationAgainstOpenAPI(t *testing.T) {
	catalog, err := ParseOperationCatalog([]byte(validOperationCatalog))
	if err != nil {
		t.Fatal(err)
	}
	spec := catalogProbeSpec()
	if err := catalog.ValidateAgainstSpec(spec); err != nil {
		t.Fatalf("ValidateAgainstSpec() error = %v", err)
	}
	delete(spec.Paths, "/api/config/namespaces/{namespace}/irregular_policys/{name}")
	if err := catalog.ValidateAgainstSpec(spec); err == nil || !strings.Contains(err.Error(), "absent from OpenAPI") {
		t.Fatalf("ValidateAgainstSpec() missing operation error = %v", err)
	}
}

func TestOperationCatalogResolvesEveryResponseOperationRole(t *testing.T) {
	raw := mutateCatalogDocument(validOperationCatalog, func(document map[string]interface{}) {
		identity := document["apiOperations"].([]interface{})[0].(map[string]interface{})
		operations := identity["operations"].([]interface{})
		for index, role := range []string{"query", "issuance", "collection", "action"} {
			operation := operations[index].(map[string]interface{})
			operation["role"] = role
			operation["terraformName"] = "probe_" + role
			operation["responseSchema"] = "probeResponse"
			if operation["method"] != "GET" {
				operation["method"] = "POST"
				operation["requestSchema"] = "probeRequest"
			}
		}
	})
	catalog, err := ParseOperationCatalog([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	spec := catalogProbeSpec()
	actionPath := spec.Paths["/api/config/namespaces/{metadata.namespace}/irregular_policys/{metadata.name}"].(map[string]interface{})
	actionPath["post"] = actionPath["put"]
	delete(actionPath, "put")
	spec.Components.Schemas["probeRequest"] = Schema{Type: "object"}
	spec.Components.Schemas["probeResponse"] = Schema{Type: "object"}
	for _, pathValue := range spec.Paths {
		for _, methodValue := range pathValue.(map[string]interface{}) {
			operation := methodValue.(map[string]interface{})
			operation["responses"] = map[string]interface{}{
				"200": map[string]interface{}{"content": map[string]interface{}{
					"application/json": map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/probeResponse"}},
				}},
			}
			if operation["requestBody"] != nil {
				operation["requestBody"] = map[string]interface{}{"content": map[string]interface{}{
					"application/json": map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/probeRequest"}},
				}}
			}
		}
	}

	operations, err := catalog.ResponseOperationsForSpec(spec)
	if err != nil {
		t.Fatalf("ResponseOperationsForSpec() error = %v", err)
	}
	if len(operations) != 4 {
		t.Fatalf("len(operations) = %d, want 4: %+v", len(operations), operations)
	}
	for index, want := range []string{"probe_action", "probe_collection", "probe_issuance", "probe_query"} {
		if operations[index].Name != want {
			t.Fatalf("operations[%d].Name = %q, want %q", index, operations[index].Name, want)
		}
	}
}

func TestOperationCatalogRejectsResponseSchemaDrift(t *testing.T) {
	raw := strings.Replace(
		validOperationCatalog,
		`"surface": "config"`,
		`"surface": "config", "role": "query", "terraformName": "probe_query", "responseSchema": "probeResponse"`,
		1,
	)
	catalog, err := ParseOperationCatalog([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	spec := catalogProbeSpec()
	operation := spec.Paths["/api/config/namespaces/{namespace}/irregular_policys"].(map[string]interface{})["get"].(map[string]interface{})
	operation["responses"] = map[string]interface{}{
		"200": map[string]interface{}{"content": map[string]interface{}{
			"application/json": map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/differentResponse"}},
		}},
	}
	if _, err := catalog.ResponseOperationsForSpec(spec); err == nil || !strings.Contains(err.Error(), "response schema") {
		t.Fatalf("ResponseOperationsForSpec() error = %v, want response-schema mismatch", err)
	}
}

func TestOperationCatalogValidationAcceptsExactDomainAfterPathCollision(t *testing.T) {
	catalog, err := ParseOperationCatalog([]byte(validOperationCatalog))
	if err != nil {
		t.Fatal(err)
	}
	exact := catalogProbeSpec()
	colliding := catalogProbeSpec()
	colliding.Paths["/api/config/namespaces/{namespace}/irregular_policys"].(map[string]interface{})["get"].(map[string]interface{})["operationId"] =
		"ves.io.schema.other.CustomAPI.List"
	if err := catalog.validateAgainstSpecs([]*Spec{colliding, exact}); err != nil {
		t.Fatalf("validateAgainstSpecs() rejected an exact domain after a colliding domain: %v", err)
	}
}

func TestOperationCatalogRequiresExactMethodAndRoleForLifecycle(t *testing.T) {
	raw := mutateCatalogDocument(validOperationCatalog, func(document map[string]interface{}) {
		identity := document["apiOperations"].([]interface{})[0].(map[string]interface{})
		operation := identity["operations"].([]interface{})[2].(map[string]interface{})
		operation["method"] = "GET"
	})
	catalog, err := ParseOperationCatalog([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	spec := catalogProbeSpec()
	operation := spec.Paths["/api/config/namespaces/{metadata.namespace}/irregular_policys"].(map[string]interface{})["post"]
	delete(spec.Paths["/api/config/namespaces/{metadata.namespace}/irregular_policys"].(map[string]interface{}), "post")
	spec.Paths["/api/config/namespaces/{metadata.namespace}/irregular_policys"].(map[string]interface{})["get"] = operation
	resolved, err := catalog.ResolveResource(spec, "probe")
	if err != nil {
		t.Fatal(err)
	}
	if resolved.HasCreate || resolved.Create != nil {
		t.Fatalf("GET operation named Create was misclassified as lifecycle create: %+v", resolved.Create)
	}
}

func TestOperationCatalogResolveResourceFailsClosed(t *testing.T) {
	catalog, err := ParseOperationCatalog([]byte(validOperationCatalog))
	if err != nil {
		t.Fatal(err)
	}
	spec := catalogProbeSpec()
	post := spec.Paths["/api/config/namespaces/{metadata.namespace}/irregular_policys"].(map[string]interface{})["post"].(map[string]interface{})
	post["operationId"] = "ves.io.schema.probe.CustomApi.Different"

	if _, err := catalog.ResolveResource(spec, "probe"); err == nil || !strings.Contains(err.Error(), "operationId") {
		t.Fatalf("ResolveResource() mismatch error = %v", err)
	}
	if _, err := catalog.ResolveResource(spec, "absent"); err == nil || !strings.Contains(err.Error(), "no apiOperations identity") {
		t.Fatalf("ResolveResource() missing identity error = %v", err)
	}
}

func TestOperationCatalogClassifiesListOnlyAPIAsNotManageable(t *testing.T) {
	raw := strings.Replace(validOperationCatalog,
		`"path": "/api/config/namespaces/{namespace}/irregular_policys/{name}"`,
		`"path": "/api/config/namespaces/{namespace}/irregular_policys"`, 1)
	raw = mutateCatalogDocument(raw, func(document map[string]interface{}) {
		identity := document["apiOperations"].([]interface{})[0].(map[string]interface{})
		operations := identity["operations"].([]interface{})
		identity["operations"] = operations[1:3]
	})
	catalog, err := ParseOperationCatalog([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	spec := catalogProbeSpec()
	delete(spec.Paths, "/api/config/namespaces/{namespace}/irregular_policys/{name}")
	spec.Paths["/api/config/namespaces/{namespace}/irregular_policys"].(map[string]interface{})["get"] =
		map[string]interface{}{"operationId": "ves.io.schema.probe.CustomApi.Get"}

	_, err = catalog.ResolveResource(spec, "probe")
	if !errors.Is(err, ErrResourceNotManageable) {
		t.Fatalf("ResolveResource() error = %v, want ErrResourceNotManageable", err)
	}
}

func TestOperationCatalogActionsUseExactCatalogReadPath(t *testing.T) {
	raw := strings.Replace(validOperationCatalog, `"apiOperations": [`, `"apiOperations": [
    {
      "apiIdentity": "ves.io.schema.registration",
      "operations": [
        {
          "method": "GET",
          "path": "/api/register/namespaces/{namespace}/objects_exact/{name}",
          "operationId": "ves.io.schema.registration.CustomAPI.Get",
          "surface": "register"
        },
        {
          "method": "POST",
          "path": "/api/register/namespaces/{namespace}/object/{name}/approve_exact",
          "operationId": "ves.io.schema.registration.CustomAPI.RegistrationApprove",
          "surface": "register",
          "requestSchema": "registrationApprovalReq"
        }
      ]
    },`, 1)
	catalog, err := ParseOperationCatalog([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	spec := catalogProbeSpec()
	spec.Components.Schemas["registrationApprovalReq"] = Schema{XF5xcAction: "approve"}
	spec.Paths["/api/register/namespaces/{namespace}/objects_exact/{name}"] = map[string]interface{}{
		"get": map[string]interface{}{"operationId": "ves.io.schema.registration.CustomAPI.Get"},
	}
	spec.Paths["/api/register/namespaces/{namespace}/object/{name}/approve_exact"] = map[string]interface{}{
		"post": map[string]interface{}{
			"operationId": "ves.io.schema.registration.CustomAPI.RegistrationApprove",
			"requestBody": map[string]interface{}{"content": map[string]interface{}{
				"application/json": map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/registrationApprovalReq"}},
			}},
		},
	}

	actions, err := catalog.ActionsForSpec(spec)
	if err != nil {
		t.Fatalf("ActionsForSpec() error = %v", err)
	}
	if len(actions) != 1 {
		t.Fatalf("len(actions) = %d, want 1: %+v", len(actions), actions)
	}
	action := actions[0]
	if action.ResourceName != "registration_approval" ||
		action.ActionPath != "/api/register/namespaces/%s/object/%s/approve_exact" ||
		action.ReadObjectPath != "/api/register/namespaces/%s/objects_exact/%s" {
		t.Fatalf("action paths were not consumed exactly: %+v", action)
	}
}

func catalogProbeSpec() *Spec {
	operation := func(operationID, requestSchema string) map[string]interface{} {
		result := map[string]interface{}{"operationId": operationID}
		if requestSchema != "" {
			result["requestBody"] = map[string]interface{}{"content": map[string]interface{}{
				"application/json": map[string]interface{}{"schema": map[string]interface{}{"$ref": "#/components/schemas/" + requestSchema}},
			}}
		}
		return result
	}
	return &Spec{
		Paths: map[string]interface{}{
			"/api/config/namespaces/{namespace}/irregular_policys": map[string]interface{}{
				"get": operation("ves.io.schema.probe.CustomApi.List", ""),
			},
			"/api/config/namespaces/{namespace}/irregular_policys/{name}": map[string]interface{}{
				"get": operation("ves.io.schema.probe.CustomApi.Get", ""),
			},
			"/api/config/namespaces/{metadata.namespace}/irregular_policys": map[string]interface{}{
				"post": operation("ves.io.schema.probe.CustomApi.Create", "probeCreateRequest"),
			},
			"/api/config/namespaces/{metadata.namespace}/irregular_policys/{metadata.name}": map[string]interface{}{
				"put": operation("ves.io.schema.probe.CustomApi.Replace", "probeReplaceRequest"),
			},
		},
		Components: Components{Schemas: map[string]Schema{}},
	}
}
