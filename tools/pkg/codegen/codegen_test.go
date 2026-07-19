// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"regexp"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestEscapeGoString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no special characters",
			input:    "simple string",
			expected: "simple string",
		},
		{
			name:     "with double quotes",
			input:    `say "hello"`,
			expected: `say \"hello\"`,
		},
		{
			name:     "with backslash",
			input:    `path\to\file`,
			expected: `path\\to\\file`,
		},
		{
			name:     "with both",
			input:    `path\to\"file"`,
			expected: `path\\to\\\"file\"`,
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EscapeGoString(tt.input)
			if got != tt.expected {
				t.Errorf("EscapeGoString(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestRegexLiteral(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple pattern uses backticks",
			input:    `^[a-z]+$`,
			expected: "`^[a-z]+$`",
		},
		{
			name:     "pattern with backslash uses backticks",
			input:    `^\d{3}-\d{4}$`,
			expected: "`" + `^\d{3}-\d{4}$` + "`",
		},
		{
			name:     "pattern with backtick uses quoted string",
			input:    "pattern`with`backtick",
			expected: `"pattern` + "`" + `with` + "`" + `backtick"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RegexLiteral(tt.input)
			if got != tt.expected {
				t.Errorf("RegexLiteral(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGetGoClientType(t *testing.T) {
	tests := []struct {
		name     string
		attr     openapi.TerraformAttribute
		expected string
	}{
		{
			name:     "string type",
			attr:     openapi.TerraformAttribute{Type: "string"},
			expected: "string",
		},
		{
			name:     "int64 type",
			attr:     openapi.TerraformAttribute{Type: "int64"},
			expected: "int64",
		},
		{
			name:     "bool type",
			attr:     openapi.TerraformAttribute{Type: "bool"},
			expected: "bool",
		},
		{
			name:     "list of strings",
			attr:     openapi.TerraformAttribute{Type: "list", ElementType: "string"},
			expected: "[]string",
		},
		{
			name:     "list of int64",
			attr:     openapi.TerraformAttribute{Type: "list", ElementType: "int64"},
			expected: "[]int64",
		},
		{
			name:     "list of unknown",
			attr:     openapi.TerraformAttribute{Type: "list", ElementType: "object"},
			expected: "[]interface{}",
		},
		{
			name:     "map type",
			attr:     openapi.TerraformAttribute{Type: "map"},
			expected: "map[string]string",
		},
		{
			name:     "block single",
			attr:     openapi.TerraformAttribute{IsBlock: true, NestedBlockType: "single"},
			expected: "map[string]interface{}",
		},
		{
			name:     "block list",
			attr:     openapi.TerraformAttribute{IsBlock: true, NestedBlockType: "list"},
			expected: "[]map[string]interface{}",
		},
		{
			name:     "unknown type",
			attr:     openapi.TerraformAttribute{Type: "unknown"},
			expected: "interface{}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetGoClientType(tt.attr)
			if got != tt.expected {
				t.Errorf("GetGoClientType() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestRenderSpecStructFields(t *testing.T) {
	tests := []struct {
		name     string
		attrs    []openapi.TerraformAttribute
		indent   string
		contains []string
		empty    bool
	}{
		{
			name:  "empty attrs",
			attrs: nil,
			empty: true,
		},
		{
			name: "metadata fields are excluded",
			attrs: []openapi.TerraformAttribute{
				{GoName: "Name", TfsdkTag: "name", JsonName: "name", Type: "string"},
			},
			empty: true,
		},
		{
			name: "string field with omitempty",
			attrs: []openapi.TerraformAttribute{
				{GoName: "Domain", TfsdkTag: "domain", JsonName: "domain", Type: "string", IsSpecField: true},
			},
			indent: "\t",
			contains: []string{
				`Domain string ` + "`" + `json:"domain,omitempty"` + "`",
			},
		},
		{
			name: "block field without omitempty",
			attrs: []openapi.TerraformAttribute{
				{GoName: "Config", TfsdkTag: "config", JsonName: "config", Type: "object", IsBlock: true, NestedBlockType: "single", IsSpecField: true},
			},
			indent: "\t",
			contains: []string{
				`Config map[string]interface{} ` + "`" + `json:"config"` + "`",
			},
		},
		{
			name: "uses tfsdk tag when json name empty",
			attrs: []openapi.TerraformAttribute{
				{GoName: "Port", TfsdkTag: "port", JsonName: "", Type: "int64", IsSpecField: true},
			},
			indent: "\t",
			contains: []string{
				`Port int64 ` + "`" + `json:"port,omitempty"` + "`",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RenderSpecStructFields(tt.attrs, tt.indent)
			if tt.empty {
				if got != "" {
					t.Errorf("expected empty string, got %q", got)
				}
				return
			}
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("RenderSpecStructFields() missing %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestRenderNestedAttributes_Empty(t *testing.T) {
	got := RenderNestedAttributes(nil, "\t")
	if got != "" {
		t.Errorf("expected empty string for nil attrs, got %q", got)
	}
}

func TestRenderNestedBlocks_NoBlocks(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{GoName: "Name", TfsdkTag: "name", Type: "string"},
	}
	got := RenderNestedBlocks(attrs, "\t")
	if got != "" {
		t.Errorf("expected empty string when no blocks, got %q", got)
	}
}

func TestCollectNestedModelTypes(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			GoName:          "Config",
			TfsdkTag:        "config",
			IsBlock:         true,
			NestedBlockType: "single",
			NestedAttributes: []openapi.TerraformAttribute{
				{GoName: "Port", TfsdkTag: "port", Type: "int64"},
			},
		},
	}

	var models []NestedModelInfo
	CollectNestedModelTypes("Test", attrs, "", &models)

	if len(models) != 1 {
		t.Fatalf("expected 1 model, got %d", len(models))
	}
	if models[0].TypeName != "TestConfigModel" {
		t.Errorf("expected type name TestConfigModel, got %s", models[0].TypeName)
	}
	if models[0].IsEmpty {
		t.Error("expected non-empty model")
	}
}

// deepComputedTree builds a list block whose Computed field sits ~4 levels deep,
// mirroring dns_zone's rr_set_group[] -> rr_set[] -> lb_record -> value.{namespace}.
func deepComputedTree() openapi.TerraformAttribute {
	return openapi.TerraformAttribute{
		GoName: "RrSetGroup", TfsdkTag: "rr_set_group", IsBlock: true, NestedBlockType: "list", IsSpecField: true,
		NestedAttributes: []openapi.TerraformAttribute{
			{
				GoName: "RrSet", TfsdkTag: "rr_set", IsBlock: true, NestedBlockType: "list",
				NestedAttributes: []openapi.TerraformAttribute{
					{
						GoName: "LbRecord", TfsdkTag: "lb_record", IsBlock: true, NestedBlockType: "single",
						NestedAttributes: []openapi.TerraformAttribute{
							{
								GoName: "Value", TfsdkTag: "value", IsBlock: true, NestedBlockType: "single",
								NestedAttributes: []openapi.TerraformAttribute{
									{GoName: "Name", TfsdkTag: "name", Type: "string"},
									{GoName: "Namespace", TfsdkTag: "namespace", Type: "string", Computed: true},
								},
							},
						},
					},
				},
			},
		},
	}
}

// A Computed+Optional scalar in a single nested block (e.g. http_health_check.use_http2)
// that the user leaves unset has an unknown planned value. The unmarshal "preserve"
// path must NOT return that unknown value — it must guard on !IsUnknown() and fall
// through to the API response / null, else apply fails with "invalid result object
// after apply". Regression test.
func TestRenderUnmarshalScalarChild_PreserveGuardsUnknown(t *testing.T) {
	var sb strings.Builder
	attr := openapi.TerraformAttribute{
		GoName: "UseHttp2", TfsdkTag: "use_http2", JsonName: "use_http2",
		Type: "bool", Optional: true,
	}
	renderUnmarshalScalarChild(&sb, "Healthcheck", attr, "blockData", "data.HTTPHealthCheck", "data.HTTPHealthCheck != nil", "single", "\t")
	got := sb.String()

	if !strings.Contains(got, "!data.HTTPHealthCheck.UseHttp2.IsUnknown()") {
		t.Errorf("preserve guard must check IsUnknown() before returning the planned value; got:\n%s", got)
	}
	// Still preserves a known explicitly-set value.
	if !strings.Contains(got, "return data.HTTPHealthCheck.UseHttp2") {
		t.Errorf("expected preserve path to return the prior value when known; got:\n%s", got)
	}
	// Still falls through to the API response.
	if !strings.Contains(got, `blockData["use_http2"].(bool)`) {
		t.Errorf("expected fallthrough to API response; got:\n%s", got)
	}
}

// Server-default oneof empty-marker members must not be populated from the API
// response on import (they cause spurious post-import drift). The flatten must
// guard the response-populate with !isImport for suppressed members, and leave
// non-suppressed members (user-intent markers) untouched.
func TestRenderUnmarshalSingleChild_ImportSuppressesServerDefault(t *testing.T) {
	mk := func(go_, tfsdk string) openapi.TerraformAttribute {
		return openapi.TerraformAttribute{GoName: go_, TfsdkTag: tfsdk, JsonName: tfsdk, IsBlock: true, NestedBlockType: "single"}
	}

	// Suppressed default marker (disable_waf) -> guarded by !isImport.
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "HTTPLoadBalancer", "", mk("DisableWaf", "disable_waf"), "apiResource.Spec", "data", "true", "single", "\t")
	got := sb.String()
	if !strings.Contains(got, "if !isImport {") {
		t.Errorf("suppressed member disable_waf must guard response-populate with !isImport; got:\n%s", got)
	}

	// Non-suppressed user-intent marker (advertise_on_public_default_vip) -> no import guard on the populate.
	var sb2 strings.Builder
	renderUnmarshalSingleChild(&sb2, "HTTPLoadBalancer", "", mk("AdvertiseOnPublicDefaultVip", "advertise_on_public_default_vip"), "apiResource.Spec", "data", "true", "single", "\t")
	got2 := sb2.String()
	// The only !isImport in a non-suppressed member is the preserve branch, not a wrapper
	// around the response-populate. Assert the populate line is not nested under an extra guard.
	if strings.Count(got2, "if !isImport {") != 0 {
		t.Errorf("non-suppressed member must not add an import-suppression guard; got:\n%s", got2)
	}

	if !isImportDefaultSuppressed("HTTPLoadBalancer", "round_robin") {
		t.Error("round_robin should be a suppressed server default for HTTPLoadBalancer")
	}
	if isImportDefaultSuppressed("HTTPLoadBalancer", "advertise_on_public_default_vip") {
		t.Error("advertise_on_public_default_vip is user intent and must NOT be suppressed")
	}
}

// #1103: plain optional empty-marker blocks the API echoes on every list element
// (origin_pool origin_servers[].labels {}, http_loadbalancer
// default_route_pools[].endpoint_subsets {}) must guard their response-populate with
// !isImport, exactly like oneof base markers — otherwise a minimal config that omits
// them drifts every plan after import. This proves the seed entries flow through to the
// generated flatten closure for these two leaves.
func TestRenderUnmarshalSingleChild_ImportSuppressesEmptyMarkerListElement_Issue1103(t *testing.T) {
	mk := func(go_, tfsdk string) openapi.TerraformAttribute {
		return openapi.TerraformAttribute{GoName: go_, TfsdkTag: tfsdk, JsonName: tfsdk, IsBlock: true, NestedBlockType: "single"}
	}
	cases := []struct {
		rc, goName, tfsdk string
	}{
		{"OriginPool", "Labels", "labels"},
		{"HTTPLoadBalancer", "EndpointSubsets", "endpoint_subsets"},
	}
	for _, c := range cases {
		var sb strings.Builder
		// Render as it appears inside a list element (positional state accessor).
		renderUnmarshalSingleChild(&sb, c.rc, "", mk(c.goName, c.tfsdk), "itemMap", "existingItems[idx]", "len(existingItems) > idx", "list", "\t")
		got := sb.String()
		if !strings.Contains(got, "if !isImport {") {
			t.Errorf("%s.%s must guard response-populate with !isImport (empty-marker import drift #1103); got:\n%s", c.rc, c.tfsdk, got)
		}
	}
}

// #1103 non-collision: seeding OriginPool.labels suppresses ONLY the origin_servers[]
// empty-marker block. The top-level metadata.labels is a types.Map rendered by a
// different path that never consults isImportDefaultSuppressed, so a map-typed "labels"
// child must NOT acquire an empty-marker import-suppression guard.
func TestRenderUnmarshalChild_MetadataLabelsMapNotSuppressed_Issue1103(t *testing.T) {
	var sb strings.Builder
	mapAttr := openapi.TerraformAttribute{GoName: "Labels", TfsdkTag: "labels", JsonName: "labels", Type: "map", ElementType: "string"}
	renderUnmarshalChild(&sb, "OriginPool", "", mapAttr, "metaMap", "", "", "single", "\t")
	got := sb.String()
	if strings.Contains(got, "EmptyModel{}") {
		t.Errorf("metadata.labels (types.Map) must not render as an empty-marker block; got:\n%s", got)
	}
	if strings.Contains(got, "if !isImport {") {
		t.Errorf("metadata.labels (types.Map) must not acquire an empty-marker import-suppression guard; got:\n%s", got)
	}
}

// A suppressed server-computed LIST block nested inside another block (e.g.
// app_firewall detection_settings.violations_view — the server materializes the
// full violation catalog whenever detection_settings is set) must not be populated
// from the API on import, or a config omitting it drifts on round-trip. The nested
// list-child renderer must emit an `if isImport { return nil }` early return for a
// suppressed leaf. Non-suppressed list children must NOT get that guard.
func TestRenderUnmarshalListChild_ImportSuppressesServerComputedList(t *testing.T) {
	mk := func(goName, tfsdk string) openapi.TerraformAttribute {
		return openapi.TerraformAttribute{
			GoName: goName, TfsdkTag: tfsdk, JsonName: tfsdk, IsBlock: true, NestedBlockType: "list",
			NestedAttributes: []openapi.TerraformAttribute{
				{GoName: "Name", TfsdkTag: "name", JsonName: "name"},
			},
		}
	}

	// Suppressed nested list (AppFirewall violations_view) -> early import return.
	var sb strings.Builder
	renderUnmarshalListChild(&sb, "AppFirewall", "", mk("ViolationsView", "violations_view"), "blockData", "data.DetectionSettings", "data.DetectionSettings != nil", "single", "\t")
	got := sb.String()
	if !strings.Contains(got, "if isImport {") {
		t.Errorf("suppressed nested list violations_view must skip populate on import; got:\n%s", got)
	}

	// Non-suppressed nested list -> no import-skip guard.
	var sb2 strings.Builder
	renderUnmarshalListChild(&sb2, "AppFirewall", "", mk("SomeUserList", "some_user_list"), "blockData", "data.DetectionSettings", "data.DetectionSettings != nil", "single", "\t")
	if strings.Contains(sb2.String(), "if isImport {") {
		t.Errorf("non-suppressed nested list must not add an import-skip guard; got:\n%s", sb2.String())
	}
}

// A suppressed non-empty server-default block (e.g. l7_ddos_protection, which the
// API returns as an empty object for a minimal LB) must not be materialized on
// import from an empty response — otherwise import creates an all-defaults block
// the user never set. The build guard must require a non-empty response on import.
func TestRenderUnmarshalTopLevelSingle_SuppressedNonEmptyRequiresContentOnImport(t *testing.T) {
	attr := openapi.TerraformAttribute{
		GoName: "L7DDOSProtection", TfsdkTag: "l7_ddos_protection", JsonName: "l7_ddos_protection",
		IsBlock: true, NestedBlockType: "single", IsSpecField: true,
		NestedAttributes: []openapi.TerraformAttribute{
			{GoName: "L7DdosActionDefault", TfsdkTag: "l7_ddos_action_default", JsonName: "l7_ddos_action_default", IsBlock: true, NestedBlockType: "single"},
		},
	}
	var sb strings.Builder
	renderUnmarshalTopLevelSingle(&sb, "HTTPLoadBalancer", attr, "\t")
	got := sb.String()
	if !strings.Contains(got, "isImport && len(blockData) > 0") {
		t.Errorf("suppressed non-empty block must require non-empty response on import; got:\n%s", got)
	}
}

// A suppressed Optional bool at its false server default must return null on
// import (config omits it) — otherwise post-import plan shows "false -> null".
func TestRenderUnmarshalScalarChild_ImportSuppressesDefaultBool(t *testing.T) {
	attr := openapi.TerraformAttribute{
		GoName: "DNSVolterraManaged", TfsdkTag: "dns_volterra_managed", JsonName: "dns_volterra_managed",
		Type: "bool", Optional: true,
	}
	var sb strings.Builder
	renderUnmarshalScalarChild(&sb, "HTTPLoadBalancer", attr, "blockData", "data.HTTP", "data.HTTP != nil", "single", "\t")
	got := sb.String()
	if !strings.Contains(got, "if isImport {") || !strings.Contains(got, "ok && !v {") {
		t.Errorf("suppressed default bool must return null on import when false; got:\n%s", got)
	}
}

// Every nested list block is modeled as types.List, regardless of whether it has a
// Computed descendant: a native Go slice cannot represent the unknown values a config may
// carry at plan time (a Computed descendant, or an element sourced from an unresolved
// reference such as the inline API crawler domains[].simple_login.password). See #1083.
func TestNestedListUsesTypesList(t *testing.T) {
	tests := []struct {
		name string
		attr openapi.TerraformAttribute
		want bool
	}{
		{
			name: "list block with computed descendant",
			attr: deepComputedTree(),
			want: true,
		},
		{
			name: "list block with no computed descendant",
			attr: openapi.TerraformAttribute{
				IsBlock: true, NestedBlockType: "list",
				NestedAttributes: []openapi.TerraformAttribute{
					{GoName: "TTL", TfsdkTag: "ttl", Type: "int64"},
					{
						GoName: "Inner", TfsdkTag: "inner", IsBlock: true, NestedBlockType: "single",
						NestedAttributes: []openapi.TerraformAttribute{
							{GoName: "Name", TfsdkTag: "name", Type: "string"},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "empty list block",
			attr: openapi.TerraformAttribute{IsBlock: true, NestedBlockType: "list"},
			want: true,
		},
		{
			name: "single nested block is not a list",
			attr: openapi.TerraformAttribute{IsBlock: true, NestedBlockType: "single"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nestedListUsesTypesList(tt.attr); got != tt.want {
				t.Errorf("nestedListUsesTypesList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// A list block with a Computed descendant must be modeled as types.List (not a native
// slice), otherwise the plugin framework cannot represent unknown values during planning.
func TestRenderNestedModelTypes_ComputedDescendantList(t *testing.T) {
	attrs := []openapi.TerraformAttribute{deepComputedTree()}
	got := RenderNestedModelTypes("Test", attrs)

	// rr_set is a list block with a Computed descendant (namespace) -> types.List
	if !strings.Contains(got, "RrSet types.List `tfsdk:\"rr_set\"`") {
		t.Errorf("expected rr_set field to be types.List, got:\n%s", got)
	}
	// rr_set_group is also a list block with a Computed descendant -> types.List (top of nested tree)
	if strings.Contains(got, "RrSet []Test") {
		t.Errorf("expected no native slice for rr_set with a Computed descendant, got:\n%s", got)
	}
}

// A nested list block with NO computed descendant is still modeled as types.List: a
// native slice cannot hold an unknown value a config may supply at plan time (e.g. an
// element field sourced from an unresolved reference). See #1083.
func TestRenderNestedModelTypes_NestedListAlwaysTypesList(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			GoName: "Outer", TfsdkTag: "outer", IsBlock: true, NestedBlockType: "single",
			NestedAttributes: []openapi.TerraformAttribute{
				{
					GoName: "Items", TfsdkTag: "items", IsBlock: true, NestedBlockType: "list",
					NestedAttributes: []openapi.TerraformAttribute{
						{GoName: "Name", TfsdkTag: "name", Type: "string"},
					},
				},
			},
		},
	}
	got := RenderNestedModelTypes("Test", attrs)
	if !strings.Contains(got, "Items types.List `tfsdk:\"items\"`") {
		t.Errorf("expected items to be types.List, got:\n%s", got)
	}
	if strings.Contains(got, "Items []TestOuterItemsModel") {
		t.Errorf("expected no native slice for nested list items, got:\n%s", got)
	}
}

