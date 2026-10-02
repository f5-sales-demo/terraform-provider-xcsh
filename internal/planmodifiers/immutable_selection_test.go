package planmodifiers

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func selectionRaw(selection string, child interface{}) tftypes.Value {
	blockType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"setting": tftypes.String}}
	rootType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"http": blockType, "https": blockType, "https_auto_cert": blockType,
	}}
	values := map[string]tftypes.Value{}
	for _, name := range []string{"http", "https", "https_auto_cert"} {
		var value interface{}
		if name == selection {
			value = map[string]tftypes.Value{"setting": tftypes.NewValue(tftypes.String, child)}
		}
		if selection == "unknown" && name == "http" {
			value = tftypes.UnknownValue
		}
		values[name] = tftypes.NewValue(blockType, value)
	}
	if selection == "null" {
		return tftypes.NewValue(rootType, nil)
	}
	return tftypes.NewValue(rootType, values)
}

func TestImmutableOneOfSelection(t *testing.T) {
	members := []string{"http", "https", "https_auto_cert"}
	tests := []struct {
		old, next           string
		oldChild, nextChild interface{}
		replace             bool
	}{
		{"null", "http", nil, "new", false},
		{"http", "null", "old", nil, false},
		{"", "", nil, nil, false},
		{"http", "http", "old", "new", false},
		{"https", "https", "old-cert", "rotated-cert", false},
		{"https_auto_cert", "https_auto_cert", "old", "new", false},
		{"http", "http", "old", tftypes.UnknownValue, false},
		{"http", "http", tftypes.UnknownValue, "new", false},
		{"unknown", "unknown", nil, nil, true},
		{"http", "unknown", "old", nil, true},
		{"unknown", "http", nil, "new", true},
	}
	for _, old := range members {
		tests = append(tests, struct {
			old, next           string
			oldChild, nextChild interface{}
			replace             bool
		}{old, "", "old", nil, true})
		tests = append(tests, struct {
			old, next           string
			oldChild, nextChild interface{}
			replace             bool
		}{"", old, nil, "new", true})
		for _, next := range members {
			if old != next {
				tests = append(tests, struct {
					old, next           string
					oldChild, nextChild interface{}
					replace             bool
				}{old, next, "old", "new", true})
			}
		}
	}
	for _, tc := range tests {
		t.Run(tc.old+"_to_"+tc.next, func(t *testing.T) {
			req := resource.ModifyPlanRequest{
				State: tfsdk.State{Raw: selectionRaw(tc.old, tc.oldChild)},
				Plan:  tfsdk.Plan{Raw: selectionRaw(tc.next, tc.nextChild)},
			}
			resp := resource.ModifyPlanResponse{}
			RequireImmutableOneOfSelection(context.Background(), req, &resp, "loadbalancer_type", members)
			if resp.Diagnostics.HasError() {
				t.Fatalf("unexpected errors: %v", resp.Diagnostics)
			}
			if (len(resp.RequiresReplace) > 0) != tc.replace {
				t.Fatalf("paths=%v, want replace=%v", resp.RequiresReplace, tc.replace)
			}
			if tc.replace {
				if len(resp.Diagnostics.Warnings()) != 1 {
					t.Fatalf("warnings=%v", resp.Diagnostics)
				}
				for _, member := range members {
					if member == tc.old || member == tc.next {
						found := false
						for _, p := range resp.RequiresReplace {
							if p.Equal(path.Root(member)) {
								found = true
							}
						}
						if !found {
							t.Fatalf("missing affected path %s: %v", member, resp.RequiresReplace)
						}
					}
				}
			} else if len(resp.Diagnostics) > 0 {
				t.Fatalf("unchanged selection warned: %v", resp.Diagnostics)
			}
		})
	}
}
