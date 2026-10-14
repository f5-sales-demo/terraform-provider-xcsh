// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"strings"
	"testing"
)

func TestPackageLifecycleRequiresStartAndTerminal(t *testing.T) {
	cases := map[string]string{
		"orphan package terminal":   `{"Action":"pass","Package":"example.invalid/provider"}`,
		"unstarted test package":    goTestEvent("run", "TestAccFixture", "") + "\n" + goTestEvent("pass", "TestAccFixture", ""),
		"package ended before test": `{"Action":"start","Package":"example.invalid/provider"}` + "\n" + `{"Action":"pass","Package":"example.invalid/provider"}` + "\n" + goTestEvent("run", "TestAccFixture", "") + "\n" + goTestEvent("pass", "TestAccFixture", ""),
	}
	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			if report := parseTestOutputReader(strings.NewReader(input), 10, true); report.Stream.Complete {
				t.Fatal("invalid package lifecycle accepted")
			}
		})
	}
}
