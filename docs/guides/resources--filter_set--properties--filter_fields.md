---
page_title: "filter_fields"
subcategory: ""
description: "filter_fields for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 3307, "body_sha256": "sha256:06f9354c6883476d4dd0b0a1ebe2707f9f76bf9e52d85ae80b71f4d56f9e7c31", "canonical_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "child_ids": ["xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "xcsh-docs:resources:filter_set:properties:filter_fields:string_field"], "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields", "parent_id": "xcsh-docs:resources:filter_set:reference", "path": "docs/guides/resources--filter_set--properties--filter_fields.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["filter_fields"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filter_fields for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# filter_fields

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md)
- [Property reference](resources--filter_set--reference.md)
- filter_fields

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of fields and their values selected by the user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("field_id"),
  validators.ConflictingListObjectAttributes("date_field",
    "filter_expression_field"),
  validators.ConflictingListObjectAttributes("date_field",
    "string_field"),
  validators.ConflictingListObjectAttributes("filter_expression_field",
    "string_field")}
```

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

Terraform syntax:

```terraform
filter_fields {
  # Configure direct properties listed below.
}
```

## Direct properties

- [date_field](resources--filter_set--properties--filter_fields--date_field.md): complete subsection reference.

<a id="schema-filter_fields--field_id"></a>

### field_id property

Type: `"string"`. Optional.

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

- [filter_expression_field](resources--filter_set--properties--filter_fields--filter_expression_field.md): complete subsection reference.

- [string_field](resources--filter_set--properties--filter_fields--string_field.md): complete subsection reference.

## Next pages

- [filter_fields.date_field](resources--filter_set--properties--filter_fields--date_field.md)
- [filter_fields.filter_expression_field](resources--filter_set--properties--filter_fields--filter_expression_field.md)
- [filter_fields.string_field](resources--filter_set--properties--filter_fields--string_field.md)
- [Property reference](resources--filter_set--reference.md)
- [xcsh_filter_set](../resources/filter_set.md)
