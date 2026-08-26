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
	if len(contract.EchoOnOperations) != 1 || contract.EchoOnOperations[0] != "replace" {
		return fmt.Errorf("%s %s concurrency token %q must declare echo_on_operations=[replace]", resourceName, envelopeName, fieldName)
	}
	return nil
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
	getSchema, _, getFound, err := ResolveEnvelopeSchema(spec, resourceName, "GetResponse")
	if err != nil {
		return nil, err
	}
	replaceSchema, _, replaceFound, err := ResolveEnvelopeSchema(spec, resourceName, "ReplaceRequest")
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
		createSchema, _, found, findErr := ResolveEnvelopeSchema(spec, resourceName, createEnvelope)
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

// ValidateGeneratedConcurrencyCoverage fails generation when a mutable resource
// would emit an unconditional replace or when a read-only/create-delete resource
// unexpectedly carries a replace token contract.
func ValidateGeneratedConcurrencyCoverage(resource *openapi.ResourceTemplate, hasReplace bool) error {
	if hasReplace && !resource.HasConcurrencyToken {
		return fmt.Errorf("%s has a Replace operation without a validated concurrency token contract", resource.Name)
	}
	if !hasReplace && resource.HasConcurrencyToken {
		return fmt.Errorf("%s declares a concurrency token without a Replace operation", resource.Name)
	}
	return nil
}
