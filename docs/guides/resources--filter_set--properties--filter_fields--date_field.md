---
page_title: "filter_fields.date_field"
subcategory: ""
description: "filter_fields.date_field for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 2027, "body_sha256": "sha256:f25244148eb74f3e963b282852362341fe3cc16fe68b1541416570743de26dcc", "canonical_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "child_ids": ["xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute"], "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "parent_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "path": "docs/guides/resources--filter_set--properties--filter_fields--date_field.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["filter_fields", "date_field"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/date_field/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filter_fields.date_field for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields.date_field

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md)
- [Property reference](resources--filter_set--reference.md)
- [filter_fields](resources--filter_set--properties--filter_fields.md)
- filter_fields.date_field

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Either an absolute time range or a relative time interval.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("absolute",
    "relative")}
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
  "x-ves-oneof-field-range_type": "[\"absolute\",\"relative\"]"
}
```

Terraform syntax:

```terraform
date_field {
  # Configure direct properties listed below.
}
```

## Direct properties

- [absolute](resources--filter_set--properties--filter_fields--date_field--absolute.md): complete subsection reference.

<a id="schema-filter_fields--date_field--relative"></a>

### relative property

Type: `"string"`. Optional.

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

- [filter_fields.date_field.absolute](resources--filter_set--properties--filter_fields--date_field--absolute.md)
- [filter_fields](resources--filter_set--properties--filter_fields.md)
- [xcsh_filter_set](../resources/filter_set.md)