// The example renderer emits a minimal valid config: identity + every required non-block
// attribute (enum-aware value), the correct provider source, and NO optional blocks.
func TestRenderResourceExampleHCL(t *testing.T) {
	rt := &openapi.ResourceTemplate{
		Description: "Manages a thing. Extra detail.",
		Attributes: []openapi.TerraformAttribute{
			{TfsdkTag: "name", Type: "string", Required: true},
			{TfsdkTag: "namespace", Type: "string", Required: true},
			{TfsdkTag: "labels", Type: "map"}, // optional -> omitted
			{TfsdkTag: "mode", Type: "string", Required: true, EnumValues: []string{"LOCAL", "GLOBAL"}},
			{TfsdkTag: "address_pool", Type: "list", ElementType: "string", Required: true},
			{TfsdkTag: "ttl", Type: "int64"},                                             // optional -> omitted
			{TfsdkTag: "cfg", IsBlock: true, NestedBlockType: "single", Required: false}, // block -> omitted
		},
	}
	got := RenderResourceExampleHCL(rt, "address_allocator", "system")

	for _, want := range []string{
		`source  = "f5-sales-demo/xcsh"`,
		`resource "xcsh_address_allocator" "example"`,
		`name      = "example-address-allocator"`,
		`namespace = "system"`,
		`mode = "LOCAL"`, // enum -> first value
		`address_pool = ["example-value"]`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("expected example to contain %q, got:\n%s", want, got)
		}
	}
	// Optional attributes and blocks must NOT appear (keeps the example minimal + valid).
	for _, unwanted := range []string{"labels", "ttl", "cfg", "f5-sales-demo/f5xc"} {
		if strings.Contains(got, unwanted) {
			t.Errorf("did not expect %q in minimal example, got:\n%s", unwanted, got)
		}
	}
}

