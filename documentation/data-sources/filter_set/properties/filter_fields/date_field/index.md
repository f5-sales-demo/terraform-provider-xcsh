---
page_title: "filter_fields.date_field"
subcategory: ""
description: "Either an absolute time range or a relative time interval."
xcsh_docs: {"aliases": ["filter fields date field"], "body_bytes": 1634, "body_sha256": "sha256:7f4ae706e93424808454edca8e2108f0d94c8aa353d420ec6893a2cbae6e5897", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field:absolute"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field", "parent_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields", "path": "documentation/data-sources/filter_set/properties/filter_fields/date_field/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3001312312012001-2312220023302031-3121301222330313-0000210131021322-0101000321101211-2213333232122011-3320213333031130-3331130100333303", "registry_path": "docs/guides/data-sources--filter_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["filter_fields", "date_field"], "schema_version": 1, "sections": [{"aliases": ["filter fields date field absolute"], "anchor": "section", "description": "Date range is for selecting a date range.", "document_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field:absolute", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filter_fields", "date_field", "absolute"], "syntax": "attribute", "type": "object"}, {"aliases": ["filter fields date field relative"], "anchor": "schema-filter_fields--date_field--relative", "description": "Exclusive with relative time duration.", "document_id": "xcsh-docs:data-sources:filter_set:properties:filter_fields:date_field", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "date_field", "relative"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/properties/filter_fields/date_field/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Either an absolute time range or a relative time interval.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["filter_setCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields.date_field

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/)
- filter_fields.date_field

<a id="section"></a>

Type: `"single"`. Computed.

Either an absolute time range or a relative time interval.

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

## Direct properties

- [absolute](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/filter_fields/date_field/absolute/): complete subsection reference.

<a id="schema-filter_fields--date_field--relative"></a>

### relative property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
