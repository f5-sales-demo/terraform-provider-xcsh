package provider_test

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/provider"
)

func TestSecuremeshSiteV2NameRejectsNamesBeyondDNS1035Limit(t *testing.T) {
	var response resource.SchemaResponse
	(&provider.SecuremeshSiteV2Resource{}).Schema(context.Background(), resource.SchemaRequest{}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
	name, ok := response.Schema.Attributes["name"].(schema.StringAttribute)
	if !ok {
		t.Fatal("SecureMesh Site V2 name must be a string attribute")
	}

	for _, test := range []struct {
		length  int
		invalid bool
	}{
		{length: 63},
		{length: 64, invalid: true},
	} {
		t.Run(strings.Repeat("a", test.length), func(t *testing.T) {
			var diagnostics validator.StringResponse
			for _, nameValidator := range name.Validators {
				nameValidator.ValidateString(context.Background(), validator.StringRequest{
					Path: path.Root("name"), ConfigValue: types.StringValue(strings.Repeat("a", test.length)),
				}, &diagnostics)
			}
			if diagnostics.Diagnostics.HasError() != test.invalid {
				t.Fatalf("length %d invalid=%t, want %t: %v", test.length, diagnostics.Diagnostics.HasError(), test.invalid, diagnostics.Diagnostics)
			}
		})
	}
}
