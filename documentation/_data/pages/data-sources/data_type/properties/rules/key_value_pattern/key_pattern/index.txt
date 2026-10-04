---
page_title: "rules.key_value_pattern.key_pattern"
subcategory: ""
description: "Test"
xcsh_docs: {"aliases": ["rules key value pattern key pattern"], "body_bytes": 3870, "body_sha256": "sha256:0fee998234a6700e1c0c87d48f19312d380e5a320810660f1c8147eb743de8a4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "parent_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern", "path": "documentation/data-sources/data_type/properties/rules/key_value_pattern/key_pattern/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1111131211022103-1022013203103023-3022022323101233-3112002102001101-3001202222010203-3132000221200102-3013013011200230-1000211123332000", "registry_path": "docs/guides/data-sources--data_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern", "key_pattern"], "schema_version": 1, "sections": [{"aliases": ["rules key value pattern key pattern exact values"], "anchor": "section", "description": "List of exact values to match.", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern", "exact_values"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules key value pattern key pattern regex value"], "anchor": "schema-rules--key_value_pattern--key_pattern--regex_value", "description": "Exclusive with Search for values matching this regular expression.", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules key value pattern key pattern substring value"], "anchor": "schema-rules--key_value_pattern--key_pattern--substring_value", "description": "Exclusive with Search for values that include this substring.", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "key_pattern", "substring_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/properties/rules/key_value_pattern/key_pattern/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Test", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["data_typeCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern.key_pattern

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/)
- rules.key_value_pattern.key_pattern

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for key pattern.

Upstream description:

Test

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

## Direct properties

- [exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/): complete subsection reference.

<a id="schema-rules--key_value_pattern--key_pattern--regex_value"></a>

### regex_value property

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Computed.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [rules.key_value_pattern.key_pattern.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
