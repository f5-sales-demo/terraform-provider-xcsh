// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TestSiteImageDataSourceSchema(t *testing.T) {
	t.Parallel()
	d := NewSiteImageDataSource()
	resp := &datasource.SchemaResponse{}
	d.Schema(context.Background(), datasource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Schema returned diagnostics: %v", resp.Diagnostics)
	}

	platformAttribute, ok := resp.Schema.Attributes["platform"].(schema.StringAttribute)
	if !ok || !platformAttribute.Required {
		t.Fatalf("platform must be a required string attribute")
	}
	for _, attributeName := range []string{"image_download_url", "image_md5_download_url"} {
		attribute, ok := resp.Schema.Attributes[attributeName].(schema.StringAttribute)
		if !ok || !attribute.Computed || !attribute.Sensitive {
			t.Errorf("%s must be a sensitive computed string attribute", attributeName)
		}
	}
}
