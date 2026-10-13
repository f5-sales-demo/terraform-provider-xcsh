// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestSyntheticStreamRetainsClassificationsWithoutSecrets(t *testing.T) {
	sentinel := "SECRET_SENTINEL \"quotes\" \x00\x01"
	stream := strings.Join([]string{
		goTestEvent("run", "TestAccPanic", ""),
		goTestEvent("output", "TestAccPanic", "panic: "+sentinel),
		goTestEvent("fail", "TestAccPanic", ""),
		goTestEvent("run", "TestAccUnknown", ""),
		goTestEvent("output", "TestAccUnknown", "unrecognized "+sentinel),
		goTestEvent("fail", "TestAccUnknown", ""),
	}, "\n")
	report := parseTestOutputReader(strings.NewReader(stream), 10, true)
	if !report.Stream.Complete || report.TotalFailed != 2 {
		t.Fatal("valid stream lost")
	}
	encoded, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(encoded, []byte("SECRET_SENTINEL")) {
		t.Fatal("JSON leaks")
	}
	for _, format := range []string{"markdown", "json"} {
		file, err := os.CreateTemp(t.TempDir(), "report")
		if err != nil {
			t.Fatal(err)
		}
		if format == "markdown" {
			outputMarkdown(file, report, true)
		} else {
			outputJSON(file, report)
		}
		file.Close()
		data, err := os.ReadFile(file.Name())
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(data), "SECRET_SENTINEL") || !strings.Contains(string(data), "unclassified") || !strings.Contains(string(data), "assertion") {
			t.Fatalf("%s classification lost or secret exposed", format)
		}
	}
}
func TestDuplicateTerminalIsIncomplete(t *testing.T) {
	stream := goTestEvent("run", "TestAccDuplicate", "") + "\n" + goTestEvent("pass", "TestAccDuplicate", "") + "\n" + goTestEvent("fail", "TestAccDuplicate", "")
	report := parseTestOutputReader(strings.NewReader(stream), 10, true)
	if report.Stream.Complete || report.Stream.InvalidEvents != 1 {
		t.Fatal("duplicate terminal accepted")
	}
}
