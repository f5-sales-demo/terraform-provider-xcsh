---
page_title: "integer_records"
subcategory: ""
description: "integer_records for xcsh_data_group."
xcsh_docs: {"aliases": [], "body_bytes": 1580, "body_sha256": "sha256:88580a63e08accd3dbff80effb58174352e9c97536e0410c6f624e060a7cff76", "canonical_id": "xcsh-docs:resources:data_group:properties:integer_records", "child_ids": [], "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_group:properties:integer_records", "parent_id": "xcsh-docs:resources:data_group:reference", "path": "docs/guides/resources--data_group--properties--integer_records.md", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["integer_records"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/properties/integer_records/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "integer_records for xcsh_data_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# integer_records

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md)
- [Property reference](resources--data_group--reference.md)
- integer_records

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for integer records.

Upstream description:

Data group with integer record List.

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
integer_records {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-integer_records--records"></a>

### records property

Type: `["map", "string"]`. Optional.

Integer records. Configuration parameter for records

Upstream description:

Configuration parameter for records

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  }
}
```

## Next pages

- [Property reference](resources--data_group--reference.md)
- [xcsh_data_group](../resources/data_group.md)
