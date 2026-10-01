---
page_title: "filter_fields"
subcategory: ""
description: "filter_fields for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 3409, "body_sha256": "sha256:371632aa280eb11dec7f55fb74147054d4fd3ea71cb13827167c6f5812e8bc88", "child_ids": ["xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field", "xcsh-docs:data-sources:filter_set:properties:filter_fields:filter_expression_field", "xcsh-docs:data-sources:filter_set:properties:filter_fields:string_field"], "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:filter_set:properties:filter_fields", "parent_id": "xcsh-docs:data-sources:filter_set:reference", "path": "documentation/data-sources/filter_set/properties/filter_fields/index.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["filter_fields"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/properties/filter_fields/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filter_fields for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- filter_fields

<a id="section"></a>

Type: `"list"`. Computed.

List of fields and their values selected by the user.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [date_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/): complete subsection reference.

<a id="schema-filter_fields--field_id"></a>

### field_id property

Type: `"string"`. Computed.

Identifier for the field that maps to some UI filter component.

Upstream description:

An identifier for the field that maps to some UI filter component.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [filter_expression_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/filter_expression_field/): complete subsection reference.

- [string_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/string_field/): complete subsection reference.

## Next pages

- [filter_fields.date_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/)
- [filter_fields.filter_expression_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/filter_expression_field/)
- [filter_fields.string_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/string_field/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
