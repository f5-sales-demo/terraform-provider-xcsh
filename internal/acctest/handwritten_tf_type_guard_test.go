// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// Regression guard for #1354.
//
// examples/, tests/ and templates/ are hand-maintained: no generator rewrites
// them, so nothing corrects a Terraform type name after codegen stops emitting
// the type behind it. That is how the addon-activation guide kept declaring
// xcsh_addon_subscription for every release after #829 deleted it — anyone
// copying the guide got "Invalid resource type" from terraform validate.
//
// The type names live in HCL as bare strings, so neither the compiler nor the
// provider's own schema-drift guards can see them. Only reading the artifacts
// and comparing against what the provider registers catches this class.
//
// Truth comes from registeredTypeNames (generated_symbol_guard_test.go), which
// calls Resources()/DataSources() on the generated provider. There is no
// parallel allowlist to keep in sync.
//
// Coverage note: the defect this guards lived in a .md guide template, not a
// .tf file — templates/guides/*.md embed literal HCL fences rather than pulling
// examples in by directive. Restricting the scan to .tf would have missed the
// original bug entirely, so terraform-tagged fences in templates/ are read too.

var (
	// Matches `resource "xcsh_foo" "bar"` and `data "xcsh_foo" "bar"`, which is
	// the only position where an unregistered type name is a hard error.
	// Reference expressions (xcsh_foo.bar.field) are deliberately not matched:
	// they cannot name a type the declaration above did not already introduce.
	handwrittenTypeDeclaration = regexp.MustCompile(`(?m)^\s*(resource|data)\s+"(xcsh_[a-z0-9_]+)"`)

	// Matches an opening ```terraform / ```hcl fence.
	terraformFenceStart = regexp.MustCompile("^\\s*```(terraform|hcl)\\s*$")
	fenceEnd            = regexp.MustCompile("^\\s*```\\s*$")
)

// handwrittenTypeReference is one declaration site, kept with its location so a
// failure names the file and line to edit rather than just the offending type.
type handwrittenTypeReference struct {
	Path     string
	Line     int
	Block    string // "resource" or "data"
	TypeName string
}

func (r handwrittenTypeReference) String() string {
	return fmt.Sprintf("%s:%d: %s %q", r.Path, r.Line, r.Block, r.TypeName)
}

// extractHandwrittenTypeReferences pulls every xcsh_* type declaration out of
// HCL source. It is separated from the filesystem walk so the non-vacuity test
// below can drive it with a synthetic document.
func extractHandwrittenTypeReferences(path, source string) []handwrittenTypeReference {
	var references []handwrittenTypeReference
	for index, line := range strings.Split(source, "\n") {
		if match := handwrittenTypeDeclaration.FindStringSubmatch(line); match != nil {
			references = append(references, handwrittenTypeReference{
				Path:     path,
				Line:     index + 1,
				Block:    match[1],
				TypeName: match[2],
			})
		}
	}
	return references
}

// extractFencedTerraform returns the terraform-tagged fenced blocks of a
// Markdown document, preserving line numbers so failures point at the template
// line a maintainer has to edit.
func extractFencedTerraform(source string) string {
	lines := strings.Split(source, "\n")
	kept := make([]string, len(lines))
	inFence := false
	for index, line := range lines {
		switch {
		case !inFence && terraformFenceStart.MatchString(line):
			inFence = true
		case inFence && fenceEnd.MatchString(line):
			inFence = false
		case inFence:
			kept[index] = line
		}
	}
	return strings.Join(kept, "\n")
}

// collectHandwrittenTypeReferences walks the hand-maintained trees and returns
// every declaration site found in them.
func collectHandwrittenTypeReferences(t *testing.T, root string) []handwrittenTypeReference {
	t.Helper()

	var references []handwrittenTypeReference
	for _, tree := range []string{"examples", "tests", "templates"} {
		treeRoot := filepath.Join(root, tree)
		if _, err := os.Stat(treeRoot); os.IsNotExist(err) {
			continue
		}
		walkErr := filepath.Walk(treeRoot, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				if info.Name() == ".terraform" {
					return filepath.SkipDir
				}
				return nil
			}
			extension := filepath.Ext(path)
			if extension != ".tf" && extension != ".md" {
				return nil
			}
			contents, readErr := os.ReadFile(path) //nolint:gosec // repository-relative walk
			if readErr != nil {
				return readErr
			}
			source := string(contents)
			if extension == ".md" {
				source = extractFencedTerraform(source)
			}
			relative, relErr := filepath.Rel(root, path)
			if relErr != nil {
				relative = path
			}
			references = append(references, extractHandwrittenTypeReferences(relative, source)...)
			return nil
		})
		if walkErr != nil {
			t.Fatalf("walk %s: %v", tree, walkErr)
		}
	}
	return references
}

