// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package main

import (
	"encoding/json"
	"os"
	"regexp"
	"sort"
	"strings"
)

var publicSchemaFields = map[string]bool{}
var diffAttributePattern = regexp.MustCompile(`(?m)^\s*[+-]\s*"([a-zA-Z0-9_.#%]+)"\s*:`)

func loadPublicSchemaFields(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var root map[string]interface{}
	if err = json.Unmarshal(data, &root); err != nil {
		return err
	}
	var walk func(map[string]interface{}, string)
	walk = func(block map[string]interface{}, prefix string) {
		if attrs, ok := block["attributes"].(map[string]interface{}); ok {
			for name := range attrs {
				publicSchemaFields[prefix+name] = true
			}
		}
		if blocks, ok := block["block_types"].(map[string]interface{}); ok {
			for name, value := range blocks {
				publicSchemaFields[prefix+name] = true
				object, _ := value.(map[string]interface{})
				child, _ := object["block"].(map[string]interface{})
				walk(child, prefix+name+".")
			}
		}
	}
	providers, _ := root["provider_schemas"].(map[string]interface{})
	for _, value := range providers {
		provider, _ := value.(map[string]interface{})
		for _, kind := range []string{"resource_schemas", "data_source_schemas"} {
			schemas, _ := provider[kind].(map[string]interface{})
			for _, value := range schemas {
				schema, _ := value.(map[string]interface{})
				block, _ := schema["block"].(map[string]interface{})
				walk(block, "")
			}
		}
	}
	return nil
}
func extractPublicSchemaFields(output string) []string {
	result := []string{}
	seen := map[string]bool{}
	for _, match := range diffAttributePattern.FindAllStringSubmatch(output, -1) {
		parts := []string{}
		for _, part := range strings.Split(match[1], ".") {
			if part == "#" || part == "%" {
				continue
			}
			numeric := true
			for _, r := range part {
				if r < '0' || r > '9' {
					numeric = false
				}
			}
			if numeric {
				continue
			}
			parts = append(parts, part)
		}
		normalized := strings.Join(parts, ".")
		if publicSchemaFields[normalized] && !seen[normalized] {
			seen[normalized] = true
			result = append(result, normalized)
		}
	}
	sort.Strings(result)
	return result
}
