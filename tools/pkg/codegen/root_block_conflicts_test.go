// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package codegen

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"strings"
	"testing"
)

func TestRootOneOfBlockConflicts(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{GoName: "HTTPHealthCheck", TfsdkTag: "http_health_check", IsBlock: true, OneOfGroup: "health_check"},
		{GoName: "TCPHealthCheck", TfsdkTag: "tcp_health_check", IsBlock: true, OneOfGroup: "health_check"},
	}
	got := RenderRootBlockConflicts(attrs)
	if !strings.Contains(got, `validators.ConfiguredObject(ctx, req.Config, path.Root("http_health_check"), &resp.Diagnostics)`) || !strings.Contains(got, `validators.ConfiguredObject(ctx, req.Config, path.Root("tcp_health_check"), &resp.Diagnostics)`) || strings.Contains(got, "!= nil") || !strings.Contains(got, "Conflicting Configuration") {
		t.Fatal(got)
	}
}

func TestRootBlockScalarConflictsRetainKnownScalarCheck(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{GoName: "Block", TfsdkTag: "block", IsBlock: true, OneOfGroup: "choice"},
		{GoName: "Marker", TfsdkTag: "marker", Type: "object", OneOfGroup: "choice"},
	}
	got := RenderRootBlockConflicts(attrs)
	if !strings.Contains(got, `validators.ConfiguredObject(ctx, req.Config, path.Root("block"), &resp.Diagnostics)`) || !strings.Contains(got, "!data.Marker.IsNull() && !data.Marker.IsUnknown()") {
		t.Fatal(got)
	}
}

func TestRootRepeatedBlockConflictsUseListPresence(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{GoName: "Blocks", TfsdkTag: "blocks", IsBlock: true, NestedBlockType: "list", OneOfGroup: "choice"},
		{GoName: "Marker", TfsdkTag: "marker", Type: "object", OneOfGroup: "choice"},
	}
	got := RenderRootBlockConflicts(attrs)
	if !strings.Contains(got, `validators.ConfiguredList(ctx, req.Config, path.Root("blocks"), &resp.Diagnostics)`) || strings.Contains(got, "!= nil") {
		t.Fatal(got)
	}
}
