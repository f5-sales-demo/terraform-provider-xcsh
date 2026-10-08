---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_device_intelligence_device_summary."
xcsh_docs: {"aliases": ["device intelligence device summary"], "body_bytes": 7582, "body_sha256": "sha256:f953ca117a5dac149a24dd967199e4b463cb8bc172e9acffa84592a660d97e76", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_summary:properties:filters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_summary:fundamentals", "path": "documentation/data-sources/device_intelligence_device_summary/properties/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_summary", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2103330111212330-1312310023110131-3130101102300031-1201100213003313-0111322033233122-3102333203020332-2002111132021031-0211302211312122", "registry_path": "docs/guides/data-sources--device_intelligence_device_summary--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["action taken"], "anchor": "schema-action_taken", "description": "Action taken or recommended based on the risk assessment.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["action_taken"], "syntax": "attribute", "type": "string"}, {"aliases": ["channel"], "anchor": "schema-channel", "description": "Channel or platform used by the device (e.g., Web, Mobile).", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["channel"], "syntax": "attribute", "type": "string"}, {"aliases": ["confidence"], "anchor": "schema-confidence", "description": "Confidence level of the risk assessment (0–100).", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["confidence"], "syntax": "attribute", "type": "number"}, {"aliases": ["detected signals"], "anchor": "schema-detected_signals", "description": "List of risk or integrity signals detected for this device.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detected_signals"], "syntax": "attribute", "type": "list"}, {"aliases": ["device id"], "anchor": "schema-device_id", "description": "DeviceID. Device identifier.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["device risk score"], "anchor": "schema-device_risk_score", "description": "Overall risk score for the device (0–100).", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_risk_score"], "syntax": "attribute", "type": "number"}, {"aliases": ["end time"], "anchor": "schema-end_time", "description": "End Time. End time of the query period.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["filters"], "anchor": "section", "description": "Global Filters. Query Global Filters.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:properties:filters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["filters"], "syntax": "attribute", "type": "object"}, {"aliases": ["high risk txn count"], "anchor": "schema-high_risk_txn_count", "description": "Number of high-risk transactions linked to the device.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["high_risk_txn_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["hosting"], "anchor": "schema-hosting", "description": "Type of hosting or network environment detected for the device.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["hosting"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip address"], "anchor": "schema-ip_address", "description": "IP Address. IP address observed for the device.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ip_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["location"], "anchor": "schema-location", "description": "Geographic coordinates associated with the device.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["location"], "syntax": "attribute", "type": "list"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace. Namespace name.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["new or returning device"], "anchor": "schema-new_or_returning_device", "description": "Indicates whether the device is new or returning.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["new_or_returning_device"], "syntax": "attribute", "type": "string"}, {"aliases": ["os"], "anchor": "schema-os", "description": "Operating system and version reported for the device.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["os"], "syntax": "attribute", "type": "string"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start Time. Start time of the query period.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["user agent"], "anchor": "schema-user_agent", "description": "User agent string reported by the device/browser.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_summary:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_agent"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_summary/properties/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Property reference for xcsh_device_intelligence_device_summary.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
EnumExtractionComplete: true
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
EnumExtractionComplete: false
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
EnumExtractionComplete: true
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
