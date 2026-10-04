package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
)

func TestHTTPDNSOwnershipMetadataGeneratesNestedReplacement(t *testing.T) {
	spec := &openapi.Spec{Components: openapi.Components{Schemas: map[string]openapi.Schema{
		"http_loadbalancerProxyTypeHttp": {Type: "object", Properties: map[string]openapi.Schema{
			"dns_volterra_managed": {Type: "boolean", XFieldMutability: "immutable"},
			"port":                 {Type: "integer"},
		}},
	}}}
	field := openapi.Schema{Ref: "#/components/schemas/http_loadbalancerProxyTypeHttp"}
	http := schema.ConvertToTerraformAttribute("http", field, false, "", spec)
	var dns, port *openapi.TerraformAttribute
	for i := range http.NestedAttributes {
		child := &http.NestedAttributes[i]
		if child.TfsdkTag == "dns_volterra_managed" {
			dns = child
		}
		if child.TfsdkTag == "port" {
			port = child
		}
	}
	if dns == nil || dns.PlanModifier != "RequiresReplace" {
		t.Fatalf("DNS metadata lost: %+v", dns)
	}
	if port == nil || port.PlanModifier == "RequiresReplace" || http.PlanModifier == "RequiresReplace" {
		t.Fatal("unrelated HTTP settings frozen")
	}
	tmpl := &openapi.ResourceTemplate{Name: "dns_ownership_probe", TitleCase: "DnsOwnershipProbe", Description: "Probe.", APIPath: "/api/config/namespaces/%s/probes", APIPathItem: "/api/config/namespaces/%s/probes/%s", HasNamespaceInPath: true, Attributes: []openapi.TerraformAttribute{http}}
	dir := t.TempDir()
	if err := GenerateResourceFile(tmpl, dir); err != nil {
		t.Fatal(err)
	}
	bytes, err := os.ReadFile(filepath.Join(dir, "dns_ownership_probe_resource.go"))
	if err != nil {
		t.Fatal(err)
	}
	text := string(bytes)
	if !strings.Contains(text, "boolplanmodifier.RequiresReplace()") || !strings.Contains(text, "resource/schema/boolplanmodifier") {
		t.Fatal("nested replacement modifier or import missing")
	}
}
