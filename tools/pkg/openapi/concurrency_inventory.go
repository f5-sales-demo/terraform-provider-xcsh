// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ConcurrencyInventory is the released, provider-wide replace-token contract.
// APIIdentity includes the exact service owner (for example
// ves.io.schema.site.API), unlike api-catalog's resource grouping.
type ConcurrencyInventory struct {
	Version       string
	EligibleCount int
	CoveredCount  int
	ExcludedCount int
	Resources     []ConcurrencyInventoryResource
	Exclusions    []ConcurrencyInventoryExclusion
	resources     map[string]ConcurrencyInventoryResource
	exclusions    map[string]ConcurrencyInventoryExclusion
}

type ConcurrencyInventoryEnvelope struct {
	Path   string `json:"path"`
	Schema string `json:"schema"`
}

type ConcurrencyInventoryResource struct {
	APIIdentity  string                       `json:"api_identity"`
	Get          ConcurrencyInventoryEnvelope `json:"get"`
	Replace      ConcurrencyInventoryEnvelope `json:"replace"`
	CreateSchema *string                      `json:"create_schema"`
	Token        string                       `json:"token"`
}

type ConcurrencyInventoryExclusion struct {
	APIIdentity string `json:"api_identity"`
	Operation   string `json:"operation"`
	Reason      string `json:"reason"`
}

type concurrencyInventoryDocument struct {
	Version       *string                          `json:"version"`
	EligibleCount *int                             `json:"eligible_count"`
	CoveredCount  *int                             `json:"covered_count"`
	ExcludedCount *int                             `json:"excluded_count"`
	Resources     *[]ConcurrencyInventoryResource  `json:"resources"`
	Exclusions    *[]ConcurrencyInventoryExclusion `json:"exclusions"`
}

// ParseConcurrencyInventoryFromDir loads concurrency_contracts.json from an
// immutable enriched-spec bundle.
func ParseConcurrencyInventoryFromDir(specDir string) (*ConcurrencyInventory, error) {
	path := filepath.Join(specDir, "concurrency_contracts.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read concurrency inventory %s: %w", path, err)
	}
	inventory, err := ParseConcurrencyInventory(data)
	if err != nil {
		return nil, fmt.Errorf("parse concurrency inventory %s: %w", path, err)
	}
	return inventory, nil
}

// ParseConcurrencyInventory strictly validates the released inventory. Counts
// and identities are part of the contract, not advisory summary fields.
func ParseConcurrencyInventory(data []byte) (*ConcurrencyInventory, error) {
	if err := rejectDuplicateJSONFields(data); err != nil {
		return nil, fmt.Errorf("decode concurrency inventory: %w", err)
	}
	var document concurrencyInventoryDocument
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("decode concurrency inventory: %w", err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return nil, fmt.Errorf("decode concurrency inventory: trailing JSON value")
	}
	if document.Version == nil || !catalogVersionRegex.MatchString(*document.Version) {
		return nil, fmt.Errorf("version is required and must match %s", catalogVersionRegex)
	}
	if document.EligibleCount == nil || document.CoveredCount == nil || document.ExcludedCount == nil ||
		document.Resources == nil || document.Exclusions == nil {
		return nil, fmt.Errorf("eligible_count, covered_count, excluded_count, resources, and exclusions are required")
	}
	if *document.EligibleCount < 0 || *document.CoveredCount < 0 || *document.ExcludedCount < 0 {
		return nil, fmt.Errorf("concurrency inventory counts must be non-negative")
	}
	if *document.EligibleCount != len(*document.Resources) || *document.CoveredCount != len(*document.Resources) ||
		*document.ExcludedCount != len(*document.Exclusions) {
		return nil, fmt.Errorf("concurrency inventory counts do not match resources and exclusions")
	}

	inventory := &ConcurrencyInventory{
		Version:       *document.Version,
		EligibleCount: *document.EligibleCount,
		CoveredCount:  *document.CoveredCount,
		ExcludedCount: *document.ExcludedCount,
		Resources:     *document.Resources,
		Exclusions:    *document.Exclusions,
		resources:     make(map[string]ConcurrencyInventoryResource, len(*document.Resources)),
		exclusions:    make(map[string]ConcurrencyInventoryExclusion, len(*document.Exclusions)),
	}
	for index, resource := range inventory.Resources {
		if !strings.HasPrefix(resource.APIIdentity, "ves.io.schema.") || resource.Get.Path == "" ||
			resource.Get.Schema == "" || resource.Replace.Path == "" || resource.Replace.Schema == "" || resource.Token != "resource_version" {
			return nil, fmt.Errorf("resources[%d] has an incomplete or unsupported contract", index)
		}
		if _, exists := inventory.resources[resource.APIIdentity]; exists {
			return nil, fmt.Errorf("resources[%d] duplicates api_identity %q", index, resource.APIIdentity)
		}
		inventory.resources[resource.APIIdentity] = resource
	}
	for index, exclusion := range inventory.Exclusions {
		if !strings.HasPrefix(exclusion.APIIdentity, "ves.io.schema.") || exclusion.Operation != "Replace" ||
			strings.TrimSpace(exclusion.Reason) == "" || exclusion.Reason != strings.TrimSpace(exclusion.Reason) {
			return nil, fmt.Errorf("exclusions[%d] has an incomplete or unsupported contract", index)
		}
		if _, exists := inventory.resources[exclusion.APIIdentity]; exists {
			return nil, fmt.Errorf("api_identity %q is both covered and excluded", exclusion.APIIdentity)
		}
		if _, exists := inventory.exclusions[exclusion.APIIdentity]; exists {
			return nil, fmt.Errorf("exclusions[%d] duplicates api_identity %q", index, exclusion.APIIdentity)
		}
		inventory.exclusions[exclusion.APIIdentity] = exclusion
	}
	return inventory, nil
}

