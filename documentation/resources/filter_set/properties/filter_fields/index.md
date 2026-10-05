---
page_title: "filter_fields"
subcategory: ""
description: "List of fields and their values selected by the user."
xcsh_docs: {"aliases": ["filter fields"], "body_bytes": 3914, "body_sha256": "sha256:b18ec3b13725855656ba6fc53b0e2787be850e915dd30e8d08d92be6c4e1d304", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "xcsh-docs:resources:filter_set:properties:filter_fields:string_field"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields", "parent_id": "xcsh-docs:resources:filter_set:reference", "path": "documentation/resources/filter_set/properties/filter_fields/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1202233312203020-2310221132102333-1202000021102222-0221311311322020-1220222202313033-2021111203221223-0120231120131101-2201031233202332", "registry_path": "docs/guides/resources--filter_set--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields:ConflictingListObjectAttributes:date_field,filter_expression_field", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields:ConflictingListObjectAttributes:date_field,string_field", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields:ConflictingListObjectAttributes:date_field,filter_expression_field", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields:ConflictingListObjectAttributes:filter_expression_field,string_field", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields:ConflictingListObjectAttributes:date_field,string_field", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:string_field", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields:ConflictingListObjectAttributes:filter_expression_field,string_field", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:string_field", "type": "conflicts"}, {"anchor": "schema-filter_fields--field_id", "enforcement": "provider-schema", "group": "filter_fields:RequiredListObjectAttributes:field_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["filter_fields"], "schema_version": 1, "sections": [{"aliases": ["filter fields date field"], "anchor": "section", "description": "Either an absolute time range or a relative time interval.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-filter_fields--date_field--relative", "enforcement": "provider-schema", "group": "filter_fields.date_field:ConflictingObjectAttributes:absolute,relative", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields.date_field:ConflictingObjectAttributes:absolute,relative", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "conflicts"}], "schema_path": ["filter_fields", "date_field"], "syntax": "block", "type": "object"}, {"aliases": ["filter fields field id"], "anchor": "schema-filter_fields--field_id", "description": "An identifier for the field that maps to some UI filter component.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "field_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["filter fields filter expression field"], "anchor": "section", "description": "Filter Expression Field.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-filter_fields--filter_expression_field--expression", "enforcement": "provider-schema", "group": "filter_fields.filter_expression_field:RequiredObjectAttributes:expression", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:filter_expression_field", "type": "requires"}], "schema_path": ["filter_fields", "filter_expression_field"], "syntax": "block", "type": "object"}, {"aliases": ["filter fields string field"], "anchor": "section", "description": "Filter String Field.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:string_field", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-filter_fields--string_field--field_values", "enforcement": "provider-schema", "group": "filter_fields.string_field:RequiredObjectAttributes:field_values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:string_field", "type": "requires"}], "schema_path": ["filter_fields", "string_field"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of fields and their values selected by the user.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/)
- filter_fields

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of fields and their values selected by the user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("field_id"),
  validators.ConflictingListObjectAttributes("date_field",
    "filter_expression_field"),
  validators.ConflictingListObjectAttributes("date_field",
    "string_field"),
  validators.ConflictingListObjectAttributes("filter_expression_field",
    "string_field")}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
filter_fields {
  # Configure direct properties listed below.
}
```

## Direct properties

- [date_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/date_field/): complete subsection reference.

<a id="schema-filter_fields--field_id"></a>

### field_id property

Type: `"string"`. Optional.

Identifier for the field that maps to some UI filter component.

Upstream description:

An identifier for the field that maps to some UI filter component.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [filter_expression_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/filter_expression_field/): complete subsection reference.

- [string_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/string_field/): complete subsection reference.

## Next pages

- [filter_fields.date_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/date_field/)
- [filter_fields.filter_expression_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/filter_expression_field/)
- [filter_fields.string_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/string_field/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/)
- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
