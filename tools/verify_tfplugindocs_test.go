// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestPublishedTfplugindocsMetadata(t *testing.T) {
	root, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	script := filepath.Join(root, "scripts", "verify-tfplugindocs.sh")
	cases := []struct {
		name, version, checksum, exit string
		pass                          bool
	}{
		{"clean pinned install", "v0.25.0", "unknown", "0", true},
		{"official release build", "v0.25.0+dirty", "7e3e9e12d913576b3c1cb4e3ac15e68b15648452dfa2a28fc11fc21bdc2fb3eb", "0", true},
		{"unknown modified build", "v0.25.0+dirty", strings.Repeat("a", 64), "0", false},
		{"wrong version", "v0.24.0", "unknown", "0", false},
		{"wrong version with official checksum", "v0.24.0+dirty", "7e3e9e12d913576b3c1cb4e3ac15e68b15648452dfa2a28fc11fc21bdc2fb3eb", "0", false},
		{"failed metadata command", "v0.25.0", "unknown", "1", false},
		{"missing module", "", "unknown", "0", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bin := t.TempDir()
			for name, body := range map[string]string{
				"tfplugindocs": "#!/bin/sh\nexit 0\n",
				"go":           "#!/bin/sh\nprintf '\\tmod\\tgithub.com/hashicorp/terraform-plugin-docs\\t%s\\th1:fixture\\n' \"$TEST_VERSION\"\nexit \"$TEST_EXIT\"\n",
				"sha256sum":    "#!/bin/sh\nprintf '%s  %s\\n' \"$TEST_CHECKSUM\" \"$1\"\n",
			} {
				if err := os.WriteFile(filepath.Join(bin, name), []byte(body), 0o700); err != nil {
					t.Fatal(err)
				}
			}
			cmd := exec.Command("bash", script)
			cmd.Env = append(os.Environ(), "PATH="+bin+string(os.PathListSeparator)+os.Getenv("PATH"), "TEST_VERSION="+tc.version, "TEST_EXIT="+tc.exit, "TEST_CHECKSUM="+tc.checksum)
			var stdout, stderr strings.Builder
			cmd.Stdout, cmd.Stderr = &stdout, &stderr
			err := cmd.Run()
			if (err == nil) != tc.pass {
				t.Fatalf("success=%v want=%v: %s", err == nil, tc.pass, stderr.String())
			}
			if tc.pass && stdout.String() != "v0.25.0\n" {
				t.Fatalf("manifest version is not canonical: %q", stdout.String())
			}
			if !tc.pass && stdout.Len() != 0 {
				t.Fatalf("failed verification emitted canonical version: %q", stdout.String())
			}
		})
	}
}
