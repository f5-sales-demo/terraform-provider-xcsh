// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func TestHTTPDatasourceFixtureChoosesAdvertisement(t *testing.T) {
	parser := hclparse.NewParser()
	file, diagnostics := parser.ParseHCL([]byte(testAccHttpLoadbalancerDataSourceConfig_basic("fixture")), "fixture.tf")
	if diagnostics.HasErrors() {
		t.Fatal(diagnostics)
	}
	resources, _, diagnostics := file.Body.PartialContent(&hcl.BodySchema{Blocks: []hcl.BlockHeaderSchema{{Type: "resource", LabelNames: []string{"type", "name"}}}})
	if diagnostics.HasErrors() {
		t.Fatal(diagnostics)
	}
	for _, block := range resources.Blocks {
		if block.Labels[0] != "xcsh_http_loadbalancer" {
			continue
		}
		body, ok := block.Body.(*hclsyntax.Body)
		if !ok {
			t.Fatal("fixture body must use native HCL")
		}
		choice, ok := body.Attributes["advertise_on_public_default_vip"]
		if !ok {
			t.Fatal("HTTP datasource prerequisite has no explicit advertisement choice")
		}
		value, diagnostics := choice.Expr.Value(nil)
		if diagnostics.HasErrors() || !value.IsWhollyKnown() || value.LengthInt() != 0 {
			t.Fatal("advertisement choice must be an empty marker")
		}
		return
	}
	t.Fatal("HTTP datasource fixture resource missing")
}
