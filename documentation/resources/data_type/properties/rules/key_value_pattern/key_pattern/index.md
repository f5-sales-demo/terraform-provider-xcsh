---
page_title: "rules.key_value_pattern.key_pattern"
subcategory: ""
description: "Test"
xcsh_docs: {"aliases": ["rules key value pattern key pattern"], "body_bytes": 4578, "body_sha256": "sha256:ae2e82cbf2d542e40f0e44d3c02126df976a302b416cc843c424526f3bfae508", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "parent_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "path": "documentation/resources/data_type/properties/rules/key_value_pattern/key_pattern/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2203231202320212-3131310130323332-0211000021112131-0210002321021002-1133200023101321-1123023102003323-3020012022301122-0300010011021213", "registry_path": "docs/guides/resources--data_type--reference--group-001.md", "relationships": [{"anchor": "schema-rules--key_value_pattern--key_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.key_value_pattern.key_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_value_pattern--key_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.key_value_pattern.key_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_value_pattern--key_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.key_value_pattern.key_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "type": "conflicts"}, {"anchor": "schema-rules--key_value_pattern--key_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.key_value_pattern.key_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.key_value_pattern.key_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.key_value_pattern.key_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern", "key_pattern"], "schema_version": 1, "sections": [{"aliases": ["exact values"], "anchor": "section", "description": "List of exact values to match.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--key_value_pattern--key_pattern--exact_values--exact_values", "enforcement": "provider-schema", "group": "rules.key_value_pattern.key_pattern.exact_values:RequiredObjectAttributes:exact_values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "type": "requires"}], "schema_path": ["rules", "key_value_pattern", "key_pattern", "exact_values"], "syntax": "block", "type": "object"}, {"aliases": ["regex value"], "anchor": "schema-rules--key_value_pattern--key_pattern--regex_value", "description": "Exclusive with Search for values matching this regular expression.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["substring value"], "anchor": "schema-rules--key_value_pattern--key_pattern--substring_value", "description": "Exclusive with Search for values that include this substring.", "document_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern", "substring_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/rules/key_value_pattern/key_pattern/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Test", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["data_typeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern.key_pattern

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/)
- rules.key_value_pattern.key_pattern

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for key pattern.

Upstream description:

Test

Provider validators and defaults (from schema source):

```go
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
key_pattern {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/): complete subsection reference.

<a id="schema-rules--key_value_pattern--key_pattern--regex_value"></a>

### regex_value property

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-rules--key_value_pattern--key_pattern--substring_value"></a>

### substring_value property

Type: `"string"`. Optional.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [rules.key_value_pattern.key_pattern.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
