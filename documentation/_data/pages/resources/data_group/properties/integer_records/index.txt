---
page_title: "integer_records"
subcategory: ""
description: "Data group with integer record List."
xcsh_docs: {"aliases": ["integer records"], "body_bytes": 2717, "body_sha256": "sha256:c886a649b9b2f7d8e185262be06e8eaab37290b1c58030135abc33babb665989", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_group:properties:integer_records", "parent_id": "xcsh-docs:resources:data_group:reference", "path": "documentation/resources/data_group/properties/integer_records/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2300211220223031-2002323232103033-1232110333132002-2103220023013000-3201132301003203-0033320000330020-1200101010000223-2113112000301222", "registry_path": "docs/guides/resources--data_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["integer_records"], "schema_version": 1, "sections": [{"aliases": ["integer records records"], "anchor": "schema-integer_records--records", "description": "Configuration parameter for records", "document_id": "xcsh-docs:resources:data_group:properties:integer_records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["integer_records", "records"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/properties/integer_records/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data group with integer record List.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# integer_records

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/)
- integer_records

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
integer_records {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-integer_records--records"></a>

### records property

Type: `["map", "string"]`. Optional.

Integer records. Configuration parameter for records

Upstream description:

Configuration parameter for records

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":4096},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"originalRules\":{\"ves.io.schema.rules.map.max_pairs\":\"4096\",\"ves.io.schema.rules.map.values.string.pattern\":\"^[a-zA-Z_][a-zA-Z0-9_]*$\"},\"values\":{\"pattern\":\"^[a-zA-Z_][a-zA-Z0-9_]*$\",\"type\":\"string\"}}")}
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
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "4096",
      "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
    },
    "values": {
      "pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/properties/)
- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
