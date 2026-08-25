// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package schema

import (
	"fmt"
	"sort"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

type concurrencyTokenContract struct {
	JSONName string
	GoName   string
}

func concurrencyTokenJSONName(contract *concurrencyTokenContract) string {
	if contract == nil {
		return ""
	}
	return contract.JSONName
}

func concurrencyTokenGoName(contract *concurrencyTokenContract) string {
	if contract == nil {
		return ""
	}
	return contract.GoName
}

func findEnvelopeSchema(spec *openapi.Spec, resourceName, suffix string) (openapi.Schema, bool, error) {
	candidates := []string{resourceName + suffix, "schema" + resourceName + suffix, "views" + resourceName + suffix}
	for _, key := range candidates {
		if schema, ok := spec.Components.Schemas[key]; ok {
			return schema, true, nil
		}
	}

	wantedSuffix := strings.ToLower(resourceName + suffix)
	matches := make([]string, 0, 1)
	for key := range spec.Components.Schemas {
		if strings.HasSuffix(strings.ToLower(key), wantedSuffix) {
			matches = append(matches, key)
		}
	}
	sort.Strings(matches)
	if len(matches) > 1 {
		return openapi.Schema{}, false, fmt.Errorf("%s has ambiguous %s envelopes: %s", resourceName, suffix, strings.Join(matches, ", "))
	}
	if len(matches) == 1 {
		return spec.Components.Schemas[matches[0]], true, nil
	}
	return openapi.Schema{}, false, nil
}

func envelopeConcurrencyToken(schema openapi.Schema) (string, openapi.Schema, bool, error) {
	names := make([]string, 0, 1)
	for name, property := range schema.Properties {
		if property.XF5XCConcurrencyToken != nil {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	if len(names) > 1 {
		return "", openapi.Schema{}, false, fmt.Errorf("multiple x-f5xc-concurrency-token fields: %s", strings.Join(names, ", "))
	}
	if len(names) == 0 {
		return "", openapi.Schema{}, false, nil
	}
	return names[0], schema.Properties[names[0]], true, nil
}

func validateConcurrencyField(resourceName, envelopeName, fieldName string, field openapi.Schema) error {
	if field.Type != "string" {
		return fmt.Errorf("%s %s concurrency token %q must be a string, got %q", resourceName, envelopeName, fieldName, field.Type)
	}
	contract := field.XF5XCConcurrencyToken
	if contract == nil {
		return fmt.Errorf("%s %s concurrency token %q is missing its extension", resourceName, envelopeName, fieldName)
	}
	if !contract.ServerAssigned {
		return fmt.Errorf("%s %s concurrency token %q must declare server_assigned=true", resourceName, envelopeName, fieldName)
	}
	for _, operation := range contract.EchoOnOperations {
		if operation == "replace" {
			return nil
		}
	}
	return fmt.Errorf("%s %s concurrency token %q must echo on replace", resourceName, envelopeName, fieldName)
}

func normalizedOperations(operations []string) string {
	normalized := append([]string(nil), operations...)
	sort.Strings(normalized)
	return strings.Join(normalized, "\x00")
}

// ExtractConcurrencyTokenContract validates and returns a client-only optimistic
// concurrency contract. A declaration on only one envelope is an error: generating
// a client that can read but not echo the token (or vice versa) would make writes unsafe.
func ExtractConcurrencyTokenContract(spec *openapi.Spec, resourceName string) (*concurrencyTokenContract, error) {
	getSchema, getFound, err := findEnvelopeSchema(spec, resourceName, "GetResponse")
	if err != nil {
		return nil, err
	}
	replaceSchema, replaceFound, err := findEnvelopeSchema(spec, resourceName, "ReplaceRequest")
	if err != nil {
		return nil, err
	}

	getName, getField, getToken, err := envelopeConcurrencyToken(getSchema)
	if err != nil {
		return nil, fmt.Errorf("%s GetResponse: %w", resourceName, err)
	}
	replaceName, replaceField, replaceToken, err := envelopeConcurrencyToken(replaceSchema)
	if err != nil {
		return nil, fmt.Errorf("%s ReplaceRequest: %w", resourceName, err)
	}
	if !getToken && !replaceToken {
		return nil, nil
	}
	if !getFound || !replaceFound || !getToken || !replaceToken {
		return nil, fmt.Errorf("%s concurrency token must exist in both GetResponse and ReplaceRequest", resourceName)
	}
	if getName != replaceName {
		return nil, fmt.Errorf("%s concurrency token field mismatch: GetResponse=%q ReplaceRequest=%q", resourceName, getName, replaceName)
	}
	if err := validateConcurrencyField(resourceName, "GetResponse", getName, getField); err != nil {
		return nil, err
	}
	if err := validateConcurrencyField(resourceName, "ReplaceRequest", replaceName, replaceField); err != nil {
		return nil, err
	}
	if normalizedOperations(getField.XF5XCConcurrencyToken.EchoOnOperations) != normalizedOperations(replaceField.XF5XCConcurrencyToken.EchoOnOperations) {
		return nil, fmt.Errorf("%s concurrency token echo_on_operations differs between GetResponse and ReplaceRequest", resourceName)
	}

	for _, createEnvelope := range []string{"CreateRequest", "CreateSpecType"} {
		createSchema, found, findErr := findEnvelopeSchema(spec, resourceName, createEnvelope)
		if findErr != nil {
			return nil, findErr
		}
		if found {
			if name, _, present, tokenErr := envelopeConcurrencyToken(createSchema); tokenErr != nil {
				return nil, fmt.Errorf("%s %s: %w", resourceName, createEnvelope, tokenErr)
			} else if present {
				return nil, fmt.Errorf("%s %s must not declare concurrency token %q", resourceName, createEnvelope, name)
			}
		}
	}

	return &concurrencyTokenContract{JSONName: getName, GoName: naming.ToResourceTypeName(getName)}, nil
}
