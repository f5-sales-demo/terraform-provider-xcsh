// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"testing"
)

func TestLogReceiverDatasourceFixtureChoosesSyslog(t *testing.T) {
	file, diags := hclparse.NewParser().ParseHCL([]byte(testAccLogReceiverDataSourceConfig_basic("fixture-ns", "fixture")), "fixture.tf")
	if diags.HasErrors() {
		t.Fatal(diags)
	}
	for _, block := range file.Body.(*hclsyntax.Body).Blocks {
		if block.Type != "resource" || block.Labels[0] != "xcsh_log_receiver" {
			continue
		}
		receiver, location := false, false
		for _, nested := range block.Body.Blocks {
			receiver = receiver || nested.Type == "syslog"
			location = location || nested.Type == "site_local"
		}
		if !receiver || !location {
			t.Fatal("log receiver requires receiver and location choices")
		}
		return
	}
	t.Fatal("log receiver fixture resource missing")
}
