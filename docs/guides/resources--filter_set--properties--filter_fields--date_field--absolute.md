---
page_title: "filter_fields.date_field.absolute"
subcategory: ""
description: "filter_fields.date_field.absolute for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 2800, "body_sha256": "sha256:8b02b68c13f3a88235b0ce5e0b3c6fac8b8cf780790d2a292ac3f026144c8dc7", "canonical_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "child_ids": [], "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "parent_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "path": "docs/guides/resources--filter_set--properties--filter_fields--date_field--absolute.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["filter_fields", "date_field", "absolute"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/date_field/absolute/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filter_fields.date_field.absolute for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields.date_field.absolute

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md)
- [Property reference](resources--filter_set--reference.md)
- [filter_fields](resources--filter_set--properties--filter_fields.md)
- [filter_fields.date_field](resources--filter_set--properties--filter_fields--date_field.md)
- filter_fields.date_field.absolute

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Date range is for selecting a date range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("end_date",
    "start_date")}
```

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

Terraform syntax:

```terraform
absolute {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-filter_fields--date_field--absolute--end_date"></a>

### end_date property

Type: `"string"`. Optional.

End Date. Contains end date.

Upstream description:

Contains end date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

<a id="schema-filter_fields--date_field--absolute--start_date"></a>

### start_date property

Type: `"string"`. Optional.

Start Date. Contains start date.

Upstream description:

Contains start date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

## Next pages

- [filter_fields.date_field](resources--filter_set--properties--filter_fields--date_field.md)
- [xcsh_filter_set](../resources/filter_set.md)
