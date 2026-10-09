---
page_title: "filter_fields.date_field.absolute"
subcategory: ""
description: "Date range is for selecting a date range."
xcsh_docs: {"aliases": ["filter fields date field absolute"], "body_bytes": 2789, "body_sha256": "sha256:e203f6d2dcf6a49e9ea88b6eae39d2f7eecf7ae824bc4697674c06dde00c4920", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "parent_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "path": "documentation/resources/filter_set/properties/filter_fields/date_field/absolute/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0021010123302132-3213302002223301-0231220023101303-0202300022203023-0202032002220103-2203201313310012-3033202331321311-3100221122333303", "registry_path": "docs/guides/resources--filter_set--reference--group-001.md", "relationships": [{"anchor": "schema-filter_fields--date_field--absolute--end_date", "enforcement": "provider-schema", "group": "filter_fields.date_field.absolute:RequiredObjectAttributes:end_date,start_date", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "requires"}, {"anchor": "schema-filter_fields--date_field--absolute--start_date", "enforcement": "provider-schema", "group": "filter_fields.date_field.absolute:RequiredObjectAttributes:end_date,start_date", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["filter_fields", "date_field", "absolute"], "schema_version": 1, "sections": [{"aliases": ["filter fields date field absolute end date"], "anchor": "schema-filter_fields--date_field--absolute--end_date", "description": "Contains end date.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "date_field", "absolute", "end_date"], "syntax": "attribute", "type": "string"}, {"aliases": ["filter fields date field absolute start date"], "anchor": "schema-filter_fields--date_field--absolute--start_date", "description": "Contains start date.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "date_field", "absolute", "start_date"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/date_field/absolute/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Date range is for selecting a date range.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["filter_setCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filter_fields.date_field.absolute

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/)
- [filter_fields](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/)
- [filter_fields.date_field](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/filter_fields/date_field/)
- filter_fields.date_field.absolute

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Date range is for selecting a date range.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("end_date",
    "start_date")}
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
absolute {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-filter_fields--date_field--absolute--end_date"></a>

### end_date property

Type: `"string"`. Optional.

End Date. Contains end date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-filter_fields--date_field--absolute--start_date"></a>

### start_date property

Type: `"string"`. Optional.

Start Date. Contains start date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
