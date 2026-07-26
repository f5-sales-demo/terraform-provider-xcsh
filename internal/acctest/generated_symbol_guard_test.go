// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"context"
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

	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/provider"
)

// Regression guard for #1351.
//
// internal/acctest is the one hand-written package whose entire job is to mirror
// symbols that codegen emits from the F5 Distributed Cloud OpenAPI specs. When
// upstream removes a schema family, codegen correctly drops the resource, its
// data source and its client methods — but nothing updates this package, so its
// mirrors silently rot.
//
// That is how main was left uncompilable: F5 removed the whole apm* family, the
// generated client lost GetAPM/DeleteAPM, and the two hand-written "xcsh_apm"
// entries in check_destroy_registry.go kept calling them.
//
// The Go compiler catches the IDENTIFIER half of that class (c.GetAPM) as long
// as CI actually compiles this package — see the `go build ./...` and `go vet
// ./...` steps in .github/workflows/_build-test.yml, and the pre-PR compile gate
// in .github/workflows/on-merge.yml, both of which this issue also fixed.
//
// The compiler is blind to the STRING half: a resource type name such as
// "xcsh_apm" is just a map key. Nothing but this guard notices when the resource
// behind it stops existing. Both halves have to be covered or the class returns
// on the next upstream removal.
//
// The provider itself is the source of truth here, not a parallel list: the sets
// below come from calling Resources()/DataSources() on the generated provider and
// asking each one for its own TypeName. There is no allowlist to keep in sync, so
// there is nothing to forget to update.

// registeredTypeNames returns the Terraform type names the generated provider
// actually registers, split into resources and data sources.
//
// Both sets are asserted non-empty: a guard that compares against an empty
// universe passes everything, which is worse than no guard at all.
func registeredTypeNames(t *testing.T) (resources, dataSources map[string]struct{}) {
	t.Helper()

	ctx := context.Background()
	p := provider.New("guard")()

	var providerMeta fwprovider.MetadataResponse
	p.Metadata(ctx, fwprovider.MetadataRequest{}, &providerMeta)
	if providerMeta.TypeName == "" {
		t.Fatal("provider Metadata returned an empty TypeName; every derived type name would be wrong")
	}

	resources = make(map[string]struct{})
	for _, newResource := range p.Resources(ctx) {
		var meta fwresource.MetadataResponse
		newResource().Metadata(ctx, fwresource.MetadataRequest{ProviderTypeName: providerMeta.TypeName}, &meta)
		if meta.TypeName == "" {
			t.Fatal("a registered resource reported an empty TypeName")
		}
		resources[meta.TypeName] = struct{}{}
	}

	dataSources = make(map[string]struct{})
	for _, newDataSource := range p.DataSources(ctx) {
		var meta fwdatasource.MetadataResponse
		newDataSource().Metadata(ctx, fwdatasource.MetadataRequest{ProviderTypeName: providerMeta.TypeName}, &meta)
		if meta.TypeName == "" {
			t.Fatal("a registered data source reported an empty TypeName")
		}
		dataSources[meta.TypeName] = struct{}{}
	}

	// Vacuity floors. The provider registers well over a hundred of each; a
	// collapse to a handful means the reflection above stopped working, and a
	// guard that no longer sees the universe must fail loudly rather than wave
	// every stale reference through.
	const minRegistered = 50
	if len(resources) < minRegistered {
		t.Fatalf("provider registered only %d resources (expected at least %d); this guard cannot be trusted", len(resources), minRegistered)
	}
	if len(dataSources) < minRegistered {
		t.Fatalf("provider registered only %d data sources (expected at least %d); this guard cannot be trusted", len(dataSources), minRegistered)
	}

	return resources, dataSources
}

// TestCheckDestroyRegistriesOnlyReferenceRegisteredResources asserts that every
// key in the hand-written CheckDestroy registries names a resource the generated
// provider still registers.
//
// A key with no resource behind it is dead weight at best; at worst — the #1351
// case — its body calls a generated client method that no longer exists and the
// package stops compiling.
func TestCheckDestroyRegistriesOnlyReferenceRegisteredResources(t *testing.T) {
	registered, _ := registeredTypeNames(t)

	registries := map[string][]string{
		"resourceVerifierRegistry": sortedKeys(resourceVerifierRegistry),
		"resourceDeleterRegistry":  sortedKeys(resourceDeleterRegistry),
	}

	for _, name := range sortedKeys(registries) {
		keys := registries[name]
		if len(keys) == 0 {
			t.Fatalf("%s is empty; this guard would be vacuous", name)
		}

		var stale []string
		for _, key := range keys {
			if _, ok := registered[key]; !ok {
				stale = append(stale, key)
			}
		}
		if len(stale) > 0 {
			sort.Strings(stale)
			t.Errorf("%s in internal/acctest/check_destroy_registry.go references %d resource type(s) the provider no longer registers: %s\n\n"+
				"Codegen drops a resource when F5 removes its schema upstream, but this registry is hand-written and codegen cannot update it. "+
				"Delete the stale entries — do not re-add the resource (#1351).",
				name, len(stale), strings.Join(stale, ", "))
		}
	}
}

