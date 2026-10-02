// Package operationsurface binds upstream API changes to public Terraform definitions.
package operationsurface

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

//go:embed testdata/v9.0.0-changes.json
var baselineFS embed.FS

var terraformNamePattern = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

type Operation struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	OperationID string `json:"operation_id"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
}

type Manifest struct {
	SchemaVersion int         `json:"schema_version"`
	SpecRelease   string      `json:"spec_release"`
	Operations    []Operation `json:"operations"`
}

type changeReport struct {
	Operations struct {
		Additions []struct {
			Method string `json:"method"`
			Path   string `json:"path"`
		} `json:"additions"`
		Removals []struct {
			Method string `json:"method"`
			Path   string `json:"path"`
		} `json:"removals"`
	} `json:"operations"`
}

func LoadAndApply(manifestPath, reportPath, expectedRelease string, catalog *openapi.OperationCatalog) (*Manifest, error) {
	if catalog == nil {
		return nil, fmt.Errorf("operation catalog is required")
	}
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("read operation surface: %w", err)
	}
	var manifest Manifest
	decoder := json.NewDecoder(strings.NewReader(string(manifestBytes)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&manifest); err != nil {
		return nil, fmt.Errorf("parse operation surface: %w", err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return nil, fmt.Errorf("parse operation surface: trailing JSON content")
	}
	if manifest.SchemaVersion != 1 || manifest.SpecRelease == "" {
		return nil, fmt.Errorf("operation surface schema_version must be 1 and spec_release is required")
	}
	if !compatibleRelease(manifest.SpecRelease, expectedRelease) {
		return nil, fmt.Errorf("operation surface spec_release %q does not match %q", manifest.SpecRelease, expectedRelease)
	}
	reportBytes, err := os.ReadFile(reportPath)
	if err != nil {
		return nil, fmt.Errorf("read upstream contract changes: %w", err)
	}
	var report changeReport
	if err := json.Unmarshal(reportBytes, &report); err != nil {
		return nil, fmt.Errorf("parse upstream contract changes: %w", err)
	}
	baselineBytes, err := baselineFS.ReadFile("testdata/v9.0.0-changes.json")
	if err != nil {
		return nil, fmt.Errorf("read operation baseline: %w", err)
	}
	var baseline changeReport
	if err := json.Unmarshal(baselineBytes, &baseline); err != nil {
		return nil, fmt.Errorf("parse operation baseline: %w", err)
	}
	expected := make(map[string]bool, len(baseline.Operations.Additions))
	for _, op := range baseline.Operations.Additions {
		expected[op.Method+" "+op.Path] = true
	}
	if len(expected) != len(baseline.Operations.Additions) {
		return nil, fmt.Errorf("upstream additions contain duplicate method/path entries")
	}
	catalogOps := make(map[string]*openapi.CatalogOperation)
	for i := range catalog.APIOperations {
		for j := range catalog.APIOperations[i].Operations {
			op := &catalog.APIOperations[i].Operations[j]
			catalogOps[op.Method+" "+op.Path] = op
		}
	}
	seen := make(map[string]bool, len(manifest.Operations))
	names := make(map[string]string)
	for i, mapped := range manifest.Operations {
		key := mapped.Method + " " + mapped.Path
		if seen[key] {
			return nil, fmt.Errorf("operation surface contains duplicate mapping %q", key)
		}
		seen[key] = true
		if !expected[key] {
			return nil, fmt.Errorf("operation surface maps non-added operation %q", key)
		}
		catalogOperation := catalogOps[key]
		if catalogOperation == nil {
			return nil, fmt.Errorf("operation surface operation %q is absent from api-catalog.json", key)
		}
		if mapped.OperationID != catalogOperation.OperationID {
			return nil, fmt.Errorf("operation surface operation %q operation ID mismatch: %q != %q", key, mapped.OperationID, catalogOperation.OperationID)
		}
		if !terraformNamePattern.MatchString(mapped.Name) {
			return nil, fmt.Errorf("operation surface operations[%d] has invalid Terraform name %q", i, mapped.Name)
		}
		switch mapped.Kind {
		case "resource":
			if mapped.Name != "cloud_user_account" {
				return nil, fmt.Errorf("operation surface resource mapping %q is not the v9 CRUD addition", mapped.Name)
			}
		case "data_source":
			if mapped.Name == "cloud_user_account" {
				break
			}
			catalogOperation.Role = "query"
			catalogOperation.TerraformName = mapped.Name
		case "action":
			catalogOperation.Role = "action"
			catalogOperation.TerraformName = mapped.Name
		case "ephemeral_resource":
			catalogOperation.Role = "ephemeral"
			catalogOperation.TerraformName = mapped.Name
		default:
			return nil, fmt.Errorf("operation surface operations[%d] has unsupported kind %q", i, mapped.Kind)
		}
		if mapped.Kind != "resource" && mapped.Name != "cloud_user_account" {
			if prior, ok := names[mapped.Name]; ok {
				return nil, fmt.Errorf("operation surface Terraform name %q maps both %s and %s", mapped.Name, prior, key)
			}
			names[mapped.Name] = key
		}
	}
	if len(seen) != len(expected) {
		missing := make([]string, 0)
		for key := range expected {
			if !seen[key] {
				missing = append(missing, key)
			}
		}
		sort.Strings(missing)
		return nil, fmt.Errorf("operation surface leaves %d upstream additions unmapped: %v", len(missing), missing)
	}
	// A later release delta may be empty, but any new operations still require
	// explicit mappings, and removed operations must disappear from the catalog.
	for _, added := range report.Operations.Additions {
		key := added.Method + " " + added.Path
		if !seen[key] {
			return nil, fmt.Errorf("current upstream addition is unmapped: %s", key)
		}
	}
	for _, removed := range report.Operations.Removals {
		key := removed.Method + " " + removed.Path
		if catalogOps[key] != nil {
			return nil, fmt.Errorf("removed upstream operation remains in api-catalog.json: %s", key)
		}
	}
	return &manifest, nil
}

// compatibleRelease keeps the reviewed baseline within its API major version.
func compatibleRelease(baseline, current string) bool {
	parse := func(value string) ([3]int, bool) {
		var result [3]int
		if !strings.HasPrefix(value, "v") {
			return result, false
		}
		parts := strings.Split(strings.TrimPrefix(value, "v"), ".")
		if len(parts) != 3 {
			return result, false
		}
		for i, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil || n < 0 || strconv.Itoa(n) != part {
				return result, false
			}
			result[i] = n
		}
		return result, true
	}
	old, ok := parse(baseline)
	if !ok {
		return false
	}
	next, ok := parse(current)
	if !ok || old[0] != next[0] {
		return false
	}
	return next[1] > old[1] || (next[1] == old[1] && next[2] >= old[2])
}
