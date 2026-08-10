// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadExpectedVersion(t *testing.T) {
	// 1. Missing file
	_, err := ReadExpectedVersion("nonexistent-file.txt")
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}

	// 2. Empty file
	tempDir := t.TempDir()
	emptyPath := filepath.Join(tempDir, "empty.txt")
	if err := os.WriteFile(emptyPath, []byte("   \n "), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	_, err = ReadExpectedVersion(emptyPath)
	if err == nil {
		t.Error("expected error for empty file, got nil")
	}

	// 3. Valid file
	validPath := filepath.Join(tempDir, "valid.txt")
	if err := os.WriteFile(validPath, []byte(" v2.1.217 \n"), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	version, err := ReadExpectedVersion(validPath)
	if err != nil {
		t.Fatalf("unexpected error for valid file: %v", err)
	}
	if version != "v2.1.217" {
		t.Errorf("expected version to be 'v2.1.217', got '%s'", version)
	}
}

func TestValidateSpecVersions(t *testing.T) {
	tests := []struct {
		name           string
		expected       string
		index          string
		catalog        string
		expectError    bool
		expectedErrSub string
	}{
		{
			name:        "exact matching versions",
			expected:    "2.1.217",
			index:       "2.1.217",
			catalog:     "2.1.217",
			expectError: false,
		},
		{
			name:        "matching with mixed v prefix",
			expected:    "v2.1.217",
			index:       "2.1.217",
			catalog:     "v2.1.217",
			expectError: false,
		},
		{
			name:           "empty expected version",
			expected:       "",
			index:          "2.1.217",
			catalog:        "2.1.217",
			expectError:    true,
			expectedErrSub: "expected version is empty",
		},
		{
			name:           "empty index version",
			expected:       "2.1.217",
			index:          "",
			catalog:        "2.1.217",
			expectError:    true,
			expectedErrSub: "index version is empty",
		},
		{
			name:           "empty catalog version",
			expected:       "2.1.217",
			index:          "2.1.217",
			catalog:        "",
			expectError:    true,
			expectedErrSub: "catalog version is empty",
		},
		{
			name:           "index version mismatch",
			expected:       "2.1.217",
			index:          "2.1.216",
			catalog:        "2.1.217",
			expectError:    true,
			expectedErrSub: "spec version mismatch: expected 2.1.217, got 2.1.216",
		},
		{
			name:           "catalog version mismatch",
			expected:       "2.1.217",
			index:          "2.1.217",
			catalog:        "2.1.215",
			expectError:    true,
			expectedErrSub: "catalog version mismatch: expected 2.1.217, got 2.1.215",
		},
		{
			name:           "matching malformed string (garbage)",
			expected:       "garbage",
			index:          "garbage",
			catalog:        "garbage",
			expectError:    true,
			expectedErrSub: "is not a valid semantic version",
		},
		{
			name:           "matching incomplete semver (major.minor)",
			expected:       "2.1",
			index:          "2.1",
			catalog:        "2.1",
			expectError:    true,
			expectedErrSub: "is not a valid semantic version",
		},
		{
			name:           "matching invalid 4-part version",
			expected:       "2.1.217.1",
			index:          "2.1.217.1",
			catalog:        "2.1.217.1",
			expectError:    true,
			expectedErrSub: "is not a valid semantic version",
		},
		{
			name:           "malformed index version",
			expected:       "2.1.217",
			index:          "invalid.version",
			catalog:        "2.1.217",
			expectError:    true,
			expectedErrSub: "is not a valid semantic version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSpecVersions(tt.expected, tt.index, tt.catalog)
			if tt.expectError {
				if err == nil {
					t.Error("expected error but got nil")
				} else if tt.expectedErrSub != "" && !contains(err.Error(), tt.expectedErrSub) {
					t.Errorf("expected error containing '%s', got '%v'", tt.expectedErrSub, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || s[0:len(substr)] == substr || len(s) > len(substr) && contains(s[1:], substr))
}
