---
page_title: "string_records"
subcategory: ""
description: "string_records for xcsh_data_group."
xcsh_docs: {"aliases": [], "body_bytes": 1449, "body_sha256": "sha256:85a84981514ea3c15855a889eee75be3b170a89e7ad1ab73637a744d9b241f52", "canonical_id": "xcsh-docs:data-sources:data_group:properties:string_records", "child_ids": [], "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_group:properties:string_records", "parent_id": "xcsh-docs:data-sources:data_group:reference", "path": "docs/guides/data-sources--data_group--properties--string_records.md", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["string_records"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/properties/string_records/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "string_records for xcsh_data_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# string_records

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md)
- [Property reference](data-sources--data_group--reference.md)
- string_records

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-string_records--records"></a>

### records property

Type: `["map", "string"]`. Computed.

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

- [Property reference](data-sources--data_group--reference.md)
- [xcsh_data_group](../data-sources/data_group.md)
