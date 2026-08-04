// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package suppress

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const defaultSuppressionComment = "Canonical per-resource server-default members suppressed on the Terraform import path. Updated from measured tenant behavior by tools/emit-import-suppressions.go. See issue #1006."

var (
	sourceResourceNamePattern  = regexp.MustCompile(`^[a-z][a-z0-9]*(?:_[a-z0-9]+)*$`)
	suppressionResourcePattern = regexp.MustCompile(`^[A-Z][A-Za-z0-9]*$`)
	memberNamePattern          = regexp.MustCompile(`^[a-z][a-z0-9]*(?:_[a-z0-9]+)*$`)
	defaultPathPattern         = regexp.MustCompile(`^spec(?:\.[a-z][a-z0-9]*(?:_[a-z0-9]+)*)+$`)
)

// EmitImportSuppressions merges newly discovered server defaults into the
// canonical measured suppression file and replaces it atomically.
func EmitImportSuppressions(inDB, outFile string) (returnErr error) {
	lock, err := acquireOutputLock(outFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := lock.Close(); err != nil {
			returnErr = errors.Join(returnErr, err)
		}
	}()

	existing, comment, err := readExistingSuppressions(outFile)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(inDB)
	if err != nil {
		return fmt.Errorf("read defaults database %s: %w", inDB, err)
	}
	if err := rejectDuplicateJSONKeys(data, inDB); err != nil {
		return err
	}
	var database Database
	if err := json.Unmarshal(data, &database); err != nil {
		return fmt.Errorf("parse defaults database %s: %w", inDB, err)
	}
	if err := validateDefaultsDatabase(data, database, inDB); err != nil {
		return err
	}

	derived := Derive(database)
	merged := Merge(existing, derived)
	if err := validateSuppressionMap(merged); err != nil {
		return fmt.Errorf("validate merged suppressions: %w", err)
	}
	out := map[string]interface{}{"_comment": comment}
	for resource, members := range merged {
		out[resource] = members
	}
	encoded, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal canonical suppressions: %w", err)
	}
	if err := atomicWriteFile(outFile, append(encoded, '\n'), 0o644); err != nil {
		return err
	}
	fmt.Printf("emit-import-suppressions: %d resources in %s (%d derived from %s)\n", len(merged), outFile, len(derived), inDB)
	return nil
}

func readExistingSuppressions(path string) (map[string][]string, string, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return map[string][]string{}, defaultSuppressionComment, nil
	}
	if err != nil {
		return nil, "", fmt.Errorf("read existing suppressions %s: %w", path, err)
	}
	return ParseCanonicalSuppressions(data, path)
}

// ParseCanonicalSuppressions strictly parses the canonical measured data. Both
// the emitter and code generator use this function so corrupt data cannot be
// accepted on one path and rejected on the other.
func ParseCanonicalSuppressions(data []byte, path string) (map[string][]string, string, error) {
	if err := rejectDuplicateJSONKeys(data, path); err != nil {
		return nil, "", err
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, "", fmt.Errorf("parse canonical suppressions %s: %w", path, err)
	}
	commentValue, ok := raw["_comment"]
	if !ok {
		return nil, "", fmt.Errorf("parse canonical suppressions %s: missing _comment", path)
	}
	var comment string
	if err := json.Unmarshal(commentValue, &comment); err != nil || strings.TrimSpace(comment) == "" {
		return nil, "", fmt.Errorf("parse canonical suppressions %s: _comment must be a non-empty string", path)
	}
	resourceNames := make([]string, 0, len(raw)-1)
	for resourceName := range raw {
		if resourceName != "_comment" {
			resourceNames = append(resourceNames, resourceName)
		}
	}
	sort.Strings(resourceNames)
	parsed := make(map[string][]string, len(resourceNames))
	for _, resourceName := range resourceNames {
		var members []string
		if err := json.Unmarshal(raw[resourceName], &members); err != nil {
			return nil, "", fmt.Errorf("parse canonical suppressions %s: resource %q must contain a string array: %w", path, resourceName, err)
		}
		parsed[resourceName] = members
	}
	if err := validateSuppressionMap(parsed); err != nil {
		return nil, "", fmt.Errorf("parse canonical suppressions %s: %w", path, err)
	}
	return parsed, comment, nil
}

