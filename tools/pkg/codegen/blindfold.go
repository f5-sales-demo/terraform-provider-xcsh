package codegen

import (
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"strings"
)

func hasBlindfoldNodes(attrs []openapi.TerraformAttribute) bool {
	cert, key := false, false
	for _, a := range attrs {
		cert = cert || a.TfsdkTag == "certificate_url"
		key = key || a.TfsdkTag == "private_key"
		if hasBlindfoldNodes(a.NestedAttributes) {
			return true
		}
	}
	return cert && key
}
func decorateBlindfoldConstructor(r *openapi.ResourceTemplate, b []byte) []byte {
	if !hasBlindfoldNodes(r.Attributes) {
		return b
	}
	old := fmt.Sprintf("return &%sResource{}", r.TitleCase)
	replacement := fmt.Sprintf("return newBlindfoldResource(&%sResource{}, %q, %q, %t)", r.TitleCase, r.APIPath, r.APIPathItem, r.Name == "certificate")
	return []byte(strings.Replace(string(b), old, replacement, 1))
}
