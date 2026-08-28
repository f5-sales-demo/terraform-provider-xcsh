// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
)

// ErrResourceNotManageable means Go's path-shape discovery found an API surface
// that does not provide both a collection and an item lifecycle. It is a normal
// provider classification result, not permission to infer the missing path.
var ErrResourceNotManageable = errors.New("API identity is not Terraform-manageable")

var (
	apiIdentityPattern   = regexp.MustCompile(`^ves\.io\.schema\.[a-z0-9_]+(?:\.[a-z0-9_]+)*$`)
	catalogVersionRegex  = regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`)
	surfacePattern       = regexp.MustCompile(`^[a-z][a-z0-9_-]*$`)
	terraformNamePattern = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
	schemaNamePattern    = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_.]*$`)
	placeholderPattern   = regexp.MustCompile(`^\{[A-Za-z_][A-Za-z0-9_.-]*\}$`)
)

// CatalogOperation is one exact operation fact published by api-specs-enriched.
// These fields are consumed as data; the provider does not reconstruct them from
// Terraform names or RPC naming conventions.
type CatalogOperation struct {
	Method         string
	Path           string
	OperationID    string
	Surface        string
	Role           string
	TerraformName  string
	RequestSchema  string
	ResponseSchema string
}

// ResolvedResponseOperation is an exact source-owned non-CRUD operation that
// generates a Terraform query, collection, issuance resource, or action.
type ResolvedResponseOperation struct {
	Name           string
	Role           string
	Method         string
	Path           string
	OperationID    string
	RequestSchema  string
	ResponseSchema string
}

// APIOperationIdentity groups the exact operations published for one
// ves.io.schema API identity.
type APIOperationIdentity struct {
	APIIdentity string
	Operations  []CatalogOperation
}

// APIExclusion records one API identity deliberately absent from APIOperations.
type APIExclusion struct {
	APIIdentity    string
	Classification string
	Reason         string
}

// OperationCatalog is the enriched operation contract published in
// api-catalog.json. Lifecycle classification remains provider-owned, while
// explicitly enriched response operations carry source-owned Terraform names
// and roles.
type OperationCatalog struct {
	Version       string
	APIOperations []APIOperationIdentity
	APIExclusions []APIExclusion
	byIdentity    map[string]int
}

// ResolvedResourceOperations contains the exact lifecycle facts selected by Go's
// provider-side resource classification.
type ResolvedResourceOperations struct {
	APIIdentity    string
	CollectionPath string
	ItemPath       string
	HasNamespace   bool
	HasCreate      bool
	Create         *CatalogOperation
	List           *CatalogOperation
	Get            *CatalogOperation
	Replace        *CatalogOperation
	Delete         *CatalogOperation
}

type operationCatalogDocument struct {
	Service       json.RawMessage    `json:"service"`
	DisplayName   json.RawMessage    `json:"displayName"`
	Version       *string            `json:"version"`
	SpecSource    json.RawMessage    `json:"specSource"`
	Auth          json.RawMessage    `json:"auth"`
	Defaults      json.RawMessage    `json:"defaults"`
	APIOperations *[]json.RawMessage `json:"apiOperations"`
	APIExclusions *[]json.RawMessage `json:"apiExclusions"`
	Categories    json.RawMessage    `json:"categories"`
}

type apiOperationIdentityWire struct {
	APIIdentity *string            `json:"apiIdentity"`
	Operations  *[]json.RawMessage `json:"operations"`
}

type catalogOperationWire struct {
	Method         *string         `json:"method"`
	Path           *string         `json:"path"`
	OperationID    *string         `json:"operationId"`
	Surface        *string         `json:"surface"`
	Role           json.RawMessage `json:"role"`
	RequestSchema  json.RawMessage `json:"requestSchema"`
	ResponseSchema json.RawMessage `json:"responseSchema"`
	TerraformName  json.RawMessage `json:"terraformName"`
}

type apiExclusionWire struct {
	APIIdentity    *string `json:"apiIdentity"`
	Classification *string `json:"classification"`
	Reason         *string `json:"reason"`
}

// ParseOperationCatalogFromDir loads the canonical api-catalog.json from a
// downloaded enriched-spec bundle.
func ParseOperationCatalogFromDir(specDir string) (*OperationCatalog, error) {
	path := filepath.Join(specDir, "api-catalog.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read operation catalog %s: %w", path, err)
	}
	catalog, err := ParseOperationCatalog(data)
	if err != nil {
		return nil, fmt.Errorf("parse operation catalog %s: %w", path, err)
	}
	return catalog, nil
}

