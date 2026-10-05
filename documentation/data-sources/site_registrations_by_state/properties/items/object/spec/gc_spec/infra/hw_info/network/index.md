---
page_title: "items.object.spec.gc_spec.infra.hw_info.network"
subcategory: ""
description: "Network. List of network devices in server."
xcsh_docs: {"aliases": ["items object spec gc spec infra hw info network"], "body_bytes": 5862, "body_sha256": "sha256:6050fcfe63e069879b747c6c42a46bd02c5ecbc1b1bad61015e7d4882fb5ccf3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/network/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0221302303121132-2023220110210002-3023302110210300-3030321132332303-0322302123102323-2222003330110110-0220310002123111-1332213202330020", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network"], "schema_version": 1, "sections": [{"aliases": ["items object spec gc spec infra hw info network driver"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--network--driver", "description": "Driver. Driver of device, eg. E1000e.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network", "driver"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info network ip address"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--network--ip_address", "description": "IP Address. IP address on interface.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network", "ip_address"], "syntax": "attribute", "type": "list"}, {"aliases": ["items object spec gc spec infra hw info network link quality"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--network--link_quality", "description": "Link quality determined by VER using different probes Unknown quality Link quality is good Link quality is poor Quality disabled. Possible values are `QUALITY_UNKNOWN`, `QUALITY_GOOD`, `QUALITY_POOR`, `QUALITY_DISABLED`. Defaults to `QUALITY_UNKNOWN`.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["QUALITY_DISABLED", "QUALITY_GOOD", "QUALITY_POOR", "QUALITY_UNKNOWN"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network", "link_quality"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info network link type"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--network--link_type", "description": "Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are `LINK_TYPE_UNKNOWN`, `LINK_TYPE_ETHERNET`, `LINK_TYPE_WIFI_802_11AC`, `LINK_TYPE_WIFI_802_11BGN`, `LINK_TYPE_4G`,", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["LINK_TYPE_4G", "LINK_TYPE_ETHERNET", "LINK_TYPE_UNKNOWN", "LINK_TYPE_WAN", "LINK_TYPE_WIFI", "LINK_TYPE_WIFI_802_11AC", "LINK_TYPE_WIFI_802_11BGN"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network", "link_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info network mac address"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--network--mac_address", "description": "MAC Address. MAC address on interface.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network", "mac_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info network name"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--network--name", "description": "Name. Name of device, eg. Eth0.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info network port"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--network--port", "description": "Port. Used port, eg. Tp.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network", "port"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra hw info network speed"], "anchor": "schema-items--object--spec--gc_spec--infra--hw_info--network--speed", "description": "Speed. Device max supported speed in Mbps.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:hw_info:network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "network", "speed"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/network/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Network. List of network devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.hw_info.network

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.network

<a id="section"></a>

Type: `"list"`. Computed.

Network. List of network devices in server.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--network--driver"></a>

### driver property

Type: `"string"`. Computed.

Driver. Driver of device, eg. E1000e.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--network--ip_address"></a>

### ip_address property

Type: `["list", "string"]`. Computed.

IP Address. IP address on interface.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--network--link_quality"></a>

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

<a id="schema-items--object--spec--gc_spec--infra--hw_info--network--link_type"></a>

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

<a id="schema-items--object--spec--gc_spec--infra--hw_info--network--mac_address"></a>

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

<a id="schema-items--object--spec--gc_spec--infra--hw_info--network--name"></a>

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

<a id="schema-items--object--spec--gc_spec--infra--hw_info--network--port"></a>

### port property

Type: `"string"`. Computed.

Port. Used port, eg. Tp.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--network--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. Device max supported speed in Mbps.

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
