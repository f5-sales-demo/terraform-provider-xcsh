// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// Tests for x-f5xc-wire-name (#1323).
//
// F5 ships several misspelled property names (blocked_sevice,
// public_advertisment, disable_lb_source_ip_persistance). The wire key must stay
// misspelled — verified live: a PUT with `blocked_sevice` returns HTTP 200 and
// round-trips, while `blocked_service` is silently ignored by the server (#1257).
// The buffer zone therefore presents the CORRECTED spelling as the property name
// and records the original key in x-f5xc-wire-name. Extraction must keep the
// Terraform-facing names (Name/TfsdkTag/GoName) on the corrected spelling while
// JsonName — the key every marshal/unmarshal emitter uses — carries the wire key.

// wireNameRefSpec provides the ioschemaEmpty component the empty-marker fixtures
// $ref, mirroring the real origin_pool advanced_options shape.
func wireNameRefSpec() *openapi.Spec {
	return &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"ioschemaEmpty": {Type: "object"},
			},
		},
	}
}

// TestWireNameDrivesJSONNameOnly asserts that x-f5xc-wire-name changes ONLY the
// JSON key, across every property shape the real annotated properties use:
// a plain scalar (public_advertisment), an allOf-wrapped $ref empty marker
// (disable_lb_source_ip_persistance) and an array of objects (blocked_sevice).
func TestWireNameDrivesJSONNameOnly(t *testing.T) {
	spec := wireNameRefSpec()

	cases := []struct {
		desc     string
		prop     string
		schema   openapi.Schema
		wantJSON string
		wantGo   string
	}{
		{
			desc:     "scalar bool (public_advertisment)",
			prop:     "public_advertisement",
			schema:   openapi.Schema{Type: "boolean", XF5xcWireName: "public_advertisment"},
			wantJSON: "public_advertisment",
			wantGo:   "PublicAdvertisement",
		},
		{
			desc: "allOf-wrapped $ref empty marker (disable_lb_source_ip_persistance)",
			prop: "disable_lb_source_ip_persistence",
			schema: openapi.Schema{
				AllOf:         []openapi.Schema{{Ref: "#/components/schemas/ioschemaEmpty"}},
				XF5xcWireName: "disable_lb_source_ip_persistance",
			},
			wantJSON: "disable_lb_source_ip_persistance",
			wantGo:   "DisableLBSourceIPPersistence",
		},
		{
			desc: "array of objects (blocked_sevice)",
			prop: "blocked_service",
			schema: openapi.Schema{
				Type:          "array",
				XF5xcWireName: "blocked_sevice",
				Items: &openapi.Schema{
					Type:       "object",
					Properties: map[string]openapi.Schema{"network_type": {Type: "string"}},
				},
			},
			wantJSON: "blocked_sevice",
			wantGo:   "BlockedService",
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			attr := ConvertToTerraformAttribute(tc.prop, tc.schema, false, "", spec)

			if attr.JsonName != tc.wantJSON {
				t.Errorf("JsonName = %q, want the wire key %q", attr.JsonName, tc.wantJSON)
			}
			// Terraform-facing identity stays on the CORRECTED property name.
			if attr.Name != tc.prop {
				t.Errorf("Name = %q, want the corrected property name %q", attr.Name, tc.prop)
			}
			if attr.TfsdkTag != tc.prop {
				t.Errorf("TfsdkTag = %q, want the corrected property name %q", attr.TfsdkTag, tc.prop)
			}
			if attr.GoName != tc.wantGo {
				t.Errorf("GoName = %q, want %q (derived from the corrected name)", attr.GoName, tc.wantGo)
			}
		})
	}
}

// TestWireNameAbsentFallsBackToPropertyName pins the 99% path: with no
// annotation, JsonName is the property name exactly as before.
func TestWireNameAbsentFallsBackToPropertyName(t *testing.T) {
	spec := wireNameRefSpec()

	for _, prop := range []string{"blocked_services", "public_advertisment", "advanced_options"} {
		attr := ConvertToTerraformAttribute(prop, openapi.Schema{Type: "string"}, false, "", spec)
		if attr.JsonName != prop {
			t.Errorf("unannotated %q: JsonName = %q, want %q", prop, attr.JsonName, prop)
		}
	}
}

// TestWireNameIsPropertyLevelOnly asserts the annotation is never inherited from
// a $ref TARGET. The wire key is a fact about the property that points at a
// component, not about the component itself: a shared component carrying the
// annotation must not rename every property that references it.
func TestWireNameIsPropertyLevelOnly(t *testing.T) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"sharedThing": {Type: "string", XF5xcWireName: "not_my_wire_name"},
			},
		},
	}

	attr := ConvertToTerraformAttribute(
		"my_property",
		openapi.Schema{Ref: "#/components/schemas/sharedThing"},
		false, "", spec,
	)
	if attr.JsonName != "my_property" {
		t.Errorf("JsonName = %q, want %q: x-f5xc-wire-name on a $ref target must not be inherited",
			attr.JsonName, "my_property")
	}
}

// TestWireNameOnNestedProperties asserts the annotation is honoured at nesting
// depth, mirroring the real shape: fleetBlockedServicesListType.blocked_sevice is
// a list-typed child INSIDE a block, and advanced_options
// .disable_lb_source_ip_persistance is an empty-marker child inside a block.
func TestWireNameOnNestedProperties(t *testing.T) {
	spec := wireNameRefSpec()

	parent := openapi.Schema{
		Type: "object",
		Properties: map[string]openapi.Schema{
			"blocked_service": {
				Type:          "array",
				XF5xcWireName: "blocked_sevice",
				Items: &openapi.Schema{
					Type: "object",
					Properties: map[string]openapi.Schema{
						"blocked_service_name": {Type: "string", XF5xcWireName: "blocked_sevice_name"},
					},
				},
			},
			"disable_lb_source_ip_persistence": {
				AllOf:         []openapi.Schema{{Ref: "#/components/schemas/ioschemaEmpty"}},
				XF5xcWireName: "disable_lb_source_ip_persistance",
			},
			"plain_child": {Type: "string"},
		},
	}

	attr := ConvertToTerraformAttribute("blocked_services", parent, false, "", spec)
	if attr.JsonName != "blocked_services" {
		t.Fatalf("parent JsonName = %q, want %q", attr.JsonName, "blocked_services")
	}

	byTag := map[string]openapi.TerraformAttribute{}
	for _, child := range attr.NestedAttributes {
		byTag[child.TfsdkTag] = child
	}

	want := map[string]string{
		"blocked_service":                  "blocked_sevice",
		"disable_lb_source_ip_persistence": "disable_lb_source_ip_persistance",
		"plain_child":                      "plain_child",
	}
	for tag, wantJSON := range want {
		child, ok := byTag[tag]
		if !ok {
			t.Fatalf("nested attribute %q missing; got %v", tag, byTag)
		}
		if child.JsonName != wantJSON {
			t.Errorf("nested %q: JsonName = %q, want %q", tag, child.JsonName, wantJSON)
		}
	}

	// One level deeper: a scalar inside the list element.
	elem := byTag["blocked_service"]
	var found bool
	for _, grandchild := range elem.NestedAttributes {
		if grandchild.TfsdkTag == "blocked_service_name" {
			found = true
			if grandchild.JsonName != "blocked_sevice_name" {
				t.Errorf("list-element scalar JsonName = %q, want %q",
					grandchild.JsonName, "blocked_sevice_name")
			}
		}
	}
	if !found {
		t.Errorf("list-element scalar blocked_service_name missing; got %+v", elem.NestedAttributes)
	}
}