// ParseOperationCatalog parses and validates the narrowed apiOperations /
// apiExclusions contract. It rejects ambiguity instead of retaining a path-
// inference fallback.
func ParseOperationCatalog(data []byte) (*OperationCatalog, error) {
	if err := rejectDuplicateJSONFields(data); err != nil {
		return nil, fmt.Errorf("decode catalog: %w", err)
	}

	var document operationCatalogDocument
	if err := decodeStrictJSONObject(data, &document, "catalog"); err != nil {
		return nil, err
	}
	if document.Version == nil || !catalogVersionRegex.MatchString(*document.Version) {
		return nil, fmt.Errorf("catalog version is required and must match %s", catalogVersionRegex)
	}
	for _, field := range []struct {
		name     string
		value    json.RawMessage
		jsonType byte
	}{
		{name: "service", value: document.Service, jsonType: '"'},
		{name: "displayName", value: document.DisplayName, jsonType: '"'},
		{name: "specSource", value: document.SpecSource, jsonType: '"'},
		{name: "auth", value: document.Auth, jsonType: '{'},
		{name: "defaults", value: document.Defaults, jsonType: '{'},
		{name: "categories", value: document.Categories, jsonType: '['},
	} {
		if err := validateRequiredCatalogTopLevel(field.name, field.value, field.jsonType); err != nil {
			return nil, err
		}
	}
	if document.APIOperations == nil || len(*document.APIOperations) == 0 {
		return nil, fmt.Errorf("catalog apiOperations must be present and non-empty")
	}
	if document.APIExclusions == nil {
		return nil, fmt.Errorf("catalog apiExclusions must be present")
	}

	catalog := &OperationCatalog{
		Version:       *document.Version,
		APIOperations: make([]APIOperationIdentity, 0, len(*document.APIOperations)),
		APIExclusions: make([]APIExclusion, 0, len(*document.APIExclusions)),
		byIdentity:    make(map[string]int, len(*document.APIOperations)),
	}
	seenMethodPath := make(map[string]string)
	seenOperationID := make(map[string]string)
	seenTerraformName := make(map[string]string)
	for index, raw := range *document.APIOperations {
		identity, err := parseAPIOperationIdentity(raw)
		if err != nil {
			return nil, fmt.Errorf("apiOperations[%d]: %w", index, err)
		}
		if _, exists := catalog.byIdentity[identity.APIIdentity]; exists {
			return nil, fmt.Errorf("apiOperations[%d]: duplicate apiIdentity %q", index, identity.APIIdentity)
		}
		for operationIndex, operation := range identity.Operations {
			methodPath := operation.Method + " " + operation.Path
			if owner, exists := seenMethodPath[methodPath]; exists {
				return nil, fmt.Errorf("apiOperations[%d].operations[%d]: duplicate method/path %q (already owned by %s)", index, operationIndex, methodPath, owner)
			}
			seenMethodPath[methodPath] = identity.APIIdentity
			if owner, exists := seenOperationID[operation.OperationID]; exists {
				return nil, fmt.Errorf("apiOperations[%d].operations[%d]: duplicate operationId %q (already owned by %s)", index, operationIndex, operation.OperationID, owner)
			}
			seenOperationID[operation.OperationID] = identity.APIIdentity
			if operation.TerraformName != "" {
				if owner, exists := seenTerraformName[operation.TerraformName]; exists {
					return nil, fmt.Errorf("apiOperations[%d].operations[%d]: duplicate terraformName %q (already owned by %s)", index, operationIndex, operation.TerraformName, owner)
				}
				seenTerraformName[operation.TerraformName] = operation.OperationID
			}
		}
		catalog.byIdentity[identity.APIIdentity] = len(catalog.APIOperations)
		catalog.APIOperations = append(catalog.APIOperations, identity)
	}

	seenExclusions := make(map[string]bool, len(*document.APIExclusions))
	for index, raw := range *document.APIExclusions {
		exclusion, err := parseAPIExclusion(raw)
		if err != nil {
			return nil, fmt.Errorf("apiExclusions[%d]: %w", index, err)
		}
		if _, exists := catalog.byIdentity[exclusion.APIIdentity]; exists {
			return nil, fmt.Errorf("apiIdentity %q is present in both apiOperations and apiExclusions", exclusion.APIIdentity)
		}
		if seenExclusions[exclusion.APIIdentity] {
			return nil, fmt.Errorf("apiExclusions[%d]: duplicate apiIdentity %q", index, exclusion.APIIdentity)
		}
		seenExclusions[exclusion.APIIdentity] = true
		catalog.APIExclusions = append(catalog.APIExclusions, exclusion)
	}

	return catalog, nil
}

