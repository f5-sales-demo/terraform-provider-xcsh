//go:build ignore
// +build ignore

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNoStaleDatasourceKeys(t *testing.T) {
	mapsToCheck := []map[string]string{
		resourceConfigs,
		skipResources,
		resourceDependencies,
	}
	
	// Add keys from boolean maps
	for key := range systemLevelResources {
		mapsToCheck[0][key] = ""
	}
	for key := range systemNamespaceResources {
		mapsToCheck[0][key] = ""
	}

	for _, m := range mapsToCheck {
		for key := range m {
			if key == "namespace" {
				continue
			}

			resourcePath := filepath.Join("../internal/provider", key+"_resource.go")
			dataSourcePath := filepath.Join("../internal/provider", key+"_data_source.go")
			
			_, errRes := os.Stat(resourcePath)
			_, errData := os.Stat(dataSourcePath)
			
			if os.IsNotExist(errRes) && os.IsNotExist(errData) {
				t.Errorf("generate-datasource-tests.go key %q has no corresponding resource or data source file", key)
			}
		}
	}
}