// An unconfigured (null or empty) nested list block must be preserved as null on normal
// Read/Create so a server-managed list the user never configured does not drift the plan
// ("Provider produced inconsistent result after apply"). Import still reads the API.
func TestUnmarshal_PreservesUnconfiguredList(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{
			GoName: "Primary", TfsdkTag: "primary", IsBlock: true, NestedBlockType: "single", IsSpecField: true,
			NestedAttributes: []openapi.TerraformAttribute{
				{
					GoName: "Grp", TfsdkTag: "grp", IsBlock: true, NestedBlockType: "list",
					NestedAttributes: []openapi.TerraformAttribute{
						{
							GoName: "Ref", TfsdkTag: "ref", IsBlock: true, NestedBlockType: "single",
							NestedAttributes: []openapi.TerraformAttribute{
								{GoName: "Uid", TfsdkTag: "uid", Type: "string", Computed: true},
							},
						},
					},
				},
			},
		},
	}
	got := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	// grp has a Computed descendant -> types.List, and must preserve null/empty prior state.
	if !strings.Contains(got, "data.Primary.Grp.IsNull() || len(data.Primary.Grp.Elements()) == 0") {
		t.Errorf("expected unconfigured-list preservation guard for grp, got:\n%s", got)
	}
	if !strings.Contains(got, "return types.ListNull(types.ObjectType{AttrTypes: TestPrimaryGrpModelAttrTypes})") {
		t.Errorf("expected canonical ListNull return for preserved grp, got:\n%s", got)
	}
}

