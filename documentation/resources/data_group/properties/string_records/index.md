---
page_title: "string_records"
subcategory: ""
description: "string_records for xcsh_data_group."
xcsh_docs: {"aliases": [], "body_bytes": 1768, "body_sha256": "sha256:89c941ef7509e606998e549efe34bbfe0311ba24c4bc5a51bf85570eeb4851d1", "child_ids": [], "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_group:properties:string_records", "parent_id": "xcsh-docs:resources:data_group:reference", "path": "documentation/resources/data_group/properties/string_records/index.md", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["string_records"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/properties/string_records/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "string_records for xcsh_data_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# string_records

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/)
- string_records

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for string records.

Upstream description:

Data group with strings record List.

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
string_records {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-string_records--records"></a>

### records property

Type: `["map", "string"]`. Optional.

String records. Configuration parameter for records

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
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/)
- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