// Identity returns the exact operation group for apiIdentity.
func (catalog *OperationCatalog) Identity(apiIdentity string) (APIOperationIdentity, bool) {
	if catalog == nil {
		return APIOperationIdentity{}, false
	}
	index, ok := catalog.byIdentity[apiIdentity]
	if !ok {
		return APIOperationIdentity{}, false
	}
	return catalog.APIOperations[index], true
}

// ResponseOperationsForSpec returns the catalog response operations that are
// present in one domain spec. Public Terraform names come only from the catalog.
func (catalog *OperationCatalog) ResponseOperationsForSpec(spec *Spec) ([]ResolvedResponseOperation, error) {
	if catalog == nil {
		return nil, fmt.Errorf("operation catalog is required")
	}
	var results []ResolvedResponseOperation
	for _, identity := range catalog.APIOperations {
		for _, operation := range identity.Operations {
			if operation.Role == "" {
				continue
			}
			present, err := catalogOperationInSpec(spec, operation)
			if err != nil {
				return nil, fmt.Errorf("%s operation %s: %w", operation.Role, operation.OperationID, err)
			}
			if !present {
				continue
			}
			if _, ok := spec.Components.Schemas[operation.ResponseSchema]; !ok {
				return nil, fmt.Errorf("%s operation %s response schema %q is absent", operation.Role, operation.OperationID, operation.ResponseSchema)
			}
			if operation.RequestSchema != "" {
				if _, ok := spec.Components.Schemas[operation.RequestSchema]; !ok {
					return nil, fmt.Errorf("%s operation %s request schema %q is absent", operation.Role, operation.OperationID, operation.RequestSchema)
				}
			}
			results = append(results, ResolvedResponseOperation{
				Name:           operation.TerraformName,
				Role:           operation.Role,
				Method:         operation.Method,
				Path:           operation.Path,
				OperationID:    operation.OperationID,
				RequestSchema:  operation.RequestSchema,
				ResponseSchema: operation.ResponseSchema,
			})
		}
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Name < results[j].Name })
	return results, nil
}

// HasResourceIdentity reports whether Go's resource name maps to at least one
// published API identity. A false result means path-shape discovery found a
// non-resource operation (for example a bulk action on a plural-looking path),
// not that the generator may invent lifecycle paths for it.
func (catalog *OperationCatalog) HasResourceIdentity(resourceName string) bool {
	if catalog == nil {
		return false
	}
	for _, identity := range catalog.APIOperations {
		parts := strings.Split(identity.APIIdentity, ".")
		if parts[len(parts)-1] == resourceName {
			return true
		}
	}
	return false
}

// ValidateAgainstSpec proves that every published operation is present in the
// canonical OpenAPI document with the same method, path, operationId, and
// request schema. Provider manageability may ignore an API identity, but the
// narrowed upstream contract itself may not drift unnoticed.
func (catalog *OperationCatalog) ValidateAgainstSpec(spec *Spec) error {
	return catalog.validateAgainstSpecs([]*Spec{spec})
}

// ValidateAgainstSpecFiles validates against the domain-organized generation
// inputs. An exact match in any domain satisfies the operation; this is required
// for the one deliberate path collision recorded by apiExclusions, where two
// domains contain different owners for the same method/path.
func (catalog *OperationCatalog) ValidateAgainstSpecFiles(paths []string) error {
	if len(paths) == 0 {
		return fmt.Errorf("at least one OpenAPI domain spec is required")
	}
	specs := make([]*Spec, 0, len(paths))
	for _, path := range paths {
		spec, err := ParseFile(path)
		if err != nil {
			return fmt.Errorf("parse OpenAPI domain %s: %w", path, err)
		}
		specs = append(specs, spec)
	}
	return catalog.validateAgainstSpecs(specs)
}

