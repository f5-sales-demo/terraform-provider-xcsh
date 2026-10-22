package schema

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"testing"
)

func TestDefaultPoolRouteDoesNotRequireExplicitOriginPools(t *testing.T) {
	s := openapi.Schema{Type: "array", Items: &openapi.Schema{Type: "object", Properties: map[string]openapi.Schema{"weight": {Type: "integer"}}}, XF5XCRequiredFor: openapi.RequiredFor{Create: true}}
	for _, tc := range []struct {
		path     string
		required bool
	}{
		{"routes.simple_route.origin_pools", false},
		{"other.origin_pools", true},
	} {
		a := ConvertToTerraformAttributeWithDepth("origin_pools", s, false, "", &openapi.Spec{}, 3, tc.path)
		if a.CreateRequired != tc.required {
			t.Fatalf("%s: CreateRequired=%v, want %v", tc.path, a.CreateRequired, tc.required)
		}
	}
}
