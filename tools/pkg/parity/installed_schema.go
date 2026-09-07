// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package parity

import (
	"encoding/json"
	"fmt"
	"sort"
)

// InstalledPaths independently extracts Terraform's exported SDK schema paths.
// It does not infer wire mappings, defaults or lifecycle equivalence, which the
// providers schema protocol does not expose.
func InstalledPaths(data []byte, address, resource string) ([]string, error) {
	type block struct {
		Attributes map[string]struct {
			Type json.RawMessage `json:"type"`
		} `json:"attributes"`
		Blocks map[string]struct {
			MaxItems int             `json:"max_items"`
			Block    json.RawMessage `json:"block"`
		} `json:"block_types"`
	}
	var document struct {
		Providers map[string]struct {
			Resources map[string]struct {
				Block json.RawMessage `json:"block"`
			} `json:"resource_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal(data, &document); err != nil {
		return nil, err
	}
	root := document.Providers[address].Resources[resource].Block
	if len(root) == 0 {
		return nil, fmt.Errorf("installed provider resource schema is missing")
	}
	paths := []string{}
	var walk func(json.RawMessage, string) error
	walk = func(raw json.RawMessage, prefix string) error {
		var value *block
		if err := json.Unmarshal(raw, &value); err != nil {
			return err
		}
		if value == nil || (prefix == "" && value.Attributes == nil && value.Blocks == nil) {
			return fmt.Errorf("installed schema block is empty or malformed")
		}
		for name, attribute := range value.Attributes {
			path := prefix + name
			var kind []json.RawMessage
			if json.Unmarshal(attribute.Type, &kind) == nil && len(kind) > 0 {
				var collection string
				if err := json.Unmarshal(kind[0], &collection); err != nil {
					return err
				}
				if collection == "list" || collection == "set" {
					path += "[]"
				}
			}
			paths = append(paths, path)
		}
		for name, nested := range value.Blocks {
			path := prefix + name
			if nested.MaxItems != 1 {
				path += "[]"
			}
			paths = append(paths, path)
			if err := walk(nested.Block, path+"."); err != nil {
				return err
			}
		}
		return nil
	}
	if err := walk(root, ""); err != nil {
		return nil, err
	}
	sort.Strings(paths)
	return paths, nil
}
