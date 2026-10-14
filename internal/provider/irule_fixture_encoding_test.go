// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"testing"
)

func TestIRuleFixturesEncodeAPIContent(t *testing.T) {
	for name, config := range map[string]string{
		"resource":   testAccIruleConfig_basic("fixture-ns", "fixture"),
		"datasource": testAccIruleDataSourceConfig_basic("fixture-ns", "fixture"),
	} {
		t.Run(name, func(t *testing.T) {
			file, diags := hclparse.NewParser().ParseHCL([]byte(config), "fixture.tf")
			if diags.HasErrors() {
				t.Fatal(diags)
			}
			for _, block := range file.Body.(*hclsyntax.Body).Blocks {
				if block.Type != "resource" || block.Labels[0] != "xcsh_irule" {
					continue
				}
				content := block.Body.Attributes["irule"]
				if content == nil {
					t.Fatal("iRule fixture content missing")
				}
				call, ok := content.Expr.(*hclsyntax.FunctionCallExpr)
				if !ok || call.Name != "base64encode" || len(call.Args) != 1 {
					t.Fatal("live API requires base64 encoded iRule content")
				}
				return
			}
			t.Fatal("iRule fixture resource missing")
		})
	}
}
