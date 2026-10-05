---
page_title: "filter_fields.date_field"
subcategory: ""
description: "Either an absolute time range or a relative time interval."
xcsh_docs: {"aliases": ["filter fields date field"], "body_bytes": 2382, "body_sha256": "sha256:6b48436ac19791bc6ea8b87e117fcb2d409fbbdab6c940deeefe4f91d18d1002", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "parent_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "path": "documentation/resources/filter_set/properties/filter_fields/date_field/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0311201023033312-3031321322112321-3013102220300333-3123311121222133-1332333132233322-0013022102112313-1223331032212111-3221231001222032", "registry_path": "docs/guides/resources--filter_set--reference--group-001.md", "relationships": [{"anchor": "schema-filter_fields--date_field--relative", "enforcement": "provider-schema", "group": "filter_fields.date_field:ConflictingObjectAttributes:absolute,relative", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields.date_field:ConflictingObjectAttributes:absolute,relative", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["filter_fields", "date_field"], "schema_version": 1, "sections": [{"aliases": ["filter fields date field absolute"], "anchor": "section", "description": "Date range is for selecting a date range.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-filter_fields--date_field--absolute--end_date", "enforcement": "provider-schema", "group": "filter_fields.date_field.absolute:RequiredObjectAttributes:end_date,start_date", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "requires"}, {"anchor": "schema-filter_fields--date_field--absolute--start_date", "enforcement": "provider-schema", "group": "filter_fields.date_field.absolute:RequiredObjectAttributes:end_date,start_date", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "requires"}], "schema_path": ["filter_fields", "date_field", "absolute"], "syntax": "block", "type": "object"}, {"aliases": ["filter fields date field relative"], "anchor": "schema-filter_fields--date_field--relative", "description": "Exclusive with relative time duration.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "date_field", "relative"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/date_field/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Either an absolute time range or a relative time interval.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields.date_field

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/)
- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/)
- filter_fields.date_field

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Either an absolute time range or a relative time interval.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("absolute",
    "relative")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-range_type": "[\"absolute\",\"relative\"]"
}
```

Terraform syntax:

```terraform
date_field {
  # Configure direct properties listed below.
}
```

## Direct properties

- [absolute](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/date_field/absolute/): complete subsection reference.

<a id="schema-filter_fields--date_field--relative"></a>

### relative property

Type: `"string"`. Optional.

Exclusive with \[absolute\] relative time duration.

Upstream description:

Exclusive with \[absolute\] relative time duration.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [filter_fields.date_field.absolute](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/date_field/absolute/)
- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/)
- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
