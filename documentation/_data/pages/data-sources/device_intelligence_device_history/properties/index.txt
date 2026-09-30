---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_device_history."
xcsh_docs: {"aliases": [], "body_bytes": 7191, "body_sha256": "sha256:935c340ad25ba2e86cc84b2cc70a541c9d3540a9178cfb0ce764329f572c0a07", "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_history:properties:filters", "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "xcsh-docs:data-sources:device_intelligence_device_history:properties:sort"], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:fundamentals", "path": "documentation/data-sources/device_intelligence_device_history/properties/index.md", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_device_history.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Next pages

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/filters/)
- [pagination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/pagination/)
- [records](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/records/)
- [sort](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/sort/)
- [xcsh_device_intelligence_device_history](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/)