// Import mode is a one-shot: the generated Read must clear the isImport private-state marker,
// otherwise every subsequent refresh re-enters import mode and drifts on server-managed fields.
func TestResourceTemplate_ClearsImportMarkerAfterImport(t *testing.T) {
	if !strings.Contains(ResourceTemplate, `resp.Private.SetKey(ctx, "isImport", nil)`) {
		t.Fatal("ResourceTemplate Read must clear the isImport marker after an import read")
	}
	if strings.Count(ResourceTemplate, `SetKey(ctx, "isImport", nil)`) != 1 {
		t.Errorf("expected exactly one isImport clear (Read only), got %d", strings.Count(ResourceTemplate, `SetKey(ctx, "isImport", nil)`))
	}
	// The clear must be guarded so it only fires on the import read.
	idx := strings.Index(ResourceTemplate, `SetKey(ctx, "isImport", nil)`)
	if idx == -1 || !strings.Contains(ResourceTemplate[:idx], "if isImport {") {
		t.Error("isImport clear must be guarded by `if isImport {`")
	}
}

// The resource template must (a) emit a guarded static string default for attributes
// carrying a StringDefault, and (b) let the namespace attribute carry a OneOf validator
// (for spec-driven fixed-namespace resources) in addition to the format validator.
func TestResourceTemplate_StringDefaultAndNamespaceOneOf(t *testing.T) {
	if !strings.Contains(ResourceTemplate, "stringdefault.StaticString(") {
		t.Error("ResourceTemplate must emit stringdefault.StaticString for StringDefault attributes")
	}
	if !strings.Contains(ResourceTemplate, `if ne .StringDefault ""`) {
		t.Error("the StringDefault emission must be guarded by a non-empty check")
	}
	// The namespace validator branch must also emit OneOf when EnumValues are present,
	// not only NamespaceValidator(). Inspect the window until the next branch.
	nsIdx := strings.Index(ResourceTemplate, `eq .TfsdkTag "namespace"`)
	if nsIdx == -1 {
		t.Fatal("namespace validator branch not found in ResourceTemplate")
	}
	window := ResourceTemplate[nsIdx:]
	if end := strings.Index(window, "else if and (eq .Type"); end != -1 {
		window = window[:end]
	}
	if !strings.Contains(window, "stringvalidator.OneOf(") {
		t.Error("namespace branch must emit stringvalidator.OneOf when EnumValues are present")
	}
}

