package schema

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestMeasuredNestedAPIDefaults(t *testing.T) {
	tests := []struct {
		parent, field, kind string
		value               any
	}{
		{"ves.io.schema.views.cdn_loadbalancer.CdnOriginPoolType", "origin_request_timeout", "string", "60s"},
		{"ves.io.schema.views.workload.ContainerType", "init_container", "boolean", false},
		{"ves.io.schema.views.workload.ImageType", "pull_policy", "string", "IMAGE_PULL_POLICY_DEFAULT"},
	}
	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			prop := openapi.Schema{Type: tt.kind}
			if tt.field == "pull_policy" {
				prop = openapi.Schema{AllOf: []openapi.Schema{{Ref: "#/components/schemas/PullPolicy"}}}
			}
			spec := &openapi.Spec{Components: openapi.Components{Schemas: map[string]openapi.Schema{
				"PullPolicy": {Type: "string", Enum: []interface{}{"IMAGE_PULL_POLICY_DEFAULT", "IMAGE_PULL_POLICY_ALWAYS"}},
			}}}
			parent := openapi.Schema{Type: "object", XVesProtoMessage: tt.parent, Properties: map[string]openapi.Schema{tt.field: prop}}
			before, _ := json.Marshal(parent)
			for _, path := range []string{"origin_pool", "service.containers", "stateful_service.containers"} {
				attr := findAttr(ExtractNestedAttributes(parent, spec, 3, path), tt.field)
				if attr == nil || !attr.Optional || !attr.Computed || !attr.ServerDefault || attr.PlanModifier != "UseStateForUnknown" || !reflect.DeepEqual(attr.Default, tt.value) {
					t.Errorf("%s omitted default is not accepted: %+v", path, attr)
				}
			}
			after, _ := json.Marshal(parent)
			if string(before) != string(after) {
				t.Fatal("conversion mutated pinned input schema")
			}
			parent.Required = []string{tt.field}
			attr := findAttr(ExtractNestedAttributes(parent, spec, 3, "required"), tt.field)
			if attr == nil || !attr.Required || attr.Computed {
				t.Fatalf("required input authority lost: %+v", attr)
			}
			parent.Required = nil
			prop.Default = "authoritative-upstream-default"
			parent.Properties[tt.field] = prop
			attr = findAttr(ExtractNestedAttributes(parent, spec, 3, "explicit"), tt.field)
			if attr == nil || attr.Default != "authoritative-upstream-default" {
				t.Fatalf("upstream default overwritten: %+v", attr)
			}
			parent.XVesProtoMessage = "ves.io.schema.synthetic.OtherType"
			parent.Properties[tt.field] = openapi.Schema{Type: tt.kind}
			attr = findAttr(ExtractNestedAttributes(parent, spec, 3, "unrelated"), tt.field)
			if attr.Computed || attr.ServerDefault {
				t.Fatalf("same-name unrelated field changed: %+v", attr)
			}
		})
	}
}
