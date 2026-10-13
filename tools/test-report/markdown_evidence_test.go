// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"os"
	"strings"
	"testing"
)

func TestMarkdownRetainsPrerequisitesAndIncompleteFailure(t *testing.T) {
	stream := goTestEvent("start", "", "") + "\n" + goTestEvent("run", "TestAccPrerequisite", "") + "\n" + goTestEvent("output", "TestAccPrerequisite", "[PREREQUISITE:operation_denied method=POST type=xcsh_namespace status=403] SECRET_SENTINEL") + "\n" + goTestEvent("skip", "TestAccPrerequisite", "")
	report := parseTestOutputReader(strings.NewReader(stream), 10, true)
	file, err := os.CreateTemp(t.TempDir(), "report")
	if err != nil {
		t.Fatal(err)
	}
	outputMarkdown(file, report, true)
	_ = file.Close()
	data, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Overall: ❌ FAIL", "operation_denied", "xcsh_namespace", "403"} {
		if !strings.Contains(string(data), want) {
			t.Errorf("Markdown lost %s", want)
		}
	}
	if strings.Contains(string(data), "SECRET_SENTINEL") {
		t.Fatal("Markdown leaks")
	}
}