func validateDefaultsDatabase(data []byte, database Database, path string) error {
	var top map[string]json.RawMessage
	if err := json.Unmarshal(data, &top); err != nil {
		return fmt.Errorf("parse defaults database %s: %w", path, err)
	}
	allowedTopLevel := map[string]bool{
		"api_endpoint": true, "discovered": true, "failed": true, "generated_at": true,
		"resources": true, "skipped": true, "total_resources": true, "version": true,
	}
	for _, field := range sortedRawKeys(top) {
		if !allowedTopLevel[field] {
			return fmt.Errorf("parse defaults database %s: unknown top-level field %q", path, field)
		}
	}
	for _, requiredField := range []string{"api_endpoint", "discovered", "failed", "generated_at", "resources", "skipped", "total_resources", "version"} {
		if top[requiredField] == nil {
			return fmt.Errorf("parse defaults database %s: missing top-level field %q", path, requiredField)
		}
	}
	if _, err := requiredNonBlankString(top, "version"); err != nil {
		return fmt.Errorf("parse defaults database %s: %w", path, err)
	}
	generatedAt, err := requiredNonBlankString(top, "generated_at")
	if err != nil {
		return fmt.Errorf("parse defaults database %s: %w", path, err)
	}
	if _, err := time.Parse(time.RFC3339, generatedAt); err != nil {
		return fmt.Errorf("parse defaults database %s: generated_at must be an RFC3339 timestamp", path)
	}
	apiEndpoint, err := requiredNonBlankString(top, "api_endpoint")
	if err != nil {
		return fmt.Errorf("parse defaults database %s: %w", path, err)
	}
	parsedEndpoint, err := url.Parse(apiEndpoint)
	if err != nil || (parsedEndpoint.Scheme != "https" && parsedEndpoint.Scheme != "http") || parsedEndpoint.Host == "" {
		return fmt.Errorf("parse defaults database %s: api_endpoint must be an absolute HTTP(S) URL", path)
	}
	if database.Resources == nil {
		return fmt.Errorf("parse defaults database %s: resources object is required", path)
	}
	var rawResources map[string]json.RawMessage
	if err := json.Unmarshal(top["resources"], &rawResources); err != nil {
		return fmt.Errorf("parse defaults database %s: resources must be an object", path)
	}
	allowedResourceFields := map[string]bool{
		"category": true, "defaults": true, "discovered_at": true, "error": true,
		"request_sent": true, "resource_name": true, "response_got": true, "skip_reason": true,
		"status": true,
	}
	allowedDefaultFields := map[string]bool{
		"default_value": true, "description": true, "is_marker_block": true, "path": true, "type": true,
	}
	statusCounts := map[string]int{"discovered": 0, "failed": 0, "skipped": 0}
	for _, resourceKey := range sortedResourceKeys(database.Resources) {
		resource := database.Resources[resourceKey]
		if !sourceResourceNamePattern.MatchString(resourceKey) || resource == nil {
			return fmt.Errorf("parse defaults database %s: resource keys must be lowercase snake_case identifiers and values must be objects", path)
		}
		var rawResource map[string]json.RawMessage
		if err := json.Unmarshal(rawResources[resourceKey], &rawResource); err != nil {
			return fmt.Errorf("parse defaults database %s: resource %q must be an object", path, resourceKey)
		}
		for _, field := range sortedRawKeys(rawResource) {
			if !allowedResourceFields[field] {
				return fmt.Errorf("parse defaults database %s: resource %q has unknown field %q", path, resourceKey, field)
			}
		}
		resourceName, err := requiredNonBlankString(rawResource, "resource_name")
		if err != nil || !sourceResourceNamePattern.MatchString(resourceName) || resourceName != resourceKey {
			return fmt.Errorf("parse defaults database %s: resource %q must have a matching lowercase snake_case resource_name", path, resourceKey)
		}
		if _, err := requiredNonBlankString(rawResource, "category"); err != nil {
			return fmt.Errorf("parse defaults database %s: resource %q must have a non-empty category string", path, resourceKey)
		}
		status, err := requiredNonBlankString(rawResource, "status")
		if err != nil || (status != "discovered" && status != "failed" && status != "skipped") {
			return fmt.Errorf("parse defaults database %s: resource %q must have a known non-empty status string", path, resourceKey)
		}
		statusCounts[status]++
		if err := validateResourceProvenance(rawResource, status); err != nil {
			return fmt.Errorf("parse defaults database %s: resource %q: %w", path, resourceKey, err)
		}
		var rawDefaults map[string]json.RawMessage
		if raw, present := rawResource["defaults"]; present {
			if err := json.Unmarshal(rawResource["defaults"], &rawDefaults); err != nil {
				return fmt.Errorf("parse defaults database %s: resource %q defaults must be an object", path, resourceKey)
			}
			if rawDefaults == nil || bytes.Equal(bytes.TrimSpace(raw), []byte("null")) {
				return fmt.Errorf("parse defaults database %s: resource %q defaults must be an object", path, resourceKey)
			}
		}
		defaultKeys := make([]string, 0, len(resource.Defaults))
		for defaultKey := range resource.Defaults {
			defaultKeys = append(defaultKeys, defaultKey)
		}
		sort.Strings(defaultKeys)
		for _, defaultKey := range defaultKeys {
			fieldDefault := resource.Defaults[defaultKey]
			if !defaultPathPattern.MatchString(defaultKey) || fieldDefault.Path != defaultKey {
				return fmt.Errorf("parse defaults database %s: resource %q has a malformed or mismatched default path", path, resourceKey)
			}
			if err := validateFieldDefault(fieldDefault); err != nil {
				return fmt.Errorf("parse defaults database %s: resource %q default %q: %w", path, resourceKey, defaultKey, err)
			}
			var rawDefault map[string]json.RawMessage
			if err := json.Unmarshal(rawDefaults[defaultKey], &rawDefault); err != nil {
				return fmt.Errorf("parse defaults database %s: resource %q default %q must be an object", path, resourceKey, defaultKey)
			}
			for _, field := range sortedRawKeys(rawDefault) {
				if !allowedDefaultFields[field] {
					return fmt.Errorf("parse defaults database %s: resource %q default %q has unknown field %q", path, resourceKey, defaultKey, field)
				}
			}
			for _, requiredField := range []string{"default_value", "path", "type"} {
				if rawDefault[requiredField] == nil {
					return fmt.Errorf("parse defaults database %s: resource %q default %q is missing %q", path, resourceKey, defaultKey, requiredField)
				}
			}
			if rawPath, err := requiredNonBlankString(rawDefault, "path"); err != nil || rawPath != defaultKey {
				return fmt.Errorf("parse defaults database %s: resource %q default %q path must be a matching string", path, resourceKey, defaultKey)
			}
			if rawType, err := requiredNonBlankString(rawDefault, "type"); err != nil || rawType != fieldDefault.Type {
				return fmt.Errorf("parse defaults database %s: resource %q default %q type must be a matching string", path, resourceKey, defaultKey)
			}
			if marker, present := rawDefault["is_marker_block"]; present {
				var markerValue bool
				if err := json.Unmarshal(marker, &markerValue); err != nil || bytes.Equal(bytes.TrimSpace(marker), []byte("null")) {
					return fmt.Errorf("parse defaults database %s: resource %q default %q is_marker_block must be a boolean", path, resourceKey, defaultKey)
				}
			}
			if _, present := rawDefault["description"]; present {
				if _, err := requiredNonBlankString(rawDefault, "description"); err != nil {
					return fmt.Errorf("parse defaults database %s: resource %q default %q description must be a non-empty string", path, resourceKey, defaultKey)
				}
			}
		}
	}
	for _, field := range []string{"discovered", "failed", "skipped", "total_resources"} {
		declared, err := nonNegativeInteger(top[field])
		if err != nil {
			return fmt.Errorf("parse defaults database %s: %s must be a non-negative integer", path, field)
		}
		measured := len(database.Resources)
		if field != "total_resources" {
			measured = statusCounts[field]
		}
		if declared != measured {
			return fmt.Errorf("parse defaults database %s: %s=%d does not match measured %d", path, field, declared, measured)
		}
	}
	return nil
}

