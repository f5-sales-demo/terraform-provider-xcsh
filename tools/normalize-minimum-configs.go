// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

// normalize-minimum-configs rewrites embedded YAML examples with the fleet's
// canonical synthetic organization values.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/docfmt"
)

type minimumConfig struct {
	ResourceName   string   `json:"resource_name"`
	RequiredFields []string `json:"required_fields"`
	ExampleFile    string   `json:"example_file"`
}

var resourceNamePattern = regexp.MustCompile(`^[a-z0-9_]+$`)

func main() {
	path := flag.String("file", "tools/minimum-configs.json", "Minimum-config file to normalize")
	flag.Parse()

	data, err := os.ReadFile(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read minimum configs: %v\n", err)
		os.Exit(1)
	}

	var configs []minimumConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		fmt.Fprintf(os.Stderr, "parse minimum configs: %v\n", err)
		os.Exit(1)
	}

	configDir := filepath.Join(filepath.Dir(*path), "minimum-configs")
	for index, config := range configs {
		if !resourceNamePattern.MatchString(config.ResourceName) {
			fmt.Fprintln(os.Stderr, "minimum config has an invalid resource name")
			os.Exit(1)
		}
		exampleName := config.ResourceName + ".yaml"
		expectedFile := filepath.ToSlash(filepath.Join("minimum-configs", exampleName))
		if config.ExampleFile != expectedFile {
			fmt.Fprintln(os.Stderr, "minimum config has an invalid example path")
			os.Exit(1)
		}
		examplePath := filepath.Join(configDir, exampleName)
		exampleData, err := os.ReadFile(examplePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read minimum-config example: %v\n", err)
			os.Exit(1)
		}
		example := docfmt.CanonicalizeIdentityLiterals(string(exampleData))
		if example != "" && example[len(example)-1] != '\n' {
			example += "\n"
		}
		if err := os.WriteFile(examplePath, []byte(example), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "write minimum-config example: %v\n", err)
			os.Exit(1)
		}
		configs[index] = config
	}

	data, err = json.MarshalIndent(configs, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal minimum configs: %v\n", err)
		os.Exit(1)
	}
	data = append(data, '\n')
	if err := os.WriteFile(*path, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write minimum configs: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Normalized minimum configurations: %s\n", *path)
}