func (catalog *OperationCatalog) validateAgainstSpecs(specs []*Spec) error {
	if catalog == nil {
		return fmt.Errorf("operation catalog is required")
	}
	for _, identity := range catalog.APIOperations {
		for _, operation := range identity.Operations {
			matched := false
			var mismatch error
			for _, spec := range specs {
				present, err := catalogOperationInSpec(spec, operation)
				if err != nil {
					mismatch = err
					continue
				}
				if present {
					matched = true
					break
				}
			}
			if matched {
				continue
			}
			if mismatch != nil {
				return fmt.Errorf("catalog operation %s: %w", operation.OperationID, mismatch)
			}
			if !matched {
				return fmt.Errorf("catalog operation %s is absent from OpenAPI at %s %s", operation.OperationID, operation.Method, operation.Path)
			}
		}
	}
	return nil
}

// ResolveResource selects exact lifecycle operations for a provider resource.
// API identity selection and lifecycle classification stay in Go; every emitted
// method/path/operationId must nonetheless match both the catalog and OpenAPI.
func (catalog *OperationCatalog) ResolveResource(spec *Spec, resourceName string) (*ResolvedResourceOperations, error) {
	if catalog == nil {
		return nil, fmt.Errorf("operation catalog is required")
	}
	var candidates []APIOperationIdentity
	for _, identity := range catalog.APIOperations {
		parts := strings.Split(identity.APIIdentity, ".")
		if parts[len(parts)-1] == resourceName {
			candidates = append(candidates, identity)
		}
	}
	if len(candidates) == 0 {
		return nil, fmt.Errorf("no apiOperations identity maps to resource %q", resourceName)
	}

	type matchedCandidate struct {
		identity APIOperationIdentity
		roles    map[string][]CatalogOperation
	}
	var matched []matchedCandidate
	anyCandidateOperationInSpec := false
	for _, candidate := range candidates {
		roles := make(map[string][]CatalogOperation)
		for _, operation := range candidate.Operations {
			present, err := catalogOperationInSpec(spec, operation)
			if err != nil {
				return nil, fmt.Errorf("resource %q operation %s: %w", resourceName, operation.OperationID, err)
			}
			if !present {
				continue
			}
			anyCandidateOperationInSpec = true
			role := catalogOperationRole(operation.OperationID)
			if expectedLifecycleMethod(role) == operation.Method {
				roles[role] = append(roles[role], operation)
			}
		}
		if len(roles) > 0 {
			matched = append(matched, matchedCandidate{identity: candidate, roles: roles})
		}
	}
	if len(matched) == 0 {
		reason := "no matching catalog operation is present in this OpenAPI domain"
		if anyCandidateOperationInSpec {
			reason = "its exact operations do not form a resource lifecycle"
		}
		return nil, fmt.Errorf("%w: resource %q: %s", ErrResourceNotManageable, resourceName, reason)
	}

	// Prefer the canonical root identity when a domain also contains an
	// operate/graph API with the same final segment (for example bgp or site).
	rootIdentity := "ves.io.schema." + resourceName
	selected := -1
	for index, candidate := range matched {
		if candidate.identity.APIIdentity == rootIdentity {
			selected = index
			break
		}
	}
	if selected == -1 {
		if len(matched) != 1 {
			identities := make([]string, 0, len(matched))
			for _, candidate := range matched {
				identities = append(identities, candidate.identity.APIIdentity)
			}
			sort.Strings(identities)
			return nil, fmt.Errorf("resource %q maps ambiguously to apiOperations identities: %s", resourceName, strings.Join(identities, ", "))
		}
		selected = 0
	}

	chosen := matched[selected]
	resolved := &ResolvedResourceOperations{APIIdentity: chosen.identity.APIIdentity}
	for _, binding := range []struct {
		role   string
		target **CatalogOperation
	}{
		{role: "Create", target: &resolved.Create},
		{role: "List", target: &resolved.List},
		{role: "Get", target: &resolved.Get},
		{role: "Replace", target: &resolved.Replace},
		{role: "Delete", target: &resolved.Delete},
	} {
		operations := chosen.roles[binding.role]
		if len(operations) > 1 {
			return nil, fmt.Errorf("resource %q apiIdentity %s has %d exact %s operations", resourceName, chosen.identity.APIIdentity, len(operations), binding.role)
		}
		if len(operations) == 1 {
			operation := operations[0]
			*binding.target = &operation
		}
	}
	if err := requireConsistentLifecyclePaths(resourceName, "collection", resolved.Create, resolved.List); err != nil {
		return nil, err
	}
	if err := requireConsistentLifecyclePaths(resourceName, "item", resolved.Get, resolved.Replace, resolved.Delete); err != nil {
		return nil, err
	}

	collection := resolved.Create
	if collection == nil {
		collection = resolved.List
	}
	item := resolved.Get
	if item == nil {
		item = resolved.Replace
	}
	if item == nil {
		item = resolved.Delete
	}
	if collection == nil || item == nil {
		return nil, fmt.Errorf("%w: resource %q apiIdentity %s does not provide both collection and item lifecycle operations", ErrResourceNotManageable, resourceName, chosen.identity.APIIdentity)
	}
	if hasNamePlaceholder(collection.Path) {
		return nil, fmt.Errorf("%w: resource %q collection operation path %q contains an item placeholder", ErrResourceNotManageable, resourceName, collection.Path)
	}
	if !hasNamePlaceholder(item.Path) {
		return nil, fmt.Errorf("%w: resource %q item operation path %q has no name placeholder", ErrResourceNotManageable, resourceName, item.Path)
	}

	resolved.CollectionPath = formatCatalogPath(collection.Path)
	resolved.ItemPath = formatCatalogPath(item.Path)
	resolved.HasNamespace = hasNamespacePlaceholder(collection.Path) || hasNamespacePlaceholder(item.Path)
	resolved.HasCreate = resolved.Create != nil
	return resolved, nil
}

