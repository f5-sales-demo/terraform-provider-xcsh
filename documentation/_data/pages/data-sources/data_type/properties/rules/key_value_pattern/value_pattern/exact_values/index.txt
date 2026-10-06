---
page_title: "rules.key_value_pattern.value_pattern.exact_values"
subcategory: ""
description: "List of exact values to match."
xcsh_docs: {"aliases": ["rules key value pattern value pattern exact values"], "body_bytes": 2261, "body_sha256": "sha256:73607b22a7c66e93465eec0680524c7d3b172bd6f9e4740d8818b464fe22eb17", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "parent_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern", "path": "documentation/data-sources/data_type/properties/rules/key_value_pattern/value_pattern/exact_values/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1303221331013302-1230301312031222-3003223303000031-2110222020203030-3332201130232200-2203133213011302-2301332203232200-0333230231221012", "registry_path": "docs/guides/data-sources--data_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "key_value_pattern", "value_pattern", "exact_values"], "schema_version": 1, "sections": [{"aliases": ["rules key value pattern value pattern exact values exact values"], "anchor": "schema-rules--key_value_pattern--value_pattern--exact_values--exact_values", "description": "List of exact values to match.", "document_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:value_pattern:exact_values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "key_value_pattern", "value_pattern", "exact_values", "exact_values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/properties/rules/key_value_pattern/value_pattern/exact_values/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of exact values to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern.value_pattern.exact_values

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/)
- [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/)
- [rules.key_value_pattern.value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/rules/key_value_pattern/value_pattern/)
- rules.key_value_pattern.value_pattern.exact_values

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Additional upstream details:

List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-rules--key_value_pattern--value_pattern--exact_values--exact_values"></a>

### exact_values property

Type: `["list", "string"]`. Computed.

Exact Values. List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
