---
page_title: "infra.hw_info.network"
subcategory: ""
description: "List of network devices in server."
xcsh_docs: {"aliases": ["infra hw info network"], "body_bytes": 7673, "body_sha256": "sha256:9e50aa676b9ef16fad8a85e402e7031dc01891b017f743ffbd6b1d0cabfe407f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "documentation/data-sources/registration/properties/infra/hw_info/network/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2002003011233221-2101002023203311-0113332333132302-0203012031033203-1302212230320112-2233130100003303-0003203001333131-1022321123001211", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "network"], "schema_version": 1, "sections": [{"aliases": ["driver"], "anchor": "schema-infra--hw_info--network--driver", "description": "Driver of device, eg. E1000e.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "network", "driver"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip address"], "anchor": "schema-infra--hw_info--network--ip_address", "description": "IP address on interface.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "network", "ip_address"], "syntax": "attribute", "type": "list"}, {"aliases": ["link quality"], "anchor": "schema-infra--hw_info--network--link_quality", "description": "Link quality determined by VER using different probes Unknown quality Link quality is good Link quality is poor Quality disabled.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "network", "link_quality"], "syntax": "attribute", "type": "string"}, {"aliases": ["link type"], "anchor": "schema-infra--hw_info--network--link_type", "description": "Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "network", "link_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["mac address"], "anchor": "schema-infra--hw_info--network--mac_address", "description": "MAC address on interface.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "network", "mac_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-infra--hw_info--network--name", "description": "Name of device, eg. Eth0.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "network", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-infra--hw_info--network--port", "description": "Used port, eg. Tp.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "network", "port"], "syntax": "attribute", "type": "string"}, {"aliases": ["speed"], "anchor": "schema-infra--hw_info--network--speed", "description": "Device max supported speed in Mbps.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "network", "speed"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of network devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.network

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- infra.hw_info.network

<a id="section"></a>

Type: `"list"`. Computed.

Network. List of network devices in server.

Upstream description:

List of network devices in server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "array",
    "maxItems": 32,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-infra--hw_info--network--driver"></a>

### driver property

Type: `"string"`. Computed.

Driver. Driver of device, eg. E1000e.

Upstream description:

Driver of device, eg. E1000e.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--network--ip_address"></a>

### ip_address property

Type: `["list", "string"]`. Computed.

IP Address. IP address on interface.

Upstream description:

IP address on interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--network--link_quality"></a>

### link_quality property

Type: `"string"`. Computed.

\[Enum: QUALITY\_UNKNOWN|QUALITY\_GOOD|QUALITY\_POOR|QUALITY\_DISABLED\] Link quality determined by
VER using different probes Unknown quality Link quality is good Link quality is poor Quality
disabled. Possible values are \`QUALITY\_UNKNOWN\`, \`QUALITY\_GOOD\`, \`QUALITY\_POOR\`,
\`QUALITY\_DISABLED\`. Defaults to \`QUALITY\_UNKNOWN\`.

Upstream description:

Link quality determined by VER using different probes

Unknown quality Link quality is good Link quality is poor Quality disabled.

Receipt-pinned upstream constraints:

```json
{
  "default": "QUALITY_UNKNOWN",
  "enum": [
    "QUALITY_UNKNOWN",
    "QUALITY_GOOD",
    "QUALITY_POOR",
    "QUALITY_DISABLED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--network--link_type"></a>

### link_type property

Type: `"string"`. Computed.

\[Enum:
LINK\_TYPE\_UNKNOWN|LINK\_TYPE\_ETHERNET|LINK\_TYPE\_WIFI\_802\_11AC|LINK\_TYPE\_WIFI\_802\_11BGN|LINK\_TYPE\_4G|LINK\_TYPE\_WIFI|LINK\_TYPE\_WAN\]
Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of
type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are
\`LINK\_TYPE\_UNKNOWN\`, \`LINK\_TYPE\_ETHERNET\`, \`LINK\_TYPE\_WIFI\_802\_11AC\`,
\`LINK\_TYPE\_WIFI\_802\_11BGN\`, \`LINK\_TYPE\_4G\`, \`LINK\_TYPE\_WIFI\`, \`LINK\_TYPE\_WAN\`.
Defaults to \`LINK\_TYPE\_UNKNOWN\`.

Upstream description:

Link type of interface determined operationally

Link type unknown Link type ethernet Wi-Fi link of type 802.11ac Wi-Fi link of type 802.11bgn Link
type 4G Wi-Fi link Wan link.

Receipt-pinned upstream constraints:

```json
{
  "default": "LINK_TYPE_UNKNOWN",
  "enum": [
    "LINK_TYPE_UNKNOWN",
    "LINK_TYPE_ETHERNET",
    "LINK_TYPE_WIFI_802_11AC",
    "LINK_TYPE_WIFI_802_11BGN",
    "LINK_TYPE_4G",
    "LINK_TYPE_WIFI",
    "LINK_TYPE_WAN"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--network--mac_address"></a>

### mac_address property

Type: `"string"`. Computed.

MAC Address. MAC address on interface.

Upstream description:

MAC address on interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "formatDescription": "MAC address (e.g., 00:1A:2B:3C:4D:5E)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--network--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of device, eg. Eth0.

Upstream description:

Name of device, eg. Eth0.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--network--port"></a>

### port property

Type: `"string"`. Computed.

Port. Used port, eg. Tp.

Upstream description:

Used port, eg. Tp.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--network--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. Device max supported speed in Mbps.

Upstream description:

Device max supported speed in Mbps.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
