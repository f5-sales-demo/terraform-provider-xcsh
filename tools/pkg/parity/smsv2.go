// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package parity builds the deterministic legacy-to-current SMSv2 contract
// matrix used to fail provider generation on an unclassified capability gap.
package parity

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

type LegacyManifest struct {
	Version      string        `json:"version"`
	Resource     string        `json:"resource"`
	SourceURL    string        `json:"source_url"`
	SourceSHA256 string        `json:"source_sha256"`
	PathCount    int           `json:"path_count"`
	Paths        []LegacyField `json:"paths"`
}

type LegacyField struct {
	Path          string      `json:"path"`
	WireKey       string      `json:"wire_key"`
	Type          string      `json:"type"`
	Cardinality   string      `json:"cardinality"`
	Required      bool        `json:"required"`
	Optional      bool        `json:"optional"`
	Computed      bool        `json:"computed"`
	ForceNew      bool        `json:"force_new"`
	Deprecated    bool        `json:"deprecated"`
	Default       interface{} `json:"default,omitempty"`
	ConflictsWith []string    `json:"conflicts_with"`
}

type CurrentManifest struct {
	Version                 string              `json:"version"`
	Resource                string              `json:"resource"`
	PathCount               int                 `json:"path_count"`
	Paths                   []CurrentField      `json:"paths"`
	ChoiceGroups            map[string][]string `json:"choice_groups"`
	DeprecatedExclusions    []string            `json:"deprecated_exclusions"`
	CurrentPlatformRemovals []string            `json:"current_platform_removals"`
}

type CurrentField struct {
	Path           string      `json:"path"`
	WireKey        string      `json:"wire_key"`
	Type           string      `json:"type"`
	Cardinality    string      `json:"cardinality"`
	Required       bool        `json:"required"`
	CreateRequired bool        `json:"create_required"`
	ReadOnly       bool        `json:"read_only"`
	WriteOnly      bool        `json:"write_only"`
	Default        interface{} `json:"default,omitempty"`
	Enum           []string    `json:"enum,omitempty"`
}

type FieldSemantics struct {
	Type          string      `json:"type"`
	Cardinality   string      `json:"cardinality"`
	Required      bool        `json:"required"`
	Optional      bool        `json:"optional"`
	Computed      bool        `json:"computed"`
	ForceNew      bool        `json:"force_new"`
	Default       interface{} `json:"default,omitempty"`
	ConflictsWith []string    `json:"conflicts_with"`
	WireKey       string      `json:"wire_key"`
	Generated     bool        `json:"generated,omitempty"`
}

type MatrixEntry struct {
	LegacyPath     string          `json:"legacy_path,omitempty"`
	CurrentPath    string          `json:"current_path,omitempty"`
	Classification string          `json:"classification"`
	Reason         string          `json:"reason"`
	Legacy         *FieldSemantics `json:"legacy,omitempty"`
	Current        *FieldSemantics `json:"current,omitempty"`
}

type Matrix struct {
	LegacyVersion    string         `json:"legacy_version"`
	CurrentVersion   string         `json:"current_version"`
	LegacySourceURL  string         `json:"legacy_source_url"`
	LegacySourceSHA  string         `json:"legacy_source_sha256"`
	LegacyPathCount  int            `json:"legacy_path_count"`
	CurrentPathCount int            `json:"current_path_count"`
	GeneratedPaths   int            `json:"generated_path_count,omitempty"`
	ClassifiedLegacy int            `json:"classified_legacy_paths"`
	Unclassified     []string       `json:"unclassified_legacy_paths"`
	Classification   map[string]int `json:"classification_counts"`
	Entries          []MatrixEntry  `json:"entries"`
}

func LoadLegacy(path string) (*LegacyManifest, error) {
	var value LegacyManifest
	if err := load(path, &value); err != nil {
		return nil, err
	}
	if value.PathCount != len(value.Paths) || value.Version != "0.11.49" || value.SourceSHA256 == "" {
		return nil, fmt.Errorf("invalid legacy SMSv2 manifest metadata")
	}
	return &value, nil
}

func LoadCurrent(path string) (*CurrentManifest, error) {
	var value CurrentManifest
	if err := load(path, &value); err != nil {
		return nil, err
	}
	if value.PathCount != len(value.Paths) || value.Version == "" || len(value.ChoiceGroups) == 0 {
		return nil, fmt.Errorf("invalid current SMSv2 parity manifest metadata")
	}
	return &value, nil
}

