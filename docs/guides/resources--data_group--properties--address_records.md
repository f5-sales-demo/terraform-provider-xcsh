---
page_title: "address_records"
subcategory: ""
description: "address_records for xcsh_data_group."
xcsh_docs: {"aliases": [], "body_bytes": 2098, "body_sha256": "sha256:94c09e38632ac789c3e01cfe410d3384f0debb96eb9df5ee503aeb02f87ef968", "canonical_id": "xcsh-docs:resources:data_group:properties:address_records", "child_ids": [], "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_group:properties:address_records", "parent_id": "xcsh-docs:resources:data_group:reference", "path": "docs/guides/resources--data_group--properties--address_records.md", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["address_records"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/properties/address_records/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "address_records for xcsh_data_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# address_records

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md)
- [Property reference](resources--data_group--reference.md)
- address_records

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: address\_records, integer\_records, string\_records\] Address Record. Data group with
address record List.

Upstream description:

Data group with address record List.

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

OneOf alternatives in this subsection:

- [address_records](resources--data_group--properties--address_records.md#section)
- [integer_records](resources--data_group--properties--integer_records.md#section)
- [string_records](resources--data_group--properties--string_records.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
address_records {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-address_records--records"></a>

### records property

Type: `["map", "string"]`. Optional.

Address records. Configuration parameter for records

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
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

## Next pages

- [Property reference](resources--data_group--reference.md)
- [xcsh_data_group](../resources/data_group.md)
