---
page_title: "rules.key_pattern"
subcategory: ""
description: "Test"
xcsh_docs: {"aliases": ["rules key pattern"], "body_bytes": 2944, "body_sha256": "sha256:24c35026ffa210c8c3dcece168e85586d642eb63454d43fe03afd44fff55b097", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:data_type:properties:rules:key_pattern:exact_values"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:properties:rules:key_pattern", "parent_id": "xcsh-docs:data-sources:data_type:properties:rules", "path": "documentation/data-sources/data_type/properties/rules/key_pattern/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0121320013020002-2121012313103223-3122202112022112-1131200013020113-3120022101203010-1120211322330200-2230300213212102-1301120233130213", "registry_path": "docs/guides/data-sources--data_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_pattern"], "schema_version": 1, "sections": [{"aliases": ["rules key pattern exact values"], "anchor": "section", "description": "List of exact values to match.", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_pattern:exact_values", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "key_pattern", "exact_values"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules key pattern regex value"], "anchor": "schema-rules--key_pattern--regex_value", "description": "Exclusive with Search for values matching this regular expression.", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_pattern", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules key pattern substring value"], "anchor": "schema-rules--key_pattern--substring_value", "description": "Exclusive with Search for values that include this substring.", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_pattern", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_pattern", "substring_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/properties/rules/key_pattern/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Test", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["data_typeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_pattern

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/)
- rules.key_pattern

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for key pattern.

Additional upstream details:

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

- [exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_pattern/exact_values/): complete subsection reference.

<a id="schema-rules--key_pattern--regex_value"></a>

### regex_value property

Type: `"string"`. Computed.

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

<a id="schema-rules--key_pattern--substring_value"></a>

### substring_value property

Type: `"string"`. Computed.

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
