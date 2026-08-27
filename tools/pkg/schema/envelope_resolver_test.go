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

func TestResolveNamespaceProfileSchemaReadOnlyIgnoresRelatedCreateEnvelopes(t *testing.T) {
	schemas := map[string]openapi.Schema{
		"schemasiteGetSpecType":           {Type: "object"},
		"viewsaws_vpc_siteCreateSpecType": {Type: "object"},
		"viewsgcp_vpc_siteCreateSpecType": {Type: "object"},
	}

	_, gotKey, found, err := ResolveNamespaceProfileSchema(schemas, "site", false)
	if err != nil || !found || gotKey != "schemasiteGetSpecType" {
		t.Fatalf("ResolveNamespaceProfileSchema() = key %q, found %v, err %v; want read-only site envelope", gotKey, found, err)
	}
}

func TestResolveNamespaceProfileSchemaMutableRejectsAmbiguousCreateEnvelopes(t *testing.T) {
	schemas := map[string]openapi.Schema{
		"one.probeCreateSpecType": {Type: "object"},
		"two.probeCreateSpecType": {Type: "object"},
		"probeGetSpecType":        {Type: "object"},
	}

	_, _, _, err := ResolveNamespaceProfileSchema(schemas, "probe", true)
	if err == nil || !strings.Contains(err.Error(), "ambiguous") {
		t.Fatalf("mutable profile resolution accepted ambiguous create envelopes: %v", err)
	}
}