// TestAcctestSourceOnlyReferencesRegisteredTypeNames widens the same check to
// every "xcsh_*" string literal in this package, not just the two registries.
//
// acctest.go, sweep.go and tracker.go all hard-code resource type names too, and
// they rot in exactly the same way; the registries were simply where it surfaced
// first because they were the only ones that also touched a generated Go symbol.
//
// Literals are collected from the AST, so names appearing only in comments (for
// example the "xcsh_my_resource" placeholder in sweep.go's doc comment) are
// correctly ignored and need no allowlist.
//
// Only literals shaped like a complete type name are considered — see
// typeNameLiteral. Format strings such as "xcsh_%s.test", which build a
// Terraform address rather than name a type, are deliberately out of scope: the
// name they produce is not visible to a static scan, and inventing a rule for
// them would trade a precise guard for a noisy one.
func TestAcctestSourceOnlyReferencesRegisteredTypeNames(t *testing.T) {
	resources, dataSources := registeredTypeNames(t)

	// go test runs each package's tests with that package's directory as the
	// working directory, so "." is this package. Resolve it to an absolute path
	// up front so failures name a path a human can act on.
	dir, err := filepath.Abs(".")
	if err != nil {
		t.Fatalf("resolving package directory: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "check_destroy_registry.go")); err != nil {
		t.Fatalf("package directory %s does not look like internal/acctest (%v); this guard must not scan the wrong tree", dir, err)
	}

	paths, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		t.Fatalf("listing Go files in %s: %v", dir, err)
	}

	fset := token.NewFileSet()

	// literal -> positions where it occurs.
	found := make(map[string][]string)
	for _, path := range paths {
		file, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
		if err != nil {
			t.Fatalf("parsing %s: %v", path, err)
		}
		ast.Inspect(file, func(n ast.Node) bool {
			lit, ok := n.(*ast.BasicLit)
			if !ok || lit.Kind != token.STRING {
				return true
			}
			value, err := strconv.Unquote(lit.Value)
			if err != nil || !typeNameLiteral(value) {
				return true
			}
			found[value] = append(found[value], fmt.Sprintf("%s:%d", filepath.Base(path), fset.Position(lit.Pos()).Line))
			return true
		})
	}

	if len(paths) == 0 {
		t.Fatalf("parsed no Go files in %s; this guard would be vacuous", dir)
	}
	if len(found) == 0 {
		t.Fatalf("found no \"xcsh_*\" string literals in %s; this guard would be vacuous", dir)
	}

	var failures []string
	for _, literal := range sortedKeys(found) {
		if _, ok := resources[literal]; ok {
			continue
		}
		if _, ok := dataSources[literal]; ok {
			continue
		}
		positions := found[literal]
		sort.Strings(positions)
		failures = append(failures, fmt.Sprintf("  %q at %s", literal, strings.Join(positions, ", ")))
	}

	if len(failures) > 0 {
		t.Errorf("internal/acctest hard-codes %d Terraform type name(s) the generated provider no longer registers:\n%s\n\n"+
			"These are hand-written mirrors of generated symbols. When F5 removes a schema upstream, codegen drops the resource "+
			"and this package has to follow — delete the stale references rather than re-adding the resource (#1351).",
			len(failures), strings.Join(failures, "\n"))
	}
}

// typeNameLiteral reports whether a string literal is shaped like a complete
// Terraform type name for this provider: the "xcsh_" prefix followed by a
// non-empty snake_case identifier, and nothing else.
//
// The shape test is what keeps the guard allowlist-free. It rejects the bare
// prefix and any literal that merely embeds a name — format strings
// ("xcsh_%s.test"), resource addresses, HCL fragments — none of which name a
// type on their own.
func typeNameLiteral(value string) bool {
	const prefix = "xcsh_"
	suffix, ok := strings.CutPrefix(value, prefix)
	if !ok || suffix == "" {
		return false
	}
	for _, r := range suffix {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '_':
		default:
			return false
		}
	}
	return true
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
