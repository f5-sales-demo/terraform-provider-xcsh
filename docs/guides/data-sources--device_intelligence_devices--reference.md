---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_devices."
xcsh_docs: {"aliases": [], "body_bytes": 5688, "body_sha256": "sha256:dfe0e9016a99d8036582a3cf045470a7cde8044496f3e4a9ede4b24bd4205dbe", "canonical_id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "child_ids": ["xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "xcsh-docs:data-sources:device_intelligence_devices:properties:filters", "xcsh-docs:data-sources:device_intelligence_devices:properties:pagination", "xcsh-docs:data-sources:device_intelligence_devices:properties:sort"], "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_devices:fundamentals", "path": "docs/guides/data-sources--device_intelligence_devices--reference.md", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md)
- Property reference

## Direct properties

- [devices](data-sources--device_intelligence_devices--properties--devices.md): complete subsection reference.

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

- [filters](data-sources--device_intelligence_devices--properties--filters.md): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. Namespace name.

- [pagination](data-sources--device_intelligence_devices--properties--pagination.md): complete subsection reference.

- [sort](data-sources--device_intelligence_devices--properties--sort.md): complete subsection reference.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `devices` | [devices](data-sources--device_intelligence_devices--properties--devices.md#section) |
| `devices.action_taken` | [devices.action_taken](data-sources--device_intelligence_devices--properties--devices.md#schema-devices--action_taken) |
| `devices.confidence` | [devices.confidence](data-sources--device_intelligence_devices--properties--devices.md#schema-devices--confidence) |
| `devices.device_id` | [devices.device_id](data-sources--device_intelligence_devices--properties--devices.md#schema-devices--device_id) |
| `devices.high_risk_txn_count` | [devices.high_risk_txn_count](data-sources--device_intelligence_devices--properties--devices.md#schema-devices--high_risk_txn_count) |
| `devices.latest_txn_id` | [devices.latest_txn_id](data-sources--device_intelligence_devices--properties--devices.md#schema-devices--latest_txn_id) |
| `devices.linked_accounts` | [devices.linked_accounts](data-sources--device_intelligence_devices--properties--devices.md#schema-devices--linked_accounts) |
| `devices.risk_score` | [devices.risk_score](data-sources--device_intelligence_devices--properties--devices.md#schema-devices--risk_score) |
| `devices.risk_signals` | [devices.risk_signals](data-sources--device_intelligence_devices--properties--devices.md#schema-devices--risk_signals) |
| `end_time` | [end_time](data-sources--device_intelligence_devices--reference.md#schema-end_time) |
| `filters` | [filters](data-sources--device_intelligence_devices--properties--filters.md#section) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_devices--properties--filters--global_filters.md#section) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_devices--properties--filters--global_filters.md#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_devices--properties--filters--global_filters.md#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_devices--properties--filters--global_filters.md#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_devices--properties--filters.md#schema-filters--region_filter) |
| `namespace` | [namespace](data-sources--device_intelligence_devices--reference.md#schema-namespace) |
| `pagination` | [pagination](data-sources--device_intelligence_devices--properties--pagination.md#section) |
| `pagination.page_number` | [pagination.page_number](data-sources--device_intelligence_devices--properties--pagination.md#schema-pagination--page_number) |
| `pagination.page_size` | [pagination.page_size](data-sources--device_intelligence_devices--properties--pagination.md#schema-pagination--page_size) |
| `sort` | [sort](data-sources--device_intelligence_devices--properties--sort.md#section) |
| `sort.key` | [sort.key](data-sources--device_intelligence_devices--properties--sort.md#schema-sort--key) |
| `sort.order` | [sort.order](data-sources--device_intelligence_devices--properties--sort.md#schema-sort--order) |
| `start_time` | [start_time](data-sources--device_intelligence_devices--reference.md#schema-start_time) |

## Next pages

- [devices](data-sources--device_intelligence_devices--properties--devices.md)
- [filters](data-sources--device_intelligence_devices--properties--filters.md)
- [pagination](data-sources--device_intelligence_devices--properties--pagination.md)
- [sort](data-sources--device_intelligence_devices--properties--sort.md)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md)
