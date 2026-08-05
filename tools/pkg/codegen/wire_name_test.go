// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
)

// Codegen tests for x-f5xc-wire-name (#1323).
//
// F5's spec misspells several property names. The wire key must stay misspelled —
// verified live: a PUT with `blocked_sevice` returns HTTP 200 and round-trips,
// while `blocked_service` is silently ignored by the server (#1257). The buffer
// zone presents the CORRECTED spelling as the property name and records the wire
// key in x-f5xc-wire-name, so generated code must split the two: the Terraform
// attribute, model field and docs use the corrected name, while EVERY JSON key —
// client struct tag, request marshal AND response unmarshal — uses the wire key.
// If only one side used the wire key the field would silently stop working.

// wireProbeSpec is a synthetic CRUD spec whose annotated properties mirror the
// real shapes of the three misspelled families, at every nesting position that
// has its own emitter:
//
//   - public_advertisement          top-level scalar (wire: public_advertisment)
//   - advertised_service            top-level LIST block (wire: advertised_sevice)
//   - blocked_services              top-level SINGLE block (unannotated parent)
//   - .blocked_service              list-of-objects child (wire: blocked_sevice)
//     — the real fleetBlockedServicesListType.blocked_sevice shape
//   - ..blocked_service_name        scalar inside a LIST ELEMENT
//   - .source_ip_persistence        nested single block (wire: ..._persistance)
//   - ..persistence_timeout         scalar inside a nested single block
//   - .disable_lb_source_ip_persistence  allOf-$ref EMPTY MARKER child, the real
//     origin_pool advanced_options shape
//
// annotate=false yields the byte-identical-output control: the same tree with
// every x-f5xc-wire-name removed.
func wireProbeSpec(annotate bool) *openapi.Spec {
	wire := func(name string) string {
		if annotate {
			return name
		}
		return ""
	}
	return &openapi.Spec{
		Components: openapi.Components{
			Schemas: map[string]openapi.Schema{
				"ioschemaEmpty": {Type: "object"},
				"zz_wire_probeCreateSpecType": {
					Type:        "object",
					Description: "Synthetic wire-name probe.",
					Properties: map[string]openapi.Schema{
						"public_advertisement": {
							Type:          "boolean",
							Description:   "Advertised publicly.",
							XF5xcWireName: wire("public_advertisment"),
						},
						"advertised_service": {
							Type:          "array",
							Description:   "Advertised services.",
							XF5xcWireName: wire("advertised_sevice"),
							Items: &openapi.Schema{
								Type: "object",
								Properties: map[string]openapi.Schema{
									"port": {Type: "integer"},
								},
							},
						},
						"blocked_services": {
							Type:        "object",
							Description: "Blocked service configuration.",
							Properties: map[string]openapi.Schema{
								"blocked_service": {
									Type:          "array",
									Description:   "Blocking or denial configuration.",
									XF5xcWireName: wire("blocked_sevice"),
									Items: &openapi.Schema{
										Type: "object",
										Properties: map[string]openapi.Schema{
											"blocked_service_name": {
												Type:          "string",
												XF5xcWireName: wire("blocked_sevice_name"),
											},
											"network_type": {Type: "string"},
										},
									},
								},
								"source_ip_persistence": {
									Type:          "object",
									Description:   "Source IP persistence.",
									XF5xcWireName: wire("source_ip_persistance"),
									Properties: map[string]openapi.Schema{
										"persistence_timeout": {
											Type:          "integer",
											XF5xcWireName: wire("persistance_timeout"),
										},
									},
								},
								"disable_lb_source_ip_persistence": {
									AllOf:         []openapi.Schema{{Ref: "#/components/schemas/ioschemaEmpty"}},
									XF5xcWireName: wire("disable_lb_source_ip_persistance"),
								},
							},
						},
					},
				},
			},
		},
	}
}

// wireProbeAPIPath is the extractAPIPath stub for the synthetic probe resource.
func wireProbeAPIPath(*openapi.Spec, string) (string, string, bool) {
	return "/api/config/namespaces/%s/zz_wire_probes", "/api/config/namespaces/%s/zz_wire_probes/%s", true
}

// wireProbeTemplate extracts the probe resource template from the fixture.
func wireProbeTemplate(t *testing.T, annotate bool) *openapi.ResourceTemplate {
	t.Helper()
	tmpl, err := schema.ExtractResourceSchema(wireProbeSpec(annotate), "zz_wire_probe", wireProbeAPIPath)
	if err != nil {
		t.Fatalf("ExtractResourceSchema: %v", err)
	}
	return tmpl
}

