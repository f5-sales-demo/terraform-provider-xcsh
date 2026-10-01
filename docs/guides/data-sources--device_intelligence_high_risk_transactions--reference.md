---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_high_risk_transactions."
xcsh_docs: {"aliases": [], "body_bytes": 4616, "body_sha256": "sha256:3c475068d8d65767205968f054cc73531a5cb4203fa81026e082a2702c9e23cf", "canonical_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:reference", "child_ids": ["xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:filters", "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results"], "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:fundamentals", "path": "docs/guides/data-sources--device_intelligence_high_risk_transactions--reference.md", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_high_risk_transactions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md)
- Property reference

## Direct properties

<a id="schema-end_time"></a>

### end_time property

Type: `"string"`. Optional.

End time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
end\_time will be evaluated to start\_time+10m If start\_time is not specified, then the end\_time
will be evaluated to &lt;current time&gt;.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](data-sources--device_intelligence_high_risk_transactions--properties--filters.md): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Optional.

Start time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
start\_time will be evaluated to end\_time-10m If end\_time is not specified, then the start\_time
will be evaluated to &lt;current time&gt;-10m.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [time_series_results](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_high_risk_transactions--reference.md#schema-end_time) |
| `filters` | [filters](data-sources--device_intelligence_high_risk_transactions--properties--filters.md#section) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_high_risk_transactions--properties--filters--global_filters.md#section) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_high_risk_transactions--properties--filters--global_filters.md#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_high_risk_transactions--properties--filters--global_filters.md#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_high_risk_transactions--properties--filters--global_filters.md#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_high_risk_transactions--properties--filters.md#schema-filters--region_filter) |
| `namespace` | [namespace](data-sources--device_intelligence_high_risk_transactions--reference.md#schema-namespace) |
| `start_time` | [start_time](data-sources--device_intelligence_high_risk_transactions--reference.md#schema-start_time) |
| `time_series_results` | [time_series_results](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results.md#section) |
| `time_series_results.series_key` | [time_series_results.series_key](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results.md#schema-time_series_results--series_key) |
| `time_series_results.time_series` | [time_series_results.time_series](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results--time_series.md#section) |
| `time_series_results.time_series.timestamp` | [time_series_results.time_series.timestamp](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results--time_series.md#schema-time_series_results--time_series--timestamp) |
| `time_series_results.time_series.value` | [time_series_results.time_series.value](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results--time_series.md#schema-time_series_results--time_series--value) |

## Next pages

- [filters](data-sources--device_intelligence_high_risk_transactions--properties--filters.md)
- [time_series_results](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results.md)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md)
