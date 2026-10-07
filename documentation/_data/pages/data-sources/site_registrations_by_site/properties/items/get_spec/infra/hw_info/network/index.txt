---
page_title: "items.get_spec.infra.hw_info.network"
subcategory: ""
description: "Network. List of network devices in server."
xcsh_docs: {"aliases": ["items get spec infra hw info network"], "body_bytes": 5027, "body_sha256": "sha256:48fedf3a5ac9482ad51e14ab3837409ca6b1dcc5ae748b816c8661385209a328", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/network/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3320213101110032-0320331201110302-1330330120311312-3211002003321111-3123133121213213-3113032130322000-2111100222302011-1223003031021223", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "network"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra hw info network driver"], "anchor": "schema-items--get_spec--infra--hw_info--network--driver", "description": "Driver. Driver of device, eg. E1000e.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "driver"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info network ip address"], "anchor": "schema-items--get_spec--infra--hw_info--network--ip_address", "description": "IP Address. IP address on interface.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "ip_address"], "syntax": "attribute", "type": "list"}, {"aliases": ["items get spec infra hw info network link quality"], "anchor": "schema-items--get_spec--infra--hw_info--network--link_quality", "description": "Link quality determined by VER using different probes Unknown quality Link quality is good Link quality is poor Quality disabled. Possible values are `QUALITY_UNKNOWN`, `QUALITY_GOOD`, `QUALITY_POOR`, `QUALITY_DISABLED`. Defaults to `QUALITY_UNKNOWN`.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["QUALITY_DISABLED", "QUALITY_GOOD", "QUALITY_POOR", "QUALITY_UNKNOWN"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "link_quality"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info network link type"], "anchor": "schema-items--get_spec--infra--hw_info--network--link_type", "description": "Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are `LINK_TYPE_UNKNOWN`, `LINK_TYPE_ETHERNET`, `LINK_TYPE_WIFI_802_11AC`, `LINK_TYPE_WIFI_802_11BGN`, `LINK_TYPE_4G`,", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["LINK_TYPE_4G", "LINK_TYPE_ETHERNET", "LINK_TYPE_UNKNOWN", "LINK_TYPE_WAN", "LINK_TYPE_WIFI", "LINK_TYPE_WIFI_802_11AC", "LINK_TYPE_WIFI_802_11BGN"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "link_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info network mac address"], "anchor": "schema-items--get_spec--infra--hw_info--network--mac_address", "description": "MAC Address. MAC address on interface.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "mac_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info network name"], "anchor": "schema-items--get_spec--infra--hw_info--network--name", "description": "Name. Name of device, eg. Eth0.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info network port"], "anchor": "schema-items--get_spec--infra--hw_info--network--port", "description": "Port. Used port, eg. Tp.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "port"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra hw info network speed"], "anchor": "schema-items--get_spec--infra--hw_info--network--speed", "description": "Speed. Device max supported speed in Mbps.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "speed"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/network/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Network. List of network devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.network

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.network

<a id="section"></a>

Type: `"list"`. Computed.

Network. List of network devices in server.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--network--driver"></a>

### driver property

Type: `"string"`. Computed.

Driver. Driver of device, eg. E1000e.

<a id="schema-items--get_spec--infra--hw_info--network--ip_address"></a>

### ip_address property

Type: `["list", "string"]`. Computed.

IP Address. IP address on interface.

<a id="schema-items--get_spec--infra--hw_info--network--link_quality"></a>

### link_quality property

Type: `"string"`. Computed.

\[Enum: QUALITY\_UNKNOWN|QUALITY\_GOOD|QUALITY\_POOR|QUALITY\_DISABLED\] Link quality determined by
VER using different probes Unknown quality Link quality is good Link quality is poor Quality
disabled. Possible values are \`QUALITY\_UNKNOWN\`, \`QUALITY\_GOOD\`, \`QUALITY\_POOR\`,
\`QUALITY\_DISABLED\`. Defaults to \`QUALITY\_UNKNOWN\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["QUALITY_DISABLED","QUALITY_GOOD","QUALITY_POOR","QUALITY_UNKNOWN"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("QUALITY_UNKNOWN",
    "QUALITY_GOOD",
    "QUALITY_POOR",
    "QUALITY_DISABLED"),
}
```

<a id="schema-items--get_spec--infra--hw_info--network--link_type"></a>

### link_type property

Type: `"string"`. Computed.

\[Enum:
LINK\_TYPE\_UNKNOWN|LINK\_TYPE\_ETHERNET|LINK\_TYPE\_WIFI\_802\_11AC|LINK\_TYPE\_WIFI\_802\_11BGN|LINK\_TYPE\_4G|LINK\_TYPE\_WIFI|LINK\_TYPE\_WAN\]
Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of
type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are
\`LINK\_TYPE\_UNKNOWN\`, \`LINK\_TYPE\_ETHERNET\`, \`LINK\_TYPE\_WIFI\_802\_11AC\`,
\`LINK\_TYPE\_WIFI\_802\_11BGN\`, \`LINK\_TYPE\_4G\`, \`LINK\_TYPE\_WIFI\`, \`LINK\_TYPE\_WAN\`.
Defaults to \`LINK\_TYPE\_UNKNOWN\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LINK_TYPE_4G","LINK_TYPE_ETHERNET","LINK_TYPE_UNKNOWN","LINK_TYPE_WAN","LINK_TYPE_WIFI","LINK_TYPE_WIFI_802_11AC","LINK_TYPE_WIFI_802_11BGN"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("LINK_TYPE_UNKNOWN",
    "LINK_TYPE_ETHERNET",
    "LINK_TYPE_WIFI_802_11AC",
    "LINK_TYPE_WIFI_802_11BGN",
    "LINK_TYPE_4G",
    "LINK_TYPE_WIFI",
    "LINK_TYPE_WAN"),
}
```

<a id="schema-items--get_spec--infra--hw_info--network--mac_address"></a>

### mac_address property

Type: `"string"`. Computed.

MAC Address. MAC address on interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(17, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--network--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of device, eg. Eth0.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-items--get_spec--infra--hw_info--network--port"></a>

### port property

Type: `"string"`. Computed.

Port. Used port, eg. Tp.

<a id="schema-items--get_spec--infra--hw_info--network--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. Device max supported speed in Mbps.
