---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_multi_account_devices."
xcsh_docs: {"aliases": [], "body_bytes": 4165, "body_sha256": "sha256:7ce21bec274bc7567184efc90c467a4295bf8a0e526e026358e7d43b2bd49ac6", "canonical_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "child_ids": ["xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters", "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices"], "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:fundamentals", "path": "docs/guides/data-sources--device_intelligence_multi_account_devices--reference.md", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_multi_account_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md)
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

- [filters](data-sources--device_intelligence_multi_account_devices--properties--filters.md): complete subsection reference.

- [multi_account_devices](data-sources--device_intelligence_multi_account_devices--properties--multi_account_devices.md): complete subsection reference.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_multi_account_devices--reference.md#schema-end_time) |
| `filters` | [filters](data-sources--device_intelligence_multi_account_devices--properties--filters.md#section) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_multi_account_devices--properties--filters--global_filters.md#section) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_multi_account_devices--properties--filters--global_filters.md#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_multi_account_devices--properties--filters--global_filters.md#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_multi_account_devices--properties--filters--global_filters.md#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_multi_account_devices--properties--filters.md#schema-filters--region_filter) |
| `multi_account_devices` | [multi_account_devices](data-sources--device_intelligence_multi_account_devices--properties--multi_account_devices.md#section) |
| `multi_account_devices.account_range` | [multi_account_devices.account_range](data-sources--device_intelligence_multi_account_devices--properties--multi_account_devices.md#schema-multi_account_devices--account_range) |
| `multi_account_devices.device_count` | [multi_account_devices.device_count](data-sources--device_intelligence_multi_account_devices--properties--multi_account_devices.md#schema-multi_account_devices--device_count) |
| `namespace` | [namespace](data-sources--device_intelligence_multi_account_devices--reference.md#schema-namespace) |
| `start_time` | [start_time](data-sources--device_intelligence_multi_account_devices--reference.md#schema-start_time) |

## Next pages

- [filters](data-sources--device_intelligence_multi_account_devices--properties--filters.md)
- [multi_account_devices](data-sources--device_intelligence_multi_account_devices--properties--multi_account_devices.md)
- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md)
