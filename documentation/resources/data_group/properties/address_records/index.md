---
page_title: "address_records"
subcategory: ""
description: "Data group with address record List."
xcsh_docs: {"aliases": ["address records"], "body_bytes": 3250, "body_sha256": "sha256:4a5325e9db82685d0f89955d02bfd129e29cbf25974eac3ff57c7e7caa8e1a18", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_group:properties:address_records", "parent_id": "xcsh-docs:resources:data_group:reference", "path": "documentation/resources/data_group/properties/address_records/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1200023100312100-3032303031001130-0022031202332330-3312333132103031-1322222120203122-1213021011122011-0003021001311301-0012003310211130", "registry_path": "docs/guides/resources--data_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["address_records"], "schema_version": 1, "sections": [{"aliases": ["address records records"], "anchor": "schema-address_records--records", "description": "Configuration parameter for records", "document_id": "xcsh-docs:resources:data_group:properties:address_records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_records", "records"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/properties/address_records/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Data group with address record List.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["data_groupCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# address_records

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/)
- address_records

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [address_records](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/address_records/#section)
- [integer_records](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/integer_records/#section)
- [string_records](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/string_records/#section)

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":4096},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"format\":\"ip-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.ip\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"4096\",\"ves.io.schema.rules.map.values.string.pattern\":\"^[0-9a-zA-Z._-]*$\"},\"values\":{\"pattern\":\"^[0-9a-zA-Z._-]*$\",\"type\":\"string\"}}")}
```

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
