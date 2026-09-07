// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestSMSv2RetainedSelectionAndDrainChoicesReachTerraform(t *testing.T) {
	var response resource.SchemaResponse
	(&SecuremeshSiteV2Resource{}).Schema(context.Background(), resource.SchemaRequest{}, &response)
	selection := response.Schema.Blocks["re_select"].(schema.SingleNestedBlock)
	if _, ok := selection.Attributes["specific_geography"].(schema.StringAttribute); !ok {
		t.Error("geographic selection retained by the current API is absent from Terraform")
	}
	upgrade := response.Schema.Blocks["upgrade_settings"].(schema.SingleNestedBlock)
	drain := upgrade.Blocks["kubernetes_upgrade_drain"].(schema.SingleNestedBlock).Blocks["enable_upgrade_drain"].(schema.SingleNestedBlock)
	if _, ok := drain.Attributes["drain_max_unavailable_node_percentage"].(schema.Int64Attribute); !ok {
		t.Error("percentage drain setting retained by the current API is absent from Terraform")
	}
}
