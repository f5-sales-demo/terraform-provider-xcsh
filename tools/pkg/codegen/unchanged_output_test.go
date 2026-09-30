// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestResourceGenerationDoesNotRewriteIdenticalOutput(t *testing.T) {
	resource := &openapi.ResourceTemplate{Name: "unchanged_probe", TitleCase: "UnchangedProbe", Description: "Probe.", APIPath: "/api/config/namespaces/%s/unchanged_probes", APIPathItem: "/api/config/namespaces/%s/unchanged_probes/%s", HasNamespaceInPath: true}
	dir := t.TempDir()
	if err := GenerateResourceFile(resource, dir); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "unchanged_probe_resource.go")
	fixed := time.Unix(1000000000, 0)
	if err := os.Chtimes(path, fixed, fixed); err != nil {
		t.Fatal(err)
	}
	before, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := GenerateResourceFile(resource, dir); err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(before, after) || !before.ModTime().Equal(after.ModTime()) {
		t.Fatal("unchanged resource was rewritten")
	}
	resource.Description = "Changed description."
	if err := GenerateResourceFile(resource, dir); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Fatal("changed resource output is empty")
	}
	final, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if final.ModTime().Equal(before.ModTime()) {
		t.Fatal("changed output was skipped")
	}
}
