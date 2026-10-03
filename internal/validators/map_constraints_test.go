package validators

import (
	"context"
	"encoding/base64"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"testing"
)

func TestMapConstraintCustomErrorBoundary(t *testing.T) {
	raw := `{"constraintType":"map","keys":{"type":"uint32-string","ranges":[[3,3],[4,4],[5,5],[300,599]]},"values":{"type":"string","maxLength":65536,"format":"uri-reference"},"cardinality":{"maxProperties":16}}`
	v := MapConstraintsValidator(raw)
	for _, tc := range []struct {
		key     string
		size    int
		invalid bool
	}{{"300", 49143, false}, {"300", 49144, true}, {"299", 1, true}, {"600", 1, true}, {"3", 1, false}} {
		value := "string:///" + base64.StdEncoding.EncodeToString(make([]byte, tc.size))
		m := types.MapValueMust(types.StringType, map[string]attr.Value{tc.key: types.StringValue(value)})
		response := &validator.MapResponse{}
		v.ValidateMap(context.Background(), validator.MapRequest{ConfigValue: m}, response)
		if response.Diagnostics.HasError() != tc.invalid {
			t.Fatalf("key=%s size=%d errors=%v", tc.key, tc.size, response.Diagnostics)
		}
	}
}
