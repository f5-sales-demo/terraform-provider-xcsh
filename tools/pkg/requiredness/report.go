// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package requiredness produces a deterministic audit of generated Terraform
// Required-to-Optional transitions across two provider/spec releases.
package requiredness

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	schemapkg "github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/schema"
)

const (
	RemovedMinimumConfigurationPromotion = "minimum_configuration_promotion_removed"
	VerifiedSingleNamespaceDefault       = "verified_single_namespace_default"
)

type Requirement string

const (
	Required Requirement = "required"
	Optional Requirement = "optional"
)

type Transition struct {
	Resource  string `json:"resource"`
	Attribute string `json:"attribute"`
	Reason    string `json:"reason"`
}

type Report struct {
	BaselineProvider string       `json:"baseline_provider"`
	BaselineSpec     string       `json:"baseline_spec"`
	CandidateSpec    string       `json:"candidate_spec"`
	Transitions      []Transition `json:"transitions"`
}

type snapshot map[string]map[string]Requirement

func boolField(lit *ast.CompositeLit, name string) bool {
	for _, elt := range lit.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		ident, ok := kv.Key.(*ast.Ident)
		if !ok || ident.Name != name {
			continue
		}
		value, ok := kv.Value.(*ast.Ident)
		return ok && value.Name == "true"
	}
	return false
}

func schemaAttributes(file *ast.File) (map[string]Requirement, error) {
	var result map[string]Requirement
	ast.Inspect(file, func(node ast.Node) bool {
		if result != nil {
			return false
		}
		lit, ok := node.(*ast.CompositeLit)
		if !ok {
			return true
		}
		sel, ok := lit.Type.(*ast.SelectorExpr)
		if !ok || sel.Sel.Name != "Schema" {
			return true
		}
		pkg, ok := sel.X.(*ast.Ident)
		if !ok || pkg.Name != "schema" {
			return true
		}
		for _, elt := range lit.Elts {
			kv, ok := elt.(*ast.KeyValueExpr)
			if !ok {
				continue
			}
			key, ok := kv.Key.(*ast.Ident)
			if !ok || key.Name != "Attributes" {
				continue
			}
			attributes, ok := kv.Value.(*ast.CompositeLit)
			if !ok {
				continue
			}
			result = make(map[string]Requirement)
			for _, attrElt := range attributes.Elts {
				attrKV, ok := attrElt.(*ast.KeyValueExpr)
				if !ok {
					continue
				}
				nameLit, ok := attrKV.Key.(*ast.BasicLit)
				if !ok || nameLit.Kind != token.STRING {
					continue
				}
				name, err := strconv.Unquote(nameLit.Value)
				if err != nil {
					continue
				}
				attrLit, ok := attrKV.Value.(*ast.CompositeLit)
				if !ok {
					continue
				}
				switch {
				case boolField(attrLit, "Required"):
					result[name] = Required
				case boolField(attrLit, "Optional"):
					result[name] = Optional
				}
			}
			return false
		}
		return true
	})
	if result == nil {
		return nil, fmt.Errorf("generated resource has no schema.Schema Attributes map")
	}
	return result, nil
}

func snapshotProvider(providerDir string) (snapshot, error) {
	files, err := filepath.Glob(filepath.Join(providerDir, "*_resource.go"))
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	result := make(snapshot, len(files))
	for _, path := range files {
		parsed, parseErr := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if parseErr != nil {
			return nil, fmt.Errorf("parsing %s: %w", path, parseErr)
		}
		attrs, attrErr := schemaAttributes(parsed)
		if attrErr != nil {
			// Helper and action resource files can share the suffix without exposing
			// a normal schema map; they do not participate in this audit.
			continue
		}
		resource := strings.TrimSuffix(filepath.Base(path), "_resource.go")
		result[resource] = attrs
	}
	return result, nil
}

func loadSchemas(specDir string) (map[string]openapi.Schema, error) {
	files, err := filepath.Glob(filepath.Join(specDir, "domains", "*.json"))
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no domain specs in %s", specDir)
	}
	sort.Strings(files)
	result := make(map[string]openapi.Schema)
	for _, path := range files {
		spec, parseErr := openapi.ParseFile(path)
		if parseErr != nil {
			return nil, fmt.Errorf("parsing %s: %w", path, parseErr)
		}
		for name, schema := range spec.Components.Schemas {
			result[name] = schema
		}
	}
	return result, nil
}

