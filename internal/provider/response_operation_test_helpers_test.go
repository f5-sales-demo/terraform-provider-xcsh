package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func responseOperationRaw(t *testing.T, model interface{}, schemaType attr.Type) tftypes.Value {
	t.Helper()
	ctx := context.Background()
	var value types.Object
	if diagnostics := tfsdk.ValueFrom(ctx, model, schemaType, &value); diagnostics.HasError() {
		t.Fatalf("encode operation config: %v", diagnostics)
	}
	raw, err := value.ToTerraformValue(ctx)
	if err != nil {
		t.Fatalf("convert operation config: %v", err)
	}
	return raw
}
