---
page_title: "integer_records"
subcategory: ""
description: "integer_records for xcsh_data_group."
xcsh_docs: {"aliases": [], "body_bytes": 1369, "body_sha256": "sha256:7ae3a64fbef090734057d60dc3ff6184f25679d8588eabf7e877807bce31e3ac", "canonical_id": "xcsh-docs:data-sources:data_group:properties:integer_records", "child_ids": [], "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_group:properties:integer_records", "parent_id": "xcsh-docs:data-sources:data_group:reference", "path": "docs/guides/data-sources--data_group--properties--integer_records.md", "provider_name": "data_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["integer_records"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/properties/integer_records/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "integer_records for xcsh_data_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# integer_records

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md)
- [Property reference](data-sources--data_group--reference.md)
- integer_records

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-integer_records--records"></a>

### records property

Type: `["map", "string"]`. Computed.

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

- [Property reference](data-sources--data_group--reference.md)
- [xcsh_data_group](../data-sources/data_group.md)
