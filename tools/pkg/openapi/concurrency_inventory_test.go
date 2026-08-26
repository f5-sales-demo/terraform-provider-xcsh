// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"strings"
	"testing"
)

const validConcurrencyInventory = `{
  "version":"2.1.225",
  "eligible_count":1,
  "covered_count":1,
  "excluded_count":1,
  "resources":[{
    "api_identity":"ves.io.schema.probe.API",
    "get":{"path":"/api/config/namespaces/{namespace}/probes/{name}","schema":"probeGetResponse"},
    "replace":{"path":"/api/config/namespaces/{metadata.namespace}/probes/{metadata.name}","schema":"probeReplaceRequest"},
    "create_schema":"probeCreateRequest",
    "token":"resource_version"
  }],
  "exclusions":[{
    "api_identity":"ves.io.schema.registration.API",
    "operation":"Replace",
    "reason":"enrollment command"
  }]
}`

func TestParseConcurrencyInventoryRejectsInvalidPermutations(t *testing.T) {
	tests := map[string]string{
		"missing counts":        strings.Replace(validConcurrencyInventory, `"eligible_count":1,`, "", 1),
		"count mismatch":        strings.Replace(validConcurrencyInventory, `"covered_count":1`, `"covered_count":2`, 1),
		"unsupported token":     strings.Replace(validConcurrencyInventory, `"resource_version"`, `"etag"`, 1),
		"empty exclusion":       strings.Replace(validConcurrencyInventory, `"enrollment command"`, `""`, 1),
		"unsupported operation": strings.Replace(validConcurrencyInventory, `"operation":"Replace"`, `"operation":"Update"`, 1),
		"unknown field":         strings.Replace(validConcurrencyInventory, `"version":"2.1.225"`, `"version":"2.1.225","extra":true`, 1),
		"duplicate field":       strings.Replace(validConcurrencyInventory, `"version":"2.1.225"`, `"version":"2.1.225","version":"2.1.225"`, 1),
	}
	for name, raw := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := ParseConcurrencyInventory([]byte(raw)); err == nil {
				t.Fatal("ParseConcurrencyInventory accepted an invalid contract")
			}
		})
	}
}

func TestConcurrencyInventoryClassifiesCoveredExcludedAndMissingReplace(t *testing.T) {
	inventory, err := ParseConcurrencyInventory([]byte(validConcurrencyInventory))
	if err != nil {
		t.Fatal(err)
	}
	covered := CatalogOperation{Method: "PUT", Path: "/api/config/namespaces/{metadata.namespace}/probes/{metadata.name}", OperationID: "ves.io.schema.probe.API.Replace", RequestSchema: "probeReplaceRequest"}
	if ok, exclusion, err := inventory.ClassifyReplace(covered); err != nil || !ok || exclusion != nil {
		t.Fatalf("covered classification = %v, %+v, %v", ok, exclusion, err)
	}
	excluded := CatalogOperation{Method: "PUT", OperationID: "ves.io.schema.registration.API.Replace"}
	if ok, exclusion, err := inventory.ClassifyReplace(excluded); err != nil || ok || exclusion == nil {
		t.Fatalf("excluded classification = %v, %+v, %v", ok, exclusion, err)
	}
	missing := CatalogOperation{Method: "PUT", OperationID: "ves.io.schema.missing.API.Replace"}
	if _, _, err := inventory.ClassifyReplace(missing); err == nil {
		t.Fatal("unclassified Replace was accepted")
	}
}
