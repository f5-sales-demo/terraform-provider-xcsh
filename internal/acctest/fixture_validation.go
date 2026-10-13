// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
	"os"
	"sort"
	"strings"
)

var acceptanceFixtures = map[string]func() string{}
var acceptanceTestFixtures = map[string][]string{}

func RegisterFixture(name string, config func() string)   { acceptanceFixtures[name] = config }
func RegisterTestFixtures(name string, fixtures []string) { acceptanceTestFixtures[name] = fixtures }

type FixtureDiagnostic struct {
	Fixture  string `json:"fixture"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Category string `json:"category"`
}
type fixtureBlockSchema struct {
	Attributes map[string]struct {
		Required bool            `json:"required"`
		Computed bool            `json:"computed"`
		Optional bool            `json:"optional"`
		Type     json.RawMessage `json:"type,omitempty"`
	} `json:"attributes"`
	Blocks map[string]struct {
		Block    fixtureBlockSchema `json:"block"`
		MinItems int                `json:"min_items"`
	} `json:"block_types"`
}
type FixtureSchema struct {
	Resources   map[string]fixtureBlockSchema
	DataSources map[string]fixtureBlockSchema
}

func LoadFixtureSchema(path string) (*FixtureSchema, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var raw struct {
		Providers map[string]struct {
			Resources map[string]struct {
				Block fixtureBlockSchema `json:"block"`
			} `json:"resource_schemas"`
			DataSources map[string]struct {
				Block fixtureBlockSchema `json:"block"`
			} `json:"data_source_schemas"`
		} `json:"provider_schemas"`
	}
	if err = json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	result := &FixtureSchema{Resources: map[string]fixtureBlockSchema{}, DataSources: map[string]fixtureBlockSchema{}}
	for _, provider := range raw.Providers {
		for name, resource := range provider.Resources {
			result.Resources[name] = resource.Block
		}
		for name, source := range provider.DataSources {
			result.DataSources[name] = source.Block
		}
	}
	if len(result.Resources) == 0 {
		return nil, fmt.Errorf("installed fixture schema is empty")
	}
	return result, nil
}
func ParseFixture(config string) (*hclsyntax.Body, error) {
	file, diags := hclsyntax.ParseConfig([]byte(config), "fixture.tf", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, fmt.Errorf("fixture syntax invalid")
	}
	return file.Body.(*hclsyntax.Body), nil
}
func (s *FixtureSchema) Validate(name, config string) []FixtureDiagnostic {
	body, err := ParseFixture(config)
	if err != nil {
		return []FixtureDiagnostic{{Fixture: name, Category: "syntax"}}
	}
	result := []FixtureDiagnostic{}
	var walk func(*hclsyntax.Body, fixtureBlockSchema, string, string)
	walk = func(body *hclsyntax.Body, schema fixtureBlockSchema, typ, path string) {
		for key, attribute := range schema.Attributes {
			if attribute.Required && body.Attributes[key] == nil {
				result = append(result, FixtureDiagnostic{name, typ, path + key, "missing_required_argument"})
			}
		}
		for key := range body.Attributes {
			attribute, ok := schema.Attributes[key]
			if !ok {
				result = append(result, FixtureDiagnostic{name, typ, path + key, "unsupported_argument"})
			} else if attribute.Computed && !attribute.Optional && !attribute.Required {
				result = append(result, FixtureDiagnostic{name, typ, path + key, "read_only_argument"})
			} else if attribute.Type != nil {
				value, diags := body.Attributes[key].Expr.Value(nil)
				if !diags.HasErrors() && value.IsKnown() && !value.IsNull() {
					var expected interface{}
					if json.Unmarshal(attribute.Type, &expected) == nil {
						kind, _ := expected.(string)
						if compound, ok := expected.([]interface{}); ok && len(compound) > 0 {
							kind, _ = compound[0].(string)
						}
						valid := true
						switch kind {
						case "string":
							valid = value.Type() == cty.String
						case "number":
							valid = value.Type() == cty.Number
						case "bool":
							valid = value.Type() == cty.Bool
						case "object", "map":
							valid = value.Type().IsObjectType() || value.Type().IsMapType()
						case "list", "set", "tuple":
							valid = value.Type().IsTupleType() || value.Type().IsListType() || value.Type().IsSetType()
						}
						if !valid {
							result = append(result, FixtureDiagnostic{name, typ, path + key, "incorrect_attribute_type"})
						}
					}
				}
			}
		}
		counts := map[string]int{}
		for _, block := range body.Blocks {
			nested, ok := schema.Blocks[block.Type]
			if !ok {
				result = append(result, FixtureDiagnostic{name, typ, path + block.Type, "unsupported_block"})
			} else {
				counts[block.Type]++
				walk(block.Body, nested.Block, typ, path+block.Type+".")
			}
		}
		for key, block := range schema.Blocks {
			if counts[key] < block.MinItems {
				result = append(result, FixtureDiagnostic{name, typ, path + key, "missing_required_block"})
			}
		}
	}
	declarations := map[string]bool{"var": true, "local": true, "path": true, "terraform": true, "count": true, "each": true, "provider": true}
	for _, block := range body.Blocks {
		if (block.Type == "resource" || block.Type == "data") && len(block.Labels) == 2 {
			key := block.Labels[0] + "." + block.Labels[1]
			if block.Type == "data" {
				key = "data." + key
			}
			declarations[key] = true
		}
	}
	for _, block := range body.Blocks {
		if (block.Type != "resource" && block.Type != "data") || len(block.Labels) != 2 || !strings.HasPrefix(block.Labels[0], "xcsh_") {
			continue
		}
		schemas := s.Resources
		if block.Type == "data" {
			schemas = s.DataSources
		}
		schema, ok := schemas[block.Labels[0]]
		if !ok {
			result = append(result, FixtureDiagnostic{name, block.Labels[0], "", "unsupported_type"})
			continue
		}
		// Terraform meta arguments are evaluated separately.
		copyBody := *block.Body
		copyBody.Attributes = map[string]*hclsyntax.Attribute{}
		for key, value := range block.Body.Attributes {
			if key != "depends_on" && key != "count" && key != "for_each" && key != "provider" {
				copyBody.Attributes[key] = value
			}
		}
		filtered := []*hclsyntax.Block{}
		for _, child := range copyBody.Blocks {
			if child.Type != "lifecycle" {
				filtered = append(filtered, child)
			}
		}
		copyBody.Blocks = filtered
		walk(&copyBody, schema, block.Labels[0], "")
		_ = hclsyntax.VisitAll(block.Body, func(node hclsyntax.Node) hcl.Diagnostics {
			expr, ok := node.(hclsyntax.Expression)
			if !ok {
				return nil
			}
			for _, traversal := range expr.Variables() {
				root := traversal.RootName()
				if declarations[root] {
					continue
				}
				key := root
				if len(traversal) > 1 {
					if attribute, ok := traversal[1].(hcl.TraverseAttr); ok {
						key += "." + attribute.Name
					}
				}
				if root == "data" && len(traversal) > 2 {
					if attribute, ok := traversal[2].(hcl.TraverseAttr); ok {
						key += "." + attribute.Name
					}
				}
				if !declarations[key] {
					result = append(result, FixtureDiagnostic{name, block.Labels[0], "", "undeclared_reference"})
				}
			}
			return nil
		})
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		return a.Type+"/"+a.Path+"/"+a.Category < b.Type+"/"+b.Path+"/"+b.Category
	})
	return result
}

type OperationRequirement struct {
	Method    string
	Type      string
	Namespace string
	Path      string
}

func FixtureRequirements(config string) ([]OperationRequirement, error) {
	body, err := ParseFixture(config)
	if err != nil {
		return nil, err
	}
	requirements := []OperationRequirement{}
	for _, block := range body.Blocks {
		if block.Type != "resource" || len(block.Labels) != 2 || !strings.HasPrefix(block.Labels[0], "xcsh_") {
			continue
		}
		typ := block.Labels[0]
		if typ == "xcsh_swagger_object" {
			requirements = append(requirements, OperationRequirement{Method: "PUT", Type: typ, Namespace: "shared", Path: "/api/object_store/namespaces/shared/stored_objects/swagger"})
			continue
		}
		route, ok := ownedCollectionRoutes[typ]
		if !ok {
			return nil, fmt.Errorf("no capability inventory for resource")
		}
		ns := "system"
		if attribute := block.Body.Attributes["namespace"]; attribute != nil {
			value, diags := attribute.Expr.Value(nil)
			if !diags.HasErrors() && value.Type() == cty.String && value.IsKnown() && !value.IsNull() {
				ns = value.AsString()
			} else {
				ns = "disposable"
			}
		}
		requirements = append(requirements, OperationRequirement{Method: "POST", Type: typ, Namespace: ns, Path: strings.ReplaceAll(route, "%s", ns)})
	}
	return requirements, nil
}
func TestOperationRequirements(testName string) ([]OperationRequirement, error) {
	fixtures, ok := acceptanceTestFixtures[testName]
	if !ok {
		return nil, fmt.Errorf("test fixture inventory missing")
	}
	result := []OperationRequirement{}
	seen := map[string]bool{}
	for _, name := range fixtures {
		config, ok := acceptanceFixtures[name]
		if !ok {
			return nil, fmt.Errorf("fixture inventory missing")
		}
		requirements, err := FixtureRequirements(config())
		if err != nil {
			return nil, err
		}
		for _, r := range requirements {
			key := r.Method + " " + r.Path
			if !seen[key] {
				seen[key] = true
				result = append(result, r)
			}
		}
	}
	return result, nil
}

func AuditAcceptanceFixtures(path string) ([]FixtureDiagnostic, error) {
	schema, err := LoadFixtureSchema(path)
	if err != nil {
		return nil, err
	}
	result := []FixtureDiagnostic{}
	names := []string{}
	for name := range acceptanceFixtures {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		result = append(result, schema.Validate(name, acceptanceFixtures[name]())...)
	}
	return result, nil
}
