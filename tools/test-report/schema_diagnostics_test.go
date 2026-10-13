// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestImportDiagnosticUsesOnlySchemaFieldNames(t *testing.T) {
	original := publicSchemaFields
	defer func() { publicSchemaFields = original }()
	publicSchemaFields = map[string]bool{"detection_settings.default_bot_setting": true}
	output := `ImportStateVerify attributes not equivalent.
- "detection_settings.0.default_bot_setting.%": "0"
+ "annotations.SECRET_SENTINEL": "PRIVATE_VALUE"
`
	f := classifyFailureFingerprint(output)
	if len(f.SchemaFields) != 1 || f.SchemaFields[0] != "detection_settings.default_bot_setting" {
		t.Fatal(f)
	}
	data, _ := json.Marshal(f)
	if strings.Contains(string(data), "SECRET_SENTINEL") || strings.Contains(string(data), "PRIVATE_VALUE") {
		t.Fatal("state leaked")
	}
}
