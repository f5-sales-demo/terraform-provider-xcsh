---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 25909, "body_sha256": "sha256:bb5cabe49a07245d65412322e895e596e1e9b6f37bc16b7ab88cc0d69c91fbe2", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules", "xcsh-docs:resources:nat_policy:properties:site", "xcsh-docs:resources:nat_policy:properties:timeouts"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:reference", "parent_id": "xcsh-docs:resources:nat_policy:fundamentals", "path": "documentation/resources/nat_policy/properties/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

Type: `"string"`. Optional.

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

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Name of the NAT Policy. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

Namespace where the NAT Policy is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/): complete subsection reference.

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-namespace) |
| `rules` | [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/#section) |
| `rules.action` | [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/#section) |
| `rules.action.dynamic` | [rules.action.dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/#section) |
| `rules.action.dynamic.elastic_ips` | [rules.action.dynamic.elastic_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/#section) |
| `rules.action.dynamic.elastic_ips.refs` | [rules.action.dynamic.elastic_ips.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#section) |
| `rules.action.dynamic.elastic_ips.refs.kind` | [rules.action.dynamic.elastic_ips.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--kind) |
| `rules.action.dynamic.elastic_ips.refs.name` | [rules.action.dynamic.elastic_ips.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--name) |
| `rules.action.dynamic.elastic_ips.refs.namespace` | [rules.action.dynamic.elastic_ips.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--namespace) |
| `rules.action.dynamic.elastic_ips.refs.tenant` | [rules.action.dynamic.elastic_ips.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--tenant) |
| `rules.action.dynamic.elastic_ips.refs.uid` | [rules.action.dynamic.elastic_ips.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--uid) |
| `rules.action.dynamic.pools` | [rules.action.dynamic.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/pools/#section) |
| `rules.action.dynamic.pools.prefixes` | [rules.action.dynamic.pools.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/pools/#schema-rules--action--dynamic--pools--prefixes) |
| `rules.action.virtual_cidr` | [rules.action.virtual_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/#schema-rules--action--virtual_cidr) |
| `rules.cloud_connect` | [rules.cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/#section) |
| `rules.cloud_connect.refs` | [rules.cloud_connect.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#section) |
| `rules.cloud_connect.refs.kind` | [rules.cloud_connect.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--kind) |
| `rules.cloud_connect.refs.name` | [rules.cloud_connect.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--name) |
| `rules.cloud_connect.refs.namespace` | [rules.cloud_connect.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--namespace) |
| `rules.cloud_connect.refs.tenant` | [rules.cloud_connect.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--tenant) |
| `rules.cloud_connect.refs.uid` | [rules.cloud_connect.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--uid) |
| `rules.criteria` | [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/#section) |
| `rules.criteria.any` | [rules.criteria.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/any/#section) |
| `rules.criteria.destination_cidr` | [rules.criteria.destination_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/#schema-rules--criteria--destination_cidr) |
| `rules.criteria.icmp` | [rules.criteria.icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/icmp/#section) |
| `rules.criteria.site_local_inside_network` | [rules.criteria.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/site_local_inside_network/#section) |
| `rules.criteria.site_local_network` | [rules.criteria.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/site_local_network/#section) |
| `rules.criteria.source_cidr` | [rules.criteria.source_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/#schema-rules--criteria--source_cidr) |
| `rules.criteria.tcp` | [rules.criteria.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/#section) |
| `rules.criteria.tcp.destination_port` | [rules.criteria.tcp.destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/#section) |
| `rules.criteria.tcp.destination_port.no_port_match` | [rules.criteria.tcp.destination_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/no_port_match/#section) |
| `rules.criteria.tcp.destination_port.port` | [rules.criteria.tcp.destination_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/#schema-rules--criteria--tcp--destination_port--port) |
| `rules.criteria.tcp.destination_port.port_ranges` | [rules.criteria.tcp.destination_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/#schema-rules--criteria--tcp--destination_port--port_ranges) |
| `rules.criteria.tcp.source_port` | [rules.criteria.tcp.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/#section) |
| `rules.criteria.tcp.source_port.no_port_match` | [rules.criteria.tcp.source_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/no_port_match/#section) |
| `rules.criteria.tcp.source_port.port` | [rules.criteria.tcp.source_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/#schema-rules--criteria--tcp--source_port--port) |
| `rules.criteria.tcp.source_port.port_ranges` | [rules.criteria.tcp.source_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/#schema-rules--criteria--tcp--source_port--port_ranges) |
| `rules.criteria.udp` | [rules.criteria.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/#section) |
| `rules.criteria.udp.destination_port` | [rules.criteria.udp.destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/destination_port/#section) |
| `rules.criteria.udp.destination_port.no_port_match` | [rules.criteria.udp.destination_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/destination_port/no_port_match/#section) |
| `rules.criteria.udp.destination_port.port` | [rules.criteria.udp.destination_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/destination_port/#schema-rules--criteria--udp--destination_port--port) |
| `rules.criteria.udp.destination_port.port_ranges` | [rules.criteria.udp.destination_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/destination_port/#schema-rules--criteria--udp--destination_port--port_ranges) |
| `rules.criteria.udp.source_port` | [rules.criteria.udp.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/#section) |
| `rules.criteria.udp.source_port.no_port_match` | [rules.criteria.udp.source_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/no_port_match/#section) |
| `rules.criteria.udp.source_port.port` | [rules.criteria.udp.source_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/#schema-rules--criteria--udp--source_port--port) |
| `rules.criteria.udp.source_port.port_ranges` | [rules.criteria.udp.source_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/#schema-rules--criteria--udp--source_port--port_ranges) |
| `rules.disable_spec` | [rules.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/disable_spec/#section) |
| `rules.enable` | [rules.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/enable/#section) |
| `rules.name` | [rules.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/#schema-rules--name) |
| `rules.node_interface` | [rules.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/#section) |
| `rules.node_interface.list` | [rules.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/#section) |
| `rules.node_interface.list.interface` | [rules.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#section) |
| `rules.node_interface.list.interface.kind` | [rules.node_interface.list.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--kind) |
| `rules.node_interface.list.interface.name` | [rules.node_interface.list.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--name) |
| `rules.node_interface.list.interface.namespace` | [rules.node_interface.list.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--namespace) |
| `rules.node_interface.list.interface.tenant` | [rules.node_interface.list.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--tenant) |
| `rules.node_interface.list.interface.uid` | [rules.node_interface.list.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--uid) |
| `rules.node_interface.list.node` | [rules.node_interface.list.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/#schema-rules--node_interface--list--node) |
| `rules.segment` | [rules.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/#section) |
| `rules.segment.refs` | [rules.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#section) |
| `rules.segment.refs.kind` | [rules.segment.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--kind) |
| `rules.segment.refs.name` | [rules.segment.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--name) |
| `rules.segment.refs.namespace` | [rules.segment.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--namespace) |
| `rules.segment.refs.tenant` | [rules.segment.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--tenant) |
| `rules.segment.refs.uid` | [rules.segment.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--uid) |
| `rules.virtual_network` | [rules.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/#section) |
| `rules.virtual_network.refs` | [rules.virtual_network.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#section) |
| `rules.virtual_network.refs.kind` | [rules.virtual_network.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--kind) |
| `rules.virtual_network.refs.name` | [rules.virtual_network.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--name) |
| `rules.virtual_network.refs.namespace` | [rules.virtual_network.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--namespace) |
| `rules.virtual_network.refs.tenant` | [rules.virtual_network.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--tenant) |
| `rules.virtual_network.refs.uid` | [rules.virtual_network.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--uid) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/#section) |
| `site.refs` | [site.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#section) |
| `site.refs.kind` | [site.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--kind) |
| `site.refs.name` | [site.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--name) |
| `site.refs.namespace` | [site.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--namespace) |
| `site.refs.tenant` | [site.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--tenant) |
| `site.refs.uid` | [site.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--uid) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
