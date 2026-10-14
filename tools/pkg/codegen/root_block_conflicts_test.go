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
	if !strings.Contains(got, "data.HTTPHealthCheck != nil && data.TCPHealthCheck != nil") || !strings.Contains(got, "Conflicting Configuration") {
		t.Fatal(got)
	}
}
