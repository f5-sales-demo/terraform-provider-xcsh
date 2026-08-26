// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

func TestResolveEnvelopeSchemaNamingFamilies(t *testing.T) {
	for _, prefix := range []string{"", "schema", "views"} {
		t.Run(prefix+"family", func(t *testing.T) {
			key := prefix + "securemesh_site_v2CreateSpecType"
			spec := &openapi.Spec{Components: openapi.Components{Schemas: map[string]openapi.Schema{
				key: {Type: "object"},
			}}}
			_, gotKey, found, err := ResolveEnvelopeSchema(spec, "securemesh_site_v2", "CreateSpecType")
			if err != nil || !found || gotKey != key {
				t.Fatalf("ResolveEnvelopeSchema() = key %q, found %v, err %v; want %q", gotKey, found, err, key)
			}
		})
	}
}

func TestResolveEnvelopeSchemaRejectsAmbiguousFallback(t *testing.T) {
	spec := &openapi.Spec{Components: openapi.Components{Schemas: map[string]openapi.Schema{
		"one.securemesh_site_v2GetResponse": {Type: "object"},
		"two.securemesh_site_v2GetResponse": {Type: "object"},
	}}}
	_, _, _, err := ResolveEnvelopeSchema(spec, "securemesh_site_v2", "GetResponse")
	if err == nil || !strings.Contains(err.Error(), "ambiguous") {
		t.Fatalf("expected ambiguous-envelope error, got %v", err)
	}
}

func TestResolveEnvelopeSchemaFullyQualifiedFallback(t *testing.T) {
	key := "ves.io.schema.securemesh_site_v2.CreateSpecType"
	spec := &openapi.Spec{Components: openapi.Components{Schemas: map[string]openapi.Schema{
		key: {Type: "object"},
	}}}
	_, gotKey, found, err := ResolveEnvelopeSchema(spec, "securemesh_site_v2", "CreateSpecType")
	if err != nil || !found || gotKey != key {
		t.Fatalf("ResolveEnvelopeSchema() = key %q, found %v, err %v; want %q", gotKey, found, err, key)
	}
}