// ActionsForSpec discovers schema-declared Terraform actions while taking both
// the action POST and sibling object GET paths exactly from apiOperations.
func (catalog *OperationCatalog) ActionsForSpec(spec *Spec) ([]ResourcePath, error) {
	if catalog == nil {
		return nil, fmt.Errorf("operation catalog is required")
	}
	var results []ResourcePath
	seen := make(map[string]bool)
	for _, identity := range catalog.APIOperations {
		for _, operation := range identity.Operations {
			if operation.Method != "POST" || operation.RequestSchema == "" {
				continue
			}
			requestSchema, ok := spec.Components.Schemas[operation.RequestSchema]
			if !ok || requestSchema.XF5xcAction == "" {
				continue
			}
			present, err := catalogOperationInSpec(spec, operation)
			if err != nil {
				return nil, fmt.Errorf("action operation %s: %w", operation.OperationID, err)
			}
			if !present {
				continue
			}

			resourceName := naming.ToSnakeCase(strings.TrimSuffix(operation.RequestSchema, "Req"))
			if seen[resourceName] {
				return nil, fmt.Errorf("action resource %q has more than one exact catalog POST", resourceName)
			}
			var read *CatalogOperation
			for _, candidate := range identity.Operations {
				if candidate.Method != "GET" || catalogOperationRole(candidate.OperationID) != "Get" || !hasNamePlaceholder(candidate.Path) {
					continue
				}
				readPresent, readErr := catalogOperationInSpec(spec, candidate)
				if readErr != nil {
					return nil, fmt.Errorf("action %q read operation %s: %w", resourceName, candidate.OperationID, readErr)
				}
				if !readPresent {
					continue
				}
				if read != nil {
					return nil, fmt.Errorf("action resource %q has more than one exact catalog object GET", resourceName)
				}
				copy := candidate
				read = &copy
			}
			if read == nil {
				return nil, fmt.Errorf("action resource %q has no exact sibling object GET in apiOperations", resourceName)
			}

			seen[resourceName] = true
			results = append(results, ResourcePath{
				ResourceName:   resourceName,
				SchemaName:     operation.RequestSchema,
				ActionValue:    requestSchema.XF5xcAction,
				ActionPath:     formatCatalogPath(operation.Path),
				ReadObjectPath: formatCatalogPath(read.Path),
			})
		}
	}
	sort.Slice(results, func(i, j int) bool { return results[i].ResourceName < results[j].ResourceName })
	return results, nil
}

func parseAPIOperationIdentity(raw json.RawMessage) (APIOperationIdentity, error) {
	var wire apiOperationIdentityWire
	if err := decodeStrictJSONObject(raw, &wire, "api operation identity"); err != nil {
		return APIOperationIdentity{}, err
	}
	if wire.APIIdentity == nil || !apiIdentityPattern.MatchString(*wire.APIIdentity) {
		return APIOperationIdentity{}, fmt.Errorf("apiIdentity is required and must match %s", apiIdentityPattern)
	}
	if wire.Operations == nil || len(*wire.Operations) == 0 {
		return APIOperationIdentity{}, fmt.Errorf("operations must be present and non-empty")
	}
	identity := APIOperationIdentity{APIIdentity: *wire.APIIdentity, Operations: make([]CatalogOperation, 0, len(*wire.Operations))}
	for index, operationRaw := range *wire.Operations {
		operation, err := parseCatalogOperation(operationRaw, identity.APIIdentity)
		if err != nil {
			return APIOperationIdentity{}, fmt.Errorf("operations[%d]: %w", index, err)
		}
		identity.Operations = append(identity.Operations, operation)
	}
	return identity, nil
}

