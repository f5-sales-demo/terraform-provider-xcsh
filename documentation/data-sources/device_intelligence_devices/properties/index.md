---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_devices."
xcsh_docs: {"aliases": [], "body_bytes": 7424, "body_sha256": "sha256:8bbd4e750de9593963ae73c7f2c6d9de0f682aa4418e9a6fcdb184a154667734", "child_ids": ["xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "xcsh-docs:data-sources:device_intelligence_devices:properties:filters", "xcsh-docs:data-sources:device_intelligence_devices:properties:pagination", "xcsh-docs:data-sources:device_intelligence_devices:properties:sort"], "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_devices:fundamentals", "path": "documentation/data-sources/device_intelligence_devices/properties/index.md", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/devices/)
- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/filters/)
- [pagination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/pagination/)
- [sort](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/sort/)
- [xcsh_device_intelligence_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/)
