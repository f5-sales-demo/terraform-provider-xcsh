// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExplicitPrerequisiteEvidenceIsSanitized(t *testing.T) {
	input := goTestEvent("run", "TestAccPrerequisite", "") + "\n" + goTestEvent("output", "TestAccPrerequisite", "[PREREQUISITE:external_credential_missing method=GET type=cloud] SECRET_SENTINEL") + "\n" + goTestEvent("skip", "TestAccPrerequisite", "")
	r := parseTestOutputReader(strings.NewReader(input), 10, true)
	if len(r.Tests) != 1 || len(r.Tests[0].Prerequisites) != 1 || r.Tests[0].Prerequisites[0].Outcome != "external_credential_missing" {
		t.Fatal("prerequisite evidence lost")
	}
	b, _ := json.Marshal(r)
	if strings.Contains(string(b), "SECRET_SENTINEL") {
		t.Fatal("prerequisite evidence leaks output")
	}
}
