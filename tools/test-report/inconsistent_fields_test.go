// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"strings"
	"testing"
)

func TestInconsistentReadbackExposesOnlySchemaPath(t *testing.T) {
	original := publicSchemaFields
	defer func() { publicSchemaFields = original }()
	publicSchemaFields = map[string]bool{"routes.route_direct_response.response_body_encoded": true}
	output := `Provider produced inconsistent result after apply: .routes[0].route_direct_response.response_body_encoded: was SECRET_SENTINEL, but now PRIVATE_VALUE.`
	f := classifyFailureFingerprint(output)
	if len(f.SchemaFields) != 1 || f.SchemaFields[0] != "routes.route_direct_response.response_body_encoded" {
		t.Fatal("readback path lost")
	}
	if strings.Contains(f.Summary, "SECRET_SENTINEL") {
		t.Fatal("payload leaked")
	}
}
