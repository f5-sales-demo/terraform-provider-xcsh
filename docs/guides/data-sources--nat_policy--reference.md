---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 19849, "body_sha256": "sha256:fb59314693944ea4fb57f6de467be1a39a094f7d9a5c85eba0a749cb5805ca44", "canonical_id": "xcsh-docs:data-sources:nat_policy:reference", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules", "xcsh-docs:data-sources:nat_policy:properties:site"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:reference", "parent_id": "xcsh-docs:data-sources:nat_policy:fundamentals", "path": "docs/guides/data-sources--nat_policy--reference.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md)
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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the NATPolicy.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the NATPolicy.

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

Namespace where the NATPolicy exists.

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

- [rules](data-sources--nat_policy--properties--rules.md): complete subsection reference.

- [site](data-sources--nat_policy--properties--site.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nat_policy--reference.md#schema-annotations) |
| `description` | [description](data-sources--nat_policy--reference.md#schema-description) |
| `id` | [id](data-sources--nat_policy--reference.md#schema-id) |
| `labels` | [labels](data-sources--nat_policy--reference.md#schema-labels) |
| `name` | [name](data-sources--nat_policy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--nat_policy--reference.md#schema-namespace) |
| `rules` | [rules](data-sources--nat_policy--properties--rules.md#section) |
| `rules.action` | [rules.action](data-sources--nat_policy--properties--rules--action.md#section) |
| `rules.action.dynamic` | [rules.action.dynamic](data-sources--nat_policy--properties--rules--action--dynamic.md#section) |
| `rules.action.dynamic.elastic_ips` | [rules.action.dynamic.elastic_ips](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips.md#section) |
| `rules.action.dynamic.elastic_ips.refs` | [rules.action.dynamic.elastic_ips.refs](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips--refs.md#section) |
| `rules.action.dynamic.elastic_ips.refs.kind` | [rules.action.dynamic.elastic_ips.refs.kind](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips--refs.md#schema-rules--action--dynamic--elastic_ips--refs--kind) |
| `rules.action.dynamic.elastic_ips.refs.name` | [rules.action.dynamic.elastic_ips.refs.name](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips--refs.md#schema-rules--action--dynamic--elastic_ips--refs--name) |
| `rules.action.dynamic.elastic_ips.refs.namespace` | [rules.action.dynamic.elastic_ips.refs.namespace](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips--refs.md#schema-rules--action--dynamic--elastic_ips--refs--namespace) |
| `rules.action.dynamic.elastic_ips.refs.tenant` | [rules.action.dynamic.elastic_ips.refs.tenant](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips--refs.md#schema-rules--action--dynamic--elastic_ips--refs--tenant) |
| `rules.action.dynamic.elastic_ips.refs.uid` | [rules.action.dynamic.elastic_ips.refs.uid](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips--refs.md#schema-rules--action--dynamic--elastic_ips--refs--uid) |
| `rules.action.dynamic.pools` | [rules.action.dynamic.pools](data-sources--nat_policy--properties--rules--action--dynamic--pools.md#section) |
| `rules.action.dynamic.pools.prefixes` | [rules.action.dynamic.pools.prefixes](data-sources--nat_policy--properties--rules--action--dynamic--pools.md#schema-rules--action--dynamic--pools--prefixes) |
| `rules.action.virtual_cidr` | [rules.action.virtual_cidr](data-sources--nat_policy--properties--rules--action.md#schema-rules--action--virtual_cidr) |
| `rules.cloud_connect` | [rules.cloud_connect](data-sources--nat_policy--properties--rules--cloud_connect.md#section) |
| `rules.cloud_connect.refs` | [rules.cloud_connect.refs](data-sources--nat_policy--properties--rules--cloud_connect--refs.md#section) |
| `rules.cloud_connect.refs.kind` | [rules.cloud_connect.refs.kind](data-sources--nat_policy--properties--rules--cloud_connect--refs.md#schema-rules--cloud_connect--refs--kind) |
| `rules.cloud_connect.refs.name` | [rules.cloud_connect.refs.name](data-sources--nat_policy--properties--rules--cloud_connect--refs.md#schema-rules--cloud_connect--refs--name) |
| `rules.cloud_connect.refs.namespace` | [rules.cloud_connect.refs.namespace](data-sources--nat_policy--properties--rules--cloud_connect--refs.md#schema-rules--cloud_connect--refs--namespace) |
| `rules.cloud_connect.refs.tenant` | [rules.cloud_connect.refs.tenant](data-sources--nat_policy--properties--rules--cloud_connect--refs.md#schema-rules--cloud_connect--refs--tenant) |
| `rules.cloud_connect.refs.uid` | [rules.cloud_connect.refs.uid](data-sources--nat_policy--properties--rules--cloud_connect--refs.md#schema-rules--cloud_connect--refs--uid) |
| `rules.criteria` | [rules.criteria](data-sources--nat_policy--properties--rules--criteria.md#section) |
| `rules.criteria.any` | [rules.criteria.any](data-sources--nat_policy--properties--rules--criteria--any.md#section) |
| `rules.criteria.destination_cidr` | [rules.criteria.destination_cidr](data-sources--nat_policy--properties--rules--criteria.md#schema-rules--criteria--destination_cidr) |
| `rules.criteria.icmp` | [rules.criteria.icmp](data-sources--nat_policy--properties--rules--criteria--icmp.md#section) |
| `rules.criteria.site_local_inside_network` | [rules.criteria.site_local_inside_network](data-sources--nat_policy--properties--rules--criteria--site_local_inside_network.md#section) |
| `rules.criteria.site_local_network` | [rules.criteria.site_local_network](data-sources--nat_policy--properties--rules--criteria--site_local_network.md#section) |
| `rules.criteria.source_cidr` | [rules.criteria.source_cidr](data-sources--nat_policy--properties--rules--criteria.md#schema-rules--criteria--source_cidr) |
| `rules.criteria.tcp` | [rules.criteria.tcp](data-sources--nat_policy--properties--rules--criteria--tcp.md#section) |
| `rules.criteria.tcp.destination_port` | [rules.criteria.tcp.destination_port](data-sources--nat_policy--properties--rules--criteria--tcp--destination_port.md#section) |
| `rules.criteria.tcp.destination_port.no_port_match` | [rules.criteria.tcp.destination_port.no_port_match](data-sources--nat_policy--properties--rules--criteria--tcp--destination_port--no_port_match.md#section) |
| `rules.criteria.tcp.destination_port.port` | [rules.criteria.tcp.destination_port.port](data-sources--nat_policy--properties--rules--criteria--tcp--destination_port.md#schema-rules--criteria--tcp--destination_port--port) |
| `rules.criteria.tcp.destination_port.port_ranges` | [rules.criteria.tcp.destination_port.port_ranges](data-sources--nat_policy--properties--rules--criteria--tcp--destination_port.md#schema-rules--criteria--tcp--destination_port--port_ranges) |
| `rules.criteria.tcp.source_port` | [rules.criteria.tcp.source_port](data-sources--nat_policy--properties--rules--criteria--tcp--source_port.md#section) |
| `rules.criteria.tcp.source_port.no_port_match` | [rules.criteria.tcp.source_port.no_port_match](data-sources--nat_policy--properties--rules--criteria--tcp--source_port--no_port_match.md#section) |
| `rules.criteria.tcp.source_port.port` | [rules.criteria.tcp.source_port.port](data-sources--nat_policy--properties--rules--criteria--tcp--source_port.md#schema-rules--criteria--tcp--source_port--port) |
| `rules.criteria.tcp.source_port.port_ranges` | [rules.criteria.tcp.source_port.port_ranges](data-sources--nat_policy--properties--rules--criteria--tcp--source_port.md#schema-rules--criteria--tcp--source_port--port_ranges) |
| `rules.criteria.udp` | [rules.criteria.udp](data-sources--nat_policy--properties--rules--criteria--udp.md#section) |
| `rules.criteria.udp.destination_port` | [rules.criteria.udp.destination_port](data-sources--nat_policy--properties--rules--criteria--udp--destination_port.md#section) |
| `rules.criteria.udp.destination_port.no_port_match` | [rules.criteria.udp.destination_port.no_port_match](data-sources--nat_policy--properties--rules--criteria--udp--destination_port--no_port_match.md#section) |
| `rules.criteria.udp.destination_port.port` | [rules.criteria.udp.destination_port.port](data-sources--nat_policy--properties--rules--criteria--udp--destination_port.md#schema-rules--criteria--udp--destination_port--port) |
| `rules.criteria.udp.destination_port.port_ranges` | [rules.criteria.udp.destination_port.port_ranges](data-sources--nat_policy--properties--rules--criteria--udp--destination_port.md#schema-rules--criteria--udp--destination_port--port_ranges) |
| `rules.criteria.udp.source_port` | [rules.criteria.udp.source_port](data-sources--nat_policy--properties--rules--criteria--udp--source_port.md#section) |
| `rules.criteria.udp.source_port.no_port_match` | [rules.criteria.udp.source_port.no_port_match](data-sources--nat_policy--properties--rules--criteria--udp--source_port--no_port_match.md#section) |
| `rules.criteria.udp.source_port.port` | [rules.criteria.udp.source_port.port](data-sources--nat_policy--properties--rules--criteria--udp--source_port.md#schema-rules--criteria--udp--source_port--port) |
| `rules.criteria.udp.source_port.port_ranges` | [rules.criteria.udp.source_port.port_ranges](data-sources--nat_policy--properties--rules--criteria--udp--source_port.md#schema-rules--criteria--udp--source_port--port_ranges) |
| `rules.disable_spec` | [rules.disable_spec](data-sources--nat_policy--properties--rules--disable_spec.md#section) |
| `rules.enable` | [rules.enable](data-sources--nat_policy--properties--rules--enable.md#section) |
| `rules.name` | [rules.name](data-sources--nat_policy--properties--rules.md#schema-rules--name) |
| `rules.node_interface` | [rules.node_interface](data-sources--nat_policy--properties--rules--node_interface.md#section) |
| `rules.node_interface.list` | [rules.node_interface.list](data-sources--nat_policy--properties--rules--node_interface--list.md#section) |
| `rules.node_interface.list.interface` | [rules.node_interface.list.interface](data-sources--nat_policy--properties--rules--node_interface--list--interface.md#section) |
| `rules.node_interface.list.interface.kind` | [rules.node_interface.list.interface.kind](data-sources--nat_policy--properties--rules--node_interface--list--interface.md#schema-rules--node_interface--list--interface--kind) |
| `rules.node_interface.list.interface.name` | [rules.node_interface.list.interface.name](data-sources--nat_policy--properties--rules--node_interface--list--interface.md#schema-rules--node_interface--list--interface--name) |
| `rules.node_interface.list.interface.namespace` | [rules.node_interface.list.interface.namespace](data-sources--nat_policy--properties--rules--node_interface--list--interface.md#schema-rules--node_interface--list--interface--namespace) |
| `rules.node_interface.list.interface.tenant` | [rules.node_interface.list.interface.tenant](data-sources--nat_policy--properties--rules--node_interface--list--interface.md#schema-rules--node_interface--list--interface--tenant) |
| `rules.node_interface.list.interface.uid` | [rules.node_interface.list.interface.uid](data-sources--nat_policy--properties--rules--node_interface--list--interface.md#schema-rules--node_interface--list--interface--uid) |
| `rules.node_interface.list.node` | [rules.node_interface.list.node](data-sources--nat_policy--properties--rules--node_interface--list.md#schema-rules--node_interface--list--node) |
| `rules.segment` | [rules.segment](data-sources--nat_policy--properties--rules--segment.md#section) |
| `rules.segment.refs` | [rules.segment.refs](data-sources--nat_policy--properties--rules--segment--refs.md#section) |
| `rules.segment.refs.kind` | [rules.segment.refs.kind](data-sources--nat_policy--properties--rules--segment--refs.md#schema-rules--segment--refs--kind) |
| `rules.segment.refs.name` | [rules.segment.refs.name](data-sources--nat_policy--properties--rules--segment--refs.md#schema-rules--segment--refs--name) |
| `rules.segment.refs.namespace` | [rules.segment.refs.namespace](data-sources--nat_policy--properties--rules--segment--refs.md#schema-rules--segment--refs--namespace) |
| `rules.segment.refs.tenant` | [rules.segment.refs.tenant](data-sources--nat_policy--properties--rules--segment--refs.md#schema-rules--segment--refs--tenant) |
| `rules.segment.refs.uid` | [rules.segment.refs.uid](data-sources--nat_policy--properties--rules--segment--refs.md#schema-rules--segment--refs--uid) |
| `rules.virtual_network` | [rules.virtual_network](data-sources--nat_policy--properties--rules--virtual_network.md#section) |
| `rules.virtual_network.refs` | [rules.virtual_network.refs](data-sources--nat_policy--properties--rules--virtual_network--refs.md#section) |
| `rules.virtual_network.refs.kind` | [rules.virtual_network.refs.kind](data-sources--nat_policy--properties--rules--virtual_network--refs.md#schema-rules--virtual_network--refs--kind) |
| `rules.virtual_network.refs.name` | [rules.virtual_network.refs.name](data-sources--nat_policy--properties--rules--virtual_network--refs.md#schema-rules--virtual_network--refs--name) |
| `rules.virtual_network.refs.namespace` | [rules.virtual_network.refs.namespace](data-sources--nat_policy--properties--rules--virtual_network--refs.md#schema-rules--virtual_network--refs--namespace) |
| `rules.virtual_network.refs.tenant` | [rules.virtual_network.refs.tenant](data-sources--nat_policy--properties--rules--virtual_network--refs.md#schema-rules--virtual_network--refs--tenant) |
| `rules.virtual_network.refs.uid` | [rules.virtual_network.refs.uid](data-sources--nat_policy--properties--rules--virtual_network--refs.md#schema-rules--virtual_network--refs--uid) |
| `site` | [site](data-sources--nat_policy--properties--site.md#section) |
| `site.refs` | [site.refs](data-sources--nat_policy--properties--site--refs.md#section) |
| `site.refs.kind` | [site.refs.kind](data-sources--nat_policy--properties--site--refs.md#schema-site--refs--kind) |
| `site.refs.name` | [site.refs.name](data-sources--nat_policy--properties--site--refs.md#schema-site--refs--name) |
| `site.refs.namespace` | [site.refs.namespace](data-sources--nat_policy--properties--site--refs.md#schema-site--refs--namespace) |
| `site.refs.tenant` | [site.refs.tenant](data-sources--nat_policy--properties--site--refs.md#schema-site--refs--tenant) |
| `site.refs.uid` | [site.refs.uid](data-sources--nat_policy--properties--site--refs.md#schema-site--refs--uid) |

## Next pages

- [rules](data-sources--nat_policy--properties--rules.md)
- [site](data-sources--nat_policy--properties--site.md)
- [xcsh_nat_policy](../data-sources/nat_policy.md)