func parseCatalogOperation(raw json.RawMessage, apiIdentity string) (CatalogOperation, error) {
	var wire catalogOperationWire
	if err := decodeStrictJSONObject(raw, &wire, "operation"); err != nil {
		return CatalogOperation{}, err
	}
	if wire.Method == nil || !map[string]bool{"DELETE": true, "GET": true, "PATCH": true, "POST": true, "PUT": true}[*wire.Method] {
		return CatalogOperation{}, fmt.Errorf("method is required and must be DELETE, GET, PATCH, POST, or PUT")
	}
	if wire.Path == nil {
		return CatalogOperation{}, fmt.Errorf("path is required")
	}
	if err := validateCatalogOperationPath(*wire.Path); err != nil {
		return CatalogOperation{}, err
	}
	if wire.OperationID == nil || !strings.HasPrefix(*wire.OperationID, apiIdentity+".") || len(*wire.OperationID) == len(apiIdentity)+1 {
		return CatalogOperation{}, fmt.Errorf("operationId is required and must begin with apiIdentity %q", apiIdentity)
	}
	if wire.Surface == nil || !surfacePattern.MatchString(*wire.Surface) {
		return CatalogOperation{}, fmt.Errorf("surface is required and must match %s", surfacePattern)
	}
	requestSchema := ""
	if len(wire.RequestSchema) > 0 {
		if bytes.Equal(bytes.TrimSpace(wire.RequestSchema), []byte("null")) {
			return CatalogOperation{}, fmt.Errorf("requestSchema must be absent or a non-empty schema name")
		}
		if err := json.Unmarshal(wire.RequestSchema, &requestSchema); err != nil || !schemaNamePattern.MatchString(requestSchema) {
			return CatalogOperation{}, fmt.Errorf("requestSchema must be absent or a non-empty schema name")
		}
	}
	responseSchema := ""
	if len(wire.ResponseSchema) > 0 {
		if bytes.Equal(bytes.TrimSpace(wire.ResponseSchema), []byte("null")) {
			return CatalogOperation{}, fmt.Errorf("responseSchema must be absent or a non-empty schema name")
		}
		if err := json.Unmarshal(wire.ResponseSchema, &responseSchema); err != nil || !schemaNamePattern.MatchString(responseSchema) {
			return CatalogOperation{}, fmt.Errorf("responseSchema must be absent or a non-empty schema name")
		}
	}
	role := ""
	if len(wire.Role) > 0 {
		if bytes.Equal(bytes.TrimSpace(wire.Role), []byte("null")) {
			return CatalogOperation{}, fmt.Errorf("role must be absent, action, collection, issuance, or query")
		}
		if err := json.Unmarshal(wire.Role, &role); err != nil || !map[string]bool{"action": true, "collection": true, "issuance": true, "query": true}[role] {
			return CatalogOperation{}, fmt.Errorf("role must be absent, action, collection, issuance, or query")
		}
	}
	terraformName := ""
	if len(wire.TerraformName) > 0 {
		if bytes.Equal(bytes.TrimSpace(wire.TerraformName), []byte("null")) || json.Unmarshal(wire.TerraformName, &terraformName) != nil || !terraformNamePattern.MatchString(terraformName) {
			return CatalogOperation{}, fmt.Errorf("terraformName must be absent or match %s", terraformNamePattern)
		}
	}
	if role != "" {
		if terraformName == "" {
			return CatalogOperation{}, fmt.Errorf("terraformName is required for %s operations", role)
		}
		if responseSchema == "" {
			return CatalogOperation{}, fmt.Errorf("responseSchema is required for %s operations", role)
		}
		if *wire.Method != "GET" && *wire.Method != "POST" {
			return CatalogOperation{}, fmt.Errorf("%s operation method must be GET or POST", role)
		}
		if role == "action" && *wire.Method != "POST" {
			return CatalogOperation{}, fmt.Errorf("action operation method must be POST")
		}
		if *wire.Method == "POST" && requestSchema == "" {
			return CatalogOperation{}, fmt.Errorf("POST %s operation requestSchema is required", role)
		}
		if *wire.Method == "GET" && requestSchema != "" {
			return CatalogOperation{}, fmt.Errorf("GET %s operation requestSchema must be absent", role)
		}
	} else if terraformName != "" {
		return CatalogOperation{}, fmt.Errorf("terraformName requires a response-operation role")
	}
	return CatalogOperation{
		Method:         *wire.Method,
		Path:           *wire.Path,
		OperationID:    *wire.OperationID,
		Surface:        *wire.Surface,
		Role:           role,
		TerraformName:  terraformName,
		RequestSchema:  requestSchema,
		ResponseSchema: responseSchema,
	}, nil
}

