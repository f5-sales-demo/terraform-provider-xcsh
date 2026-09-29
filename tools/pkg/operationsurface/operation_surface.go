// Package operationsurface binds upstream API changes to public Terraform definitions.
package operationsurface

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

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
	if manifest.SpecRelease != expectedRelease {
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
	expected := make(map[string]bool, len(report.Operations.Additions))
	for _, op := range report.Operations.Additions {
		expected[op.Method+" "+op.Path] = true
	}
	if len(expected) != len(report.Operations.Additions) {
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
	for _, removed := range report.Operations.Removals {
		key := removed.Method + " " + removed.Path
		if catalogOps[key] != nil {
			return nil, fmt.Errorf("removed upstream operation remains in api-catalog.json: %s", key)
		}
	}
	return &manifest, nil
}