// wireProbeRenders returns the renders that consume JsonName — the client struct
// tags, the request marshal, the response unmarshal and the computed read-back —
// concatenated in a stable order. These are the complete set of emitters that
// decide a JSON key, so byte-identity here IS the "unannotated output is
// unchanged" claim.
func wireProbeRenders(tmpl *openapi.ResourceTemplate) string {
	var sb strings.Builder
	sb.WriteString("// ---- struct fields ----\n")
	sb.WriteString(RenderSpecStructFields(tmpl.Attributes, "\t"))
	sb.WriteString("// ---- marshal (create) ----\n")
	sb.WriteString(RenderSpecMarshalCodeForCreate(tmpl.Attributes, "\t", tmpl.TitleCase))
	sb.WriteString("// ---- marshal (update) ----\n")
	sb.WriteString(RenderSpecMarshalCode(tmpl.Attributes, "\t", tmpl.TitleCase))
	sb.WriteString("// ---- unmarshal ----\n")
	sb.WriteString(RenderSpecUnmarshalCode(tmpl.Attributes, "\t", tmpl.TitleCase))
	sb.WriteString("// ---- computed read-back ----\n")
	sb.WriteString(RenderCreateComputedFieldsCode(tmpl.Attributes, "\t"))
	return sb.String()
}

// TestWireNameRoundTripsInBothDirections asserts that, for every annotated
// property, the WIRE key is the JSON key on both the request-construction and the
// read-back path, while the CORRECTED name is what Terraform and the Go model
// present. It also asserts the inverse in both directions: the corrected name is
// never used as a JSON key, and the wire key never leaks into a tfsdk tag, schema
// attribute name or Go identifier.
func TestWireNameRoundTripsInBothDirections(t *testing.T) {
	tmpl := wireProbeTemplate(t, true)

	structFields := RenderSpecStructFields(tmpl.Attributes, "\t")
	marshal := RenderSpecMarshalCodeForCreate(tmpl.Attributes, "\t", tmpl.TitleCase)
	unmarshal := RenderSpecUnmarshalCode(tmpl.Attributes, "\t", tmpl.TitleCase)
	tfSchema := RenderNestedAttributes(tmpl.Attributes, "\t") +
		RenderNestedBlocks(tmpl.Attributes, "\t") +
		RenderNestedModelTypes(tmpl.TitleCase, tmpl.Attributes) +
		RenderBlockFields(tmpl.TitleCase, tmpl.Attributes)

	cases := []struct {
		desc      string
		corrected string
		wire      string
		// topLevel properties additionally carry a client struct JSON tag.
		topLevel bool
	}{
		{desc: "top-level scalar", corrected: "public_advertisement", wire: "public_advertisment", topLevel: true},
		{desc: "top-level list block", corrected: "advertised_service", wire: "advertised_sevice", topLevel: true},
		{desc: "nested list block", corrected: "blocked_service", wire: "blocked_sevice"},
		{desc: "scalar in list element", corrected: "blocked_service_name", wire: "blocked_sevice_name"},
		{desc: "nested single block", corrected: "source_ip_persistence", wire: "source_ip_persistance"},
		{desc: "scalar in nested single block", corrected: "persistence_timeout", wire: "persistance_timeout"},
		{desc: "nested empty marker", corrected: "disable_lb_source_ip_persistence", wire: "disable_lb_source_ip_persistance"},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			wireKey := `"` + tc.wire + `"`
			correctedKey := `"` + tc.corrected + `"`

			// Direction 1: request construction marshals under the wire key.
			if !strings.Contains(marshal, wireKey) {
				t.Errorf("marshal must emit the wire key %s; got:\n%s", wireKey, marshal)
			}
			if strings.Contains(marshal, correctedKey) {
				t.Errorf("marshal must NOT emit the corrected name %s as a JSON key (the server silently ignores it); got:\n%s",
					correctedKey, marshal)
			}

			// Direction 2: read-back unmarshals from the wire key.
			if !strings.Contains(unmarshal, wireKey) {
				t.Errorf("unmarshal must read the wire key %s; got:\n%s", wireKey, unmarshal)
			}
			if strings.Contains(unmarshal, correctedKey) {
				t.Errorf("unmarshal must NOT read the corrected name %s (the API never returns it); got:\n%s",
					correctedKey, unmarshal)
			}

			// The Terraform surface keeps the corrected name, and the wire
			// misspelling never reaches a user-visible identifier.
			if !strings.Contains(tfSchema, correctedKey) {
				t.Errorf("Terraform schema/model must present the corrected name %s; got:\n%s", correctedKey, tfSchema)
			}
			if strings.Contains(tfSchema, tc.wire) {
				t.Errorf("Terraform schema/model must not leak the wire misspelling %q; got:\n%s", tc.wire, tfSchema)
			}

			if tc.topLevel {
				if !strings.Contains(structFields, `json:"`+tc.wire) {
					t.Errorf("client struct tag must use the wire key %q; got:\n%s", tc.wire, structFields)
				}
				if strings.Contains(structFields, `json:"`+tc.corrected) {
					t.Errorf("client struct tag must not use the corrected name %q; got:\n%s", tc.corrected, structFields)
				}
			}
		})
	}
}

