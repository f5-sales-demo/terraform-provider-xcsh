---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_device_history."
xcsh_docs: {"aliases": ["device intelligence device history"], "body_bytes": 6634, "body_sha256": "sha256:c5995adc88114568947aa79cc8b91ff55cd5ee42c64b49460a8710662d07132f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_history:properties:filters", "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "xcsh-docs:data-sources:device_intelligence_device_history:properties:sort"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:fundamentals", "path": "documentation/data-sources/device_intelligence_device_history/properties/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1113221232103010-3013003312221121-0011022102212302-3131012031331210-0020303211131110-3123123112212302-2300102012010332-2102313110312030", "registry_path": "docs/guides/data-sources--device_intelligence_device_history--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["device id"], "anchor": "schema-device_id", "description": "DeviceID. Device identifier.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["end time"], "anchor": "schema-end_time", "description": "End Time. End time of the query period.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters"], "anchor": "section", "description": "Global Filters. Query Global Filters.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:filters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filters"], "syntax": "attribute", "type": "object"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. Namespace name.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["pagination"], "anchor": "section", "description": "Pagination for Request with number and size.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["pagination"], "syntax": "attribute", "type": "object"}, {"aliases": ["records"], "anchor": "section", "description": "Records. List of activity records.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["records"], "syntax": "attribute", "type": "object"}, {"aliases": ["sort"], "anchor": "section", "description": "Sort Option. Query Result Sort Option.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:sort", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sort"], "syntax": "attribute", "type": "object"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start Time. Start time of the query period.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_device_intelligence_device_history.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_device_history](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/)
- Property reference

## Direct properties

<a id="schema-device_id"></a>

### device_id property

Type: `"string"`. Required.

DeviceID. Device identifier.

<a id="schema-end_time"></a>

### end_time property

Type: `"string"`. Optional.

End Time. End time of the query period.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/filters/): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. Namespace name.

- [pagination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/pagination/): complete subsection reference.

- [records](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/): complete subsection reference.

- [sort](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/sort/): complete subsection reference.

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Optional.

Start Time. Start time of the query period.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `device_id` | [device_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/#schema-device_id) |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/#schema-end_time) |
| `filters` | [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/filters/#section) |
| `filters.global_filters` | [filters.global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/filters/global_filters/#section) |
| `filters.global_filters.key` | [filters.global_filters.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/filters/global_filters/#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/filters/global_filters/#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/filters/global_filters/#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/filters/#schema-filters--region_filter) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/#schema-namespace) |
| `pagination` | [pagination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/pagination/#section) |
| `pagination.page_number` | [pagination.page_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/pagination/#schema-pagination--page_number) |
| `pagination.page_size` | [pagination.page_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/pagination/#schema-pagination--page_size) |
| `records` | [records](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/#section) |
| `records.action_taken` | [records.action_taken](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/#schema-records--action_taken) |
| `records.detected_signals` | [records.detected_signals](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/#schema-records--detected_signals) |
| `records.endpoint_label` | [records.endpoint_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/#schema-records--endpoint_label) |
| `records.risk_score` | [records.risk_score](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/#schema-records--risk_score) |
| `records.timestamp` | [records.timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/#schema-records--timestamp) |
| `records.txn_id` | [records.txn_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/#schema-records--txn_id) |
| `records.url` | [records.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/#schema-records--url) |
| `sort` | [sort](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/sort/#section) |
| `sort.key` | [sort.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/sort/#schema-sort--key) |
| `sort.order` | [sort.order](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/sort/#schema-sort--order) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/#schema-start_time) |