// The recursive emitters must reach the deep list block and convert it: marshal via
// ElementsAs, unmarshal via ListValueFrom, referencing the deep model's AttrTypes.
func TestRecursiveEmitters_DeepListConversion(t *testing.T) {
	attrs := []openapi.TerraformAttribute{deepComputedTree()}

	marshal := RenderSpecMarshalCode(attrs, "\t", "Test")
	if !strings.Contains(marshal, "RrSet.ElementsAs(ctx,") {
		t.Errorf("expected marshal to ElementsAs the deep rr_set types.List, got:\n%s", marshal)
	}

	unmarshal := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if !strings.Contains(unmarshal, "types.ListValueFrom(ctx, types.ObjectType{AttrTypes: TestRrSetGroupRrSetModelAttrTypes}") {
		t.Errorf("expected unmarshal to ListValueFrom the deep rr_set with its AttrTypes, got:\n%s", unmarshal)
	}
}

// RenderRequirementPreflights emits, for each declared prerequisite, an apply-time
// guard: nil-check the triggering block, LIST the requirement's collection in the
// resource namespace, and fail fast with the remediation message when it is empty.
// This is the shipped-binary enforcement of x-f5xc-requires.
func TestRenderRequirementPreflights_CSD(t *testing.T) {
	pf := []openapi.RequirementPreflight{{
		WhenField:   "client_side_defense",
		WhenGoField: "ClientSideDefense",
		ListPath:    "/api/shape/csd/namespaces/%s/protected_domains",
		Requires:    "client_side_defense requires a same-namespace protected_domain",
		ErrorTitle:  "Client-Side Defense prerequisite missing",
		ErrorDetail: `no protected_domain in namespace %s; create an xcsh_protected_domain (the API says "Failed to get CSD JS Configuration")`,
	}}
	got := RenderRequirementPreflights(pf, "r")

	for _, want := range []string{
		"if data.ClientSideDefense != nil {",
		`fmt.Sprintf("/api/shape/csd/namespaces/%s/protected_domains", data.Namespace.ValueString())`,
		"r.client.Get(ctx,",
		"Items []map[string]interface{} `json:\"items\"`",
		"len(preflightResp.Items) == 0",
		`resp.Diagnostics.AddError(`,
		`"Client-Side Defense prerequisite missing"`,
		"return",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("preflight code missing %q, got:\n%s", want, got)
		}
	}
	// The detail string carries embedded quotes and a %s verb; it must be emitted as a
	// valid, correctly-escaped Go literal (via strconv.Quote), not spliced raw.
	if strings.Contains(got, `Configuration")`) && !strings.Contains(got, `Configuration\")`) {
		t.Errorf("error_detail quotes must be escaped in the generated literal, got:\n%s", got)
	}
}

// No declared preflights -> no emitted code (so unaffected resources are byte-identical).
func TestRenderRequirementPreflights_Empty(t *testing.T) {
	if got := RenderRequirementPreflights(nil, "r"); strings.TrimSpace(got) != "" {
		t.Errorf("want empty output for no preflights, got:\n%s", got)
	}
}

// A nested string attribute carrying the etld_plus_one flag must emit the eTLD+1
// validator (the top-level path is exercised by regeneration of protected_domain).
func TestRenderNestedAttributes_ETLDPlusOne(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{GoName: "Domain", TfsdkTag: "domain", Type: "string", ETLDPlusOne: true},
	}
	got := RenderNestedAttributes(attrs, "\t")
	if !strings.Contains(got, "validators.ETLDPlusOneValidator()") {
		t.Errorf("expected ETLDPlusOneValidator for etld_plus_one attribute, got:\n%s", got)
	}
}

