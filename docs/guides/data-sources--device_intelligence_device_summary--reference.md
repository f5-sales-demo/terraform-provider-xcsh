---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_device_summary."
xcsh_docs: {"aliases": [], "body_bytes": 6344, "body_sha256": "sha256:2a0eb28f632544c8211ae8550f3be3b190674cdc73cba820768505ce5b17e05e", "canonical_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_summary:properties:filters"], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_summary:fundamentals", "path": "docs/guides/data-sources--device_intelligence_device_summary--reference.md", "provider_name": "device_intelligence_device_summary", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_summary/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_device_summary.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md)
- Property reference

## Direct properties

<a id="schema-action_taken"></a>

### action_taken property

Type: `"string"`. Computed.

Action taken or recommended based on the risk assessment.

<a id="schema-channel"></a>

### channel property

Type: `"string"`. Computed.

Channel or platform used by the device (e.g., Web, Mobile).

<a id="schema-confidence"></a>

### confidence property

Type: `"number"`. Computed.

Confidence level of the risk assessment (0–100).

<a id="schema-detected_signals"></a>

### detected_signals property

Type: `["list", "string"]`. Computed.

List of risk or integrity signals detected for this device.

<a id="schema-device_id"></a>

### device_id property

Type: `"string"`. Required.

DeviceID. Device identifier.

<a id="schema-device_risk_score"></a>

### device_risk_score property

Type: `"number"`. Computed.

Overall risk score for the device (0–100).

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

- [filters](data-sources--device_intelligence_device_summary--properties--filters.md): complete subsection reference.

<a id="schema-high_risk_txn_count"></a>

### high_risk_txn_count property

Type: `"string"`. Computed.

Number of high-risk transactions linked to the device.

<a id="schema-hosting"></a>

### hosting property

Type: `"string"`. Computed.

Type of hosting or network environment detected for the device.

<a id="schema-ip_address"></a>

### ip_address property

Type: `"string"`. Computed.

IP Address. IP address observed for the device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`),
    ""),
}
```

<a id="schema-location"></a>

### location property

Type: `["list", "number"]`. Computed.

Geographic coordinates associated with the device.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="schema-new_or_returning_device"></a>

### new_or_returning_device property

Type: `"string"`. Computed.

Indicates whether the device is new or returning.

<a id="schema-os"></a>

### os property

Type: `"string"`. Computed.

Operating system and version reported for the device.

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

<a id="schema-user_agent"></a>

### user_agent property

Type: `"string"`. Computed.

User agent string reported by the device/browser.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action_taken` | [action_taken](data-sources--device_intelligence_device_summary--reference.md#schema-action_taken) |
| `channel` | [channel](data-sources--device_intelligence_device_summary--reference.md#schema-channel) |
| `confidence` | [confidence](data-sources--device_intelligence_device_summary--reference.md#schema-confidence) |
| `detected_signals` | [detected_signals](data-sources--device_intelligence_device_summary--reference.md#schema-detected_signals) |
| `device_id` | [device_id](data-sources--device_intelligence_device_summary--reference.md#schema-device_id) |
| `device_risk_score` | [device_risk_score](data-sources--device_intelligence_device_summary--reference.md#schema-device_risk_score) |
| `end_time` | [end_time](data-sources--device_intelligence_device_summary--reference.md#schema-end_time) |
| `filters` | [filters](data-sources--device_intelligence_device_summary--properties--filters.md#section) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_device_summary--properties--filters--global_filters.md#section) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_device_summary--properties--filters--global_filters.md#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_device_summary--properties--filters--global_filters.md#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_device_summary--properties--filters--global_filters.md#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_device_summary--properties--filters.md#schema-filters--region_filter) |
| `high_risk_txn_count` | [high_risk_txn_count](data-sources--device_intelligence_device_summary--reference.md#schema-high_risk_txn_count) |
| `hosting` | [hosting](data-sources--device_intelligence_device_summary--reference.md#schema-hosting) |
| `ip_address` | [ip_address](data-sources--device_intelligence_device_summary--reference.md#schema-ip_address) |
| `location` | [location](data-sources--device_intelligence_device_summary--reference.md#schema-location) |
| `namespace` | [namespace](data-sources--device_intelligence_device_summary--reference.md#schema-namespace) |
| `new_or_returning_device` | [new_or_returning_device](data-sources--device_intelligence_device_summary--reference.md#schema-new_or_returning_device) |
| `os` | [os](data-sources--device_intelligence_device_summary--reference.md#schema-os) |
| `start_time` | [start_time](data-sources--device_intelligence_device_summary--reference.md#schema-start_time) |
| `user_agent` | [user_agent](data-sources--device_intelligence_device_summary--reference.md#schema-user_agent) |

## Next pages

- [filters](data-sources--device_intelligence_device_summary--properties--filters.md)
- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md)