func findCreateSchema(schemas map[string]openapi.Schema, resource string) (openapi.Schema, error) {
	resolved, _, found, err := schemapkg.ResolveEnvelopeSchemaFromSchemas(schemas, resource, "CreateSpecType")
	if err != nil {
		return openapi.Schema{}, err
	}
	if !found {
		return openapi.Schema{}, fmt.Errorf("resource %s has no CreateSpecType schema", resource)
	}
	return resolved, nil
}

func terraformName(property string) string {
	name := naming.ToSnakeCaseTerraform(property)
	switch strings.ToLower(property) {
	case "description":
		return "description_spec"
	case "disable":
		return "disable_spec"
	default:
		return name
	}
}

func propertyByTerraformName(schema openapi.Schema, attribute string) (openapi.Schema, bool) {
	for property, value := range schema.Properties {
		if terraformName(property) == attribute {
			return value, true
		}
	}
	return openapi.Schema{}, false
}

func minimumConfigurationContains(schema openapi.Schema, attribute string) bool {
	for _, field := range schemapkg.ParseMinConfigRequiredFields(schema.XF5XCMinimumConfiguration) {
		// The removed promotion compared its raw, spec.-stripped field directly
		// with TfsdkTag. Preserve that exact historical behavior in the audit,
		// including the old description/description_spec collision it exposed.
		if field == attribute {
			return true
		}
	}
	return false
}

func explicitlyRequired(schema openapi.Schema, attribute string) bool {
	property, ok := propertyByTerraformName(schema, attribute)
	if !ok {
		return false
	}
	for _, required := range schema.Required {
		if terraformName(required) == attribute {
			return true
		}
	}
	return property.XVesRequired == "true" || property.XF5XCRequiredFor.Create
}

func Compare(baselineProviderDir, candidateProviderDir, baselineSpecDir, candidateSpecDir string, metadata Report) (Report, error) {
	baseline, err := snapshotProvider(baselineProviderDir)
	if err != nil {
		return Report{}, err
	}
	candidate, err := snapshotProvider(candidateProviderDir)
	if err != nil {
		return Report{}, err
	}
	baselineSchemas, err := loadSchemas(baselineSpecDir)
	if err != nil {
		return Report{}, err
	}
	candidateSchemas, err := loadSchemas(candidateSpecDir)
	if err != nil {
		return Report{}, err
	}

	resources := make([]string, 0, len(baseline))
	for resource := range baseline {
		resources = append(resources, resource)
	}
	sort.Strings(resources)
	transitions := make([]Transition, 0)
	for _, resource := range resources {
		candidateAttrs, ok := candidate[resource]
		if !ok {
			continue
		}
		attributes := make([]string, 0, len(baseline[resource]))
		for attribute := range baseline[resource] {
			attributes = append(attributes, attribute)
		}
		sort.Strings(attributes)
		for _, attribute := range attributes {
			if baseline[resource][attribute] != Required || candidateAttrs[attribute] != Optional {
				continue
			}
			baselineCreate, createErr := findCreateSchema(baselineSchemas, resource)
			if createErr != nil {
				return Report{}, createErr
			}
			candidateCreate, createErr := findCreateSchema(candidateSchemas, resource)
			if createErr != nil {
				return Report{}, createErr
			}
			if attribute == "namespace" && candidateCreate.XF5XCNamespaceProfile != nil &&
				candidateCreate.XF5XCNamespaceProfile.Constraint != nil &&
				candidateCreate.XF5XCNamespaceProfile.Constraint.Enforced &&
				len(candidateCreate.XF5XCNamespaceProfile.Constraint.Allowed) == 1 {
				transitions = append(transitions, Transition{
					Resource: resource, Attribute: attribute, Reason: VerifiedSingleNamespaceDefault,
				})
				continue
			}
			if !minimumConfigurationContains(baselineCreate, attribute) {
				return Report{}, fmt.Errorf("unexpected Required-to-Optional transition %s.%s: absent from baseline minimum configuration", resource, attribute)
			}
			if explicitlyRequired(baselineCreate, attribute) {
				return Report{}, fmt.Errorf("invalid Required-to-Optional transition %s.%s: baseline contract genuinely required it for create", resource, attribute)
			}
			if explicitlyRequired(candidateCreate, attribute) {
				return Report{}, fmt.Errorf("invalid Required-to-Optional transition %s.%s: candidate contract still requires it for create", resource, attribute)
			}
			transitions = append(transitions, Transition{Resource: resource, Attribute: attribute, Reason: RemovedMinimumConfigurationPromotion})
		}
	}
	metadata.Transitions = transitions
	return metadata, nil
}

func Write(path string, report Report) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o644)
}
