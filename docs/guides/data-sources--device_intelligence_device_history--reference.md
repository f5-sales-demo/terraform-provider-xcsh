---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_device_history."
xcsh_docs: {"aliases": [], "body_bytes": 5453, "body_sha256": "sha256:f175a69ecb84088834d6dfaab5b173999626bb6d41a03e3d06b246e99c2d30ea", "canonical_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_history:properties:filters", "xcsh-docs:data-sources:device_intelligence_device_history:properties:pagination", "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "xcsh-docs:data-sources:device_intelligence_device_history:properties:sort"], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:fundamentals", "path": "docs/guides/data-sources--device_intelligence_device_history--reference.md", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_device_history.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md)
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

- [filters](data-sources--device_intelligence_device_history--properties--filters.md): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. Namespace name.

- [pagination](data-sources--device_intelligence_device_history--properties--pagination.md): complete subsection reference.

- [records](data-sources--device_intelligence_device_history--properties--records.md): complete subsection reference.

- [sort](data-sources--device_intelligence_device_history--properties--sort.md): complete subsection reference.

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
| `device_id` | [device_id](data-sources--device_intelligence_device_history--reference.md#schema-device_id) |
| `end_time` | [end_time](data-sources--device_intelligence_device_history--reference.md#schema-end_time) |
| `filters` | [filters](data-sources--device_intelligence_device_history--properties--filters.md#section) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_device_history--properties--filters--global_filters.md#section) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_device_history--properties--filters--global_filters.md#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_device_history--properties--filters--global_filters.md#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_device_history--properties--filters--global_filters.md#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_device_history--properties--filters.md#schema-filters--region_filter) |
| `namespace` | [namespace](data-sources--device_intelligence_device_history--reference.md#schema-namespace) |
| `pagination` | [pagination](data-sources--device_intelligence_device_history--properties--pagination.md#section) |
| `pagination.page_number` | [pagination.page_number](data-sources--device_intelligence_device_history--properties--pagination.md#schema-pagination--page_number) |
| `pagination.page_size` | [pagination.page_size](data-sources--device_intelligence_device_history--properties--pagination.md#schema-pagination--page_size) |
| `records` | [records](data-sources--device_intelligence_device_history--properties--records.md#section) |
| `records.action_taken` | [records.action_taken](data-sources--device_intelligence_device_history--properties--records.md#schema-records--action_taken) |
| `records.detected_signals` | [records.detected_signals](data-sources--device_intelligence_device_history--properties--records.md#schema-records--detected_signals) |
| `records.endpoint_label` | [records.endpoint_label](data-sources--device_intelligence_device_history--properties--records.md#schema-records--endpoint_label) |
| `records.risk_score` | [records.risk_score](data-sources--device_intelligence_device_history--properties--records.md#schema-records--risk_score) |
| `records.timestamp` | [records.timestamp](data-sources--device_intelligence_device_history--properties--records.md#schema-records--timestamp) |
| `records.txn_id` | [records.txn_id](data-sources--device_intelligence_device_history--properties--records.md#schema-records--txn_id) |
| `records.url` | [records.url](data-sources--device_intelligence_device_history--properties--records.md#schema-records--url) |
| `sort` | [sort](data-sources--device_intelligence_device_history--properties--sort.md#section) |
| `sort.key` | [sort.key](data-sources--device_intelligence_device_history--properties--sort.md#schema-sort--key) |
| `sort.order` | [sort.order](data-sources--device_intelligence_device_history--properties--sort.md#schema-sort--order) |
| `start_time` | [start_time](data-sources--device_intelligence_device_history--reference.md#schema-start_time) |

## Next pages

- [filters](data-sources--device_intelligence_device_history--properties--filters.md)
- [pagination](data-sources--device_intelligence_device_history--properties--pagination.md)
- [records](data-sources--device_intelligence_device_history--properties--records.md)
- [sort](data-sources--device_intelligence_device_history--properties--sort.md)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md)
