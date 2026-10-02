---
page_title: "items.get_spec.infra.hw_info.network"
subcategory: ""
description: "Network. List of network devices in server."
xcsh_docs: {"aliases": ["items get spec infra hw info network"], "body_bytes": 4607, "body_sha256": "sha256:886dac83c28ea823c6947d0816038c54308d164d3f8d6544d911f4dd5af5022c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/network/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3320213101110032-0320331201110302-1330330120311312-3211002003321111-3123133121213213-3113032130322000-2111100222302011-1223003031021223", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "network"], "schema_version": 1, "sections": [{"aliases": ["driver"], "anchor": "schema-items--get_spec--infra--hw_info--network--driver", "description": "Driver. Driver of device, eg. E1000e.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "driver"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip address"], "anchor": "schema-items--get_spec--infra--hw_info--network--ip_address", "description": "IP Address. IP address on interface.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "ip_address"], "syntax": "attribute", "type": "list"}, {"aliases": ["link quality"], "anchor": "schema-items--get_spec--infra--hw_info--network--link_quality", "description": "Link quality determined by VER using different probes Unknown quality Link quality is good Link quality is poor Quality disabled. Possible values are `QUALITY_UNKNOWN`, `QUALITY_GOOD`, `QUALITY_POOR`, `QUALITY_DISABLED`. Defaults to `QUALITY_UNKNOWN`.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "link_quality"], "syntax": "attribute", "type": "string"}, {"aliases": ["link type"], "anchor": "schema-items--get_spec--infra--hw_info--network--link_type", "description": "Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are `LINK_TYPE_UNKNOWN`, `LINK_TYPE_ETHERNET`, `LINK_TYPE_WIFI_802_11AC`, `LINK_TYPE_WIFI_802_11BGN`, `LINK_TYPE_4G`,", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "link_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["mac address"], "anchor": "schema-items--get_spec--infra--hw_info--network--mac_address", "description": "MAC Address. MAC address on interface.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "mac_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-items--get_spec--infra--hw_info--network--name", "description": "Name. Name of device, eg. Eth0.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-items--get_spec--infra--hw_info--network--port", "description": "Port. Used port, eg. Tp.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "port"], "syntax": "attribute", "type": "string"}, {"aliases": ["speed"], "anchor": "schema-items--get_spec--infra--hw_info--network--speed", "description": "Speed. Device max supported speed in Mbps.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "network", "speed"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Network. List of network devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
