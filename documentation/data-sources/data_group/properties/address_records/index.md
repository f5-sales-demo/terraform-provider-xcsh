---
page_title: "address_records"
subcategory: ""
description: "Data group with address record List."
xcsh_docs: {"aliases": ["address records"], "body_bytes": 2897, "body_sha256": "sha256:bc86157be8edbaf96dc5db44692409a4eb3428e1eb7dc68fe278b0403be99243", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_group:properties:address_records", "parent_id": "xcsh-docs:data-sources:data_group:reference", "path": "documentation/data-sources/data_group/properties/address_records/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0000231131200300-2003120323302033-2322221110023130-3302222302303321-3302001123203133-2101100110313300-0212130202203331-3021200300120022", "registry_path": "docs/guides/data-sources--data_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["address_records"], "schema_version": 1, "sections": [{"aliases": ["address records records"], "anchor": "schema-address_records--records", "description": "Configuration parameter for records", "document_id": "xcsh-docs:data-sources:data_group:properties:address_records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_records", "records"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/properties/address_records/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data group with address record List.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 4096
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "format": "ip-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.ip": "true",
      "ves.io.schema.rules.map.max_pairs": "4096",
      "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
    },
    "values": {
      "pattern": "^[0-9a-zA-Z._-]*$",
      "type": "string"
    }
  },
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
