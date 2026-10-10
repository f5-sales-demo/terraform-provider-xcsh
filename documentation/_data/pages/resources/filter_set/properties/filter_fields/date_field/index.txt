---
page_title: "filter_fields.date_field"
subcategory: ""
description: "Either an absolute time range or a relative time interval."
xcsh_docs: {"aliases": ["filter fields date field"], "body_bytes": 1937, "body_sha256": "sha256:62fa27dcb4ebd886914e91b2678b01ac2b6f5c5218d3cdd48587cc9e62231c41", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "parent_id": "xcsh-docs:resources:filter_set:properties:filter_fields", "path": "documentation/resources/filter_set/properties/filter_fields/date_field/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0311201023033312-3031321322112321-3013102220300333-3123311121222133-1332333132233322-0013022102112313-1223331032212111-3221231001222032", "registry_path": "docs/guides/resources--filter_set--reference--group-001.md", "relationships": [{"anchor": "schema-filter_fields--date_field--relative", "enforcement": "provider-schema", "group": "filter_fields.date_field:ConflictingObjectAttributes:absolute,relative", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "filter_fields.date_field:ConflictingObjectAttributes:absolute,relative", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["filter_fields", "date_field"], "schema_version": 1, "sections": [{"aliases": ["filter fields date field absolute"], "anchor": "section", "description": "Date range is for selecting a date range.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-filter_fields--date_field--absolute--end_date", "enforcement": "provider-schema", "group": "filter_fields.date_field.absolute:RequiredObjectAttributes:end_date,start_date", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "requires"}, {"anchor": "schema-filter_fields--date_field--absolute--start_date", "enforcement": "provider-schema", "group": "filter_fields.date_field.absolute:RequiredObjectAttributes:end_date,start_date", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field:absolute", "type": "requires"}], "schema_path": ["filter_fields", "date_field", "absolute"], "syntax": "block", "type": "object"}, {"aliases": ["filter fields date field relative"], "anchor": "schema-filter_fields--date_field--relative", "description": "Exclusive with relative time duration.", "document_id": "xcsh-docs:resources:filter_set:properties:filter_fields:date_field", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filter_fields", "date_field", "relative"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/filter_fields/date_field/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Either an absolute time range or a relative time interval.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["filter_setCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
