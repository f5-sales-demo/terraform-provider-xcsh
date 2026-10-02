---
page_title: "address_records"
subcategory: ""
description: "Data group with address record List."
xcsh_docs: {"aliases": ["address records"], "body_bytes": 2356, "body_sha256": "sha256:664ae7b8c45c9894bd8766e2944efd4852bcb8f7c578eae78e3bf1ebac85c796", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_group:properties:address_records", "parent_id": "xcsh-docs:data-sources:data_group:reference", "path": "documentation/data-sources/data_group/properties/address_records/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0000231131200300-2003120323302033-2322221110023130-3302222302303321-3302001123203133-2101100110313300-0212130202203331-3021200300120022", "registry_path": "docs/guides/data-sources--data_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["address_records"], "schema_version": 1, "sections": [{"aliases": ["records"], "anchor": "schema-address_records--records", "description": "Configuration parameter for records", "document_id": "xcsh-docs:data-sources:data_group:properties:address_records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_records", "records"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/properties/address_records/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data group with address record List.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# address_records

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/properties/)
- address_records

<a id="section"></a>

Type: `"single"`. Computed.

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

- [address_records](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/properties/address_records/#section)
- [integer_records](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/properties/integer_records/#section)
- [string_records](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/properties/string_records/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-address_records--records"></a>

### records property

Type: `["map", "string"]`. Computed.

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/properties/)
- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/)