func load(path string, value interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, value); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	return nil
}

func BuildSMSv2Matrix(legacy *LegacyManifest, current *CurrentManifest) (*Matrix, error) {
	return buildSMSv2Matrix(legacy, current, nil)
}

// BuildSMSv2MatrixFromTerraform builds the matrix with practitioner-facing
// semantics from the generated Terraform attribute IR. The current manifest
// remains the authoritative API field inventory, but cannot falsely claim
// required/current parity for a field the generator emitted as optional.
func BuildSMSv2MatrixFromTerraform(legacy *LegacyManifest, current *CurrentManifest, attrs []openapi.TerraformAttribute) (*Matrix, error) {
	generated := flattenTerraformAttributes(attrs)
	return buildSMSv2Matrix(legacy, current, generated)
}

func buildSMSv2Matrix(legacy *LegacyManifest, current *CurrentManifest, generated map[string]openapi.TerraformAttribute) (*Matrix, error) {
	currentByPath := make(map[string]CurrentField, len(current.Paths))
	conflicts := make(map[string][]string)
	for _, members := range current.ChoiceGroups {
		for _, member := range members {
			for _, peer := range members {
				if peer != member {
					conflicts[member] = append(conflicts[member], terraformPath(peer))
				}
			}
			sort.Strings(conflicts[member])
		}
	}
	for _, field := range current.Paths {
		currentByPath[field.Path] = field
	}

	matrix := &Matrix{
		LegacyVersion:    legacy.Version,
		CurrentVersion:   current.Version,
		LegacySourceURL:  legacy.SourceURL,
		LegacySourceSHA:  legacy.SourceSHA256,
		LegacyPathCount:  legacy.PathCount,
		CurrentPathCount: current.PathCount,
		Classification:   map[string]int{},
		Unclassified:     make([]string, 0),
		Entries:          make([]MatrixEntry, 0, len(legacy.Paths)+len(current.Paths)),
	}
	if generated != nil {
		for path := range generated {
			if path != "metadata.id" && path != "metadata.uid" {
				matrix.GeneratedPaths++
			}
		}
	}
	consumedCurrent := make(map[string]bool, len(current.Paths))
	for _, field := range legacy.Paths {
		entry := classifyLegacyField(field, currentByPath, conflicts, generated)
		if entry.Classification == "" {
			matrix.Unclassified = append(matrix.Unclassified, field.Path)
			continue
		}
		if entry.CurrentPath != "" {
			consumedCurrent[apiPath(entry.CurrentPath)] = true
		}
		matrix.Classification[entry.Classification]++
		matrix.Entries = append(matrix.Entries, entry)
		matrix.ClassifiedLegacy++
	}
	for _, field := range current.Paths {
		if field.Path == "metadata" || field.Path == "spec" || consumedCurrent[field.Path] {
			continue
		}
		classification := "current_only"
		reason := "capability added after legacy v0.11.49"
		if generated != nil {
			if _, present := generated[field.Path]; !present {
				classification = "source_only_not_generated"
				reason = "current API field is not present in the generated Terraform schema"
			}
		}
		entry := MatrixEntry{
			CurrentPath:    terraformPath(field.Path),
			Classification: classification,
			Reason:         reason,
			Current:        resolvedCurrentSemantics(field, conflicts[field.Path], generated),
		}
		matrix.Classification[entry.Classification]++
		matrix.Entries = append(matrix.Entries, entry)
	}
	sort.Strings(matrix.Unclassified)
	sort.Slice(matrix.Entries, func(i, j int) bool {
		left := matrix.Entries[i].LegacyPath + "\x00" + matrix.Entries[i].CurrentPath
		right := matrix.Entries[j].LegacyPath + "\x00" + matrix.Entries[j].CurrentPath
		return left < right
	})
	if len(matrix.Unclassified) != 0 {
		return matrix, fmt.Errorf("SMSv2 parity has %d unclassified legacy paths", len(matrix.Unclassified))
	}
	return matrix, nil
}