func parseAPIExclusion(raw json.RawMessage) (APIExclusion, error) {
	var wire apiExclusionWire
	if err := decodeStrictJSONObject(raw, &wire, "API exclusion"); err != nil {
		return APIExclusion{}, err
	}
	if wire.APIIdentity == nil || !apiIdentityPattern.MatchString(*wire.APIIdentity) {
		return APIExclusion{}, fmt.Errorf("apiIdentity is required and must match %s", apiIdentityPattern)
	}
	if wire.Classification == nil || !surfacePattern.MatchString(*wire.Classification) {
		return APIExclusion{}, fmt.Errorf("classification is required and must match %s", surfacePattern)
	}
	if wire.Reason == nil || *wire.Reason == "" || strings.TrimSpace(*wire.Reason) != *wire.Reason {
		return APIExclusion{}, fmt.Errorf("reason is required and must not have surrounding whitespace")
	}
	return APIExclusion{APIIdentity: *wire.APIIdentity, Classification: *wire.Classification, Reason: *wire.Reason}, nil
}

func catalogOperationInSpec(spec *Spec, operation CatalogOperation) (bool, error) {
	pathValue, ok := spec.Paths[operation.Path]
	if !ok {
		return false, nil
	}
	pathItem, ok := pathValue.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("OpenAPI path %q is not an object", operation.Path)
	}
	methodValue, ok := pathItem[strings.ToLower(operation.Method)]
	if !ok {
		return false, nil
	}
	method, ok := methodValue.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("OpenAPI %s %s is not an object", operation.Method, operation.Path)
	}
	operationID, _ := method["operationId"].(string)
	if operationID != operation.OperationID {
		return false, fmt.Errorf("OpenAPI operationId = %q, catalog operationId = %q", operationID, operation.OperationID)
	}
	if operation.RequestSchema != "" {
		ref := requestBodySchemaRef(method)
		if GetRefName(ref) != operation.RequestSchema {
			return false, fmt.Errorf("OpenAPI request schema = %q, catalog requestSchema = %q", GetRefName(ref), operation.RequestSchema)
		}
	}
	if operation.ResponseSchema != "" {
		ref, err := responseBodySchemaRef(method)
		if err != nil {
			return false, err
		}
		if GetRefName(ref) != operation.ResponseSchema {
			return false, fmt.Errorf("OpenAPI response schema = %q, catalog response schema = %q", GetRefName(ref), operation.ResponseSchema)
		}
	}
	return true, nil
}

func responseBodySchemaRef(operation map[string]interface{}) (string, error) {
	responses, ok := operation["responses"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("OpenAPI responses are absent or not an object")
	}
	refs := make(map[string]bool)
	for status, raw := range responses {
		if len(status) != 3 || status[0] != '2' {
			continue
		}
		response, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		content, _ := response["content"].(map[string]interface{})
		jsonContent, _ := content["application/json"].(map[string]interface{})
		schema, _ := jsonContent["schema"].(map[string]interface{})
		ref, _ := schema["$ref"].(string)
		if ref != "" {
			refs[ref] = true
		}
	}
	if len(refs) == 0 {
		return "", fmt.Errorf("OpenAPI has no successful JSON response schema")
	}
	if len(refs) > 1 {
		return "", fmt.Errorf("OpenAPI has ambiguous successful JSON response schemas")
	}
	for ref := range refs {
		return ref, nil
	}
	return "", fmt.Errorf("OpenAPI has no successful JSON response schema")
}

func catalogOperationRole(operationID string) string {
	if index := strings.LastIndex(operationID, "."); index >= 0 {
		return operationID[index+1:]
	}
	return operationID
}

