---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_device_summary."
xcsh_docs: {"aliases": [], "body_bytes": 7797, "body_sha256": "sha256:e5341ef1b2c4c4a3fde6603117116719a83df33f9a27bf6e3f98d5cbbc5ce151", "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_summary:properties:filters"], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_summary:fundamentals", "path": "documentation/data-sources/device_intelligence_device_summary/properties/index.md", "provider_name": "device_intelligence_device_summary", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_summary/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_device_intelligence_device_summary.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/)
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

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/filters/): complete subsection reference.

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
| `action_taken` | [action_taken](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-action_taken) |
| `channel` | [channel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-channel) |
| `confidence` | [confidence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-confidence) |
| `detected_signals` | [detected_signals](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-detected_signals) |
| `device_id` | [device_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-device_id) |
| `device_risk_score` | [device_risk_score](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-device_risk_score) |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-end_time) |
| `filters` | [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/filters/#section) |
| `filters.global_filters` | [filters.global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/filters/global_filters/#section) |
| `filters.global_filters.key` | [filters.global_filters.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/filters/global_filters/#schema-filters--global_filters--key) |
| `filters.global_filters.op` | [filters.global_filters.op](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/filters/global_filters/#schema-filters--global_filters--op) |
| `filters.global_filters.values` | [filters.global_filters.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/filters/global_filters/#schema-filters--global_filters--values) |
| `filters.region_filter` | [filters.region_filter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/filters/#schema-filters--region_filter) |
| `high_risk_txn_count` | [high_risk_txn_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-high_risk_txn_count) |
| `hosting` | [hosting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-hosting) |
| `ip_address` | [ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-ip_address) |
| `location` | [location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-location) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-namespace) |
| `new_or_returning_device` | [new_or_returning_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-new_or_returning_device) |
| `os` | [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-os) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-start_time) |
| `user_agent` | [user_agent](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/#schema-user_agent) |

## Next pages

- [filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/filters/)
- [xcsh_device_intelligence_device_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/)
