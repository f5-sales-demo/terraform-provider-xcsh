// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"
)

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errors.New("SECRET_SENTINEL") }

func TestStreamCompletenessFailsClosed(t *testing.T) {
	cases := map[string]string{
		"malformed":           "SECRET_SENTINEL\n",
		"invalid action":      `{"Action":"bogus","Package":"example.invalid/provider"}`,
		"missing terminal":    goTestEvent("run", "TestAccInterrupted", ""),
		"orphan terminal":     goTestEvent("pass", "TestAccOrphan", ""),
		"package failure":     `{"Action":"start","Package":"example.invalid/provider"}` + "\n" + `{"Action":"fail","Package":"example.invalid/provider"}`,
		"package interrupted": `{"Action":"start","Package":"example.invalid/provider"}`,
		"oversized":           strings.Repeat("x", 11*1024*1024),
		"empty":               "",
	}
	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			report := parseTestOutputReader(strings.NewReader(input), 10, true)
			if report.Stream.Complete {
				t.Fatal("unsafe stream accepted")
			}
			assertTenantSafe(t, reportJSON(t, report), []string{"SECRET_SENTINEL"})
		})
	}
	report := parseTestOutputReader(io.MultiReader(strings.NewReader(goTestEvent("run", "TestAccRead", "")+"\n"), failingReader{}), 10, true)
	if report.Stream.Complete || !report.Stream.ScannerError || report.TotalFailed != 1 {
		t.Fatal("read failure or unfinished test lost")
	}
}
func reportJSON(t *testing.T, r *Report) string {
	t.Helper()
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestReproducedCleanupFingerprint(t *testing.T) {
	for _, s := range []string{"no API verifier registered for release-surface resource type", "no deleter registered for release-surface resource type"} {
		f := classifyFailureFingerprint(s + " SECRET_SENTINEL")
		if f.Class != FailureClassCleanup || f.Summary != "Missing cleanup verifier" {
			t.Fatalf("got %#v", f)
		}
	}
}