func validateResourceProvenance(raw map[string]json.RawMessage, status string) error {
	validateObject := func(field string, required bool) error {
		value, present := raw[field]
		if !present {
			if required {
				return fmt.Errorf("%s is required", field)
			}
			return nil
		}
		var object map[string]interface{}
		if err := json.Unmarshal(value, &object); err != nil || object == nil || len(object) == 0 {
			return fmt.Errorf("%s must be a non-empty object", field)
		}
		return nil
	}
	validateAbsent := func(fields ...string) error {
		for _, field := range fields {
			if _, present := raw[field]; present {
				return fmt.Errorf("%s must be absent for status %s", field, status)
			}
		}
		return nil
	}

	switch status {
	case "discovered":
		discoveredAt, err := requiredNonBlankString(raw, "discovered_at")
		if err != nil {
			return fmt.Errorf("discovered_at must be a non-empty RFC3339 string")
		}
		if _, err := time.Parse(time.RFC3339, discoveredAt); err != nil {
			return fmt.Errorf("discovered_at must be a non-empty RFC3339 string")
		}
		if err := validateObject("request_sent", true); err != nil {
			return err
		}
		if err := validateObject("response_got", true); err != nil {
			return err
		}
		return validateAbsent("error", "skip_reason")
	case "failed":
		if _, err := requiredNonBlankString(raw, "error"); err != nil {
			return fmt.Errorf("error must be a non-empty string")
		}
		if err := validateObject("request_sent", false); err != nil {
			return err
		}
		return validateAbsent("defaults", "discovered_at", "response_got", "skip_reason")
	case "skipped":
		if _, err := requiredNonBlankString(raw, "skip_reason"); err != nil {
			return fmt.Errorf("skip_reason must be a non-empty string")
		}
		return validateAbsent("defaults", "discovered_at", "error", "request_sent", "response_got")
	default:
		return fmt.Errorf("unknown status")
	}
}

