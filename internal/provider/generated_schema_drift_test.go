// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Regression guard for #1351, third layer.
//
// The acceptance tests in this package are hand-written and they name generated
// schema attributes as STRINGS — resource.TestCheckResourceAttr(resourceName,
// "brute_force_detection_settings.max_login_failures", ...). Two things conspire
// to make that rot invisible:
//
//  1. The compiler sees a string, not an attribute, so a removed or renamed
//     attribute type-checks perfectly.
//  2. Every TestAcc* case is gated on TF_ACC, so CI compiles these files and then
//     skips every assertion in them. Nothing ever evaluates the string.
//
// Upstream restructured tenant_configuration in the same window that removed the
// apm* family: basic_configuration became tenant_details and
// brute_force_detection_settings became brute_force_detection. The generated
// resource followed; this package's acceptance test did not, and stayed green
// (skipped) while asserting on attributes that no longer exist.
//
// This guard evaluates those strings without a tenant, credentials or TF_ACC: it
// reads the top-level attribute and block names straight out of the generated
// Schema method and requires every checked path to start with one of them.

// frameworkMetaNames are addressable on every resource without appearing in its
// Schema, so they can never be drift.
var frameworkMetaNames = map[string]bool{
	"id":         true,
	"timeouts":   true,
	"depends_on": true,
	"lifecycle":  true,
	"count":      true,
	"for_each":   true,
	"provider":   true,
}

// attrCheckFuncs are the terraform-plugin-testing helpers whose second argument
// is a resource attribute path.
var attrCheckFuncs = map[string]bool{
	"TestCheckResourceAttr":           true,
	"TestCheckResourceAttrSet":        true,
	"TestCheckNoResourceAttr":         true,
	"TestCheckResourceAttrPtr":        true,
	"TestCheckTypeSetElemAttr":        true,
	"TestCheckTypeSetElemNestedAttrs": true,
}

// topLevelSchemaNames returns the names directly under Attributes and Blocks of
// the schema.Schema literal returned by the file's Schema method.
//
// The names are read from the AST rather than matched by indentation so the guard
// does not quietly stop finding anything the next time the generator's formatting
// changes — a guard that silently sees an empty schema would wave every stale
// path through.
func topLevelSchemaNames(t *testing.T, path string) map[string]bool {
	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parsing %s: %v", path, err)
	}

	var schemaLit *ast.CompositeLit
	ast.Inspect(file, func(n ast.Node) bool {
		if schemaLit != nil {
			return false
		}
		fn, ok := n.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "Schema" || fn.Recv == nil {
			return true
		}
		ast.Inspect(fn, func(inner ast.Node) bool {
			if schemaLit != nil {
				return false
			}
			lit, ok := inner.(*ast.CompositeLit)
			if !ok {
				return true
			}
			if sel, ok := lit.Type.(*ast.SelectorExpr); ok && sel.Sel.Name == "Schema" {
				schemaLit = lit
				return false
			}
			return true
		})
		return true
	})

	if schemaLit == nil {
		return nil
	}

	names := make(map[string]bool)
	for _, elt := range schemaLit.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := kv.Key.(*ast.Ident)
		if !ok || (key.Name != "Attributes" && key.Name != "Blocks") {
			continue
		}
		mapLit, ok := kv.Value.(*ast.CompositeLit)
		if !ok {
			continue
		}
		for _, entry := range mapLit.Elts {
			entryKV, ok := entry.(*ast.KeyValueExpr)
			if !ok {
				continue
			}
			lit, ok := entryKV.Key.(*ast.BasicLit)
			if !ok || lit.Kind != token.STRING {
				continue
			}
			name, err := strconv.Unquote(lit.Value)
			if err != nil {
				continue
			}
			names[name] = true
		}
	}
	return names
}

