---
page_title: "filter_fields.string_field"
subcategory: ""
description: "filter_fields.string_field for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 1515, "body_sha256": "sha256:50471f5d0357e3a1cd1c8f4900ce3d93132a7375fd2b55b80760888115f679fb", "canonical_id": "xcsh-docs:resources:filter_set:properties:filter_fields:string_field", "child_ids": [], "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields:string_field", "parent_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "path": "docs/guides/resources--filter_set--properties--filter_fields--string_field.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["filter_fields", "string_field"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/string_field/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filter_fields.string_field for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# filter_fields.string_field

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md)
- [Property reference](resources--filter_set--reference.md)
- [filter_fields](resources--filter_set--properties--filter_fields.md)
- filter_fields.string_field

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Filter String Field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("field_values")}
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
string_field {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-filter_fields--string_field--field_values"></a>

### field_values property

Type: `["list", "string"]`. Optional.

String Value(s). Field specification or configuration

Upstream description:

Field specification or configuration

Receipt-pinned upstream constraints:

```json
{
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