// TestHandwrittenArtifactsOnlyDeclareRegisteredTypes asserts that every xcsh_*
// resource and data source declared in the hand-maintained trees is a type the
// generated provider registers.
func TestHandwrittenArtifactsOnlyDeclareRegisteredTypes(t *testing.T) {
	root := testRepositoryRoot(t)
	resources, dataSources := registeredTypeNames(t)

	references := collectHandwrittenTypeReferences(t, root)

	// Vacuity floor: the guide and example trees declare well over a hundred
	// type names. Finding almost none means the walk or the pattern broke, and
	// a guard that sees nothing must fail rather than wave everything through.
	const minReferences = 50
	if len(references) < minReferences {
		t.Fatalf("found only %d hand-written type declarations (expected at least %d); this guard cannot be trusted", len(references), minReferences)
	}

	var unregistered []string
	for _, reference := range references {
		registry := resources
		if reference.Block == "data" {
			registry = dataSources
		}
		if _, ok := registry[reference.TypeName]; !ok {
			unregistered = append(unregistered, reference.String())
		}
	}

	if len(unregistered) > 0 {
		sort.Strings(unregistered)
		t.Fatalf("hand-written artifacts declare %d Terraform type(s) the provider does not register:\n  %s",
			len(unregistered), strings.Join(unregistered, "\n  "))
	}
}

// TestHandwrittenTypeGuardDetectsUnregisteredType is the mutation proof for the
// guard above: it feeds a type name the provider cannot register through the
// same extraction path and asserts it is surfaced. Without this, an extractor
// that silently matched nothing would leave the guard green forever.
func TestHandwrittenTypeGuardDetectsUnregisteredType(t *testing.T) {
	resources, dataSources := registeredTypeNames(t)

	// Names are split from the prefix for the reason documented on
	// TestFencedTerraformExtractionIgnoresProse below.
	const prefix = "xcsh_"
	synthetic := fmt.Sprintf(
		"\nresource %q \"mutant\" {\n  name = \"mutant\"\n}\n\ndata %q \"mutant\" {\n  name = \"mutant\"\n}\n",
		prefix+"definitely_not_a_registered_type", prefix+"also_not_registered")

	references := extractHandwrittenTypeReferences("synthetic.tf", synthetic)
	if len(references) != 2 {
		t.Fatalf("extractor found %d declarations in the synthetic document, want 2: %v", len(references), references)
	}

	for _, reference := range references {
		registry := resources
		if reference.Block == "data" {
			registry = dataSources
		}
		if _, ok := registry[reference.TypeName]; ok {
			t.Fatalf("%s is somehow registered; pick a name the provider cannot ship", reference.TypeName)
		}
	}

	// And the converse: a type the provider does register must NOT be flagged,
	// or the guard would fail on every healthy tree.
	var registeredExample string
	for name := range resources {
		registeredExample = name
		break
	}
	healthy := extractHandwrittenTypeReferences("healthy.tf",
		fmt.Sprintf("resource %q \"ok\" {\n  name = \"ok\"\n}\n", registeredExample))
	if len(healthy) != 1 {
		t.Fatalf("extractor found %d declarations in the healthy document, want 1", len(healthy))
	}
	if _, ok := resources[healthy[0].TypeName]; !ok {
		t.Fatalf("extractor mangled a registered type name: %q", healthy[0].TypeName)
	}
}

// TestFencedTerraformExtractionIgnoresProse guards the Markdown path: only
// terraform-tagged fences are read, so prose that happens to look like HCL, and
// fenced blocks in other languages, cannot produce phantom findings.
//
// The fixture type names are assembled from the prefix instead of being written
// out whole, and that is load-bearing rather than stylistic.
// TestAcctestSourceOnlyReferencesRegisteredTypeNames walks this package's string
// literals and requires every bare "xcsh_*" name to be one the provider
// registers (#1351). A synthetic name spelled out in full is exactly what that
// guard is built to catch, so the two would deadlock: this test needs names the
// provider cannot register, and that one forbids them. Splitting the prefix
// satisfies both — "xcsh_" alone has an empty suffix and is not a type-name
// literal. Do not "simplify" these back into single strings.
func TestFencedTerraformExtractionIgnoresProse(t *testing.T) {
	const prefix = "xcsh_"
	proseOnly, bashFence, realFence := prefix+"prose_only", prefix+"bash_fence", prefix+"real_fence"

	document := fmt.Sprintf("Prose mentioning resource %q \"x\" inline.\n", proseOnly) +
		fmt.Sprintf("\n```bash\nresource %q \"x\" {}\n```\n", bashFence) +
		fmt.Sprintf("\n```terraform\nresource %q \"x\" {}\n```\n", realFence)

	references := extractHandwrittenTypeReferences("doc.md", extractFencedTerraform(document))
	if len(references) != 1 {
		t.Fatalf("expected exactly one declaration from the terraform fence, got %d: %v", len(references), references)
	}
	if references[0].TypeName != realFence {
		t.Fatalf("read the wrong fence: %q", references[0].TypeName)
	}
}