// The Delete template must retry transient referential BAD_REQUEST (a resource briefly
// still referenced during teardown) with a bounded, context-aware loop, while leaving
// NOT_FOUND/404 (already deleted) and 501 (unsupported) as terminal.
func TestResourceTemplate_DeleteRetriesTransient400(t *testing.T) {
	for _, want := range []string{
		"for attempt := 0; ; attempt++ {",
		`strings.Contains(msg, "400") || strings.Contains(msg, "BAD_REQUEST")`,
		"attempt >= 5",
		"time.After(5 * time.Second)",
		"case <-ctx.Done():",
	} {
		if !strings.Contains(ResourceTemplate, want) {
			t.Errorf("Delete template missing retry construct %q", want)
		}
	}
	// The transient guard must exclude the terminal conditions so they aren't retried.
	if !strings.Contains(ResourceTemplate, `!strings.Contains(msg, "NOT_FOUND") && !strings.Contains(msg, "404") && !strings.Contains(msg, "501")`) {
		t.Error("transient guard must exclude NOT_FOUND/404/501")
	}
}

// Resources with create-only, API-unreadable fields carry those fields in the import ID
// (namespace/name/<field>...) so a round-trip import is drift-free. The ImportState
// template must parse the extra segments and set the attributes.
func TestResourceTemplate_ImportIDExtraFields(t *testing.T) {
	for _, want := range []string{
		"{{- if .ImportIDExtraFields}}",
		"len(parts) != {{add 2 (len .ImportIDExtraFields)}}",
		"{{- range $i, $f := .ImportIDExtraFields}}",
		`path.Root("{{$f}}"), parts[{{add 2 $i}}]`,
	} {
		if !strings.Contains(ResourceTemplate, want) {
			t.Errorf("ImportState template missing extra-fields construct %q", want)
		}
	}
}

// #1079 part 2: the read-back for an object-reference nested block must reconstruct
// from the API response (so Computed-only tenant/uid/kind become known), NOT preserve
// the planned value (which carries an unknown tenant -> "invalid result object after
// apply"). Non-reference single blocks keep the drift-preserving behavior.
func TestRenderUnmarshalSingleChild_ObjectRefReadsFromAPI(t *testing.T) {
	ref := openapi.TerraformAttribute{
		GoName: "MaliciousUserMitigation", JsonName: "malicious_user_mitigation", TfsdkTag: "malicious_user_mitigation",
		NestedAttributes: []openapi.TerraformAttribute{
			{GoName: "Name", TfsdkTag: "name", JsonName: "name", Type: "string"},
			{GoName: "Namespace", TfsdkTag: "namespace", JsonName: "namespace", Type: "string"},
			{GoName: "Tenant", TfsdkTag: "tenant", JsonName: "tenant", Type: "string"},
		},
	}
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "R", "EnableChallengeMaliciousUserMitigation", ref,
		"blockData", "data.EnableChallenge", "data.EnableChallenge != nil", "single", "\t")
	out := sb.String()
	if strings.Contains(out, "return data.EnableChallenge.MaliciousUserMitigation") {
		t.Errorf("object-reference block must NOT preserve the planned value (carries unknown tenant):\n%s", out)
	}
	if !strings.Contains(out, `MaliciousUserMitigationData["tenant"]`) {
		t.Errorf("object-reference block read-back must read tenant from the API response:\n%s", out)
	}
}

// A non-reference single block (no tenant child) keeps the drift-preserving early return.
func TestRenderUnmarshalSingleChild_NonRefPreserves(t *testing.T) {
	nonRef := openapi.TerraformAttribute{
		GoName: "Policy", JsonName: "policy", TfsdkTag: "policy",
		NestedAttributes: []openapi.TerraformAttribute{
			{GoName: "CookieExpiry", TfsdkTag: "cookie_expiry", JsonName: "cookie_expiry", Type: "int64"},
		},
	}
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "R", "Policy", nonRef,
		"blockData", "data", "data != nil", "single", "\t")
	if !strings.Contains(sb.String(), "return data.Policy") {
		t.Errorf("non-reference single block must keep the drift-preserving early return:\n%s", sb.String())
	}
}

// #1091: a single nested block that CONTAINS an object-reference descendant at any
// depth (e.g. custom_api_auth_discovery -> api_discovery_ref) must NOT preserve the
// planned value either — the planned api_discovery_ref.tenant is unknown (Computed-only),
// so preserving the whole parent carries that unknown and yields "invalid result object
// after apply". The parent must reconstruct from the API response so the nested ref's
// tenant becomes known. #1080 only covered blocks that ARE references, not ones nesting one.
func TestRenderUnmarshalSingleChild_NestedObjectRefReconstructs(t *testing.T) {
	parent := openapi.TerraformAttribute{
		GoName: "CustomAPIAuthDiscovery", JsonName: "custom_api_auth_discovery", TfsdkTag: "custom_api_auth_discovery",
		IsBlock: true, NestedBlockType: "single",
		NestedAttributes: []openapi.TerraformAttribute{
			{
				GoName: "APIDiscoveryRef", JsonName: "api_discovery_ref", TfsdkTag: "api_discovery_ref",
				IsBlock: true, NestedBlockType: "single",
				NestedAttributes: []openapi.TerraformAttribute{
					{GoName: "Name", TfsdkTag: "name", JsonName: "name", Type: "string"},
					{GoName: "Namespace", TfsdkTag: "namespace", JsonName: "namespace", Type: "string"},
					{GoName: "Tenant", TfsdkTag: "tenant", JsonName: "tenant", Type: "string"},
				},
			},
		},
	}
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "R", "EnableAPIDiscoveryCustomAPIAuthDiscovery", parent,
		"blockData", "data.EnableAPIDiscovery", "data.EnableAPIDiscovery != nil", "single", "\t")
	out := sb.String()
	if strings.Contains(out, "return data.EnableAPIDiscovery.CustomAPIAuthDiscovery") {
		t.Errorf("a block nesting an object reference must NOT preserve the planned value (carries unknown nested tenant):\n%s", out)
	}
	if !strings.Contains(out, `APIDiscoveryRefData["tenant"]`) {
		t.Errorf("the nested object-reference read-back must read tenant from the API response:\n%s", out)
	}
}

