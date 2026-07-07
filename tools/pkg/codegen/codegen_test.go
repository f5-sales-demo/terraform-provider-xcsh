// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
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

func TestBlockHasComputedDescendant(t *testing.T) {
	tests := []struct {
		name string
		attr openapi.TerraformAttribute
		want bool
	}{
		{
			name: "direct computed child",
			attr: openapi.TerraformAttribute{
				IsBlock: true, NestedBlockType: "list",
				NestedAttributes: []openapi.TerraformAttribute{
					{GoName: "Uid", TfsdkTag: "uid", Type: "string", Computed: true},
				},
			},
			want: true,
		},
		{
			name: "computed deep at depth >= 3",
			attr: deepComputedTree(),
			want: true,
		},
		{
			name: "no computed descendant",
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
			want: false,
		},
		{
			name: "no nested attributes",
			attr: openapi.TerraformAttribute{IsBlock: true, NestedBlockType: "list"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blockHasComputedDescendant(tt.attr); got != tt.want {
				t.Errorf("blockHasComputedDescendant() = %v, want %v", got, tt.want)
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

// A list block with NO computed descendant keeps the native slice representation.
func TestRenderNestedModelTypes_NoComputedKeepsSlice(t *testing.T) {
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
	if !strings.Contains(got, "Items []TestOuterItemsModel `tfsdk:\"items\"`") {
		t.Errorf("expected items to remain a native slice, got:\n%s", got)
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