func expectedLifecycleMethod(role string) string {
	return map[string]string{
		"Create":  "POST",
		"List":    "GET",
		"Get":     "GET",
		"Replace": "PUT",
		"Delete":  "DELETE",
	}[role]
}

func requireConsistentLifecyclePaths(resourceName, kind string, operations ...*CatalogOperation) error {
	want := ""
	for _, operation := range operations {
		if operation == nil {
			continue
		}
		path := formatCatalogPath(operation.Path)
		if want == "" {
			want = path
			continue
		}
		if path != want {
			return fmt.Errorf("resource %q has inconsistent exact %s paths %q and %q", resourceName, kind, want, path)
		}
	}
	return nil
}

func hasNamespacePlaceholder(path string) bool {
	return strings.Contains(path, "{namespace}") || strings.Contains(path, "{metadata.namespace}")
}

func hasNamePlaceholder(path string) bool {
	return strings.Contains(path, "{name}") || strings.Contains(path, "{metadata.name}")
}

func formatCatalogPath(path string) string {
	return strings.NewReplacer(
		"{metadata.namespace}", "%s",
		"{namespace}", "%s",
		"{metadata.name}", "%s",
		"{name}", "%s",
	).Replace(path)
}

func validateCatalogOperationPath(path string) error {
	if path == "" || strings.TrimSpace(path) != path || !strings.HasPrefix(path, "/") || strings.HasPrefix(path, "//") {
		return fmt.Errorf("path must be an absolute path without surrounding whitespace")
	}
	if strings.ContainsAny(path, "?#\\") {
		return fmt.Errorf("path must not contain a query, fragment, or backslash")
	}
	for _, character := range path {
		if unicode.IsSpace(character) || unicode.IsControl(character) {
			return fmt.Errorf("path must not contain whitespace or control characters")
		}
	}
	for _, segment := range strings.Split(path[1:], "/") {
		if segment == "" || segment == "." || segment == ".." {
			return fmt.Errorf("path contains an empty or dot segment")
		}
		if strings.ContainsAny(segment, "{}") && !placeholderPattern.MatchString(segment) {
			return fmt.Errorf("path contains a malformed placeholder segment %q", segment)
		}
	}
	return nil
}

func validateRequiredCatalogTopLevel(name string, raw json.RawMessage, jsonType byte) error {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 || trimmed[0] != jsonType {
		return fmt.Errorf("catalog %s is required and has the wrong JSON type", name)
	}
	if jsonType == '"' {
		var value string
		if err := json.Unmarshal(trimmed, &value); err != nil || value == "" || strings.TrimSpace(value) != value {
			return fmt.Errorf("catalog %s must be a non-empty string without surrounding whitespace", name)
		}
	}
	return nil
}

func decodeStrictJSONObject(raw []byte, target interface{}, field string) error {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 || trimmed[0] != '{' {
		return fmt.Errorf("%s must be an object", field)
	}
	decoder := json.NewDecoder(bytes.NewReader(trimmed))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("decode %s: %w", field, err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return fmt.Errorf("decode %s: trailing JSON value", field)
	}
	return nil
}

func rejectDuplicateJSONFields(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := walkJSONValue(decoder, "$"); err != nil {
		return err
	}
	if _, err := decoder.Token(); err != io.EOF {
		return fmt.Errorf("trailing JSON value")
	}
	return nil
}

func walkJSONValue(decoder *json.Decoder, path string) error {
	token, err := decoder.Token()
	if err != nil {
		return err
	}
	delimiter, ok := token.(json.Delim)
	if !ok {
		return nil
	}
	switch delimiter {
	case '{':
		seen := make(map[string]bool)
		for decoder.More() {
			keyToken, err := decoder.Token()
			if err != nil {
				return err
			}
			key, ok := keyToken.(string)
			if !ok {
				return fmt.Errorf("object key at %s is not a string", path)
			}
			if seen[key] {
				return fmt.Errorf("duplicate field %q at %s", key, path)
			}
			seen[key] = true
			if err := walkJSONValue(decoder, path+"."+key); err != nil {
				return err
			}
		}
	case '[':
		for index := 0; decoder.More(); index++ {
			if err := walkJSONValue(decoder, fmt.Sprintf("%s[%d]", path, index)); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unexpected JSON delimiter %q at %s", delimiter, path)
	}
	if _, err := decoder.Token(); err != nil {
		return err
	}
	return nil
}