func classifyLegacyField(field LegacyField, current map[string]CurrentField, conflicts map[string][]string, generated map[string]openapi.TerraformAttribute) MatrixEntry {
	entry := MatrixEntry{LegacyPath: field.Path, Legacy: legacySemantics(field)}
	if field.Deprecated || hasPathPrefix(field.Path, "log_receiver") || hasPathPrefix(field.Path, "private_adn") || hasPathPrefix(field.Path, "rseries") {
		entry.Classification = "deprecated_exclusion"
		entry.Reason = "legacy field is deprecated and intentionally not restored"
		return entry
	}

	target, reason, classification := mappedCurrentPath(field.Path)
	if target == "" {
		target = apiPath(field.Path)
	}
	if currentField, found := current[target]; found {
		entry.CurrentPath = terraformPath(target)
		entry.Current = resolvedCurrentSemantics(currentField, conflicts[target], generated)
		_, generatedFieldPresent := generated[target]
		if generated != nil && !generatedFieldPresent {
			entry.Classification = "generator_gap"
			entry.Reason = "legacy capability exists in the current API but is absent from the generated Terraform schema"
		} else if classification != "" {
			entry.Classification = classification
			entry.Reason = reason
		} else if semanticallyEqual(entry.Legacy, entry.Current) {
			entry.Classification = "current_parity"
			entry.Reason = "path and normalized Terraform semantics match"
		} else {
			entry.Classification = "modernized_semantics"
			entry.Reason = "capability retained with enriched requiredness, validation, cardinality, default, conflict, or wire semantics"
		}
		return entry
	}
	if classification == "current_platform_removal" {
		entry.Classification = classification
		entry.Reason = reason
		return entry
	}
	if strings.HasPrefix(field.Path, "aws.managed") {
		entry.Classification = "modernized_semantics"
		entry.Reason = "provider-managed AWS lifecycle is represented by xcsh_aws_vpc_site"
		return entry
	}
	if strings.HasPrefix(field.Path, "azure.managed") {
		entry.Classification = "modernized_semantics"
		entry.Reason = "provider-managed Azure lifecycle is represented by xcsh_azure_vnet_site"
		return entry
	}
	if strings.HasPrefix(field.Path, "gcp.managed") {
		entry.Classification = "modernized_semantics"
		entry.Reason = "provider-managed GCP lifecycle is represented by xcsh_gcp_vpc_site"
		return entry
	}
	return entry
}

func mappedCurrentPath(path string) (string, string, string) {
	if strings.Contains(path, ".network_option.segment_network") {
		return "spec.segment_vrf[].segment_network", "per-interface Segment selection is represented by the named top-level Segment VRF contract", "modernized_semantics"
	}
	if strings.Contains(path, ".static_ipv6_address.cluster_static_ip.interface_ip_map.") {
		parent := path[:strings.Index(path, ".static_ipv6_address.cluster_static_ip.interface_ip_map.")] + ".static_ipv6_address.cluster_static_ip.interface_ip_map"
		return apiPath(parent), "legacy map-entry implementation details are represented as one Terraform map", "modernized_semantics"
	}
	if strings.Contains(path, "blocked_services.blocked_sevice") {
		return apiPath(strings.Replace(path, "blocked_sevice", "blocked_service", 1)), "corrected legacy blocked_sevice spelling while preserving the current blocked_service wire key", "modernized_semantics"
	}
	for legacy, current := range map[string]string{
		"local_vrf.sli_config.nameserver_v6":           "local_vrf.sli_config.nameserver",
		"local_vrf.sli_config.secondary_nameserver_v6": "local_vrf.sli_config.secondary_nameserver",
		"local_vrf.sli_config.vip_v6":                  "local_vrf.sli_config.vip",
		"local_vrf.slo_config.nameserver_v6":           "local_vrf.slo_config.nameserver",
		"local_vrf.slo_config.secondary_nameserver_v6": "local_vrf.slo_config.secondary_nameserver",
		"local_vrf.slo_config.vip_v6":                  "local_vrf.slo_config.vip",
	} {
		if path == legacy {
			return apiPath(current), "address-family-specific field is represented by the current IP-address field", "modernized_semantics"
		}
	}
	if path == "segment_vrf[].segment_config.nameserver_v6" || path == "segment_vrf[].segment_config.secondary_nameserver_v6" {
		return "", "controlled current-platform probe accepted but stripped the field; the release evidence records the removal", "current_platform_removal"
	}
	return "", "", ""
}

func legacySemantics(field LegacyField) *FieldSemantics {
	fieldType, cardinality := normalizeLegacyType(field.Type, field.Cardinality)
	return &FieldSemantics{
		Type: fieldType, Cardinality: cardinality, Required: field.Required, Optional: field.Optional,
		Computed: field.Computed, ForceNew: field.ForceNew, Default: field.Default,
		ConflictsWith: append([]string(nil), field.ConflictsWith...), WireKey: field.WireKey,
	}
}