// #41 (SP3 API Protection): a single block that CONTAINS an object reference on one
// arm (a "spine" block, e.g. client_matcher whose deep ip_matcher.prefix_sets arm is a
// reference) must reconstruct only the reference arm from the API while PRESERVING its
// off-spine Optional markers/scalars from the planned state. Reconstructing the whole
// block materializes server-echoed defaults the plan omitted (any_client:{},
// invert_matcher:false) -> "Provider produced inconsistent result after apply: was
// absent/null, now present/false". The reference arm must still read its Computed tenant
// from the API.
func TestRenderUnmarshalSingleChild_SpinePreservesOffSpineLeaves(t *testing.T) {
	clientMatcher := openapi.TerraformAttribute{
		GoName: "ClientMatcher", JsonName: "client_matcher", TfsdkTag: "client_matcher",
		IsBlock: true, NestedBlockType: "single",
		NestedAttributes: []openapi.TerraformAttribute{
			// off-spine empty-marker oneof member
			{GoName: "AnyClient", JsonName: "any_client", TfsdkTag: "any_client", IsBlock: true, NestedBlockType: "single"},
			// off-spine Optional scalar
			{GoName: "InvertMatcher", JsonName: "invert_matcher", TfsdkTag: "invert_matcher", Type: "bool", Optional: true},
			// spine: ip_matcher -> prefix_sets (an object reference with a tenant child)
			{
				GoName: "IpMatcher", JsonName: "ip_matcher", TfsdkTag: "ip_matcher",
				IsBlock: true, NestedBlockType: "single",
				NestedAttributes: []openapi.TerraformAttribute{
					{
						GoName: "PrefixSets", JsonName: "prefix_sets", TfsdkTag: "prefix_sets",
						IsBlock: true, NestedBlockType: "single",
						NestedAttributes: []openapi.TerraformAttribute{
							{GoName: "Name", TfsdkTag: "name", JsonName: "name", Type: "string"},
							{GoName: "Tenant", TfsdkTag: "tenant", JsonName: "tenant", Type: "string"},
						},
					},
				},
			},
		},
	}
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "R", "ApiEndpointRulesClientMatcher", clientMatcher,
		"itemMap", "apiEndpointRulesItem", "len(existing) > i", "single", "\t")
	out := sb.String()
	if !strings.Contains(out, "return apiEndpointRulesItem.ClientMatcher.AnyClient") {
		t.Errorf("off-spine empty-marker must preserve the planned value (not materialize any_client from the API):\n%s", out)
	}
	if !strings.Contains(out, "return apiEndpointRulesItem.ClientMatcher.InvertMatcher") {
		t.Errorf("off-spine Optional scalar must preserve the planned value (not materialize invert_matcher from the API):\n%s", out)
	}
	if !strings.Contains(out, `PrefixSetsData["tenant"]`) {
		t.Errorf("the spine's object-reference arm must still reconstruct its tenant from the API:\n%s", out)
	}
}

// #45 (SP4 API Testing): a list nested inside a LIST element (e.g. api_testing
// domains[].credentials[]) must ALSO thread prior-state positionally, not only lists
// inside single blocks. Without it, credential elements get stateBase="" and their
// markers/secrets (standard, api_key.value) reconstruct from the API on apply/import.
func TestRenderUnmarshalListChild_ThreadsInsideListElement(t *testing.T) {
	list := openapi.TerraformAttribute{
		GoName: "Credentials", JsonName: "credentials", TfsdkTag: "credentials",
		IsBlock: true, NestedBlockType: "list",
		NestedAttributes: []openapi.TerraformAttribute{
			{GoName: "Standard", JsonName: "standard", TfsdkTag: "standard", IsBlock: true, NestedBlockType: "single"},
		},
	}
	var sb strings.Builder
	renderUnmarshalListChild(&sb, "R", "DomainsCredentials", list,
		"domMap", "existingDomains[i]", "len(existingDomains) > i", "list", "\t")
	out := sb.String()
	if !strings.Contains(out, "existingDomains[i].Credentials.ElementsAs(ctx, &CredentialsExisting") {
		t.Errorf("list nested in a list element must load prior-state elements for threading:\n%s", out)
	}
	if !strings.Contains(out, "CredentialsExisting[CredentialsIdx].Standard") {
		t.Errorf("list-in-list-element child marker must preserve the planned value positionally:\n%s", out)
	}
}

// #45 (SP4 API Testing): an empty-marker oneof member that is a direct child of a
// LIST element (e.g. api_testing.domains[].credentials[].standard — the server-default
// credentials_choice base marker) must preserve the PLANNED value (presence AND
// absence) on the apply path when prior-state is threaded, exactly like a single-block
// child. The old list-container branch only preserved presence (returned &Empty{} when
// state had it) and otherwise fell through to the API populate, so a plan that omits
// the marker while the server echoes it drifts ("was absent, now present"). Import
// still reads the API.
func TestRenderUnmarshalSingleChild_ListEmptyMarkerPreservesAbsence(t *testing.T) {
	marker := openapi.TerraformAttribute{
		GoName: "Standard", JsonName: "standard", TfsdkTag: "standard",
		IsBlock: true, NestedBlockType: "single",
	}
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "R", "CredentialsStandard", marker,
		"credMap", "existingCreds[i]", "len(existingCreds) > i", "list", "\t")
	out := sb.String()
	if !strings.Contains(out, "return existingCreds[i].Standard") {
		t.Errorf("list-element empty marker must preserve the planned value (return stateBase.Field), not materialize the server echo:\n%s", out)
	}
	if strings.Contains(out, "existingCreds[i].Standard != nil") {
		t.Errorf("list-element empty marker must not use the presence-only guard (that drops absence):\n%s", out)
	}
}

