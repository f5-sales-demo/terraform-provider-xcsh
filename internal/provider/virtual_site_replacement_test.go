// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestVirtualSiteCreationFieldsRequireReplacement(t *testing.T) {
	var response resource.SchemaResponse
	(&VirtualSiteResource{}).Schema(context.Background(), resource.SchemaRequest{}, &response)
	selector := response.Schema.Blocks["site_selector"].(schema.SingleNestedBlock)
	if len(selector.PlanModifiers) == 0 {
		t.Fatal("virtual-site selector is absent from ReplaceSpecType and must require recreation")
	}
	siteType := response.Schema.Attributes["site_type"].(schema.StringAttribute)
	found := false
	for _, modifier := range siteType.PlanModifiers {
		if modifier.Description(context.Background()) == "If the value of this attribute changes, Terraform will destroy and recreate the resource." {
			found = true
		}
	}
	if !found {
		t.Fatal("virtual-site site_type must require recreation")
	}
}
