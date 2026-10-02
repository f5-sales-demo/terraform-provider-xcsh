---
page_title: "address_records"
subcategory: ""
description: "Data group with address record List."
xcsh_docs: {"aliases": ["address records"], "body_bytes": 2459, "body_sha256": "sha256:1b03439ced17359482de1ba9c32da9843e18cf20bfb6e0e7aa1b1cc9858e8e76", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_group:properties:address_records", "parent_id": "xcsh-docs:resources:data_group:reference", "path": "documentation/resources/data_group/properties/address_records/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1200023100312100-3032303031001130-0022031202332330-3312333132103031-1322222120203122-1213021011122011-0003021001311301-0012003310211130", "registry_path": "docs/guides/resources--data_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["address_records"], "schema_version": 1, "sections": [{"aliases": ["records"], "anchor": "schema-address_records--records", "description": "Configuration parameter for records", "document_id": "xcsh-docs:resources:data_group:properties:address_records", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address_records", "records"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/properties/address_records/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data group with address record List.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["data_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/)
- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
