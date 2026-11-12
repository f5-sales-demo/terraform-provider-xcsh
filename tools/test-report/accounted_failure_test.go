// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"strings"
	"testing"
)

func TestAccountedPackageFailureHasCompleteEvidence(t *testing.T) {
	stream := strings.Join([]string{
		goTestEvent("start", "", ""),
		goTestEvent("run", "TestAccFixture", ""),
		goTestEvent("output", "TestAccFixture", "assertion failed"),
		goTestEvent("fail", "TestAccFixture", ""),
		goTestEvent("fail", "", ""),
	}, "\n")
	report := parseTestOutputReader(strings.NewReader(stream), 10, true)
	if !report.Stream.Complete || report.TotalFailed != 1 || report.Stream.PackageFailures != 1 {
		t.Fatalf("terminal failed stream was mislabeled: %+v", report.Stream)
	}
}

func TestUnexplainedPackageFailureIsNotCoveredByAnotherPackage(t *testing.T) {
	stream := strings.Join([]string{
		goTestEvent("start", "", ""),
		goTestEvent("run", "TestAccFixture", ""),
		goTestEvent("fail", "TestAccFixture", ""),
		goTestEvent("fail", "", ""),
		`{"Action":"start","Package":"other"}`,
		`{"Action":"fail","Package":"other"}`,
	}, "\n")
	report := parseTestOutputReader(strings.NewReader(stream), 10, true)
	if report.Stream.Complete || report.Stream.PackageFailures != 2 {
		t.Fatal("unexplained package failure accepted")
	}
}
