// Package releasesurface loads the exact public SMSv2 provider contract.
package releasesurface

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
)

type Surface struct {
	Resources   []string `json:"resources"`
	DataSources []string `json:"data_sources"`
	Actions     []string `json:"actions"`
	Functions   []string `json:"functions"`
}

// FilterResults makes the public provider surface equal the manifest. It also
// supplies read-only entries for intentionally hand-written SMSv2 data sources.
func (s *Surface) FilterResults(results []openapi.GenerationResult) []openapi.GenerationResult {
	kept := make(map[string]openapi.GenerationResult)
	for _, result := range results {
		switch {
		case result.IsTerraformAction && s.AllowsAction(result.ResourceName):
			kept[result.ResourceName] = result
		case result.IsReadOnly && s.AllowsDataSource(result.ResourceName):
			kept[result.ResourceName] = result
		case result.IsAction && s.AllowsResource(result.ResourceName):
			kept[result.ResourceName] = result
		case !result.IsReadOnly && !result.IsAction && !result.IsTerraformAction && s.AllowsResource(result.ResourceName):
			if !s.AllowsDataSource(result.ResourceName) {
				result.IsAction = true
			}
			kept[result.ResourceName] = result
		}
	}
	for _, name := range s.DataSources {
		if _, exists := kept[name]; !exists {
			kept[name] = openapi.GenerationResult{ResourceName: name, Success: true, IsReadOnly: true}
		}
	}
	filtered := make([]openapi.GenerationResult, 0, len(kept))
	for _, result := range kept {
		filtered = append(filtered, result)
	}
	return filtered
}

func Load(path string) (*Surface, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read release surface: %w", err)
	}
	var surface Surface
	if err := json.Unmarshal(data, &surface); err != nil {
		return nil, fmt.Errorf("parse release surface: %w", err)
	}
	if err := surface.validate(); err != nil {
		return nil, err
	}
	return &surface, nil
}

func (s *Surface) validate() error {
	for _, group := range []struct {
		name   string
		values []string
	}{
		{"resources", s.Resources}, {"data_sources", s.DataSources}, {"actions", s.Actions}, {"functions", s.Functions},
	} {
		seen := map[string]struct{}{}
		for _, value := range group.values {
			if value == "" {
				return fmt.Errorf("release surface %s contains an empty name", group.name)
			}
			if _, exists := seen[value]; exists {
				return fmt.Errorf("release surface %s contains duplicate %q", group.name, value)
			}
			seen[value] = struct{}{}
		}
	}
	return nil
}

func (s *Surface) AllowsResource(name string) bool   { return contains(s.Resources, name) }
func (s *Surface) AllowsDataSource(name string) bool { return contains(s.DataSources, name) }
func (s *Surface) AllowsAction(name string) bool     { return contains(s.Actions, name) }
func (s *Surface) AllowsFunction(name string) bool   { return contains(s.Functions, name) }
func contains(values []string, value string) bool {
	for _, candidate := range values {
		if candidate == value {
			return true
		}
	}
	return false
}
