package provider

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestSiteRegistrationCertifiedHardwareSchema(t *testing.T) {
	var response datasource.SchemaResponse
	(&SiteRegistrationDataSource{}).Schema(context.Background(), datasource.SchemaRequest{}, &response)
	attribute, ok := response.Schema.Attributes["certified_hw"].(schema.StringAttribute)
	if !ok || !attribute.Computed || attribute.Optional || attribute.Required {
		t.Fatalf("certified_hw must be a computed-only string: %#v", response.Schema.Attributes["certified_hw"])
	}
}

func TestSiteRegistrationCertifiedHardwareResponse(t *testing.T) {
	for _, test := range []struct {
		name, wire string
		expected   types.String
	}{
		{"reported", `{"name":"r-test","get_spec":{"infra":{"certified_hw":"site-v2-generic-chw","provider_ref":"AWS","instance_id":"synthetic-instance"}}}`, types.StringValue("site-v2-generic-chw")},
		{"omitted", `{"name":"r-test","get_spec":{"infra":{"provider_ref":"AWS","instance_id":"synthetic-instance"}}}`, types.StringNull()},
	} {
		t.Run(test.name, func(t *testing.T) {
			var registration client.RegistrationListItem
			if err := json.Unmarshal([]byte(test.wire), &registration); err != nil {
				t.Fatal(err)
			}
			state := SiteRegistrationDataSourceModel{}
			applyRegistrationMatch(&state, &registration)
			if !state.CertifiedHW.Equal(test.expected) {
				t.Fatalf("certified_hw = %v, want %v", state.CertifiedHW, test.expected)
			}
			if state.ProviderType.ValueString() != "AWS" || state.InstanceID.ValueString() != "synthetic-instance" {
				t.Fatal("unrelated registration identity changed")
			}
		})
	}
}
