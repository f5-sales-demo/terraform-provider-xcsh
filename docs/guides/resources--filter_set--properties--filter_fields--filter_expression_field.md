---
page_title: "filter_fields.filter_expression_field"
subcategory: ""
description: "filter_fields.filter_expression_field for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 1974, "body_sha256": "sha256:7dcd63e90602f8ede500604eb3a11a90aad062d62be90a941f3d594d9fd37e86", "canonical_id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "child_ids": [], "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "parent_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "path": "docs/guides/resources--filter_set--properties--filter_fields--filter_expression_field.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["filter_fields", "filter_expression_field"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/filter_expression_field/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filter_fields.filter_expression_field for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# filter_fields.filter_expression_field

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md)
- [Property reference](resources--filter_set--reference.md)
- [filter_fields](resources--filter_set--properties--filter_fields.md)
- filter_fields.filter_expression_field

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Filter Expression Field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expression")}
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
filter_expression_field {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-filter_fields--filter_expression_field--expression"></a>

### expression property

Type: `"string"`. Optional.

Expression is a Kubernetes style label expression for selections, but differs in that it allows
special characters in the keys and values.

Upstream description:

Expression is a Kubernetes style label expression for selections, but differs in that it allows
special characters in the keys and values.

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

## Next pages

- [filter_fields](resources--filter_set--properties--filter_fields.md)
- [xcsh_filter_set](../resources/filter_set.md)
