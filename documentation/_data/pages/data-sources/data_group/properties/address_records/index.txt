---
page_title: "address_records"
subcategory: ""
description: "Data group with address record List."
xcsh_docs: {"aliases": ["address records"], "body_bytes": 2542, "body_sha256": "sha256:b8fd31b575b4f512b9173966d37a6fd5fe78d5533fc12cecddbdaf004cbd2d79", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_group:properties:address_records", "parent_id": "xcsh-docs:data-sources:data_group:reference", "path": "documentation/data-sources/data_group/properties/address_records/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0000231131200300-2003120323302033-2322221110023130-3302222302303321-3302001123203133-2101100110313300-0212130202203331-3021200300120022", "registry_path": "docs/guides/data-sources--data_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["address_records"], "schema_version": 1, "sections": [{"aliases": ["address records records"], "anchor": "schema-address_records--records", "description": "Configuration parameter for records", "document_id": "xcsh-docs:data-sources:data_group:properties:address_records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_records", "records"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/properties/address_records/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data group with address record List.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["data_groupCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
