---
page_title: "filter_fields.filter_expression_field"
subcategory: ""
description: "Filter Expression Field."
xcsh_docs: {"aliases": ["filter fields filter expression field"], "body_bytes": 1960, "body_sha256": "sha256:8aeca2fd52c93048c2dca5af482f1d8c254a288f5c8875077b39e601e5b42bf0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "parent_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "path": "documentation/resources/filter_set/properties/filter_fields/filter_expression_field/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1013212232102003-0023331222000202-3103321201011201-1312112112231332-1102300321031130-3222323231013203-3222002300201330-2013102110132122", "registry_path": "docs/guides/resources--filter_set--reference--group-001.md", "relationships": [{"anchor": "schema-filter_fields--filter_expression_field--expression", "enforcement": "provider-schema", "group": "filter_fields.filter_expression_field:RequiredObjectAttributes:expression", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["filter_fields", "filter_expression_field"], "schema_version": 1, "sections": [{"aliases": ["filter fields filter expression field expression"], "anchor": "schema-filter_fields--filter_expression_field--expression", "description": "Expression is a Kubernetes style label expression for selections, but differs in that it allows special characters in the keys and values.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "filter_expression_field", "expression"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/filter_expression_field/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Filter Expression Field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["filter_setCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields.filter_expression_field

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/)
- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/)
- filter_fields.filter_expression_field

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Filter Expression Field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expression")}
```

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
filter_expression_field {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-filter_fields--filter_expression_field--expression"></a>

### expression property

Type: `"string"`. Optional.

Expression is a Kubernetes style label expression for selections, but differs in that it allows
special characters in the keys and values.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-17T10:10:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
