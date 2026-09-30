---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 33609, "body_sha256": "sha256:f392950245ed4a86d87c74aae386a78056a1a821a7147f4ffc405b01e0b080ba", "canonical_id": "xcsh-docs:data-sources:network_interface:reference", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:dedicated_interface", "xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "xcsh-docs:data-sources:network_interface:properties:tunnel_interface"], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:reference", "parent_id": "xcsh-docs:data-sources:network_interface:fundamentals", "path": "docs/guides/data-sources--network_interface--reference.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
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

- [dedicated_interface](data-sources--network_interface--properties--dedicated_interface.md): complete subsection reference.

- [dedicated_management_interface](data-sources--network_interface--properties--dedicated_management_interface.md): complete subsection reference.

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

- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md): complete subsection reference.

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

- [layer2_interface](data-sources--network_interface--properties--layer2_interface.md): complete subsection reference.

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

- [tunnel_interface](data-sources--network_interface--properties--tunnel_interface.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_interface--reference.md#schema-annotations) |
| `dedicated_interface` | [dedicated_interface](data-sources--network_interface--properties--dedicated_interface.md#section) |
| `dedicated_interface.cluster` | [dedicated_interface.cluster](data-sources--network_interface--properties--dedicated_interface--cluster.md#section) |
| `dedicated_interface.device` | [dedicated_interface.device](data-sources--network_interface--properties--dedicated_interface.md#schema-dedicated_interface--device) |
| `dedicated_interface.is_primary` | [dedicated_interface.is_primary](data-sources--network_interface--properties--dedicated_interface--is_primary.md#section) |
| `dedicated_interface.monitor` | [dedicated_interface.monitor](data-sources--network_interface--properties--dedicated_interface--monitor.md#section) |
| `dedicated_interface.monitor_disabled` | [dedicated_interface.monitor_disabled](data-sources--network_interface--properties--dedicated_interface--monitor_disabled.md#section) |
| `dedicated_interface.mtu` | [dedicated_interface.mtu](data-sources--network_interface--properties--dedicated_interface.md#schema-dedicated_interface--mtu) |
| `dedicated_interface.node` | [dedicated_interface.node](data-sources--network_interface--properties--dedicated_interface.md#schema-dedicated_interface--node) |
| `dedicated_interface.not_primary` | [dedicated_interface.not_primary](data-sources--network_interface--properties--dedicated_interface--not_primary.md#section) |
| `dedicated_interface.priority` | [dedicated_interface.priority](data-sources--network_interface--properties--dedicated_interface.md#schema-dedicated_interface--priority) |
| `dedicated_management_interface` | [dedicated_management_interface](data-sources--network_interface--properties--dedicated_management_interface.md#section) |
| `dedicated_management_interface.cluster` | [dedicated_management_interface.cluster](data-sources--network_interface--properties--dedicated_management_interface--cluster.md#section) |
| `dedicated_management_interface.device` | [dedicated_management_interface.device](data-sources--network_interface--properties--dedicated_management_interface.md#schema-dedicated_management_interface--device) |
| `dedicated_management_interface.mtu` | [dedicated_management_interface.mtu](data-sources--network_interface--properties--dedicated_management_interface.md#schema-dedicated_management_interface--mtu) |
| `dedicated_management_interface.node` | [dedicated_management_interface.node](data-sources--network_interface--properties--dedicated_management_interface.md#schema-dedicated_management_interface--node) |
| `description` | [description](data-sources--network_interface--reference.md#schema-description) |
| `ethernet_interface` | [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md#section) |
| `ethernet_interface.cluster` | [ethernet_interface.cluster](data-sources--network_interface--properties--ethernet_interface--cluster.md#section) |
| `ethernet_interface.device` | [ethernet_interface.device](data-sources--network_interface--properties--ethernet_interface.md#schema-ethernet_interface--device) |
| `ethernet_interface.dhcp_client` | [ethernet_interface.dhcp_client](data-sources--network_interface--properties--ethernet_interface--dhcp_client.md#section) |
| `ethernet_interface.dhcp_server` | [ethernet_interface.dhcp_server](data-sources--network_interface--properties--ethernet_interface--dhcp_server.md#section) |
| `ethernet_interface.dhcp_server.automatic_from_end` | [ethernet_interface.dhcp_server.automatic_from_end](data-sources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_end.md#section) |
| `ethernet_interface.dhcp_server.automatic_from_start` | [ethernet_interface.dhcp_server.automatic_from_start](data-sources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_start.md#section) |
| `ethernet_interface.dhcp_server.dhcp_networks` | [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md#section) |
| `ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [ethernet_interface.dhcp_server.dhcp_networks.dgw_address](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md#schema-ethernet_interface--dhcp_server--dhcp_networks--dgw_address) |
| `ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [ethernet_interface.dhcp_server.dhcp_networks.dns_address](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md#schema-ethernet_interface--dhcp_server--dhcp_networks--dns_address) |
| `ethernet_interface.dhcp_server.dhcp_networks.first_address` | [ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks--first_address.md#section) |
| `ethernet_interface.dhcp_server.dhcp_networks.last_address` | [ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks--last_address.md#section) |
| `ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [ethernet_interface.dhcp_server.dhcp_networks.network_prefix](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md#schema-ethernet_interface--dhcp_server--dhcp_networks--network_prefix) |
| `ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [ethernet_interface.dhcp_server.dhcp_networks.pool_settings](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md#schema-ethernet_interface--dhcp_server--dhcp_networks--pool_settings) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools` | [ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks--pools.md#section) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-ethernet_interface--dhcp_server--dhcp_networks--pools--end_ip) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-ethernet_interface--dhcp_server--dhcp_networks--pools--exclude) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks--pools.md#schema-ethernet_interface--dhcp_server--dhcp_networks--pools--start_ip) |
| `ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks--same_as_dgw.md#section) |
| `ethernet_interface.dhcp_server.dhcp_option82_tag` | [ethernet_interface.dhcp_server.dhcp_option82_tag](data-sources--network_interface--properties--ethernet_interface--dhcp_server.md#schema-ethernet_interface--dhcp_server--dhcp_option82_tag) |
| `ethernet_interface.dhcp_server.fixed_ip_map` | [ethernet_interface.dhcp_server.fixed_ip_map](data-sources--network_interface--properties--ethernet_interface--dhcp_server.md#schema-ethernet_interface--dhcp_server--fixed_ip_map) |
| `ethernet_interface.dhcp_server.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map](data-sources--network_interface--properties--ethernet_interface--dhcp_server--interface_ip_map.md#section) |
| `ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](data-sources--network_interface--properties--ethernet_interface--dhcp_server--interface_ip_map.md#schema-ethernet_interface--dhcp_server--interface_ip_map--interface_ip_map) |
| `ethernet_interface.ipv6_auto_config` | [ethernet_interface.ipv6_auto_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config.md#section) |
| `ethernet_interface.ipv6_auto_config.host` | [ethernet_interface.ipv6_auto_config.host](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--host.md#section) |
| `ethernet_interface.ipv6_auto_config.router` | [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config` | [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md#schema-ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md#schema-ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--first_address.md#section) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--last_address.md#section) |
| `ethernet_interface.ipv6_auto_config.router.network_prefix` | [ethernet_interface.ipv6_auto_config.router.network_prefix](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md#schema-ethernet_interface--ipv6_auto_config--router--network_prefix) |
| `ethernet_interface.ipv6_auto_config.router.stateful` | [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful.md#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--automatic_from_end.md#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--automatic_from_start.md#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#schema-ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--network_prefix) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks.md#schema-ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pool_settings) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#schema-ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools.md#schema-ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip) |
| `ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful.md#schema-ethernet_interface--ipv6_auto_config--router--stateful--fixed_ip_map) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map.md#section) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map.md#schema-ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map) |
| `ethernet_interface.is_primary` | [ethernet_interface.is_primary](data-sources--network_interface--properties--ethernet_interface--is_primary.md#section) |
| `ethernet_interface.monitor` | [ethernet_interface.monitor](data-sources--network_interface--properties--ethernet_interface--monitor.md#section) |
| `ethernet_interface.monitor_disabled` | [ethernet_interface.monitor_disabled](data-sources--network_interface--properties--ethernet_interface--monitor_disabled.md#section) |
| `ethernet_interface.mtu` | [ethernet_interface.mtu](data-sources--network_interface--properties--ethernet_interface.md#schema-ethernet_interface--mtu) |
| `ethernet_interface.no_ipv6_address` | [ethernet_interface.no_ipv6_address](data-sources--network_interface--properties--ethernet_interface--no_ipv6_address.md#section) |
| `ethernet_interface.node` | [ethernet_interface.node](data-sources--network_interface--properties--ethernet_interface.md#schema-ethernet_interface--node) |
| `ethernet_interface.not_primary` | [ethernet_interface.not_primary](data-sources--network_interface--properties--ethernet_interface--not_primary.md#section) |
| `ethernet_interface.priority` | [ethernet_interface.priority](data-sources--network_interface--properties--ethernet_interface.md#schema-ethernet_interface--priority) |
| `ethernet_interface.site_local_inside_network` | [ethernet_interface.site_local_inside_network](data-sources--network_interface--properties--ethernet_interface--site_local_inside_network.md#section) |
| `ethernet_interface.site_local_network` | [ethernet_interface.site_local_network](data-sources--network_interface--properties--ethernet_interface--site_local_network.md#section) |
| `ethernet_interface.static_ip` | [ethernet_interface.static_ip](data-sources--network_interface--properties--ethernet_interface--static_ip.md#section) |
| `ethernet_interface.static_ip.cluster_static_ip` | [ethernet_interface.static_ip.cluster_static_ip](data-sources--network_interface--properties--ethernet_interface--static_ip--cluster_static_ip.md#section) |
| `ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--network_interface--properties--ethernet_interface--static_ip--cluster_static_ip.md#schema-ethernet_interface--static_ip--cluster_static_ip--interface_ip_map) |
| `ethernet_interface.static_ip.node_static_ip` | [ethernet_interface.static_ip.node_static_ip](data-sources--network_interface--properties--ethernet_interface--static_ip--node_static_ip.md#section) |
| `ethernet_interface.static_ip.node_static_ip.default_gw` | [ethernet_interface.static_ip.node_static_ip.default_gw](data-sources--network_interface--properties--ethernet_interface--static_ip--node_static_ip.md#schema-ethernet_interface--static_ip--node_static_ip--default_gw) |
| `ethernet_interface.static_ip.node_static_ip.dns_server` | [ethernet_interface.static_ip.node_static_ip.dns_server](data-sources--network_interface--properties--ethernet_interface--static_ip--node_static_ip.md#schema-ethernet_interface--static_ip--node_static_ip--dns_server) |
| `ethernet_interface.static_ip.node_static_ip.ip_address` | [ethernet_interface.static_ip.node_static_ip.ip_address](data-sources--network_interface--properties--ethernet_interface--static_ip--node_static_ip.md#schema-ethernet_interface--static_ip--node_static_ip--ip_address) |
| `ethernet_interface.static_ipv6_address` | [ethernet_interface.static_ipv6_address](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address.md#section) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip` | [ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--cluster_static_ip.md#section) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--cluster_static_ip.md#schema-ethernet_interface--static_ipv6_address--cluster_static_ip--interface_ip_map) |
| `ethernet_interface.static_ipv6_address.node_static_ip` | [ethernet_interface.static_ipv6_address.node_static_ip](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--node_static_ip.md#section) |
| `ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [ethernet_interface.static_ipv6_address.node_static_ip.default_gw](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-ethernet_interface--static_ipv6_address--node_static_ip--default_gw) |
| `ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [ethernet_interface.static_ipv6_address.node_static_ip.dns_server](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-ethernet_interface--static_ipv6_address--node_static_ip--dns_server) |
| `ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [ethernet_interface.static_ipv6_address.node_static_ip.ip_address](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--node_static_ip.md#schema-ethernet_interface--static_ipv6_address--node_static_ip--ip_address) |
| `ethernet_interface.storage_network` | [ethernet_interface.storage_network](data-sources--network_interface--properties--ethernet_interface--storage_network.md#section) |
| `ethernet_interface.untagged` | [ethernet_interface.untagged](data-sources--network_interface--properties--ethernet_interface--untagged.md#section) |
| `ethernet_interface.vlan_id` | [ethernet_interface.vlan_id](data-sources--network_interface--properties--ethernet_interface.md#schema-ethernet_interface--vlan_id) |
| `id` | [id](data-sources--network_interface--reference.md#schema-id) |
| `labels` | [labels](data-sources--network_interface--reference.md#schema-labels) |
| `layer2_interface` | [layer2_interface](data-sources--network_interface--properties--layer2_interface.md#section) |
| `layer2_interface.l2sriov_interface` | [layer2_interface.l2sriov_interface](data-sources--network_interface--properties--layer2_interface--l2sriov_interface.md#section) |
| `layer2_interface.l2sriov_interface.device` | [layer2_interface.l2sriov_interface.device](data-sources--network_interface--properties--layer2_interface--l2sriov_interface.md#schema-layer2_interface--l2sriov_interface--device) |
| `layer2_interface.l2sriov_interface.untagged` | [layer2_interface.l2sriov_interface.untagged](data-sources--network_interface--properties--layer2_interface--l2sriov_interface--untagged.md#section) |
| `layer2_interface.l2sriov_interface.vlan_id` | [layer2_interface.l2sriov_interface.vlan_id](data-sources--network_interface--properties--layer2_interface--l2sriov_interface.md#schema-layer2_interface--l2sriov_interface--vlan_id) |
| `layer2_interface.l2vlan_interface` | [layer2_interface.l2vlan_interface](data-sources--network_interface--properties--layer2_interface--l2vlan_interface.md#section) |
| `layer2_interface.l2vlan_interface.device` | [layer2_interface.l2vlan_interface.device](data-sources--network_interface--properties--layer2_interface--l2vlan_interface.md#schema-layer2_interface--l2vlan_interface--device) |
| `layer2_interface.l2vlan_interface.vlan_id` | [layer2_interface.l2vlan_interface.vlan_id](data-sources--network_interface--properties--layer2_interface--l2vlan_interface.md#schema-layer2_interface--l2vlan_interface--vlan_id) |
| `layer2_interface.l2vlan_slo_interface` | [layer2_interface.l2vlan_slo_interface](data-sources--network_interface--properties--layer2_interface--l2vlan_slo_interface.md#section) |
| `layer2_interface.l2vlan_slo_interface.vlan_id` | [layer2_interface.l2vlan_slo_interface.vlan_id](data-sources--network_interface--properties--layer2_interface--l2vlan_slo_interface.md#schema-layer2_interface--l2vlan_slo_interface--vlan_id) |
| `name` | [name](data-sources--network_interface--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--network_interface--reference.md#schema-namespace) |
| `tunnel_interface` | [tunnel_interface](data-sources--network_interface--properties--tunnel_interface.md#section) |
| `tunnel_interface.mtu` | [tunnel_interface.mtu](data-sources--network_interface--properties--tunnel_interface.md#schema-tunnel_interface--mtu) |
| `tunnel_interface.node` | [tunnel_interface.node](data-sources--network_interface--properties--tunnel_interface.md#schema-tunnel_interface--node) |
| `tunnel_interface.priority` | [tunnel_interface.priority](data-sources--network_interface--properties--tunnel_interface.md#schema-tunnel_interface--priority) |
| `tunnel_interface.site_local_inside_network` | [tunnel_interface.site_local_inside_network](data-sources--network_interface--properties--tunnel_interface--site_local_inside_network.md#section) |
| `tunnel_interface.site_local_network` | [tunnel_interface.site_local_network](data-sources--network_interface--properties--tunnel_interface--site_local_network.md#section) |
| `tunnel_interface.static_ip` | [tunnel_interface.static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip.md#section) |
| `tunnel_interface.static_ip.cluster_static_ip` | [tunnel_interface.static_ip.cluster_static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip--cluster_static_ip.md#section) |
| `tunnel_interface.static_ip.cluster_static_ip.interface_ip_map` | [tunnel_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--network_interface--properties--tunnel_interface--static_ip--cluster_static_ip.md#schema-tunnel_interface--static_ip--cluster_static_ip--interface_ip_map) |
| `tunnel_interface.static_ip.node_static_ip` | [tunnel_interface.static_ip.node_static_ip](data-sources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md#section) |
| `tunnel_interface.static_ip.node_static_ip.default_gw` | [tunnel_interface.static_ip.node_static_ip.default_gw](data-sources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md#schema-tunnel_interface--static_ip--node_static_ip--default_gw) |
| `tunnel_interface.static_ip.node_static_ip.dns_server` | [tunnel_interface.static_ip.node_static_ip.dns_server](data-sources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md#schema-tunnel_interface--static_ip--node_static_ip--dns_server) |
| `tunnel_interface.static_ip.node_static_ip.ip_address` | [tunnel_interface.static_ip.node_static_ip.ip_address](data-sources--network_interface--properties--tunnel_interface--static_ip--node_static_ip.md#schema-tunnel_interface--static_ip--node_static_ip--ip_address) |
| `tunnel_interface.tunnel` | [tunnel_interface.tunnel](data-sources--network_interface--properties--tunnel_interface--tunnel.md#section) |
| `tunnel_interface.tunnel.name` | [tunnel_interface.tunnel.name](data-sources--network_interface--properties--tunnel_interface--tunnel.md#schema-tunnel_interface--tunnel--name) |
| `tunnel_interface.tunnel.namespace` | [tunnel_interface.tunnel.namespace](data-sources--network_interface--properties--tunnel_interface--tunnel.md#schema-tunnel_interface--tunnel--namespace) |
| `tunnel_interface.tunnel.tenant` | [tunnel_interface.tunnel.tenant](data-sources--network_interface--properties--tunnel_interface--tunnel.md#schema-tunnel_interface--tunnel--tenant) |

## Next pages

- [dedicated_interface](data-sources--network_interface--properties--dedicated_interface.md)
- [dedicated_management_interface](data-sources--network_interface--properties--dedicated_management_interface.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- [layer2_interface](data-sources--network_interface--properties--layer2_interface.md)
- [tunnel_interface](data-sources--network_interface--properties--tunnel_interface.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
