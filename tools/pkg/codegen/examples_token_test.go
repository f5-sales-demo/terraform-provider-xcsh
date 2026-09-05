// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestRenderTokenResourceExampleUsesSiteBoundJWT(t *testing.T) {
	hcl := RenderResourceExampleHCL(&openapi.ResourceTemplate{}, "token", "system")
	for _, want := range []string{
		"type      = 1",
		"site_name = \"example-securemesh-site\"",
	} {
		if !strings.Contains(hcl, want) {
			t.Fatalf("token example is missing %q", want)
		}
	}
}