func currentSemantics(field CurrentField, conflicts []string) *FieldSemantics {
	fieldType := field.Type
	if fieldType == "array" {
		fieldType = "list"
	}
	required := field.CreateRequired && !field.ReadOnly
	optional := !required && !field.ReadOnly
	computed := field.ReadOnly
	defaultValue := field.Default
	if field.Path == "metadata.namespace" {
		required, optional, computed, defaultValue = false, true, true, "system"
	}
	return &FieldSemantics{
		Type: fieldType, Cardinality: field.Cardinality, Required: required, Optional: optional,
		Computed: computed, ForceNew: field.Path == "metadata.name" || field.Path == "metadata.namespace",
		Default: defaultValue, ConflictsWith: append([]string(nil), conflicts...), WireKey: field.WireKey,
		Generated: true,
	}
}

func resolvedCurrentSemantics(field CurrentField, conflicts []string, generated map[string]openapi.TerraformAttribute) *FieldSemantics {
	semantics := currentSemantics(field, conflicts)
	if generated == nil {
		return semantics
	}
	attr := generated[field.Path]
	if _, ok := generated[field.Path]; !ok {
		semantics.Required = false
		semantics.Optional = false
		semantics.Computed = false
		semantics.ForceNew = false
		semantics.Default = nil
		semantics.Generated = false
		return semantics
	}
	semantics.Required = attr.Required || attr.CreateRequired
	semantics.Optional = attr.Optional && !attr.CreateRequired
	semantics.Computed = attr.Computed
	semantics.ForceNew = attr.PlanModifier == "RequiresReplace"
	semantics.Default = attr.Default
	if attr.StringDefault != "" {
		semantics.Default = attr.StringDefault
	}
	if attr.JsonName != "" {
		semantics.WireKey = attr.JsonName
	}
	return semantics
}

func flattenTerraformAttributes(attrs []openapi.TerraformAttribute) map[string]openapi.TerraformAttribute {
	result := make(map[string]openapi.TerraformAttribute)
	metadata := map[string]bool{
		"annotations": true, "description": true, "disable": true,
		"labels": true, "name": true, "namespace": true,
	}
	var walk func([]openapi.TerraformAttribute, string)
	walk = func(fields []openapi.TerraformAttribute, prefix string) {
		for _, attr := range fields {
			name := attr.JsonName
			if name == "" {
				name = attr.TfsdkTag
			}
			if name == "" {
				name = attr.Name
			}
			if name == "" {
				continue
			}
			path := name
			if prefix == "" {
				if metadata[name] && !attr.IsSpecField {
					path = "metadata." + name
				} else {
					path = "spec." + name
				}
			} else {
				path = prefix + "." + name
			}
			if attr.Type == "list" || (attr.IsBlock && attr.NestedBlockType == "list") {
				path += "[]"
			}
			result[path] = attr
			childPrefix := path
			walk(attr.NestedAttributes, childPrefix)
		}
	}
	walk(attrs, "")
	return result
}

func normalizeLegacyType(fieldType, cardinality string) (string, string) {
	if cardinality == "single_block" {
		return "object", "single"
	}
	if fieldType == "map" {
		return "object", "single"
	}
	return fieldType, cardinality
}

func semanticallyEqual(left, right *FieldSemantics) bool {
	leftCopy, rightCopy := *left, *right
	// Generated is reporting provenance, not practitioner-facing field
	// semantics; do not let it turn otherwise equal legacy/current contracts
	// into a modernization.
	leftCopy.Generated = false
	rightCopy.Generated = false
	leftJSON, _ := json.Marshal(&leftCopy)
	rightJSON, _ := json.Marshal(&rightCopy)
	return string(leftJSON) == string(rightJSON)
}

func apiPath(terraform string) string {
	for _, metadata := range []string{"annotations", "description", "disable", "labels", "name", "namespace"} {
		if terraform == metadata {
			return "metadata." + terraform
		}
	}
	return "spec." + terraform
}

func terraformPath(api string) string {
	return strings.TrimPrefix(strings.TrimPrefix(api, "metadata."), "spec.")
}

func hasPathPrefix(path, prefix string) bool {
	return path == prefix || strings.HasPrefix(path, prefix+".") || strings.HasPrefix(path, prefix+"[]")
}
