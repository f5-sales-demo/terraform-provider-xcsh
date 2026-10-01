---
page_title: "filter_fields.date_field"
subcategory: ""
description: "filter_fields.date_field for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 2118, "body_sha256": "sha256:23b06e11f4d735d24ce9ece4f3423707608ed55ccc83fd7352590a58e9bc5f7a", "child_ids": ["xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field:absolute"], "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field", "parent_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields", "path": "documentation/data-sources/filter_set/properties/filter_fields/date_field/index.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["filter_fields", "date_field"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/properties/filter_fields/date_field/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filter_fields.date_field for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields.date_field

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/)
- filter_fields.date_field

<a id="section"></a>

Type: `"single"`. Computed.

Either an absolute time range or a relative time interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-range_type": "[\"absolute\",\"relative\"]"
}
```

## Direct properties

- [absolute](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/absolute/): complete subsection reference.

<a id="schema-filter_fields--date_field--relative"></a>

### relative property

Type: `"string"`. Computed.

Exclusive with \[absolute\] relative time duration.

Upstream description:

Exclusive with \[absolute\] relative time duration.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [filter_fields.date_field.absolute](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/absolute/)
- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/)
- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
