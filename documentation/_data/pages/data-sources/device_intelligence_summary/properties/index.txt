---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_summary."
xcsh_docs: {"aliases": ["device intelligence summary"], "body_bytes": 7023, "body_sha256": "sha256:22dad2da375fd557a978d4cbd2aad279945375dac4208f3e08208cb4f2577eda", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_summary:properties:filters"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_summary:fundamentals", "path": "documentation/data-sources/device_intelligence_summary/properties/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_summary", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0130130001220323-3031111221132001-3222102300212300-0203021003111000-2213312212330210-2310230230130330-0213320113111122-2313223002130011", "registry_path": "docs/guides/data-sources--device_intelligence_summary--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["end time"], "anchor": "schema-end_time", "description": "End time of the query period Format: unix_timestamp|RFC 3339 Optional: If not specified, then the end_time will be evaluated to start_time+10m If start_time is not specified, then the end_time will be evaluated to <current time>.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters"], "anchor": "section", "description": "Global Filters. Query Global Filters.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:properties:filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filters"], "syntax": "attribute", "type": "object"}, {"aliases": ["high risk device count"], "anchor": "schema-high_risk_device_count", "description": "Number of devices classified as high risk.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["high_risk_device_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["high risk device rate"], "anchor": "schema-high_risk_device_rate", "description": "High-risk devices as a percentage of total devices.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["high_risk_device_rate"], "syntax": "attribute", "type": "number"}, {"aliases": ["high risk txn count"], "anchor": "schema-high_risk_txn_count", "description": "Number of transactions classified as high risk.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["high_risk_txn_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["high risk txn rate"], "anchor": "schema-high_risk_txn_rate", "description": "High-risk transactions as a percentage of total transactions.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["high_risk_txn_rate"], "syntax": "attribute", "type": "number"}, {"aliases": ["multi acc high risk device count"], "anchor": "schema-multi_acc_high_risk_device_count", "description": "Number of high-risk devices associated with multiple accounts.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["multi_acc_high_risk_device_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. Namespace name.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start time of the query period Format: unix_timestamp|RFC 3339 Optional: If not specified, then the start_time will be evaluated to end_time-10m If end_time is not specified, then the start_time will be evaluated to <current time>-10m.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["total device count"], "anchor": "schema-total_device_count", "description": "Total Device Count. Total number of devices observed.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["total_device_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["total evaluated txn count"], "anchor": "schema-total_evaluated_txn_count", "description": "Total Evaluated Transaction Count. Total number of transactions evaluated.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["total_evaluated_txn_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["total multi acc device count"], "anchor": "schema-total_multi_acc_device_count", "description": "Total number of devices associated with multiple accounts.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["total_multi_acc_device_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["total txn count"], "anchor": "schema-total_txn_count", "description": "Total Transaction Count. Total number of transactions.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["total_txn_count"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_summary/properties/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Property reference for xcsh_device_intelligence_summary.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": [], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
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
