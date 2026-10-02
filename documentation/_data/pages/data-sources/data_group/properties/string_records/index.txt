---
page_title: "string_records"
subcategory: ""
description: "Data group with strings record List."
xcsh_docs: {"aliases": ["string records"], "body_bytes": 1657, "body_sha256": "sha256:95f8fa0a7cb886a663f5d07b8ee6cd2ee3005e93509835ea3e94e90a637c980e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_group:properties:string_records", "parent_id": "xcsh-docs:data-sources:data_group:reference", "path": "documentation/data-sources/data_group/properties/string_records/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3000201011020031-0300333202003013-1032212333232010-1130110032033211-3013220313313233-3311123021230233-3301233232230021-1302012323321013", "registry_path": "docs/guides/data-sources--data_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["string_records"], "schema_version": 1, "sections": [{"aliases": ["records"], "anchor": "schema-string_records--records", "description": "Configuration parameter for records", "document_id": "xcsh-docs:data-sources:data_group:properties:string_records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["string_records", "records"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/properties/string_records/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data group with strings record List.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["data_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# string_records

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/properties/)
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/properties/)
- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/)
