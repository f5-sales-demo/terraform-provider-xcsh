---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_interface."
xcsh_docs: {"aliases": ["network interface"], "body_bytes": 40304, "body_sha256": "sha256:175e1772088f3c8a9c835dfbdffc903f3318493c0135ce12d98322c1a6a484cb", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:dedicated_interface", "xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "xcsh-docs:data-sources:network_interface:properties:tunnel_interface"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:reference", "parent_id": "xcsh-docs:data-sources:network_interface:fundamentals", "path": "documentation/data-sources/network_interface/properties/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210", "registry_path": "docs/guides/data-sources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:network_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["dedicated interface"], "anchor": "section", "description": "Dedicated Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dedicated_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated management interface"], "anchor": "section", "description": "Dedicated Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dedicated_management_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:network_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["ethernet interface"], "anchor": "section", "description": "Ethernet Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ethernet_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:network_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:network_interface:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["layer2 interface"], "anchor": "section", "description": "Layer2 Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["layer2_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:network_interface:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:network_interface:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tunnel interface"], "anchor": "section", "description": "Tunnel Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tunnel_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_network_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/): complete subsection reference.

- [dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the NetworkInterface.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the NetworkInterface.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the NetworkInterface exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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

- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/#schema-annotations) |
| `dedicated_interface` | [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/#section) |
| `dedicated_interface.cluster` | [dedicated_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/cluster/#section) |
| `dedicated_interface.device` | [dedicated_interface.device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/#schema-dedicated_interface--device) |
| `dedicated_interface.is_primary` | [dedicated_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/is_primary/#section) |
| `dedicated_interface.monitor` | [dedicated_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/monitor/#section) |
| `dedicated_interface.monitor_disabled` | [dedicated_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/monitor_disabled/#section) |
| `dedicated_interface.mtu` | [dedicated_interface.mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/#schema-dedicated_interface--mtu) |
| `dedicated_interface.node` | [dedicated_interface.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/#schema-dedicated_interface--node) |
| `dedicated_interface.not_primary` | [dedicated_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/not_primary/#section) |
| `dedicated_interface.priority` | [dedicated_interface.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/#schema-dedicated_interface--priority) |
| `dedicated_management_interface` | [dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/#section) |
| `dedicated_management_interface.cluster` | [dedicated_management_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/cluster/#section) |
| `dedicated_management_interface.device` | [dedicated_management_interface.device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/#schema-dedicated_management_interface--device) |
| `dedicated_management_interface.mtu` | [dedicated_management_interface.mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/#schema-dedicated_management_interface--mtu) |
| `dedicated_management_interface.node` | [dedicated_management_interface.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/#schema-dedicated_management_interface--node) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/#schema-description) |
| `ethernet_interface` | [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/#section) |
| `ethernet_interface.cluster` | [ethernet_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/cluster/#section) |
| `ethernet_interface.device` | [ethernet_interface.device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/#schema-ethernet_interface--device) |
| `ethernet_interface.dhcp_client` | [ethernet_interface.dhcp_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_client/#section) |
| `ethernet_interface.dhcp_server` | [ethernet_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/#section) |
| `ethernet_interface.dhcp_server.automatic_from_end` | [ethernet_interface.dhcp_server.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/automatic_from_end/#section) |
| `ethernet_interface.dhcp_server.automatic_from_start` | [ethernet_interface.dhcp_server.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/automatic_from_start/#section) |
| `ethernet_interface.dhcp_server.dhcp_networks` | [ethernet_interface.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/#section) |
| `ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [ethernet_interface.dhcp_server.dhcp_networks.dgw_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/#schema-ethernet_interface--dhcp_server--dhcp_networks--dgw_address) |
| `ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [ethernet_interface.dhcp_server.dhcp_networks.dns_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/#schema-ethernet_interface--dhcp_server--dhcp_networks--dns_address) |
| `ethernet_interface.dhcp_server.dhcp_networks.first_address` | [ethernet_interface.dhcp_server.dhcp_networks.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/first_address/#section) |
| `ethernet_interface.dhcp_server.dhcp_networks.last_address` | [ethernet_interface.dhcp_server.dhcp_networks.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/last_address/#section) |
| `ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [ethernet_interface.dhcp_server.dhcp_networks.network_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/#schema-ethernet_interface--dhcp_server--dhcp_networks--network_prefix) |
| `ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [ethernet_interface.dhcp_server.dhcp_networks.pool_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/#schema-ethernet_interface--dhcp_server--dhcp_networks--pool_settings) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools` | [ethernet_interface.dhcp_server.dhcp_networks.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/pools/#section) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/pools/#schema-ethernet_interface--dhcp_server--dhcp_networks--pools--end_ip) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/pools/#schema-ethernet_interface--dhcp_server--dhcp_networks--pools--exclude) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/pools/#schema-ethernet_interface--dhcp_server--dhcp_networks--pools--start_ip) |
| `ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/same_as_dgw/#section) |
| `ethernet_interface.dhcp_server.dhcp_option82_tag` | [ethernet_interface.dhcp_server.dhcp_option82_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/#schema-ethernet_interface--dhcp_server--dhcp_option82_tag) |
| `ethernet_interface.dhcp_server.fixed_ip_map` | [ethernet_interface.dhcp_server.fixed_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/#schema-ethernet_interface--dhcp_server--fixed_ip_map) |
| `ethernet_interface.dhcp_server.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/interface_ip_map/#section) |
| `ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/interface_ip_map/#schema-ethernet_interface--dhcp_server--interface_ip_map--interface_ip_map) |
| `ethernet_interface.ipv6_auto_config` | [ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/#section) |
| `ethernet_interface.ipv6_auto_config.host` | [ethernet_interface.ipv6_auto_config.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/host/#section) |
| `ethernet_interface.ipv6_auto_config.router` | [ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config` | [ethernet_interface.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/#schema-ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/#schema-ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/first_address/#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/last_address/#section) |
| `ethernet_interface.ipv6_auto_config.router.network_prefix` | [ethernet_interface.ipv6_auto_config.router.network_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/#schema-ethernet_interface--ipv6_auto_config--router--network_prefix) |
| `ethernet_interface.ipv6_auto_config.router.stateful` | [ethernet_interface.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_end/#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_start/#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/#schema-ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--network_prefix) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/#schema-ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pool_settings) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/pools/#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/pools/#schema-ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/pools/#schema-ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip) |
| `ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/#schema-ethernet_interface--ipv6_auto_config--router--stateful--fixed_ip_map) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/interface_ip_map/#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/interface_ip_map/#schema-ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map) |
| `ethernet_interface.is_primary` | [ethernet_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/is_primary/#section) |
| `ethernet_interface.monitor` | [ethernet_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/monitor/#section) |
| `ethernet_interface.monitor_disabled` | [ethernet_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/monitor_disabled/#section) |
| `ethernet_interface.mtu` | [ethernet_interface.mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/#schema-ethernet_interface--mtu) |
| `ethernet_interface.no_ipv6_address` | [ethernet_interface.no_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/no_ipv6_address/#section) |
| `ethernet_interface.node` | [ethernet_interface.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/#schema-ethernet_interface--node) |
| `ethernet_interface.not_primary` | [ethernet_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/not_primary/#section) |
| `ethernet_interface.priority` | [ethernet_interface.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/#schema-ethernet_interface--priority) |
| `ethernet_interface.site_local_inside_network` | [ethernet_interface.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/site_local_inside_network/#section) |
| `ethernet_interface.site_local_network` | [ethernet_interface.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/site_local_network/#section) |
| `ethernet_interface.static_ip` | [ethernet_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/#section) |
| `ethernet_interface.static_ip.cluster_static_ip` | [ethernet_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/cluster_static_ip/#section) |
| `ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/cluster_static_ip/#schema-ethernet_interface--static_ip--cluster_static_ip--interface_ip_map) |
| `ethernet_interface.static_ip.node_static_ip` | [ethernet_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/#section) |
| `ethernet_interface.static_ip.node_static_ip.default_gw` | [ethernet_interface.static_ip.node_static_ip.default_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/#schema-ethernet_interface--static_ip--node_static_ip--default_gw) |
| `ethernet_interface.static_ip.node_static_ip.dns_server` | [ethernet_interface.static_ip.node_static_ip.dns_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/#schema-ethernet_interface--static_ip--node_static_ip--dns_server) |
| `ethernet_interface.static_ip.node_static_ip.ip_address` | [ethernet_interface.static_ip.node_static_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/#schema-ethernet_interface--static_ip--node_static_ip--ip_address) |
| `ethernet_interface.static_ipv6_address` | [ethernet_interface.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/#section) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip` | [ethernet_interface.static_ipv6_address.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/cluster_static_ip/#section) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/cluster_static_ip/#schema-ethernet_interface--static_ipv6_address--cluster_static_ip--interface_ip_map) |
| `ethernet_interface.static_ipv6_address.node_static_ip` | [ethernet_interface.static_ipv6_address.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/node_static_ip/#section) |
| `ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [ethernet_interface.static_ipv6_address.node_static_ip.default_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/node_static_ip/#schema-ethernet_interface--static_ipv6_address--node_static_ip--default_gw) |
| `ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [ethernet_interface.static_ipv6_address.node_static_ip.dns_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/node_static_ip/#schema-ethernet_interface--static_ipv6_address--node_static_ip--dns_server) |
| `ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [ethernet_interface.static_ipv6_address.node_static_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/node_static_ip/#schema-ethernet_interface--static_ipv6_address--node_static_ip--ip_address) |
| `ethernet_interface.storage_network` | [ethernet_interface.storage_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/storage_network/#section) |
| `ethernet_interface.untagged` | [ethernet_interface.untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/untagged/#section) |
| `ethernet_interface.vlan_id` | [ethernet_interface.vlan_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/#schema-ethernet_interface--vlan_id) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/#schema-labels) |
| `layer2_interface` | [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/#section) |
| `layer2_interface.l2sriov_interface` | [layer2_interface.l2sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2sriov_interface/#section) |
| `layer2_interface.l2sriov_interface.device` | [layer2_interface.l2sriov_interface.device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2sriov_interface/#schema-layer2_interface--l2sriov_interface--device) |
| `layer2_interface.l2sriov_interface.untagged` | [layer2_interface.l2sriov_interface.untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2sriov_interface/untagged/#section) |
| `layer2_interface.l2sriov_interface.vlan_id` | [layer2_interface.l2sriov_interface.vlan_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2sriov_interface/#schema-layer2_interface--l2sriov_interface--vlan_id) |
| `layer2_interface.l2vlan_interface` | [layer2_interface.l2vlan_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/#section) |
| `layer2_interface.l2vlan_interface.device` | [layer2_interface.l2vlan_interface.device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/#schema-layer2_interface--l2vlan_interface--device) |
| `layer2_interface.l2vlan_interface.vlan_id` | [layer2_interface.l2vlan_interface.vlan_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/#schema-layer2_interface--l2vlan_interface--vlan_id) |
| `layer2_interface.l2vlan_slo_interface` | [layer2_interface.l2vlan_slo_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_slo_interface/#section) |
| `layer2_interface.l2vlan_slo_interface.vlan_id` | [layer2_interface.l2vlan_slo_interface.vlan_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_slo_interface/#schema-layer2_interface--l2vlan_slo_interface--vlan_id) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/#schema-namespace) |
| `tunnel_interface` | [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/#section) |
| `tunnel_interface.mtu` | [tunnel_interface.mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/#schema-tunnel_interface--mtu) |
| `tunnel_interface.node` | [tunnel_interface.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/#schema-tunnel_interface--node) |
| `tunnel_interface.priority` | [tunnel_interface.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/#schema-tunnel_interface--priority) |
| `tunnel_interface.site_local_inside_network` | [tunnel_interface.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/site_local_inside_network/#section) |
| `tunnel_interface.site_local_network` | [tunnel_interface.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/site_local_network/#section) |
| `tunnel_interface.static_ip` | [tunnel_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/#section) |
| `tunnel_interface.static_ip.cluster_static_ip` | [tunnel_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/cluster_static_ip/#section) |
| `tunnel_interface.static_ip.cluster_static_ip.interface_ip_map` | [tunnel_interface.static_ip.cluster_static_ip.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/cluster_static_ip/#schema-tunnel_interface--static_ip--cluster_static_ip--interface_ip_map) |
| `tunnel_interface.static_ip.node_static_ip` | [tunnel_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/#section) |
| `tunnel_interface.static_ip.node_static_ip.default_gw` | [tunnel_interface.static_ip.node_static_ip.default_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/#schema-tunnel_interface--static_ip--node_static_ip--default_gw) |
| `tunnel_interface.static_ip.node_static_ip.dns_server` | [tunnel_interface.static_ip.node_static_ip.dns_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/#schema-tunnel_interface--static_ip--node_static_ip--dns_server) |
| `tunnel_interface.static_ip.node_static_ip.ip_address` | [tunnel_interface.static_ip.node_static_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/#schema-tunnel_interface--static_ip--node_static_ip--ip_address) |
| `tunnel_interface.tunnel` | [tunnel_interface.tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/tunnel/#section) |
| `tunnel_interface.tunnel.name` | [tunnel_interface.tunnel.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/tunnel/#schema-tunnel_interface--tunnel--name) |
| `tunnel_interface.tunnel.namespace` | [tunnel_interface.tunnel.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/tunnel/#schema-tunnel_interface--tunnel--namespace) |
| `tunnel_interface.tunnel.tenant` | [tunnel_interface.tunnel.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/tunnel/#schema-tunnel_interface--tunnel--tenant) |

## Next pages

- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_interface/)
- [dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/)
- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
