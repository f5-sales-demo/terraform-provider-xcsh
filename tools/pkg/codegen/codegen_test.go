// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
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

func TestGenerateResourceFilePreservesExplicitPlanModifierImports(t *testing.T) {
	tmpl := &openapi.ResourceTemplate{
		Name:               "zz_plan_modifier_import_probe",
		TitleCase:          "ZzPlanModifierImportProbe",
		Description:        "Probe.",
		APIPath:            "/api/config/namespaces/%s/zz_plan_modifier_import_probes",
		APIPathItem:        "/api/config/namespaces/%s/zz_plan_modifier_import_probes/%s",
		HasNamespaceInPath: true,
		Attributes: []openapi.TerraformAttribute{
			{
				Name: "ranges", GoName: "Ranges", TfsdkTag: "ranges", JsonName: "ranges",
				Type: "list", ElementType: "string", Required: true, PlanModifier: "RequiresReplace",
			},
		},
	}

	dir := t.TempDir()
	if err := GenerateResourceFile(tmpl, dir); err != nil {
		t.Fatalf("GenerateResourceFile: %v", err)
	}
	content, err := os.ReadFile(filepath.Join(dir, "zz_plan_modifier_import_probe_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	generated := string(content)
	for _, want := range []string{
		"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier",
		"listplanmodifier.RequiresReplace()",
	} {
		if !strings.Contains(generated, want) {
			t.Fatalf("generated resource is missing %q:\n%s", want, generated)
		}
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

	// Suppressed any-depth default marker (common_buffering, a route advanced_options
	// oneof-base default) -> guarded by !isImport. (disable_waf is root-only scoped;
	// see TestRenderUnmarshalSingleChild_RootOnlySuppression_Issue1145.)
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "HTTPLoadBalancer", "", mk("CommonBuffering", "common_buffering"), "apiResource.Spec", "data", "true", "single", "\t")
	got := sb.String()
	if !strings.Contains(got, "if !isImport {") {
		t.Errorf("suppressed member common_buffering must guard response-populate with !isImport; got:\n%s", got)
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

// #1145: a leaf that is a server-default at the resource ROOT (suppressed there so a bare
// LB imports clean) but a user-DECLARED oneof arm when nested must be suppressed ONLY at the
// root. http_loadbalancer disable_waf is the case: at the LB root it is the WAF oneof default;
// inside routes[].{simple_route}.advanced_options it is a per-route "disable WAF for this
// route" choice. Because suppression matches by leaf name at any depth, the nested declared
// disable_waf was stripped on import and drifted (+ disable_waf {}) — forcing custom-routes
// coverage to drop route-level waf_mode=disable (CR-3). Root-only scope fixes it: the nested
// renderer (renderUnmarshalSingleChild, which also renders single blocks inside list elements
// like routes[]) must NOT guard a root-only leaf, so the declared nested marker reads back and
// round-trips. renderUnmarshalTopLevelSingle (root) still suppresses it. A genuinely any-depth
// marker (common_buffering, a route advanced_options oneof-base default) stays guarded nested.
func TestRenderUnmarshalSingleChild_RootOnlySuppression_Issue1145(t *testing.T) {
	mk := func(go_, tfsdk string) openapi.TerraformAttribute {
		return openapi.TerraformAttribute{GoName: go_, TfsdkTag: tfsdk, JsonName: tfsdk, IsBlock: true, NestedBlockType: "single"}
	}

	// Root-only leaf nested inside a route (container="list") -> NOT guarded: the declared
	// per-route disable_waf must read back from the API so it round-trips on import.
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "HTTPLoadBalancer", "RoutesSimpleRouteAdvancedOptions", mk("DisableWaf", "disable_waf"), "advMap", "existingItems[idx]", "len(existingItems) > idx", "list", "\t")
	if strings.Contains(sb.String(), "if !isImport {") {
		t.Errorf("root-only leaf disable_waf must NOT be import-suppressed when nested; got:\n%s", sb.String())
	}

	// Genuinely any-depth suppressed marker (common_buffering) -> STILL guarded nested.
	var sb2 strings.Builder
	renderUnmarshalSingleChild(&sb2, "HTTPLoadBalancer", "RoutesSimpleRouteAdvancedOptions", mk("CommonBuffering", "common_buffering"), "advMap", "existingItems[idx]", "len(existingItems) > idx", "list", "\t")
	if !strings.Contains(sb2.String(), "if !isImport {") {
		t.Errorf("any-depth marker common_buffering must remain import-suppressed when nested; got:\n%s", sb2.String())
	}

	// The root suppression is unchanged: at the LB root, disable_waf is an empty-marker
	// server default and renderUnmarshalTopLevelSingle emits NO populate for it.
	var sb3 strings.Builder
	renderUnmarshalTopLevelSingle(&sb3, "HTTPLoadBalancer", mk("DisableWaf", "disable_waf"), "\t")
	if strings.Contains(sb3.String(), "EmptyModel{}") {
		t.Errorf("root disable_waf must stay suppressed (no populate) at the resource root; got:\n%s", sb3.String())
	}

	// Helper contract: disable_waf is root-only; common_buffering / round_robin are not.
	if !isSuppressionRootOnly("HTTPLoadBalancer", "disable_waf") {
		t.Error("disable_waf must be root-only-scoped for HTTPLoadBalancer")
	}
	if isSuppressionRootOnly("HTTPLoadBalancer", "common_buffering") {
		t.Error("common_buffering is suppressed at any depth and must NOT be root-only-scoped")
	}
	if isSuppressionRootOnly("HTTPLoadBalancer", "round_robin") {
		t.Error("round_robin is suppressed at any depth and must NOT be root-only-scoped")
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
		// #1244: same class on securemesh_site_v2 — the server materializes labels {} on
		// every azure.not_managed.node_list[].interface_list[] element.
		{"SecuremeshSiteV2", "Labels", "labels"},
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

// #1103 / #1244 non-collision: seeding a resource's "labels" suppresses ONLY the nested
// empty-marker blocks (origin_pool origin_servers[].labels, securemesh_site_v2's 13
// *EmptyModel labels). The top-level metadata `labels` is a types.Map emitted by static
// ResourceTemplate text that never consults isImportDefaultSuppressed, and it is written
// BEFORE the isImport marker is even read — so it must never acquire the suppression
// guard. Render a real resource file (TitleCase SecuremeshSiteV2, so the seed is live)
// carrying both shapes and assert the two paths differ: metadata unguarded, nested guarded.
func TestResourceTemplate_MetadataLabelsMapNotImportSuppressed_Issue1244(t *testing.T) {
	marker := func(go_, tfsdk string) openapi.TerraformAttribute {
		return openapi.TerraformAttribute{GoName: go_, TfsdkTag: tfsdk, JsonName: tfsdk, IsBlock: true, NestedBlockType: "single"}
	}
	tmpl := &openapi.ResourceTemplate{
		Name:               "zz_labels_probe",
		TitleCase:          "SecuremeshSiteV2",
		Description:        "Probe.",
		HasNamespaceInPath: true,
		APIPath:            "/api/config/namespaces/%s/zz_labels_probes",
		APIPathItem:        "/api/config/namespaces/%s/zz_labels_probes/%s",
		Attributes: []openapi.TerraformAttribute{
			{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string", JsonName: "name", Required: true},
			{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string", JsonName: "namespace", Required: true},
			// spec.azure.node_list[].labels {} — the nested empty-marker shape the #1244 seed targets.
			{
				Name: "azure", GoName: "Azure", TfsdkTag: "azure", JsonName: "azure",
				IsBlock: true, NestedBlockType: "single", IsSpecField: true, Optional: true,
				NestedAttributes: []openapi.TerraformAttribute{{
					GoName: "NodeList", TfsdkTag: "node_list", JsonName: "node_list",
					IsBlock: true, NestedBlockType: "list", Optional: true,
					NestedAttributes: []openapi.TerraformAttribute{
						{GoName: "NodeName", TfsdkTag: "node_name", JsonName: "node_name", Type: "string", Optional: true},
						marker("Labels", "labels"),
					},
				}},
			},
		},
	}
	dir := t.TempDir()
	if err := GenerateResourceFile(tmpl, dir); err != nil {
		t.Fatalf("GenerateResourceFile: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(dir, "zz_labels_probe_resource.go"))
	if err != nil {
		t.Fatalf("reading rendered resource: %v", err)
	}
	got := string(b)

	// Isolate the Read body — the only place the API response overwrites state.
	readIdx := strings.Index(got, ") Read(ctx context.Context")
	if readIdx == -1 {
		t.Fatal("rendered resource has no Read method")
	}
	read := got[readIdx:]
	if end := strings.Index(read, ") Update(ctx context.Context"); end != -1 {
		read = read[:end]
	}

	// ---- top-level metadata labels/annotations: present, and NOT import-guarded ----
	metaStart := strings.Index(read, "priorLabelsEmpty :=")
	metaEnd := strings.Index(read, "isImport := false")
	if metaStart == -1 || metaEnd == -1 || metaEnd <= metaStart {
		t.Fatalf("cannot locate the metadata read-back region in the rendered Read:\n%s", read)
	}
	meta := read[metaStart:metaEnd]
	for _, want := range []string{
		"data.Labels = types.MapValueMust(types.StringType, nil)",
		"data.Labels = types.MapNull(types.StringType)",
		"data.Annotations = types.MapNull(types.StringType)",
	} {
		if !strings.Contains(meta, want) {
			t.Errorf("metadata read-back must emit %q; region was:\n%s", want, meta)
		}
	}
	// The region ends where isImport is declared, so any mention of it here means the
	// metadata map read-back was pulled onto the import-suppression path (#1244 collision).
	if strings.Contains(meta, "isImport") {
		t.Errorf("metadata labels/annotations read-back must not be import-guarded (#1244 non-collision); region was:\n%s", meta)
	}
	if strings.Contains(meta, "EmptyModel") {
		t.Errorf("metadata labels must stay a types.Map, never an empty-marker block; region was:\n%s", meta)
	}

	// ---- nested labels {} inside node_list[]: IS import-guarded by the #1244 seed ----
	nestedStart := strings.Index(read, "Labels: func() *SecuremeshSiteV2EmptyModel {")
	if nestedStart == -1 {
		t.Fatalf("rendered Read has no nested labels empty-marker closure:\n%s", read)
	}
	nested := read[nestedStart:]
	if end := strings.Index(nested, "}(),"); end != -1 {
		nested = nested[:end]
	}
	if !strings.Contains(nested, "if !isImport {") {
		t.Errorf("nested node_list[].labels {} must guard its response-populate with !isImport (#1244); closure was:\n%s", nested)
	}
}

func TestGenerateResourceFileSynchronizesPlanModifierImportsFromFinalIR(t *testing.T) {
	resource := &openapi.ResourceTemplate{
		Name:                 "zz_plan_modifier_import_probe",
		TitleCase:            "ZzPlanModifierImportProbe",
		Description:          "Probe.",
		APIPath:              "/api/config/namespaces/%s/zz_plan_modifier_import_probes",
		APIPathItem:          "/api/config/namespaces/%s/zz_plan_modifier_import_probes/%s",
		UsesBoolPlanModifier: true,
		Attributes: []openapi.TerraformAttribute{
			{
				Name: "items", GoName: "Items", TfsdkTag: "items", JsonName: "items",
				Type: "list", ElementType: "string", Required: true, PlanModifier: "RequiresReplace",
			},
		},
	}
	dir := t.TempDir()
	if err := GenerateResourceFile(resource, dir); err != nil {
		t.Fatalf("GenerateResourceFile: %v", err)
	}
	if !resource.UsesListPlanModifier || resource.UsesBoolPlanModifier {
		t.Fatalf("final plan-modifier usage did not replace stale template flags: %+v", resource)
	}
	body, err := os.ReadFile(filepath.Join(dir, "zz_plan_modifier_import_probe_resource.go"))
	if err != nil {
		t.Fatalf("read generated resource: %v", err)
	}
	generated := string(body)
	for _, want := range []string{
		`"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"`,
		"listplanmodifier.RequiresReplace()",
	} {
		if !strings.Contains(generated, want) {
			t.Errorf("generated resource is missing %q:\n%s", want, generated)
		}
	}
	if strings.Contains(generated, "boolplanmodifier") {
		t.Fatalf("generated resource retained a stale bool plan-modifier import:\n%s", generated)
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

func TestRenderResponseOperationExampleHCL(t *testing.T) {
	rt := &openapi.ResourceTemplate{Attributes: []openapi.TerraformAttribute{
		{TfsdkTag: "namespace", Type: "string", Required: true},
		{TfsdkTag: "name", Type: "string", Required: true},
		{TfsdkTag: "version", Type: "string", Required: true},
		{TfsdkTag: "force", Type: "bool"},
	}}
	got := RenderResponseOperationExampleHCL(rt, "site_upgrade_sw", "action")
	for _, want := range []string{
		`required_version = ">= 1.14"`,
		`action "xcsh_site_upgrade_sw" "example"`,
		"  config {",
		`namespace = "example-value"`,
		`name = "example-value"`,
		`version = "example-value"`,
		"convergence is asynchronous",
		"does not reconcile a site's pinned software_settings",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("expected action example to contain %q, got:\n%s", want, got)
		}
	}
	if strings.Contains(got, "force =") {
		t.Errorf("optional force must be omitted so its false default is exercised, got:\n%s", got)
	}

	data := RenderResponseOperationExampleHCL(&openapi.ResourceTemplate{Attributes: []openapi.TerraformAttribute{
		{TfsdkTag: "provider_ref", Type: "string", Required: true},
		{TfsdkTag: "image_download_url", Type: "string", Computed: true, Sensitive: true},
	}}, "site_image", "data_source")
	if !strings.Contains(data, `provider_ref = "example-value"`) || strings.Contains(data, "image_download_url =") {
		t.Fatalf("response data-source example did not follow required/computed schema:\n%s", data)
	}
	for _, want := range []string{
		`output "site_image_result"`,
		`value = data.xcsh_site_image.example`,
		`sensitive = true`,
	} {
		if !strings.Contains(data, want) {
			t.Fatalf("response data-source example is missing %q:\n%s", want, data)
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
	got, err := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	// grp has a Computed descendant -> types.List, and must preserve null/empty prior state.
	if !strings.Contains(got, "data.Primary.Grp.IsNull() || len(data.Primary.Grp.Elements()) == 0") {
		t.Errorf("expected unconfigured-list preservation guard for grp, got:\n%s", got)
	}
	if !strings.Contains(got, "return types.ListNull(types.ObjectType{AttrTypes: TestPrimaryGrpModelAttrTypes})") {
		t.Errorf("expected canonical ListNull return for preserved grp, got:\n%s", got)
	}
}

// #1286: the metadata read-back must distinguish "absent from config" from
// "declared empty". A config that declares `labels = {}` (or `annotations = {}`)
// gets no labels back from the F5 XC API, and nulling the attribute on every Read
// makes a plain post-apply plan drift `+ labels = {}` forever — no import involved.
// So when the API returns no entries: keep a KNOWN EMPTY prior value as an empty
// map, and null it only when the prior value is itself null (never declared) or
// unknown. A prior value with entries must still be nulled, so genuine
// out-of-band deletion still shows as drift.
func TestResourceTemplate_PreservesConfigDeclaredEmptyMetadataMap_Issue1286(t *testing.T) {
	// Render a real resource file: the assertion is on emitted Go, not template text.
	tmpl := &openapi.ResourceTemplate{
		Name:               "zz_metadata_probe",
		TitleCase:          "ZZMetadataProbe",
		Description:        "Probe.",
		HasNamespaceInPath: true,
		APIPath:            "/api/config/namespaces/%s/zz_metadata_probes",
		APIPathItem:        "/api/config/namespaces/%s/zz_metadata_probes/%s",
		Attributes: []openapi.TerraformAttribute{
			{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string", JsonName: "name", Required: true},
			{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string", JsonName: "namespace", Required: true},
		},
	}
	dir := t.TempDir()
	if err := GenerateResourceFile(tmpl, dir); err != nil {
		t.Fatalf("GenerateResourceFile: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(dir, "zz_metadata_probe_resource.go"))
	if err != nil {
		t.Fatalf("reading rendered resource: %v", err)
	}
	got := string(b)

	// Isolate the Read body — the only place the API response overwrites metadata maps.
	readIdx := strings.Index(got, ") Read(ctx context.Context")
	if readIdx == -1 {
		t.Fatal("rendered resource has no Read method")
	}
	read := got[readIdx:]
	if end := strings.Index(read, ") Update(ctx context.Context"); end != -1 {
		read = read[:end]
	}

	for _, attr := range []string{"Labels", "Annotations"} {
		nullAssign := "data." + attr + " = types.MapNull(types.StringType)"
		emptyAssign := "data." + attr + " = types.MapValueMust(types.StringType, nil)"
		if !strings.Contains(read, emptyAssign) {
			t.Errorf("Read must preserve a config-declared empty %s as an empty map (#1286); expected %q in:\n%s", attr, emptyAssign, read)
		}
		if !strings.Contains(read, nullAssign) {
			t.Errorf("Read must still null %s when it was never declared (#1286); expected %q in:\n%s", attr, nullAssign, read)
		}
		// Every null assignment must be gated on the prior value being null/unknown —
		// an unguarded null is exactly the #1286 drift.
		for _, guard := range []string{
			"data." + attr + ".IsNull()",
			"data." + attr + ".IsUnknown()",
			"len(data." + attr + ".Elements()) == 0",
		} {
			if !strings.Contains(read, guard) {
				t.Errorf("Read must gate the %s null-vs-empty decision on the prior value (#1286); expected %q in:\n%s", attr, guard, read)
			}
		}
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

	marshal, err := RenderSpecMarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}
	if !strings.Contains(marshal, "RrSet.ElementsAs(ctx,") {
		t.Errorf("expected marshal to ElementsAs the deep rr_set types.List, got:\n%s", marshal)
	}

	unmarshal, err := RenderSpecUnmarshalCode(attrs, "\t", "Test")
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
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

func TestRenderNestedAttributes_FormatValidators(t *testing.T) {
	for format, want := range map[string]string{
		"mac-address": "validators.MACValidator()",
		"ipv4":        "validators.IPv4Validator()",
		"ipv6":        "validators.IPv6Validator()",
		"ip":          "validators.IPValidator()",
		"cidr":        "validators.CIDRValidator()",
	} {
		got := RenderNestedAttributes([]openapi.TerraformAttribute{
			{GoName: "F", TfsdkTag: "f", Type: "string", Format: format}}, "\t")
		if !strings.Contains(got, want) {
			t.Errorf("format %q: expected %q, got:\n%s", format, want, got)
		}
	}
	// Unknown/other formats must not emit any format validator.
	got := RenderNestedAttributes([]openapi.TerraformAttribute{
		{GoName: "F", TfsdkTag: "f", Type: "string", Format: "uri"}}, "\t")
	if strings.Contains(got, "validators.") {
		t.Errorf("format %q: expected no format validator, got:\n%s", "uri", got)
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

func TestRenderUnmarshalSingleChild_ComputedDescendantReconstructs(t *testing.T) {
	block := openapi.TerraformAttribute{
		GoName: "RateLimiter", JsonName: "rate_limiter", TfsdkTag: "rate_limiter",
		IsBlock: true, NestedBlockType: "single",
		NestedAttributes: []openapi.TerraformAttribute{
			{GoName: "PeriodMultiplier", TfsdkTag: "period_multiplier", JsonName: "period_multiplier", Type: "int64", Optional: true, Computed: true},
		},
	}
	var sb strings.Builder
	renderUnmarshalSingleChild(&sb, "R", "RateLimitRateLimiter", block,
		"blockData", "data.RateLimit", "data.RateLimit != nil", "single", "\t")
	out := sb.String()
	if strings.Contains(out, "return data.RateLimit.RateLimiter\n") {
		t.Errorf("block with a Computed descendant must not preserve an unknown planned value:\n%s", out)
	}
	if !strings.Contains(out, `RateLimiterData["period_multiplier"]`) {
		t.Errorf("block with a Computed descendant must reconstruct the leaf from the API response:\n%s", out)
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

// TestGenerateClientTypes_ExposeUID verifies that when a resource opts in to
// ExposeUID, the generated client type carries a SystemMetadata field and a
// per-resource *SystemMetadata type with a uid; and that a resource without
// ExposeUID is unchanged (no SystemMetadata anywhere).
func TestGenerateClientTypes_ExposeUID(t *testing.T) {
	base := &openapi.ResourceTemplate{
		Name:               "token",
		TitleCase:          "Token",
		APIPath:            "/api/register/namespaces/%s/tokens",
		APIPathItem:        "/api/register/namespaces/%s/tokens/%s",
		HasNamespaceInPath: true,
	}

	render := func(exposeUID bool) string {
		r := *base
		r.ExposeUID = exposeUID
		dir := t.TempDir()
		if err := GenerateClientTypes(&r, dir); err != nil {
			t.Fatalf("GenerateClientTypes error: %v", err)
		}
		data, err := os.ReadFile(filepath.Join(dir, "token_types.go"))
		if err != nil {
			t.Fatalf("reading generated file: %v", err)
		}
		return string(data)
	}

	with := render(true)
	// Substrings chosen to survive gofmt column alignment of struct fields.
	for _, want := range []string{
		"SystemMetadata *TokenSystemMetadata",
		"json:\"system_metadata,omitempty\"",
		"type TokenSystemMetadata struct",
		"UID string",
		"json:\"uid,omitempty\"",
	} {
		if !strings.Contains(with, want) {
			t.Errorf("ExposeUID=true output missing %q; got:\n%s", want, with)
		}
	}

	without := render(false)
	if strings.Contains(without, "SystemMetadata") {
		t.Errorf("ExposeUID=false output must not mention SystemMetadata; got:\n%s", without)
	}
}

// TestActionResourceApprove verifies the action-resource codegen: Create issues
// the action POST to the singular path with state=APPROVED, Read does a lenient
// GET on the pluralized sibling path with 404 -> remove-from-state, Delete is a
// no-op, there is no PUT/update and no data-source companion, every settable
// attribute forces replace, and the client types file carries the request struct.
func TestActionResourceApprove(t *testing.T) {
	tmpl := &openapi.ResourceTemplate{
		Name:               "registration_approval",
		TitleCase:          "RegistrationApproval",
		Description:        "Approve a registration.",
		HasNamespaceInPath: true,
		HasStringDefaults:  true,
		IsAction:           true,
		ActionPath:         "/api/register/namespaces/%s/registration/%s/approve",
		ActionState:        "APPROVED",
		ReadObjectPath:     "/api/register/namespaces/%s/registrations/%s",
		Attributes: []openapi.TerraformAttribute{
			{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string", JsonName: "namespace", Required: true, PlanModifier: "RequiresReplace"},
			{Name: "backup_connected_region", GoName: "BackupConnectedRegion", TfsdkTag: "backup_connected_region", Type: "string", JsonName: "backup_connected_region", Optional: true, PlanModifier: "RequiresReplace"},
			{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string", JsonName: "name", Required: true, PlanModifier: "RequiresReplace"},
			{Name: "state", GoName: "State", TfsdkTag: "state", Type: "string", JsonName: "state", Optional: true, Computed: true, StringDefault: "APPROVED", PlanModifier: "RequiresReplace"},
		},
	}

	outDir := t.TempDir()
	clientDir := t.TempDir()
	if err := GenerateActionResource(tmpl, outDir, clientDir); err != nil {
		t.Fatalf("GenerateActionResource error: %v", err)
	}

	resourceBytes, err := os.ReadFile(filepath.Join(outDir, "registration_approval_resource.go"))
	if err != nil {
		t.Fatalf("reading resource file: %v", err)
	}
	res := string(resourceBytes)
	typesBytes, err := os.ReadFile(filepath.Join(clientDir, "registration_approval_types.go"))
	if err != nil {
		t.Fatalf("reading types file: %v", err)
	}
	types := string(typesBytes)

	// Create → action POST on the singular approve path, body state = APPROVED.
	for _, want := range []string{"r.client.Post(ctx,", "/api/register/namespaces/%s/registration/%s/approve", `"APPROVED"`} {
		if !strings.Contains(res, want) {
			t.Errorf("resource missing Create marker %q; got:\n%s", want, res)
		}
	}
	// Read → lenient GET on the pluralized sibling path; 404 → remove from state.
	for _, want := range []string{"r.client.GetLenient(ctx,", "/api/register/namespaces/%s/registrations/%s", "resp.State.RemoveResource(ctx)"} {
		if !strings.Contains(res, want) {
			t.Errorf("resource missing Read marker %q; got:\n%s", want, res)
		}
	}
	// Delete is a no-op — no client Delete call.
	if strings.Contains(res, "r.client.Delete") {
		t.Errorf("action Delete must be a no-op (no client Delete call); got:\n%s", res)
	}
	// No in-place update — no PUT.
	if strings.Contains(res, "r.client.Put") {
		t.Errorf("action resource must not issue an Update/PUT; got:\n%s", res)
	}
	// Every settable schema attribute forces replace (name, namespace, passport, state).
	if n := strings.Count(res, "RequiresReplace()"); n < 4 {
		t.Errorf("expected >=4 RequiresReplace(), got %d;\n%s", n, res)
	}
	// No data-source companion for an action resource.
	if _, statErr := os.Stat(filepath.Join(outDir, "registration_approval_data_source.go")); statErr == nil {
		t.Error("action resource must not generate a data source companion")
	}
	// Client types file carries the request struct.
	if !strings.Contains(types, "type RegistrationApproval struct") {
		t.Errorf("types file missing request struct; got:\n%s", types)
	}
}

// The approve POST must carry the registration's own passport: without it F5
// answers HTTP 500 "Validation approval: Passport is required" and the resource
// can never create (#1355). This drives the WHOLE pipeline — the real approve
// spec shape through ExtractActionResourceSchema into the generated Go — so it
// covers the extractor, the server-derived declaration and both templates.
func TestActionResourceCreateSendsServerDerivedPassport(t *testing.T) {
	spec, action := actionApproveSpecForCodegen()

	tmpl, err := schema.ExtractActionResourceSchema(spec, action)
	if err != nil {
		t.Fatalf("ExtractActionResourceSchema: %v", err)
	}

	outDir := t.TempDir()
	clientDir := t.TempDir()
	if err := GenerateActionResource(tmpl, outDir, clientDir); err != nil {
		t.Fatalf("GenerateActionResource error: %v", err)
	}
	resBytes, err := os.ReadFile(filepath.Join(outDir, "registration_approval_resource.go"))
	if err != nil {
		t.Fatalf("reading resource file: %v", err)
	}
	res := string(resBytes)
	typeBytes, err := os.ReadFile(filepath.Join(clientDir, "registration_approval_types.go"))
	if err != nil {
		t.Fatalf("reading types file: %v", err)
	}
	types := string(typeBytes)

	// The request struct must be able to carry the passport, and must carry it
	// as an opaque value: the server accepts only its own object echoed back, so
	// it may not be flattened into a string.
	if !regexp.MustCompile(`\bPassport\s+interface\{\}`).MatchString(types) {
		t.Errorf("request struct must carry Passport as an opaque interface{} value; got:\n%s", types)
	}
	if !strings.Contains(types, `json:"passport,omitempty"`) {
		t.Errorf("request struct Passport must marshal to the wire key \"passport\"; got:\n%s", types)
	}

	// Create must read the sibling registration and assign the passport from it.
	create := funcBody(t, res, "Create")
	for _, want := range []string{
		"r.client.GetLenient(ctx,",
		"/api/register/namespaces/%s/registrations/%s",
		"client.LookupNestedField(",
		"body.Passport =",
	} {
		if !strings.Contains(create, want) {
			t.Errorf("Create missing server-derived passport marker %q; got:\n%s", want, create)
		}
	}
	// At least one declared lookup path must address the registration's passport.
	if !regexp.MustCompile(`"[a-z_.]*\.passport"`).MatchString(create) {
		t.Errorf("Create must look the passport up by path in the registration read; got:\n%s", create)
	}
	// Order matters: the read has to happen BEFORE the approve POST.
	readAt := strings.Index(create, "r.client.GetLenient(ctx,")
	postAt := strings.Index(create, "r.client.Post(ctx,")
	if readAt < 0 || postAt < 0 || readAt > postAt {
		t.Errorf("Create must read the registration before POSTing the approve (readAt=%d postAt=%d); got:\n%s", readAt, postAt, create)
	}
	// A missing passport must fail loudly rather than repeat the silent 500.
	if !strings.Contains(create, "resp.Diagnostics.AddError") || !strings.Contains(create, "passport") {
		t.Errorf("Create must raise a diagnostic naming the passport when it cannot be derived; got:\n%s", create)
	}

	// Server-derived is not user input: no passport attribute, no model field.
	if strings.Contains(res, `"passport": schema.StringAttribute`) {
		t.Errorf("passport must not be exposed as a user-settable attribute; got:\n%s", res)
	}
	if regexp.MustCompile(`\bPassport\s+types\.String\b`).MatchString(res) {
		t.Errorf("passport must not be a Terraform model field; got:\n%s", res)
	}
}

// funcBody returns the source of the named method on the generated resource,
// from its `func (r *…) <name>(` line to the next top-level `func ` line.
func funcBody(t *testing.T, src, name string) string {
	t.Helper()
	start := regexp.MustCompile(`(?m)^func \(r \*[A-Za-z]+Resource\) ` + name + `\(`).FindStringIndex(src)
	if start == nil {
		t.Fatalf("method %s not found in generated source:\n%s", name, src)
	}
	rest := src[start[1]:]
	if next := regexp.MustCompile(`(?m)^func `).FindStringIndex(rest); next != nil {
		return rest[:next[0]]
	}
	return rest
}

// actionApproveSpecForCodegen mirrors the REAL registration approve request: a
// camelCase "registrationApprovalReq" carrying x-f5xc-action, scalar string
// props, object props (annotations, labels), non-enum $ref props (passport,
// tunnel_type) and a $ref-enum state, plus the approve POST path and the plural
// sibling GET path.
func actionApproveSpecForCodegen() (*openapi.Spec, openapi.ResourcePath) {
	spec := &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"registrationApprovalReq": {
					Type:        "object",
					XF5xcAction: "approve",
					Properties: map[string]openapi.Schema{
						"name":                    {Type: "string"},
						"backup_connected_region": {Type: "string"},
						"connected_region":        {Type: "string"},
						"preferred_active_re":     {Type: "string"},
						"annotations":             {Type: "object"},
						"labels":                  {Type: "object"},
						"passport":                {AllOf: []openapi.Schema{{Ref: "#/components/schemas/registrationPassport"}}},
						"tunnel_type":             {AllOf: []openapi.Schema{{Ref: "#/components/schemas/schemaSiteToSiteTunnelType"}}},
						"state":                   {Ref: "#/components/schemas/registrationObjectState"},
					},
				},
			},
		},
		Paths: map[string]interface{}{
			"/api/register/namespaces/{namespace}/registration/{name}/approve": map[string]interface{}{
				"post": map[string]interface{}{
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/registrationApprovalReq",
								},
							},
						},
					},
				},
			},
			"/api/register/namespaces/{namespace}/registrations/{name}": map[string]interface{}{
				"get": map[string]interface{}{},
			},
		},
	}
	action := openapi.ResourcePath{
		ResourceName:   "registration_approval",
		SchemaName:     "registrationApprovalReq",
		ActionValue:    "approve",
		ActionPath:     "/api/register/namespaces/%s/registration/%s/approve",
		ReadObjectPath: "/api/register/namespaces/%s/registrations/%s",
	}
	return spec, action
}

// repoRootFromTest returns the module root by walking up from this test file's
// location (…/tools/pkg/codegen/codegen_test.go -> module root).
func repoRootFromTest(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	// file = <root>/tools/pkg/codegen/codegen_test.go
	root := filepath.Join(filepath.Dir(file), "..", "..", "..")
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err != nil {
		t.Fatalf("could not locate module root (no go.mod at %s): %v", root, err)
	}
	return root
}

// approveCompileSpec mirrors the REAL registration-approval action spec shape:
// a camelCase request-body component ("registrationApprovalReq") carrying
// x-f5xc-action, a mix of scalar string props, object props (annotations,
// labels), non-enum $ref props (passport, tunnel_type) and a $ref-enum `state`,
// and an approve POST path whose {namespace} segment is NOT a body property.
func approveCompileSpec() *openapi.Spec {
	return &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"zzCompileProbeReq": {
					Type:        "object",
					XF5xcAction: "approve",
					Properties: map[string]openapi.Schema{
						"name":                    {Type: "string"},
						"backup_connected_region": {Type: "string"},
						"annotations":             {Type: "object"},
						"labels":                  {Type: "object"},
						"passport":                {AllOf: []openapi.Schema{{Ref: "#/components/schemas/registrationPassport"}}},
						"tunnel_type":             {AllOf: []openapi.Schema{{Ref: "#/components/schemas/schemaSiteToSiteTunnelType"}}},
						"state":                   {Ref: "#/components/schemas/registrationObjectState"},
					},
				},
			},
		},
		Paths: map[string]interface{}{
			"/api/register/namespaces/{namespace}/compile_probe/{name}/approve": map[string]interface{}{
				"post": map[string]interface{}{
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"$ref": "#/components/schemas/zzCompileProbeReq",
								},
							},
						},
					},
				},
			},
			"/api/register/namespaces/{namespace}/compile_probes/{name}": map[string]interface{}{
				"get": map[string]interface{}{},
			},
		},
	}
}

// TestActionResourceCompiles is the compile-level guard for action-resource
// codegen. It runs the full pipeline from an in-memory spec whose shape mirrors
// the real registration-approval action (camelCase Req schema; {namespace} path
// param that is NOT a body property; object + non-enum $ref props that must be
// excluded), then WRITES the generated Go into the module tree and runs
// `go build` to prove the output actually COMPILES.
//
// The fixture uses a SYNTHETIC resource name ("zz_compile_probe") rather than the
// real "registration_approval" on purpose: the real resource ships in
// internal/provider + internal/client once generated, and its compilation is
// covered by the normal `go build ./...` / `go test ./internal/...`. Reusing the
// real name here would make this guard write (and refuse to clobber) the
// committed internal/client/registration_approval_types.go, so it would fail
// permanently after the first regen. A synthetic name can never collide, so the
// guard stays live on every branch while still exercising the exact generator
// paths (snake_case derivation, {namespace} injection, object/$ref exclusion).
//
// This is the gap that previously let a broken generator ship: the earlier tests
// only asserted rendered substrings and never compiled the result, so the
// `data.Namespace undefined` type error (namespace was never emitted as a model
// field) went unnoticed. The generated resource file imports internal/client and
// internal/validators, which — by Go's internal-package visibility rule — are
// importable only from within this module; a bare t.TempDir() outside the module
// cannot resolve them. So the throwaway resource package is created UNDER
// internal/ (unique dir), the client request struct is dropped into
// internal/client, both are built, and both are removed on cleanup. A dir whose
// package clause is `provider` builds fine via `go build ./internal/<dir>/`
// without colliding with the real provider package.
func TestActionResourceCompiles(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping go build compile check in -short mode")
	}

	spec := approveCompileSpec()

	// 1. Schema extraction from the exact operation-catalog action: exact
	// attribute set, object/$ref props excluded.
	action := openapi.ResourcePath{
		ResourceName:   "zz_compile_probe",
		SchemaName:     "zzCompileProbeReq",
		ActionValue:    "approve",
		ActionPath:     "/api/register/namespaces/%s/compile_probe/%s/approve",
		ReadObjectPath: "/api/register/namespaces/%s/compile_probes/%s",
	}
	tmpl, err := schema.ExtractActionResourceSchema(spec, action)
	if err != nil {
		t.Fatalf("ExtractActionResourceSchema: %v", err)
	}
	got := map[string]openapi.TerraformAttribute{}
	for _, a := range tmpl.Attributes {
		got[a.TfsdkTag] = a
	}
	for _, excluded := range []string{"annotations", "labels", "passport", "tunnel_type"} {
		if _, ok := got[excluded]; ok {
			t.Errorf("object/$ref prop %q must be excluded from action attributes", excluded)
		}
	}
	if ns, ok := got["namespace"]; !ok || !ns.Required {
		t.Errorf("namespace must be present and Required (injected path param); got %+v ok=%v", got["namespace"], ok)
	}

	// 3. Generate into the module tree and compile.
	root := repoRootFromTest(t)
	clientDir := filepath.Join(root, "internal", "client")
	clientTypes := filepath.Join(clientDir, "zz_compile_probe_types.go")
	if _, err := os.Stat(clientTypes); err == nil {
		t.Fatalf("refusing to clobber existing %s", clientTypes)
	}
	provDir, err := os.MkdirTemp(filepath.Join(root, "internal"), "zz_actioncompile_")
	if err != nil {
		t.Fatalf("mkdir temp provider dir: %v", err)
	}
	t.Cleanup(func() {
		os.RemoveAll(provDir)
		os.Remove(clientTypes)
	})

	if err := GenerateActionResource(tmpl, provDir, clientDir); err != nil {
		t.Fatalf("GenerateActionResource: %v", err)
	}

	resBytes, err := os.ReadFile(filepath.Join(provDir, "zz_compile_probe_resource.go"))
	if err != nil {
		t.Fatalf("reading generated resource: %v", err)
	}
	res := string(resBytes)
	// Model must carry a Namespace field (the bug: it did not). gofmt aligns
	// struct fields, so match name + type across arbitrary whitespace.
	if !regexp.MustCompile(`\bNamespace\s+types\.String\b`).MatchString(res) {
		t.Errorf("generated model missing Namespace field:\n%s", res)
	}
	// snake_case TF type name, not camelCase struct/typename.
	if !strings.Contains(res, `+ "_zz_compile_probe"`) {
		t.Errorf("generated resource missing snake_case type name _zz_compile_probe:\n%s", res)
	}
	if strings.Contains(res, "Zzcompileprobe") {
		t.Errorf("generated resource contains camelCase-collapsed name 'Zzcompileprobe' (snake_case derivation failed):\n%s", res)
	}
	// Excluded props must not appear as model fields.
	for _, bad := range []string{"Labels", "Passport", "TunnelType", "Annotations"} {
		if regexp.MustCompile(`\b` + bad + `\s+types\.String\b`).MatchString(res) {
			t.Errorf("generated model contains excluded field %q:\n%s", bad, res)
		}
	}

	buildTarget := "./internal/" + filepath.Base(provDir) + "/"
	cmd := exec.Command("go", "build", buildTarget)
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generated action resource failed to compile (%v):\n%s", err, out)
	}
}

func TestRenderNestedAttributes_Int64MinZero(t *testing.T) {
	attrs := []openapi.TerraformAttribute{
		{GoName: "Priority", TfsdkTag: "priority", Type: "int64", Minimum: 0, HasMinimum: true, Maximum: 255, HasMaximum: true},
	}
	got := RenderNestedAttributes(attrs, "\t")
	if !strings.Contains(got, "int64validator.Between(0, 255)") {
		t.Errorf("expected int64validator.Between(0, 255), got:\n%s", got)
	}
}

func TestResourceTemplate_TopLevelFormatAndZeroBoundValidators(t *testing.T) {
	tmpl := &openapi.ResourceTemplate{
		Name: "zz_top_level_validator_probe", TitleCase: "TopLevelValidatorProbe",
		Description: "Probe.", APIPath: "/api/config/zz_top_level_validator_probes",
		APIPathItem:             "/api/config/zz_top_level_validator_probes/%s",
		HasInt64RangeValidators: true,
		Attributes: []openapi.TerraformAttribute{
			{Name: "address", GoName: "Address", TfsdkTag: "address", JsonName: "address", Type: "string", Optional: true, Format: "ipv4", IsSpecField: true},
			{Name: "count", GoName: "Count", TfsdkTag: "count", JsonName: "count", Type: "int64", Optional: true, HasMinimum: true, Minimum: 0, HasMaximum: true, Maximum: 10, IsSpecField: true},
		},
	}
	dir := t.TempDir()
	if err := GenerateResourceFile(tmpl, dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "zz_top_level_validator_probe_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	got := string(data)
	for _, want := range []string{"validators.IPv4Validator()", "int64validator.Between(0, 10)"} {
		if !strings.Contains(got, want) {
			t.Fatalf("rendered top-level schema missing %q:\n%s", want, got)
		}
	}
}

func TestRenderNestedAttributes_DiscontinuousInt64Range(t *testing.T) {
	got := RenderNestedAttributes([]openapi.TerraformAttribute{{
		Name: "mtu", GoName: "MTU", TfsdkTag: "mtu", JsonName: "mtu", Type: "int64", Optional: true,
		Int64RangeSpans: []openapi.Int64RangeSpan{{Minimum: 0, Maximum: 0}, {Minimum: 512, Maximum: 16384}},
	}}, "\t")
	for _, want := range []string{
		"validators.Int64RangeSetValidator(",
		"validators.Int64Range{Minimum: 0, Maximum: 0}",
		"validators.Int64Range{Minimum: 512, Maximum: 16384}",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("rendered nested schema missing %q:\n%s", want, got)
		}
	}
}

func TestRenderConditionalRequiredValidators(t *testing.T) {
	child := openapi.TerraformAttribute{TfsdkTag: "required_child", CreateRequired: true}
	for _, test := range []struct {
		name  string
		block openapi.TerraformAttribute
		want  string
	}{
		{"single", openapi.TerraformAttribute{NestedBlockType: "single", NestedAttributes: []openapi.TerraformAttribute{child}}, `Validators: []validator.Object{validators.RequiredObjectAttributes("required_child")}`},
		{"list", openapi.TerraformAttribute{NestedBlockType: "list", NestedAttributes: []openapi.TerraformAttribute{child}}, `Validators: []validator.List{validators.RequiredListObjectAttributes("required_child")}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := RenderConditionalRequiredValidators(test.block, "\t")
			if !strings.Contains(got, test.want) {
				t.Fatalf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestRenderConditionalRequiredValidators_ServerMaterializedViolationsView(t *testing.T) {
	block := openapi.TerraformAttribute{
		NestedBlockType: "single",
		NestedAttributes: []openapi.TerraformAttribute{{
			TfsdkTag:       "violations_view",
			CreateRequired: true,
		}},
	}
	if got := RenderConditionalRequiredValidators(block, "\t"); got != "" {
		t.Fatalf("server-materialized violations_view must not be required in configuration: %q", got)
	}
}

// #1391: the discovery-label filter is a convergence device for the resource Read. It
// must be driven by FiltersDiscoveredSiteLabels there, and must never be switched on in
// a data source: a data source cannot propose deleting a label, so filtering there buys
// no convergence and only hides hw-model and friends from a caller that asked for the
// object. A local review caught exactly that regression before this test existed.
func TestTemplates_DiscoveryLabelFilterIsResourceOnly_Issue1391(t *testing.T) {
	// #1398 changed the call to filterSystemLabelsOwning, which takes the same
	// per-resource flag plus the set of keys the configuration declared. The invariant
	// this test protects is unchanged and still asserted: the flag drives the filter.
	resourceCall := "filterSystemLabelsOwning(apiResource.Metadata.Labels, {{.FiltersDiscoveredSiteLabels}}, ownedLabels)"
	if !strings.Contains(ResourceTemplate, resourceCall) {
		t.Errorf("the resource template must drive the discovery filter from the per-resource flag; expected %q", resourceCall)
	}

	for name, tmpl := range map[string]string{
		"DataSourceTemplate":         DataSourceTemplate,
		"ReadOnlyDataSourceTemplate": ReadOnlyDataSourceTemplate,
	} {
		if strings.Contains(tmpl, "{{.FiltersDiscoveredSiteLabels}}") {
			t.Errorf("%s must not filter discovery labels: a data source has no plan to converge, "+
				"and hiding them breaks callers reading labels[\"hw-model\"]", name)
		}
		if !strings.Contains(tmpl, "filterSystemLabels(resource.Metadata.Labels, false)") {
			t.Errorf("%s must call filterSystemLabels with siteDiscovery=false", name)
		}
	}
}

// #1396: F5 XC replaces metadata.labels on write, so a PUT built only from the
// configuration erases the platform's discovery labels — measured live, a site holding
// all six came back holding none after an apply. Read stashes them in private state and
// Update sends them back. Both halves are gated on the per-resource flag, so a resource
// F5 XC does not decorate gets no new API surface and no unused import.
func TestResourceTemplate_PreservesPlatformLabelsAcrossAWrite_Issue1396(t *testing.T) {
	render := func(t *testing.T, decorated bool) string {
		t.Helper()
		tmpl := &openapi.ResourceTemplate{
			Name:                        "zz_label_probe",
			TitleCase:                   "ZZLabelProbe",
			Description:                 "Probe.",
			HasNamespaceInPath:          true,
			APIPath:                     "/api/config/namespaces/%s/zz_label_probes",
			APIPathItem:                 "/api/config/namespaces/%s/zz_label_probes/%s",
			FiltersDiscoveredSiteLabels: decorated,
			Attributes: []openapi.TerraformAttribute{
				{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string", JsonName: "name", Required: true},
				{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string", JsonName: "namespace", Required: true},
			},
		}
		dir := t.TempDir()
		if err := GenerateResourceFile(tmpl, dir); err != nil {
			t.Fatalf("GenerateResourceFile(decorated=%v): %v", decorated, err)
		}
		b, err := os.ReadFile(filepath.Join(dir, "zz_label_probe_resource.go"))
		if err != nil {
			t.Fatalf("reading rendered resource: %v", err)
		}
		return string(b)
	}

	decorated := render(t, true)
	for _, want := range []string{
		// Fetched immediately before the write, never carried from an earlier Read: a
		// saved plan can be arbitrarily old, and writing back a stale serial number or
		// OS version would replace live inventory with wrong inventory that Read hides.
		"current, currentErr := r.client.GetZZLabelProbe(ctx,",
		"Unable to Read Current Labels Before Update",
		"preservedPlatformLabels(current.Metadata.Labels)",
		"mergePreservedLabels(",
		// A discovery-named label the configuration sets cannot converge, so say so
		// instead of planning the same addition forever (#1398).
		"isDiscoveredSiteLabel(key)",
		"Label Authored by F5 Distributed Cloud",
		// Backstop for a labels map that is unknown at validate time.
		"discoveredSiteLabelKeys(apiResource.Metadata.Labels)",
		"discoveredSiteLabelKeys(createReq.Metadata.Labels)",
	} {
		if !strings.Contains(decorated, want) {
			t.Errorf("a decorated resource must preserve platform labels across a write (#1396); missing %q", want)
		}
	}

	// A failed pre-write fetch must abort. Proceeding would send a payload with no
	// platform labels in it, which is exactly the erasure this exists to prevent.
	fetchIdx := strings.Index(decorated, "Unable to Read Current Labels Before Update")
	if fetchIdx == -1 || !strings.Contains(decorated[fetchIdx:fetchIdx+900], "return") {
		t.Error("a failed pre-write label fetch must return, not fall through to the write (#1396)")
	}

	// The merge must not be nested inside the "configuration set some labels" branch:
	// the case that erased the fleet's labels is a configuration with NO labels block.
	mergeIdx := strings.Index(decorated, "mergePreservedLabels(")
	guardIdx := strings.Index(decorated, "if !data.Labels.IsNull() {")
	if mergeIdx == -1 || guardIdx == -1 || mergeIdx < guardIdx {
		t.Fatal("could not locate the Update label marshalling to check the merge is unguarded")
	}
	between := decorated[guardIdx:mergeIdx]
	if !strings.Contains(between, "apiResource.Metadata.Labels = labels\n\t}") {
		t.Error("the platform-label merge must run whether or not the configuration sets labels (#1396): " +
			"it appears before the !data.Labels.IsNull() block closes, so it is nested inside the branch " +
			"that is NOT the case which erases them")
	}

	undecorated := render(t, false)
	for _, unwanted := range []string{"preservedPlatformLabels", "mergePreservedLabels", "isDiscoveredSiteLabel", "discoveredSiteLabelKeys"} {
		if strings.Contains(undecorated, unwanted) {
			t.Errorf("a resource F5 XC does not decorate must get none of the preservation mechanism; found %q", unwanted)
		}
	}
}

// #1398: a configuration must be able to own a label in the ves.io/ namespace. That
// works only if the set of configuration-declared keys is recorded in private state at
// Create and Update — the only points where the configuration is visible — and consulted
// in Read, which sees prior state and cannot otherwise tell an owned label from one the
// platform added.
func TestTemplates_RecordOwnedLabelKeys_Issue1398(t *testing.T) {
	record := "encodeOwnedLabelKeys(configLabelKeys("
	if got := strings.Count(ResourceTemplate, record); got != 2 {
		t.Errorf("expected ownership to be recorded in exactly Create and Update, found %d call(s) to %q", got, record)
	}
	if !strings.Contains(ResourceTemplate, "req.Private.GetKey(ctx, ownedLabelKeysPrivateKey)") {
		t.Error("Read must consult the recorded ownership; without it a ves.io/ label never converges")
	}

	// The ordering that matters. Update merges the platform's own discovery labels into
	// the outgoing map (#1391); recording ownership AFTER that merge would file those
	// platform labels as configuration-owned, so Read would stop filtering them and the
	// plan would propose deleting them again — the exact bug #1391 fixed.
	updateIdx := strings.Index(ResourceTemplate, "func (r *{{.TitleCase}}Resource) Update(")
	if updateIdx < 0 {
		t.Fatal("Update method not found in the resource template")
	}
	update := ResourceTemplate[updateIdx:]
	recordIdx := strings.Index(update, record)
	mergeIdx := strings.Index(update, "mergePreservedLabels(")
	if recordIdx < 0 {
		t.Fatal("Update does not record owned label keys")
	}
	if mergeIdx >= 0 && recordIdx > mergeIdx {
		t.Error("Update records ownership after mergePreservedLabels; the platform's own " +
			"discovery labels would be recorded as configuration-owned, reintroducing #1391")
	}

	// Data sources have no plan to converge and no private state to read, so they must
	// keep using the plain filter.
	for name, tmpl := range map[string]string{
		"DataSourceTemplate":         DataSourceTemplate,
		"ReadOnlyDataSourceTemplate": ReadOnlyDataSourceTemplate,
	} {
		if strings.Contains(tmpl, "ownedLabelKeys") {
			t.Errorf("%s must not consult label ownership: it has no plan to converge", name)
		}
	}
}

// #1398, second round: private state is persisted even when Create or Update returns an
// error. terraform-plugin-framework v1.17.0 copies createResp.Private / updateResp.Private
// into the response (server_createresource.go:150, server_updateresource.go:165) BEFORE
// its `if resp.Diagnostics.HasError() { return }`.
//
// So recording ownership before the API call is unsafe. Remove an owned ves.io/ label,
// have the update fail, and private state records "owns nothing" while the server still
// holds the label: the next Read filters it out, the plan shows nothing, and the label is
// stranded on the server where Terraform can never see or remove it.
//
// The keys must still be CAPTURED before mergePreservedLabels — that ordering is asserted
// separately — but they may only be WRITTEN once the API call has succeeded.
func TestTemplates_OwnershipIsPersistedOnlyAfterTheAPICall_Issue1398(t *testing.T) {
	for _, tc := range []struct{ method, apiCall string }{
		{"Create", "r.client.Create{{.TitleCase}}(ctx, createReq)"},
		{"Update", "r.client.Update{{.TitleCase}}(ctx, apiResource)"},
	} {
		t.Run(tc.method, func(t *testing.T) {
			start := strings.Index(ResourceTemplate, "func (r *{{.TitleCase}}Resource) "+tc.method+"(")
			if start < 0 {
				t.Fatalf("%s method not found", tc.method)
			}
			body := ResourceTemplate[start:]

			apiIdx := strings.Index(body, tc.apiCall)
			setIdx := strings.Index(body, "Private.SetKey(ctx, ownedLabelKeysPrivateKey")
			if apiIdx < 0 {
				t.Fatalf("%s: API call anchor %q not found — this guard would silently pass", tc.method, tc.apiCall)
			}
			if setIdx < 0 {
				t.Fatalf("%s: ownership is never written to private state", tc.method)
			}
			if setIdx < apiIdx {
				t.Errorf("%s writes ownership to private state before the API call. Private state "+
					"survives an error, so a failed write would strand the label on the server, "+
					"invisible to Terraform.", tc.method)
			}
		})
	}
}

func TestConcurrencyTokenGenerationIsClientOnlyAndPrivate(t *testing.T) {
	tmpl := &openapi.ResourceTemplate{
		Name:                     "zz_token_probe",
		TitleCase:                "ZZTokenProbe",
		Description:              "Probe.",
		HasNamespaceInPath:       true,
		APIPath:                  "/api/config/namespaces/%s/zz_token_probes",
		APIPathItem:              "/api/config/namespaces/%s/zz_token_probes/%s",
		HasConcurrencyToken:      true,
		ConcurrencyTokenJSONName: "resource_version",
		ConcurrencyTokenGoName:   "ResourceVersion",
		Attributes: []openapi.TerraformAttribute{
			{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string", Required: true},
			{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string", Required: true},
			{Name: "id", GoName: "ID", TfsdkTag: "id", Type: "string", Computed: true},
		},
	}
	dir := t.TempDir()
	if err := GenerateResourceFile(tmpl, dir); err != nil {
		t.Fatalf("GenerateResourceFile: %v", err)
	}
	if err := GenerateClientTypes(tmpl, dir); err != nil {
		t.Fatalf("GenerateClientTypes: %v", err)
	}
	resourceBytes, err := os.ReadFile(filepath.Join(dir, "zz_token_probe_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	clientBytes, err := os.ReadFile(filepath.Join(dir, "zz_token_probe_types.go"))
	if err != nil {
		t.Fatal(err)
	}
	resourceSource := string(resourceBytes)
	clientSource := string(clientBytes)

	if !strings.Contains(clientSource, "ResourceVersion") || !strings.Contains(clientSource, "`json:\"resource_version,omitempty\"`") {
		t.Fatal("client model is missing the conditional resource_version field")
	}
	for _, forbidden := range []string{`tfsdk:"resource_version"`, `"resource_version": schema.`, "data.ResourceVersion"} {
		if strings.Contains(resourceSource, forbidden) {
			t.Fatalf("concurrency token leaked into Terraform schema/state: %q", forbidden)
		}
	}
	for _, want := range []string{
		"concurrencyTokenPrivateKey",
		"_, err := r.client.CreateZZTokenProbe(ctx, createReq)",
		"apiResource, err := r.client.GetZZTokenProbe(ctx",
		"req.Private.GetKey(ctx, concurrencyTokenPrivateKey)",
		"encodeConcurrencyToken(apiResource.ResourceVersion)",
		"apiResource.ResourceVersion = concurrencyToken",
		"Stale Configuration",
		"update of zz_token_probe",
		"Refresh Required Before Update",
	} {
		if !strings.Contains(resourceSource, want) {
			t.Errorf("generated resource is missing concurrency behavior %q", want)
		}
	}
	if strings.Contains(resourceSource, "apiResource, err := r.client.CreateZZTokenProbe(ctx, createReq)") {
		t.Fatal("concurrency-token Create must discard the response that its mandatory GET replaces")
	}
	if strings.Contains(resourceSource, "apiResource = fetched") {
		t.Fatal("metadata-only Update must not assign a GET response it has no state fields to consume")
	}

	updateStart := strings.Index(resourceSource, "func (r *ZZTokenProbeResource) Update(")
	if updateStart < 0 {
		t.Fatal("generated Update not found")
	}
	update := resourceSource[updateStart:]
	privateRead := strings.Index(update, "req.Private.GetKey(ctx, concurrencyTokenPrivateKey)")
	put := strings.Index(update, "r.client.UpdateZZTokenProbe(ctx, apiResource)")
	if privateRead < 0 || put < 0 || privateRead > put {
		t.Fatal("Update must read the prior private token before PUT")
	}
	if got := strings.Count(update, "r.client.UpdateZZTokenProbe(ctx, apiResource)"); got != 1 {
		t.Fatalf("generated conflict path can execute %d PUT calls, want exactly one", got)
	}
	if strings.Count(update[:put], "GetZZTokenProbe(ctx") != 0 {
		t.Fatal("token-bearing Update must not refresh the object/token immediately before PUT")
	}
	setToken := strings.Index(update, "resp.Private.SetKey(ctx, concurrencyTokenPrivateKey")
	readback := strings.Index(update, "fetched, fetchErr := r.client.GetZZTokenProbe")
	if setToken < 0 || readback < 0 || setToken < readback {
		t.Fatal("new private token must be recorded only after successful update readback")
	}

	// SecureMesh preserves platform-authored labels with a pre-write GET. That GET
	// must never replace the token selected from private state: doing so would adopt
	// remote changes the reviewed plan did not contain.
	tmpl.FiltersDiscoveredSiteLabels = true
	decoratedDir := t.TempDir()
	if err := GenerateResourceFile(tmpl, decoratedDir); err != nil {
		t.Fatalf("GenerateResourceFile(decorated): %v", err)
	}
	decoratedBytes, err := os.ReadFile(filepath.Join(decoratedDir, "zz_token_probe_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	decoratedUpdate := string(decoratedBytes)
	decoratedUpdate = decoratedUpdate[strings.Index(decoratedUpdate, "func (r *ZZTokenProbeResource) Update("):]
	privateRead = strings.Index(decoratedUpdate, "req.Private.GetKey(ctx, concurrencyTokenPrivateKey)")
	labelRead := strings.Index(decoratedUpdate, "current, currentErr := r.client.GetZZTokenProbe")
	put = strings.Index(decoratedUpdate, "r.client.UpdateZZTokenProbe(ctx, apiResource)")
	if privateRead < 0 || labelRead < 0 || put < 0 || privateRead > labelRead || labelRead > put {
		t.Fatal("decorated Update must select its private token before the label-preservation GET and PUT")
	}
	if strings.Contains(decoratedUpdate[:put], "current.ResourceVersion") {
		t.Fatal("label-preservation GET must not replace the private concurrency token")
	}
}

func TestExcludedReplaceGeneratesForceReplacementAndNoPut(t *testing.T) {
	tmpl := &openapi.ResourceTemplate{
		Name:               "zz_enrollment_probe",
		TitleCase:          "ZZEnrollmentProbe",
		Description:        "Probe.",
		HasNamespaceInPath: true,
		APIPath:            "/api/register/namespaces/%s/probes",
		APIPathItem:        "/api/register/namespaces/%s/probes/%s",
		UpdateDisabled:     true,
		Attributes: []openapi.TerraformAttribute{
			{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string", Required: true, PlanModifier: "RequiresReplace"},
			{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string", Required: true, PlanModifier: "RequiresReplace"},
			{Name: "value", GoName: "Value", TfsdkTag: "value", Type: "string", Optional: true, PlanModifier: "RequiresReplace"},
			{Name: "id", GoName: "ID", TfsdkTag: "id", Type: "string", Computed: true},
		},
	}
	dir := t.TempDir()
	if err := GenerateResourceFile(tmpl, dir); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(filepath.Join(dir, "zz_enrollment_probe_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	source := string(content)
	if !strings.Contains(source, "Update Not Supported") {
		t.Fatal("excluded Replace does not fail closed")
	}
	if strings.Contains(source, "UpdateZZEnrollmentProbe(ctx") {
		t.Fatal("excluded Replace generated a PUT client call")
	}
}

func TestRenderMarshalOmitsComputedOnlyNestedFields(t *testing.T) {
	attrs := []openapi.TerraformAttribute{{
		Name: "interfaces", GoName: "Interfaces", TfsdkTag: "interfaces", JsonName: "interfaces",
		IsSpecField: true, IsBlock: true, NestedBlockType: "list",
		NestedAttributes: []openapi.TerraformAttribute{
			{Name: "name", GoName: "Name", TfsdkTag: "name", JsonName: "name", Type: "string", Optional: true},
			{Name: "is_management", GoName: "IsManagement", TfsdkTag: "is_management", JsonName: "is_management", Type: "bool", Computed: true},
			{Name: "is_primary", GoName: "IsPrimary", TfsdkTag: "is_primary", JsonName: "is_primary", Type: "bool", Computed: true},
		},
	}}
	got, err := RenderSpecMarshalCode(attrs, "\t", "Probe")
	if err != nil {
		t.Fatalf("RenderSpecMarshalCode: %v", err)
	}
	if !strings.Contains(got, `["name"]`) {
		t.Fatal("marshal omitted the configurable sibling")
	}
	for _, forbidden := range []string{`["is_management"]`, `["is_primary"]`} {
		if strings.Contains(got, forbidden) {
			t.Fatalf("marshal included read-only nested field %s", forbidden)
		}
	}
}