// TestReadOnlyDataSourceHonorsWireName pins the read-only template separately
// from the CRUD renderers above. The site contract exposes a corrected property
// name while declaring the historical, misspelled API key in x-f5xc-wire-name;
// response lookup must use only that declared wire key.
func TestReadOnlyDataSourceHonorsWireName(t *testing.T) {
	tmpl := &openapi.ResourceTemplate{
		Name:               "wire_probe",
		TitleCase:          "WireProbe",
		Description:        "Synthetic read-only wire-name probe.",
		APIPath:            "/api/web/namespaces/%s/sites/%s",
		HasNamespaceInPath: true,
		Attributes: []openapi.TerraformAttribute{
			{
				Name:        "volterra_software_override",
				TfsdkTag:    "volterra_software_override",
				GoName:      "VolterraSoftwareOverride",
				JsonName:    "volterra_software_overide",
				Description: "Synthetic wire-name field.",
				Computed:    true,
			},
		},
	}

	outDir := t.TempDir()
	if err := GenerateReadOnlyDataSource(tmpl, outDir); err != nil {
		t.Fatalf("GenerateReadOnlyDataSource: %v", err)
	}
	generated, err := os.ReadFile(filepath.Join(outDir, "wire_probe_data_source.go"))
	if err != nil {
		t.Fatalf("read generated data source: %v", err)
	}
	got := string(generated)
	for _, want := range []string{
		`resource.Spec["volterra_software_overide"]`,
		`"volterra_software_overide" is the API wire key declared by x-f5xc-wire-name for Terraform attribute "volterra_software_override"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("generated read-only data source missing %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, `resource.Spec["volterra_software_override"]`) {
		t.Errorf("generated read-only data source must not look up the corrected Terraform name as an API key:\n%s", got)
	}
}

// TestWireNameAbsentOutputUnchanged is the regression guard for the 99% of
// properties that carry no annotation: the JSON-key-deciding renders must be
// BYTE-IDENTICAL to what the generator produced before x-f5xc-wire-name existed.
// The golden was captured from the pre-change generator, so any drift here means
// the new code path is no longer inert when the annotation is absent.
func TestWireNameAbsentOutputUnchanged(t *testing.T) {
	golden := filepath.Join("testdata", "wire_name_absent.golden")
	want, err := os.ReadFile(golden) // #nosec G304 -- fixed testdata path
	if err != nil {
		t.Fatalf("reading golden %s: %v", golden, err)
	}

	got := wireProbeRenders(wireProbeTemplate(t, false))
	if got != string(want) {
		t.Errorf("unannotated output drifted from the pre-change generator (%s).\n--- got ---\n%s\n--- want ---\n%s",
			golden, got, want)
	}
}

// TestWireNameResourceCompiles is the compile-level guard (#1271's lesson: the
// substring assertions above cannot catch a generator that emits code which does
// not build). It runs the full pipeline on the annotated fixture, writes the
// generated Go into the module tree and runs `go build`.
//
// The fixture uses the SYNTHETIC resource name "zz_wire_probe": no such resource
// exists in the specs, so this guard can never clobber a committed generated file
// and can never start failing once a real artifact exists.
//
// The generated files are written to their REAL destinations (internal/provider
// and internal/client) and removed on cleanup. A throwaway package under
// internal/ is not enough: a generated CRUD resource calls package-level helpers
// that live in internal/provider (filterSystemLabels,
// filterSystemManagedRrSetGroups, ...), so only compiling it as part of that
// package proves the output actually builds. Both paths are refused if a file is
// already there, so the guard can never overwrite committed output.
func TestWireNameResourceCompiles(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping go build compile check in -short mode")
	}

	tmpl := wireProbeTemplate(t, true)

	root := repoRootFromTest(t)
	provDir := filepath.Join(root, "internal", "provider")
	clientDir := filepath.Join(root, "internal", "client")
	resourceFile := filepath.Join(provDir, "zz_wire_probe_resource.go")
	clientTypes := filepath.Join(clientDir, "zz_wire_probe_types.go")
	for _, path := range []string{resourceFile, clientTypes} {
		if _, err := os.Stat(path); err == nil {
			t.Fatalf("refusing to clobber existing %s", path)
		}
	}
	t.Cleanup(func() {
		os.Remove(resourceFile)
		os.Remove(clientTypes)
	})

	if err := GenerateResourceFile(tmpl, provDir); err != nil {
		t.Fatalf("GenerateResourceFile: %v", err)
	}
	if err := GenerateClientTypes(tmpl, clientDir); err != nil {
		t.Fatalf("GenerateClientTypes: %v", err)
	}

	resBytes, err := os.ReadFile(resourceFile)
	if err != nil {
		t.Fatalf("reading generated resource: %v", err)
	}
	res := string(resBytes)
	// The generated file must carry BOTH spellings: the corrected one as the
	// tfsdk tag and the misspelled one as the wire key.
	if !strings.Contains(res, `tfsdk:"public_advertisement"`) {
		t.Errorf("generated resource missing corrected tfsdk tag public_advertisement:\n%s", res)
	}
	if !strings.Contains(res, `"public_advertisment"`) {
		t.Errorf("generated resource missing wire key public_advertisment:\n%s", res)
	}

	cmd := exec.Command("go", "build", "./internal/provider/", "./internal/client/")
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generated wire-name resource failed to compile (%v):\n%s", err, out)
	}
}
