package schema

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestExtractImmutableOneOfSelection(t *testing.T) {
	spec, apiPath := systemOnlySpec("selection_probe")
	var contract openapi.Schema
	if err := json.Unmarshal([]byte(`{"type":"object","x-f5xc-immutable-oneof-groups":{"loadbalancer_type":["http","https","https_auto_cert"]},"properties":{"http":{"type":"object","properties":{"port":{"type":"integer"}}},"https":{"type":"object","properties":{"certificate":{"type":"string"}}},"https_auto_cert":{"type":"object","properties":{"add_hsts":{"type":"boolean"}}}}}`), &contract); err != nil {
		t.Fatal(err)
	}
	spec.Components.Schemas["selection_probeCreateSpecType"] = contract
	result, err := ExtractResourceSchema(spec, "selection_probe", apiPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string][]string{"loadbalancer_type": {"http", "https", "https_auto_cert"}}
	if !reflect.DeepEqual(result.ImmutableOneOfGroups, expected) {
		t.Fatalf("groups=%v", result.ImmutableOneOfGroups)
	}
	for _, member := range expected["loadbalancer_type"] {
		attr := findAttr(result.Attributes, member)
		if attr == nil || attr.PlanModifier == "RequiresReplace" || !strings.Contains(attr.Description, "selection requires recreation") {
			t.Fatalf("incorrect selection-only contract for %s: %+v", member, attr)
		}
		if strings.Contains(attr.Description, "Default:") {
			t.Fatal("recommendation must not claim a default")
		}
	}
}

func TestExtractImmutableOneOfRejectsInvalidMembers(t *testing.T) {
	for _, members := range [][]string{{"http"}, {"http", "http"}, {"http", "missing"}} {
		spec, apiPath := systemOnlySpec("selection_probe")
		contract := spec.Components.Schemas["selection_probeCreateSpecType"]
		contract.Properties["http"] = openapi.Schema{Type: "object", Properties: map[string]openapi.Schema{"port": {Type: "integer"}}}
		contract.XF5XCImmutableOneOfGroups = map[string][]string{"type": members}
		spec.Components.Schemas["selection_probeCreateSpecType"] = contract
		if _, err := ExtractResourceSchema(spec, "selection_probe", apiPath); err == nil {
			t.Fatalf("accepted invalid members %v", members)
		}
	}
}
