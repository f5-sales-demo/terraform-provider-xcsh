// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var strictSemverRegex = regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9.]+)?$`)

// ReadExpectedVersion reads the version string from a given file path.
func ReadExpectedVersion(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("read expected version: %w", err)
	}
	v := strings.TrimSpace(string(bytes))
	if v == "" {
		return "", fmt.Errorf("expected version file is empty")
	}
	return v, nil
}

// ValidateSpecVersions asserts that the versions from tools/spec-version.txt,
// index.json, and api-catalog.json match in a fail-closed, prefix-insensitive,
// and strict semver manner.
func ValidateSpecVersions(expectedVersion, indexVersion, catalogVersion string) error {
	cleanExpected := strings.TrimPrefix(strings.TrimSpace(expectedVersion), "v")
	cleanExpected = strings.TrimPrefix(cleanExpected, "V")
	if cleanExpected == "" {
		return fmt.Errorf("expected version is empty")
	}
	if !strictSemverRegex.MatchString(cleanExpected) {
		return fmt.Errorf("expected version %q is not a valid semantic version", expectedVersion)
	}

	cleanIndex := strings.TrimPrefix(strings.TrimSpace(indexVersion), "v")
	cleanIndex = strings.TrimPrefix(cleanIndex, "V")
	if cleanIndex == "" {
		return fmt.Errorf("index version is empty")
	}
	if !strictSemverRegex.MatchString(cleanIndex) {
		return fmt.Errorf("index version %q is not a valid semantic version", indexVersion)
	}
	if cleanIndex != cleanExpected {
		return fmt.Errorf("spec version mismatch: expected %s, got %s", expectedVersion, indexVersion)
	}

	cleanCatalog := strings.TrimPrefix(strings.TrimSpace(catalogVersion), "v")
	cleanCatalog = strings.TrimPrefix(cleanCatalog, "V")
	if cleanCatalog == "" {
		return fmt.Errorf("catalog version is empty")
	}
	if !strictSemverRegex.MatchString(cleanCatalog) {
		return fmt.Errorf("catalog version %q is not a valid semantic version", catalogVersion)
	}
	if cleanCatalog != cleanExpected {
		return fmt.Errorf("catalog version mismatch: expected %s, got %s", expectedVersion, catalogVersion)
	}

	return nil
}
