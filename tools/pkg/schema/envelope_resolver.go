// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"fmt"
	"sort"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

// ResolveEnvelopeSchema resolves a resource envelope across the canonical bare,
// schema-prefixed, and views-prefixed naming families. The deterministic suffix
// fallback supports fully-qualified generated names while rejecting ambiguity.
func ResolveEnvelopeSchema(spec *openapi.Spec, resourceName, suffix string) (openapi.Schema, string, bool, error) {
	return ResolveEnvelopeSchemaFromSchemas(spec.Components.Schemas, resourceName, suffix)
}

// ResolveNamespaceProfileSchema selects only an envelope that can describe the
// provider surface being generated. Read-only API identities must not search
// CreateSpecType: a domain document may contain many related child-resource
// create envelopes whose names all end in the read-only parent's name.
func ResolveNamespaceProfileSchema(schemas map[string]openapi.Schema, resourceName string, hasCreate bool) (openapi.Schema, string, bool, error) {
	if hasCreate {
		resolved, key, found, err := ResolveEnvelopeSchemaFromSchemas(schemas, resourceName, "CreateSpecType")
		if err != nil || found {
			return resolved, key, found, err
		}
	}
	return ResolveEnvelopeSchemaFromSchemas(schemas, resourceName, "GetSpecType")
}

// ResolveEnvelopeSchemaFromSchemas is the map-based form used by audits that
// combine schemas from all released domain documents.
func ResolveEnvelopeSchemaFromSchemas(schemas map[string]openapi.Schema, resourceName, suffix string) (openapi.Schema, string, bool, error) {
	for _, key := range []string{
		resourceName + suffix,
		"schema" + resourceName + suffix,
		"views" + resourceName + suffix,
	} {
		if resolved, ok := schemas[key]; ok {
			return resolved, key, true, nil
		}
	}

	wanted := strings.ToLower(resourceName + suffix)
	qualifiedWanted := strings.ToLower(resourceName + "." + suffix)
	matches := make([]string, 0, 1)
	for key := range schemas {
		lowerKey := strings.ToLower(key)
		if strings.HasSuffix(lowerKey, wanted) || strings.HasSuffix(lowerKey, qualifiedWanted) {
			matches = append(matches, key)
		}
	}
	sort.Strings(matches)
	if len(matches) > 1 {
		return openapi.Schema{}, "", false, fmt.Errorf(
			"%s has ambiguous %s envelopes: %s", resourceName, suffix, strings.Join(matches, ", "),
		)
	}
	if len(matches) == 1 {
		return schemas[matches[0]], matches[0], true, nil
	}
	return openapi.Schema{}, "", false, nil
}