// ClassifyReplace returns whether an exact Replace operation is token-covered
// or evidence-backed excluded. An unclassified operation is always an error.
func (inventory *ConcurrencyInventory) ClassifyReplace(operation CatalogOperation) (bool, *ConcurrencyInventoryExclusion, error) {
	if operation.Method != "PUT" || !strings.HasSuffix(operation.OperationID, ".Replace") {
		return false, nil, fmt.Errorf("operation %q is not an exact PUT Replace", operation.OperationID)
	}
	identity := strings.TrimSuffix(operation.OperationID, ".Replace")
	if resource, ok := inventory.resources[identity]; ok {
		if resource.Replace.Path != operation.Path || resource.Replace.Schema != operation.RequestSchema {
			return false, nil, fmt.Errorf("covered Replace %s disagrees with api-catalog", operation.OperationID)
		}
		return true, nil, nil
	}
	if exclusion, ok := inventory.exclusions[identity]; ok {
		copy := exclusion
		return false, &copy, nil
	}
	return false, nil, fmt.Errorf("Replace operation %s has no concurrency contract or exclusion", operation.OperationID)
}

// ValidateAgainstCatalog proves every catalog Replace is classified and every
// published inventory entry corresponds to exact catalog Get/Replace facts.
func (inventory *ConcurrencyInventory) ValidateAgainstCatalog(catalog *OperationCatalog) error {
	if inventory == nil || catalog == nil {
		return fmt.Errorf("concurrency inventory and operation catalog are required")
	}
	if inventory.Version != catalog.Version {
		return fmt.Errorf("concurrency inventory version %s does not match catalog version %s", inventory.Version, catalog.Version)
	}
	seen := make(map[string]bool, len(inventory.resources)+len(inventory.exclusions))
	getOperations := make(map[string]CatalogOperation)
	for _, group := range catalog.APIOperations {
		for _, operation := range group.Operations {
			if operation.Method == "GET" && strings.HasSuffix(operation.OperationID, ".Get") {
				getOperations[strings.TrimSuffix(operation.OperationID, ".Get")] = operation
			}
		}
	}
	for _, group := range catalog.APIOperations {
		for _, operation := range group.Operations {
			if operation.Method != "PUT" || !strings.HasSuffix(operation.OperationID, ".Replace") {
				continue
			}
			covered, exclusion, err := inventory.ClassifyReplace(operation)
			if err != nil {
				return err
			}
			identity := strings.TrimSuffix(operation.OperationID, ".Replace")
			seen[identity] = true
			if !covered {
				if exclusion == nil {
					return fmt.Errorf("Replace operation %s has no classification", operation.OperationID)
				}
				continue
			}
			get, ok := getOperations[identity]
			resource := inventory.resources[identity]
			if !ok || get.Path != resource.Get.Path || get.ResponseSchema != resource.Get.Schema {
				return fmt.Errorf("covered Replace identity %s has no matching catalog Get", identity)
			}
		}
	}
	missing := make([]string, 0)
	for identity := range inventory.resources {
		if !seen[identity] {
			missing = append(missing, identity)
		}
	}
	for identity := range inventory.exclusions {
		if !seen[identity] {
			missing = append(missing, identity)
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		return fmt.Errorf("concurrency inventory identities have no catalog Replace: %s", strings.Join(missing, ", "))
	}
	return nil
}
