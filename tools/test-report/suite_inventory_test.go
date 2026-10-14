// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"strings"
	"testing"
)

func TestSuiteInventoryRejectsMissingOrUnexpectedTests(t *testing.T) {
	stream := goTestEvent("start", "", "") + "\n" + goTestEvent("run", "TestAccFixture", "") + "\n" + goTestEvent("pass", "TestAccFixture", "") + "\n" + goTestEvent("pass", "", "")
	for _, names := range [][]string{{"TestAccFixture", "TestAccMissing"}, {"TestAccOther"}} {
		report := parseTestOutputReader(strings.NewReader(stream), 10, true)
		verifySuiteInventory(report, names)
		if report.Stream.Complete {
			t.Fatal("partial suite accepted")
		}
	}
	report := parseTestOutputReader(strings.NewReader(stream), 10, true)
	verifySuiteInventory(report, []string{"TestAccFixture"})
	if !report.Stream.Complete {
		t.Fatal("complete suite rejected")
	}
}