func requiredNonBlankString(values map[string]json.RawMessage, field string) (string, error) {
	raw, present := values[field]
	if !present {
		return "", fmt.Errorf("%s is required", field)
	}
	var value string
	if err := json.Unmarshal(raw, &value); err != nil || strings.TrimSpace(value) == "" {
		return "", fmt.Errorf("%s must be a non-empty string", field)
	}
	return value, nil
}

func nonNegativeInteger(raw json.RawMessage) (int, error) {
	if bytes.Equal(bytes.TrimSpace(raw), []byte("null")) {
		return 0, fmt.Errorf("null is not an integer")
	}
	var value int
	if err := json.Unmarshal(raw, &value); err != nil || value < 0 {
		return 0, fmt.Errorf("not a non-negative integer")
	}
	return value, nil
}

func validateFieldDefault(fieldDefault FieldDefault) error {
	validType := false
	switch fieldDefault.Type {
	case "array":
		_, validType = fieldDefault.DefaultValue.([]interface{})
	case "bool":
		_, validType = fieldDefault.DefaultValue.(bool)
	case "null":
		validType = fieldDefault.DefaultValue == nil
	case "number":
		_, validType = fieldDefault.DefaultValue.(float64)
	case "object":
		_, validType = fieldDefault.DefaultValue.(map[string]interface{})
	case "string":
		_, validType = fieldDefault.DefaultValue.(string)
	default:
		return fmt.Errorf("unknown type %q", fieldDefault.Type)
	}
	if !validType {
		return fmt.Errorf("type %q does not match default_value", fieldDefault.Type)
	}
	object, isObject := fieldDefault.DefaultValue.(map[string]interface{})
	isEmptyObject := isObject && len(object) == 0
	if fieldDefault.IsMarkerBlock != isEmptyObject {
		return fmt.Errorf("is_marker_block must be true exactly for an empty object default")
	}
	return nil
}