// checkedAttributePaths returns every literal attribute path asserted in a test
// file, mapped to the positions where it appears. Paths built with fmt.Sprintf or
// string concatenation are not literals and are out of scope.
func checkedAttributePaths(t *testing.T, path string) map[string][]string {
	t.Helper()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parsing %s: %v", path, err)
	}

	paths := make(map[string][]string)
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok || len(call.Args) < 2 {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || !attrCheckFuncs[sel.Sel.Name] {
			return true
		}
		// The first argument addresses the resource under test; only calls that
		// name it through an identifier are considered, which is how every test
		// in this package is written.
		if _, ok := call.Args[0].(*ast.Ident); !ok {
			return true
		}
		lit, ok := call.Args[1].(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return true
		}
		value, err := strconv.Unquote(lit.Value)
		if err != nil || value == "" || strings.Contains(value, "%") {
			return true
		}
		paths[value] = append(paths[value], fmt.Sprintf("%s:%d", filepath.Base(path), fset.Position(lit.Pos()).Line))
		return true
	})
	return paths
}

func TestProviderTestsOnlyCheckAttributesTheGeneratedSchemaDeclares(t *testing.T) {
	dir, err := filepath.Abs(".")
	if err != nil {
		t.Fatalf("resolving package directory: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "provider.go")); err != nil {
		t.Fatalf("package directory %s does not look like internal/provider (%v); this guard must not scan the wrong tree", dir, err)
	}

	// Each test file is paired with the generated source whose schema it asserts
	// on, by the naming convention the generator itself uses.
	suffixes := map[string]string{
		"_resource_test.go":      "_resource.go",
		"_resource_mock_test.go": "_resource.go",
		"_data_source_test.go":   "_data_source.go",
	}

	testFiles, err := filepath.Glob(filepath.Join(dir, "*_test.go"))
	if err != nil {
		t.Fatalf("listing test files in %s: %v", dir, err)
	}
	if len(testFiles) == 0 {
		t.Fatalf("found no test files in %s; this guard would be vacuous", dir)
	}

	var failures []string
	paired, checked := 0, 0

	for _, testFile := range testFiles {
		base := filepath.Base(testFile)

		var sourceFile string
		// Longest suffix first: _resource_mock_test.go also ends in _test.go.
		for _, suffix := range []string{"_resource_mock_test.go", "_resource_test.go", "_data_source_test.go"} {
			if !strings.HasSuffix(base, suffix) {
				continue
			}
			sourceFile = filepath.Join(dir, strings.TrimSuffix(base, suffix)+suffixes[suffix])
			break
		}
		if sourceFile == "" {
			continue
		}
		if _, err := os.Stat(sourceFile); err != nil {
			// No generated counterpart: either a hand-maintained utility data
			// source or a test whose resource was removed. The type-name guards
			// in internal/acctest cover the latter; nothing to compare here.
			continue
		}

		schemaNames := topLevelSchemaNames(t, sourceFile)
		if len(schemaNames) == 0 {
			t.Fatalf("extracted no top-level schema names from %s; the AST walk has stopped working and this guard would be vacuous", sourceFile)
		}
		paired++

		for _, attrPath := range sortedPathKeys(checkedAttributePaths(t, testFile)) {
			checked++
			segment := strings.SplitN(attrPath, ".", 2)[0]
			if frameworkMetaNames[segment] || schemaNames[segment] {
				continue
			}
			failures = append(failures, fmt.Sprintf("  %s checks %q, but %s declares no top-level %q",
				base, attrPath, filepath.Base(sourceFile), segment))
		}
	}

	if paired == 0 {
		t.Fatal("paired no test file with a generated source; this guard would be vacuous")
	}
	if checked == 0 {
		t.Fatal("found no literal attribute paths to check; this guard would be vacuous")
	}

	if len(failures) > 0 {
		sort.Strings(failures)
		t.Errorf("%d hand-written attribute check(s) name a schema attribute the generated provider no longer declares:\n%s\n\n"+
			"Acceptance tests are skipped without TF_ACC, so these strings are never evaluated in CI and rot silently. "+
			"Update them to the current generated schema (#1351).", len(failures), strings.Join(failures, "\n"))
	}
}

func sortedPathKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