// #41 (SP3 API Protection): a list block nested inside a configured single block (e.g.
// api_protection_rules.api_endpoint_rules[]) must thread the prior-state elements
// positionally into element children, mirroring the top-level list renderer, so element
// Optional markers/scalars preserve the planned value on Read/Create instead of
// materializing server-echoed defaults. Import still reads the API.
func TestRenderUnmarshalListChild_PreservesElementStatePositionally(t *testing.T) {
	list := openapi.TerraformAttribute{
		GoName: "ApiEndpointRules", JsonName: "api_endpoint_rules", TfsdkTag: "api_endpoint_rules",
		IsBlock: true, NestedBlockType: "list",
		NestedAttributes: []openapi.TerraformAttribute{
			{
				GoName: "ApiEndpointMethod", JsonName: "api_endpoint_method", TfsdkTag: "api_endpoint_method",
				IsBlock: true, NestedBlockType: "single",
				NestedAttributes: []openapi.TerraformAttribute{
					{GoName: "InvertMatcher", JsonName: "invert_matcher", TfsdkTag: "invert_matcher", Type: "bool", Optional: true},
				},
			},
		},
	}
	var sb strings.Builder
	renderUnmarshalListChild(&sb, "R", "ApiProtectionRulesApiEndpointRules", list,
		"apiProtectionRulesData", "data.ApiProtectionRules", "data.ApiProtectionRules != nil", "single", "\t")
	out := sb.String()
	if !strings.Contains(out, "data.ApiProtectionRules.ApiEndpointRules.ElementsAs(ctx, &ApiEndpointRulesExisting") {
		t.Errorf("nested list must load prior-state elements from the parent state for positional preservation:\n%s", out)
	}
	if !strings.Contains(out, "ApiEndpointRulesExisting[ApiEndpointRulesIdx].ApiEndpointMethod.InvertMatcher") {
		t.Errorf("nested list element leaf must preserve the planned value positionally:\n%s", out)
	}
}

// Coverage Batch F (#61): a single nested block whose child is another single nested
// block with the SAME GoName (F5 XC "view ref" wrappers like http_loadbalancer {
// http_loadbalancer { name } }) must not collide the marshal map variable. The old
// code named both maps "<GoName>Map", so the inner `:= make` shadowed the outer and
// emitted `XMap["http_loadbalancer"] = XMap` (self-reference) while the OUTER map sent
// to the API stayed empty {} — dropping the LB association (live 400). The map var
// must be unique per nesting level (childPath-based), so the outer map receives the
// inner map, not itself.
func TestRenderMarshalBlock_NestedSameNameNoShadow(t *testing.T) {
	ref := openapi.TerraformAttribute{
		GoName: "HTTPLoadBalancer", TfsdkTag: "http_loadbalancer", JsonName: "http_loadbalancer",
		IsBlock: true, NestedBlockType: "single",
		NestedAttributes: []openapi.TerraformAttribute{
			{GoName: "Name", TfsdkTag: "name", JsonName: "name"},
			{GoName: "Namespace", TfsdkTag: "namespace", JsonName: "namespace"},
		},
	}
	outer := openapi.TerraformAttribute{
		GoName: "HTTPLoadBalancer", TfsdkTag: "http_loadbalancer", JsonName: "http_loadbalancer",
		IsBlock: true, NestedBlockType: "single",
		NestedAttributes: []openapi.TerraformAttribute{ref},
	}
	var sb strings.Builder
	renderMarshalBlock(&sb, "AppAPIGroup", "", outer, "data.HTTPLoadBalancer", "createReq.Spec", "\t", false)
	got := sb.String()

	// No map may be assigned to its own key (the shadow self-reference bug).
	selfRef := regexp.MustCompile(`(\w+)\["[a-z_]+"\] = (\w+)`)
	for _, m := range selfRef.FindAllStringSubmatch(got, -1) {
		if m[1] == m[2] {
			t.Errorf("marshal emits self-referential map assignment %q (shadowed nested block); got:\n%s", m[0], got)
		}
	}
	// The two nested maps must be declared with DISTINCT identifiers.
	decl := regexp.MustCompile(`(\w+) := make\(map\[string\]interface\{\}\)`)
	names := map[string]bool{}
	for _, m := range decl.FindAllStringSubmatch(got, -1) {
		if names[m[1]] {
			t.Errorf("duplicate marshal map var %q (shadow); got:\n%s", m[1], got)
		}
		names[m[1]] = true
	}
}

// #1129: meaningful-zero int64 leaves (signature_id, where 0 = "all signatures") must read a
// returned 0 back faithfully — the generated read must DROP the `v != 0` guard for them, while
// every other int64 field keeps it (0 = unset). Regression test for both the allowlisted leaf
// and a control leaf.
func TestRenderUnmarshalScalarChild_MeaningfulZeroInt64_Issue1129(t *testing.T) {
	// signature_id on HTTPLoadBalancer (a list-element leaf) drops the v != 0 guard.
	var sig strings.Builder
	sigAttr := openapi.TerraformAttribute{
		GoName: "SignatureID", TfsdkTag: "signature_id", JsonName: "signature_id", Type: "int64", Optional: true,
	}
	renderUnmarshalScalarChild(&sig, "HTTPLoadBalancer", sigAttr, "m", "", "", "list", "\t")
	got := sig.String()
	if !strings.Contains(got, `if v, ok := m["signature_id"].(float64); ok {`) {
		t.Errorf("signature_id read must drop the `v != 0` guard (#1129); got:\n%s", got)
	}
	if strings.Contains(got, "ok && v != 0") {
		t.Errorf("signature_id read must NOT keep the `v != 0` guard (#1129); got:\n%s", got)
	}

	// A control int64 leaf keeps the v != 0 guard (0 = unset for the common case).
	var ctl strings.Builder
	ctlAttr := openapi.TerraformAttribute{
		GoName: "Timeout", TfsdkTag: "timeout", JsonName: "timeout", Type: "int64", Optional: true,
	}
	renderUnmarshalScalarChild(&ctl, "HTTPLoadBalancer", ctlAttr, "m", "", "", "list", "\t")
	if !strings.Contains(ctl.String(), "ok && v != 0") {
		t.Errorf("non-meaningful-zero int64 (timeout) must keep the `v != 0` guard; got:\n%s", ctl.String())
	}
}