func validateSuppressionMap(suppressions map[string][]string) error {
	if len(suppressions) == 0 {
		return fmt.Errorf("at least one measured resource is required")
	}
	resources := make([]string, 0, len(suppressions))
	for resource := range suppressions {
		resources = append(resources, resource)
	}
	sort.Strings(resources)
	for _, resource := range resources {
		members := suppressions[resource]
		if !suppressionResourcePattern.MatchString(resource) || len(members) == 0 {
			return fmt.Errorf("resource names must be Go-style identifiers and suppression lists must be non-empty")
		}
		seen := make(map[string]bool, len(members))
		for _, member := range members {
			if !memberNamePattern.MatchString(member) || seen[member] {
				return fmt.Errorf("resource %q contains a malformed or duplicate member %q", resource, member)
			}
			seen[member] = true
		}
	}
	return nil
}

func sortedRawKeys(values map[string]json.RawMessage) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedResourceKeys(values map[string]*ResourceResult) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func rejectDuplicateJSONKeys(data []byte, path string) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	token, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("parse JSON %s: %w", path, err)
	}
	if err := walkJSONValue(decoder, token, path); err != nil {
		return err
	}
	if _, err := decoder.Token(); err != io.EOF {
		if err == nil {
			return fmt.Errorf("parse JSON %s: trailing JSON value", path)
		}
		return fmt.Errorf("parse JSON %s: %w", path, err)
	}
	return nil
}

func walkJSONValue(decoder *json.Decoder, token json.Token, path string) error {
	delimiter, ok := token.(json.Delim)
	if !ok {
		return nil
	}
	switch delimiter {
	case '{':
		seen := map[string]bool{}
		for decoder.More() {
			keyToken, err := decoder.Token()
			if err != nil {
				return fmt.Errorf("parse JSON %s: %w", path, err)
			}
			key, ok := keyToken.(string)
			if !ok {
				return fmt.Errorf("parse JSON %s: object key is not a string", path)
			}
			if seen[key] {
				return fmt.Errorf("parse JSON %s: duplicate object key %q", path, key)
			}
			seen[key] = true
			valueToken, err := decoder.Token()
			if err != nil {
				return fmt.Errorf("parse JSON %s: %w", path, err)
			}
			if err := walkJSONValue(decoder, valueToken, path); err != nil {
				return err
			}
		}
	case '[':
		for decoder.More() {
			valueToken, err := decoder.Token()
			if err != nil {
				return fmt.Errorf("parse JSON %s: %w", path, err)
			}
			if err := walkJSONValue(decoder, valueToken, path); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("parse JSON %s: unexpected delimiter %q", path, delimiter)
	}
	if _, err := decoder.Token(); err != nil {
		return fmt.Errorf("parse JSON %s: %w", path, err)
	}
	return nil
}

func atomicWriteFile(path string, data []byte, mode os.FileMode) error {
	dir := filepath.Dir(path)
	temporary, err := os.CreateTemp(dir, "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return fmt.Errorf("create temporary suppression file for %s: %w", path, err)
	}
	temporaryPath := temporary.Name()
	defer func() {
		_ = temporary.Close()
		if temporaryPath != "" {
			_ = os.Remove(temporaryPath)
		}
	}()
	if err := temporary.Chmod(mode); err != nil {
		return fmt.Errorf("set temporary suppression file mode for %s: %w", path, err)
	}
	if _, err := temporary.Write(data); err != nil {
		return fmt.Errorf("write temporary suppression file for %s: %w", path, err)
	}
	if err := temporary.Sync(); err != nil {
		return fmt.Errorf("sync temporary suppression file for %s: %w", path, err)
	}
	if err := temporary.Close(); err != nil {
		return fmt.Errorf("close temporary suppression file for %s: %w", path, err)
	}
	if err := replaceFileAndSyncDirectory(temporaryPath, path, dir); err != nil {
		return fmt.Errorf("replace and sync suppression file %s: %w", path, err)
	}
	temporaryPath = ""
	return nil
}
