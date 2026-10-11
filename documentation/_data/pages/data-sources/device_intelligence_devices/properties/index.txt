---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_devices."
xcsh_docs: {"aliases": ["device intelligence devices"], "body_bytes": 6810, "body_sha256": "sha256:5ddce40cbe0a503477d98538935783128519812158a6c41d9aaf919e00aa2cb3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "xcsh-docs:data-sources:device_intelligence_devices:properties:filters", "xcsh-docs:data-sources:device_intelligence_devices:properties:pagination", "xcsh-docs:data-sources:device_intelligence_devices:properties:sort"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_devices:fundamentals", "path": "documentation/data-sources/device_intelligence_devices/properties/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1010221202313233-3310120333232013-0103200231213012-2012130133200323-1120003102201202-2222210221033201-2030231010113300-1231012332201303", "registry_path": "docs/guides/data-sources--device_intelligence_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["devices"], "anchor": "section", "description": "Devices. List of devices for this page.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["devices"], "syntax": "attribute", "type": "object"}, {"aliases": ["end time"], "anchor": "schema-end_time", "description": "End time of the query period Format: unix_timestamp|RFC 3339 Optional: If not specified, then the end_time will be evaluated to start_time+10m If start_time is not specified, then the end_time will be evaluated to <current time>.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters"], "anchor": "section", "description": "Global Filters. Query Global Filters.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:filters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filters"], "syntax": "attribute", "type": "object"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. Namespace name.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["pagination"], "anchor": "section", "description": "Pagination for Request with number and size.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:pagination", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["pagination"], "syntax": "attribute", "type": "object"}, {"aliases": ["sort"], "anchor": "section", "description": "Sort Option. Query Result Sort Option.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:sort", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sort"], "syntax": "attribute", "type": "object"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start time of the query period Format: unix_timestamp|RFC 3339 Optional: If not specified, then the start_time will be evaluated to end_time-10m If end_time is not specified, then the start_time will be evaluated to <current time>-10m.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_device_intelligence_devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/)
- Property reference

## Direct properties

- [devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/): complete subsection reference.

<a id="schema-end_time"></a>

### end_time property

Type: `"string"`. Optional.

End time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
end\_time will be evaluated to start\_time+10m If start\_time is not specified, then the end\_time
will be evaluated to &lt;current time&gt;.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. Namespace name.

- [pagination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/pagination/): complete subsection reference.

- [sort](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/sort/): complete subsection reference.

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Optional.

Start time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
start\_time will be evaluated to end\_time-10m If end\_time is not specified, then the start\_time
will be evaluated to &lt;current time&gt;-10m.

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
| `devices` | [devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#section) |
| `devices.action_taken` | [devices.action_taken](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#schema-devices--action_taken) |
| `devices.confidence` | [devices.confidence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#schema-devices--confidence) |
| `devices.device_id` | [devices.device_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#schema-devices--device_id) |
| `devices.high_risk_txn_count` | [devices.high_risk_txn_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#schema-devices--high_risk_txn_count) |
| `devices.latest_txn_id` | [devices.latest_txn_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#schema-devices--latest_txn_id) |
| `devices.linked_accounts` | [devices.linked_accounts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#schema-devices--linked_accounts) |
| `devices.risk_score` | [devices.risk_score](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#schema-devices--risk_score) |
| `devices.risk_signals` | [devices.risk_signals](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/#schema-devices--risk_signals) |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/#schema-end_time) |
| `filters` | [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/#section) |
| `filters.global_filters` | [filters.global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/global_filters/#section) |
| `filters.global_filters.key` | [filters.global_filters.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/global_filters/#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/global_filters/#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/global_filters/#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/#schema-filters--region_filter) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/#schema-namespace) |
| `pagination` | [pagination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/pagination/#section) |
| `pagination.page_number` | [pagination.page_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/pagination/#schema-pagination--page_number) |
| `pagination.page_size` | [pagination.page_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/pagination/#schema-pagination--page_size) |
| `sort` | [sort](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/sort/#section) |
| `sort.key` | [sort.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/sort/#schema-sort--key) |
| `sort.order` | [sort.order](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/sort/#schema-sort--order) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/#schema-start_time) |
