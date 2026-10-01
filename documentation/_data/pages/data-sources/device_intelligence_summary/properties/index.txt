---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_summary."
xcsh_docs: {"aliases": [], "body_bytes": 7023, "body_sha256": "sha256:22dad2da375fd557a978d4cbd2aad279945375dac4208f3e08208cb4f2577eda", "child_ids": ["xcsh-docs:data-sources:device_intelligence_summary:properties:filters"], "collection_id": "xcsh-docs:data-sources:device_intelligence_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_summary:fundamentals", "path": "documentation/data-sources/device_intelligence_summary/properties/index.md", "provider_name": "device_intelligence_summary", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_summary/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_summary.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/)
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

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/): complete subsection reference.

<a id="schema-high_risk_device_count"></a>

### high_risk_device_count property

Type: `"string"`. Computed.

Number of devices classified as high risk.

<a id="schema-high_risk_device_rate"></a>

### high_risk_device_rate property

Type: `"number"`. Computed.

High-risk devices as a percentage of total devices.

<a id="schema-high_risk_txn_count"></a>

### high_risk_txn_count property

Type: `"string"`. Computed.

Number of transactions classified as high risk.

<a id="schema-high_risk_txn_rate"></a>

### high_risk_txn_rate property

Type: `"number"`. Computed.

High-risk transactions as a percentage of total transactions.

<a id="schema-multi_acc_high_risk_device_count"></a>

### multi_acc_high_risk_device_count property

Type: `"string"`. Computed.

Number of high-risk devices associated with multiple accounts.

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

<a id="schema-total_device_count"></a>

### total_device_count property

Type: `"string"`. Computed.

Total Device Count. Total number of devices observed.

<a id="schema-total_evaluated_txn_count"></a>

### total_evaluated_txn_count property

Type: `"string"`. Computed.

Total Evaluated Transaction Count. Total number of transactions evaluated.

<a id="schema-total_multi_acc_device_count"></a>

### total_multi_acc_device_count property

Type: `"string"`. Computed.

Total number of devices associated with multiple accounts.

<a id="schema-total_txn_count"></a>

### total_txn_count property

Type: `"string"`. Computed.

Total Transaction Count. Total number of transactions.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-end_time) |
| `filters` | [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/#section) |
| `filters.global_filters` | [filters.global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/global_filters/#section) |
| `filters.global_filters.key` | [filters.global_filters.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/global_filters/#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/global_filters/#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/global_filters/#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/#schema-filters--region_filter) |
| `high_risk_device_count` | [high_risk_device_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-high_risk_device_count) |
| `high_risk_device_rate` | [high_risk_device_rate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-high_risk_device_rate) |
| `high_risk_txn_count` | [high_risk_txn_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-high_risk_txn_count) |
| `high_risk_txn_rate` | [high_risk_txn_rate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-high_risk_txn_rate) |
| `multi_acc_high_risk_device_count` | [multi_acc_high_risk_device_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-multi_acc_high_risk_device_count) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-namespace) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-start_time) |
| `total_device_count` | [total_device_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-total_device_count) |
| `total_evaluated_txn_count` | [total_evaluated_txn_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-total_evaluated_txn_count) |
| `total_multi_acc_device_count` | [total_multi_acc_device_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-total_multi_acc_device_count) |
| `total_txn_count` | [total_txn_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/#schema-total_txn_count) |

## Next pages

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/)
- [xcsh_device_intelligence_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/)
