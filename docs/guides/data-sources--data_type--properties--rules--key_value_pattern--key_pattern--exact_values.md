---
page_title: "rules.key_value_pattern.key_pattern.exact_values"
subcategory: ""
description: "rules.key_value_pattern.key_pattern.exact_values for xcsh_data_type."
xcsh_docs: {"aliases": [], "body_bytes": 2229, "body_sha256": "sha256:91e6f21f94fdc3f757c2874f0633f0c2662b5e3f1205574d69be120dda2739fa", "canonical_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "child_ids": [], "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values", "parent_id": "xcsh-docs:data-sources:data_type:properties:rules:key_value_pattern:key_pattern", "path": "docs/guides/data-sources--data_type--properties--rules--key_value_pattern--key_pattern--exact_values.md", "provider_name": "data_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "key_value_pattern", "key_pattern", "exact_values"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.key_value_pattern.key_pattern.exact_values for xcsh_data_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern.key_pattern.exact_values

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md)
- [Property reference](data-sources--data_type--reference.md)
- [rules](data-sources--data_type--properties--rules.md)
- [rules.key_value_pattern](data-sources--data_type--properties--rules--key_value_pattern.md)
- [rules.key_value_pattern.key_pattern](data-sources--data_type--properties--rules--key_value_pattern--key_pattern.md)
- rules.key_value_pattern.key_pattern.exact_values

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

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

<a id="schema-rules--key_value_pattern--key_pattern--exact_values--exact_values"></a>

### exact_values property

Type: `["list", "string"]`. Computed.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [rules.key_value_pattern.key_pattern](data-sources--data_type--properties--rules--key_value_pattern--key_pattern.md)
- [xcsh_data_type](../data-sources/data_type.md)
