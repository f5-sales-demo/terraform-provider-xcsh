---
page_title: "rules.key_value_pattern.value_pattern"
subcategory: ""
description: "Test"
xcsh_docs: {"aliases": ["rules key value pattern value pattern"], "body_bytes": 3999, "body_sha256": "sha256:20c67e3efff58fdc076524e65857e582d81dd1d419d73ce7dcfddf956e40e727", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern", "parent_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "path": "documentation/resources/data_type/properties/rules/key_value_pattern/value_pattern/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2131001001031100-0201133233332332-0233303120233132-1233011300212301-0311220031301113-0110122312110133-3030231331120303-0232211111230102", "registry_path": "docs/guides/resources--data_type--reference--group-001.md", "relationships": [{"anchor": "schema-rules--key_value_pattern--value_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.key_value_pattern.value_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_value_pattern--value_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.key_value_pattern.value_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_value_pattern--value_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.key_value_pattern.value_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_value_pattern--value_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.key_value_pattern.value_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.key_value_pattern.value_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.key_value_pattern.value_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern", "value_pattern"], "schema_version": 1, "sections": [{"aliases": ["rules key value pattern value pattern exact values"], "anchor": "section", "description": "List of exact values to match.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--key_value_pattern--value_pattern--exact_values--exact_values", "enforcement": "provider-schema", "group": "rules.key_value_pattern.value_pattern.exact_values:RequiredObjectAttributes:exact_values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "type": "requires"}], "schema_path": ["rules", "key_value_pattern", "value_pattern", "exact_values"], "syntax": "block", "type": "object"}, {"aliases": ["rules key value pattern value pattern regex value"], "anchor": "schema-rules--key_value_pattern--value_pattern--regex_value", "description": "Exclusive with Search for values matching this regular expression.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "value_pattern", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules key value pattern value pattern substring value"], "anchor": "schema-rules--key_value_pattern--value_pattern--substring_value", "description": "Exclusive with Search for values that include this substring.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:value_pattern", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "value_pattern", "substring_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/rules/key_value_pattern/value_pattern/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Test", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["data_typeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern.value_pattern

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/)
- rules.key_value_pattern.value_pattern

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for value pattern.

Additional upstream details:

Test

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
value_pattern {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/value_pattern/exact_values/): complete subsection reference.

<a id="schema-rules--key_value_pattern--value_pattern--regex_value"></a>

### regex_value property

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-rules--key_value_pattern--value_pattern--substring_value"></a>

### substring_value property

Type: `"string"`. Optional.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```
