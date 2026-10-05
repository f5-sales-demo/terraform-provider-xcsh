---
page_title: "filter_fields"
subcategory: ""
description: "List of fields and their values selected by the user."
xcsh_docs: {"aliases": ["filter fields"], "body_bytes": 3409, "body_sha256": "sha256:c9e405c8aa16dff453fbb40d4cda8a010a922faa58974a06a6aecd049971cbae", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field", "xcsh-docs:data-sources:filter_set:properties:filter_fields:filter_expression_field", "xcsh-docs:data-sources:filter_set:properties:filter_fields:string_field"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:filter_set:properties:filter_fields", "parent_id": "xcsh-docs:data-sources:filter_set:reference", "path": "documentation/data-sources/filter_set/properties/filter_fields/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1231321122330232-3132213321230321-3232300130032302-3000132223002232-0211003002102322-1100023311310211-3223030201031001-2310130113233331", "registry_path": "docs/guides/data-sources--filter_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["filter_fields"], "schema_version": 1, "sections": [{"aliases": ["filter fields date field"], "anchor": "section", "description": "Either an absolute time range or a relative time interval.", "document_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filter_fields", "date_field"], "syntax": "attribute", "type": "object"}, {"aliases": ["filter fields field id"], "anchor": "schema-filter_fields--field_id", "description": "An identifier for the field that maps to some UI filter component.", "document_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "field_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["filter fields filter expression field"], "anchor": "section", "description": "Filter Expression Field.", "document_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:filter_expression_field", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filter_fields", "filter_expression_field"], "syntax": "attribute", "type": "object"}, {"aliases": ["filter fields string field"], "anchor": "section", "description": "Filter String Field.", "document_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:string_field", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filter_fields", "string_field"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/properties/filter_fields/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of fields and their values selected by the user.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- filter_fields

<a id="section"></a>

Type: `"list"`. Computed.

List of fields and their values selected by the user.

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

## Direct properties

- [date_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/): complete subsection reference.

<a id="schema-filter_fields--field_id"></a>

### field_id property

Type: `"string"`. Computed.

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

- [filter_expression_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/filter_expression_field/): complete subsection reference.

- [string_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/string_field/): complete subsection reference.

## Next pages

- [filter_fields.date_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/)
- [filter_fields.filter_expression_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/filter_expression_field/)
- [filter_fields.string_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/string_field/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
