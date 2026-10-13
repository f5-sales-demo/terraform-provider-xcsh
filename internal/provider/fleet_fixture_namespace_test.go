// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"testing"
)

func TestFleetDatasourceUsesRequiredSystemNamespace(t *testing.T) {
	file, diags := hclparse.NewParser().ParseHCL([]byte(testAccFleetDataSourceConfig_basic("fixture")), "fixture.tf")
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	for _, block := range file.Body.(*hclsyntax.Body).Blocks {
		if block.Type != "resource" || block.Labels[0] != "xcsh_fleet" {
			continue
		}
		ns, diags := block.Body.Attributes["namespace"].Expr.Value(nil)
		if diags.HasErrors() || !ns.IsWhollyKnown() || ns.AsString() != "system" {
			t.Fatal("fleet API requires system namespace")
		}
		return
	}
	t.Fatal("fleet fixture missing")
}
