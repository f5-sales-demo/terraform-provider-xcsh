---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 20983, "body_sha256": "sha256:7cfa80a93b1666e85dd1f6f35a86a0f484e901ecd540103d55def7dacb742b79", "canonical_id": "xcsh-docs:data-sources:bgp:reference", "child_ids": ["xcsh-docs:data-sources:bgp:properties:bgp_parameters", "xcsh-docs:data-sources:bgp:properties:peers", "xcsh-docs:data-sources:bgp:properties:where"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:reference", "parent_id": "xcsh-docs:data-sources:bgp:fundamentals", "path": "docs/guides/data-sources--bgp--reference.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
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

- [bgp_parameters](data-sources--bgp--properties--bgp_parameters.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the BGP.

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

Name of the BGP.

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

Namespace where the BGP exists.

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

- [peers](data-sources--bgp--properties--peers.md): complete subsection reference.

- [where](data-sources--bgp--properties--where.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp--reference.md#schema-annotations) |
| `bgp_parameters` | [bgp_parameters](data-sources--bgp--properties--bgp_parameters.md#section) |
| `bgp_parameters.asn` | [bgp_parameters.asn](data-sources--bgp--properties--bgp_parameters.md#schema-bgp_parameters--asn) |
| `bgp_parameters.from_site` | [bgp_parameters.from_site](data-sources--bgp--properties--bgp_parameters--from_site.md#section) |
| `bgp_parameters.ip_address` | [bgp_parameters.ip_address](data-sources--bgp--properties--bgp_parameters.md#schema-bgp_parameters--ip_address) |
| `bgp_parameters.local_address` | [bgp_parameters.local_address](data-sources--bgp--properties--bgp_parameters--local_address.md#section) |
| `description` | [description](data-sources--bgp--reference.md#schema-description) |
| `id` | [id](data-sources--bgp--reference.md#schema-id) |
| `labels` | [labels](data-sources--bgp--reference.md#schema-labels) |
| `name` | [name](data-sources--bgp--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--bgp--reference.md#schema-namespace) |
| `peers` | [peers](data-sources--bgp--properties--peers.md#section) |
| `peers.bfd_disabled` | [peers.bfd_disabled](data-sources--bgp--properties--peers--bfd_disabled.md#section) |
| `peers.bfd_enabled` | [peers.bfd_enabled](data-sources--bgp--properties--peers--bfd_enabled.md#section) |
| `peers.bfd_enabled.multiplier` | [peers.bfd_enabled.multiplier](data-sources--bgp--properties--peers--bfd_enabled.md#schema-peers--bfd_enabled--multiplier) |
| `peers.bfd_enabled.receive_interval_milliseconds` | [peers.bfd_enabled.receive_interval_milliseconds](data-sources--bgp--properties--peers--bfd_enabled.md#schema-peers--bfd_enabled--receive_interval_milliseconds) |
| `peers.bfd_enabled.transmit_interval_milliseconds` | [peers.bfd_enabled.transmit_interval_milliseconds](data-sources--bgp--properties--peers--bfd_enabled.md#schema-peers--bfd_enabled--transmit_interval_milliseconds) |
| `peers.disable_spec` | [peers.disable_spec](data-sources--bgp--properties--peers--disable_spec.md#section) |
| `peers.ebgp_multihop_disabled` | [peers.ebgp_multihop_disabled](data-sources--bgp--properties--peers--ebgp_multihop_disabled.md#section) |
| `peers.ebgp_multihop_enabled` | [peers.ebgp_multihop_enabled](data-sources--bgp--properties--peers--ebgp_multihop_enabled.md#section) |
| `peers.external` | [peers.external](data-sources--bgp--properties--peers--external.md#section) |
| `peers.external.address` | [peers.external.address](data-sources--bgp--properties--peers--external.md#schema-peers--external--address) |
| `peers.external.address_ipv6` | [peers.external.address_ipv6](data-sources--bgp--properties--peers--external.md#schema-peers--external--address_ipv6) |
| `peers.external.asn` | [peers.external.asn](data-sources--bgp--properties--peers--external.md#schema-peers--external--asn) |
| `peers.external.default_gateway` | [peers.external.default_gateway](data-sources--bgp--properties--peers--external--default_gateway.md#section) |
| `peers.external.default_gateway_v6` | [peers.external.default_gateway_v6](data-sources--bgp--properties--peers--external--default_gateway_v6.md#section) |
| `peers.external.disable_spec` | [peers.external.disable_spec](data-sources--bgp--properties--peers--external--disable_spec.md#section) |
| `peers.external.disable_v6` | [peers.external.disable_v6](data-sources--bgp--properties--peers--external--disable_v6.md#section) |
| `peers.external.external_connector` | [peers.external.external_connector](data-sources--bgp--properties--peers--external--external_connector.md#section) |
| `peers.external.family_inet` | [peers.external.family_inet](data-sources--bgp--properties--peers--external--family_inet.md#section) |
| `peers.external.family_inet.disable_spec` | [peers.external.family_inet.disable_spec](data-sources--bgp--properties--peers--external--family_inet--disable_spec.md#section) |
| `peers.external.family_inet.enable` | [peers.external.family_inet.enable](data-sources--bgp--properties--peers--external--family_inet--enable.md#section) |
| `peers.external.family_inet.enable.aggregation` | [peers.external.family_inet.enable.aggregation](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation.md#section) |
| `peers.external.family_inet.enable.aggregation.ip_prefix` | [peers.external.family_inet.enable.aggregation.ip_prefix](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation.md#schema-peers--external--family_inet--enable--aggregation--ip_prefix) |
| `peers.external.family_inet.enable.aggregation.options` | [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation--options.md#section) |
| `peers.external.family_inet.enable.aggregation.options.summary_only` | [peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--bgp--properties--peers--external--family_inet--enable--aggregation--options--summary_only.md#section) |
| `peers.external.from_site` | [peers.external.from_site](data-sources--bgp--properties--peers--external--from_site.md#section) |
| `peers.external.from_site_v6` | [peers.external.from_site_v6](data-sources--bgp--properties--peers--external--from_site_v6.md#section) |
| `peers.external.interface` | [peers.external.interface](data-sources--bgp--properties--peers--external--interface.md#section) |
| `peers.external.interface.name` | [peers.external.interface.name](data-sources--bgp--properties--peers--external--interface.md#schema-peers--external--interface--name) |
| `peers.external.interface.namespace` | [peers.external.interface.namespace](data-sources--bgp--properties--peers--external--interface.md#schema-peers--external--interface--namespace) |
| `peers.external.interface.tenant` | [peers.external.interface.tenant](data-sources--bgp--properties--peers--external--interface.md#schema-peers--external--interface--tenant) |
| `peers.external.interface_list` | [peers.external.interface_list](data-sources--bgp--properties--peers--external--interface_list.md#section) |
| `peers.external.interface_list.interfaces` | [peers.external.interface_list.interfaces](data-sources--bgp--properties--peers--external--interface_list--interfaces.md#section) |
| `peers.external.interface_list.interfaces.name` | [peers.external.interface_list.interfaces.name](data-sources--bgp--properties--peers--external--interface_list--interfaces.md#schema-peers--external--interface_list--interfaces--name) |
| `peers.external.interface_list.interfaces.namespace` | [peers.external.interface_list.interfaces.namespace](data-sources--bgp--properties--peers--external--interface_list--interfaces.md#schema-peers--external--interface_list--interfaces--namespace) |
| `peers.external.interface_list.interfaces.tenant` | [peers.external.interface_list.interfaces.tenant](data-sources--bgp--properties--peers--external--interface_list--interfaces.md#schema-peers--external--interface_list--interfaces--tenant) |
| `peers.external.md5_auth_key` | [peers.external.md5_auth_key](data-sources--bgp--properties--peers--external.md#schema-peers--external--md5_auth_key) |
| `peers.external.no_authentication` | [peers.external.no_authentication](data-sources--bgp--properties--peers--external--no_authentication.md#section) |
| `peers.external.port` | [peers.external.port](data-sources--bgp--properties--peers--external.md#schema-peers--external--port) |
| `peers.external.subnet_begin_offset` | [peers.external.subnet_begin_offset](data-sources--bgp--properties--peers--external.md#schema-peers--external--subnet_begin_offset) |
| `peers.external.subnet_begin_offset_v6` | [peers.external.subnet_begin_offset_v6](data-sources--bgp--properties--peers--external.md#schema-peers--external--subnet_begin_offset_v6) |
| `peers.external.subnet_end_offset` | [peers.external.subnet_end_offset](data-sources--bgp--properties--peers--external.md#schema-peers--external--subnet_end_offset) |
| `peers.external.subnet_end_offset_v6` | [peers.external.subnet_end_offset_v6](data-sources--bgp--properties--peers--external.md#schema-peers--external--subnet_end_offset_v6) |
| `peers.label` | [peers.label](data-sources--bgp--properties--peers.md#schema-peers--label) |
| `peers.metadata` | [peers.metadata](data-sources--bgp--properties--peers--metadata.md#section) |
| `peers.metadata.description_spec` | [peers.metadata.description_spec](data-sources--bgp--properties--peers--metadata.md#schema-peers--metadata--description_spec) |
| `peers.metadata.name` | [peers.metadata.name](data-sources--bgp--properties--peers--metadata.md#schema-peers--metadata--name) |
| `peers.passive_mode_disabled` | [peers.passive_mode_disabled](data-sources--bgp--properties--peers--passive_mode_disabled.md#section) |
| `peers.passive_mode_enabled` | [peers.passive_mode_enabled](data-sources--bgp--properties--peers--passive_mode_enabled.md#section) |
| `peers.routing_policies` | [peers.routing_policies](data-sources--bgp--properties--peers--routing_policies.md#section) |
| `peers.routing_policies.route_policy` | [peers.routing_policies.route_policy](data-sources--bgp--properties--peers--routing_policies--route_policy.md#section) |
| `peers.routing_policies.route_policy.all_nodes` | [peers.routing_policies.route_policy.all_nodes](data-sources--bgp--properties--peers--routing_policies--route_policy--all_nodes.md#section) |
| `peers.routing_policies.route_policy.inbound` | [peers.routing_policies.route_policy.inbound](data-sources--bgp--properties--peers--routing_policies--route_policy--inbound.md#section) |
| `peers.routing_policies.route_policy.node_name` | [peers.routing_policies.route_policy.node_name](data-sources--bgp--properties--peers--routing_policies--route_policy--node_name.md#section) |
| `peers.routing_policies.route_policy.node_name.node` | [peers.routing_policies.route_policy.node_name.node](data-sources--bgp--properties--peers--routing_policies--route_policy--node_name.md#schema-peers--routing_policies--route_policy--node_name--node) |
| `peers.routing_policies.route_policy.object_refs` | [peers.routing_policies.route_policy.object_refs](data-sources--bgp--properties--peers--routing_policies--route_policy--object_refs.md#section) |
| `peers.routing_policies.route_policy.object_refs.kind` | [peers.routing_policies.route_policy.object_refs.kind](data-sources--bgp--properties--peers--routing_policies--route_policy--object_refs.md#schema-peers--routing_policies--route_policy--object_refs--kind) |
| `peers.routing_policies.route_policy.object_refs.name` | [peers.routing_policies.route_policy.object_refs.name](data-sources--bgp--properties--peers--routing_policies--route_policy--object_refs.md#schema-peers--routing_policies--route_policy--object_refs--name) |
| `peers.routing_policies.route_policy.object_refs.namespace` | [peers.routing_policies.route_policy.object_refs.namespace](data-sources--bgp--properties--peers--routing_policies--route_policy--object_refs.md#schema-peers--routing_policies--route_policy--object_refs--namespace) |
| `peers.routing_policies.route_policy.object_refs.tenant` | [peers.routing_policies.route_policy.object_refs.tenant](data-sources--bgp--properties--peers--routing_policies--route_policy--object_refs.md#schema-peers--routing_policies--route_policy--object_refs--tenant) |
| `peers.routing_policies.route_policy.object_refs.uid` | [peers.routing_policies.route_policy.object_refs.uid](data-sources--bgp--properties--peers--routing_policies--route_policy--object_refs.md#schema-peers--routing_policies--route_policy--object_refs--uid) |
| `peers.routing_policies.route_policy.outbound` | [peers.routing_policies.route_policy.outbound](data-sources--bgp--properties--peers--routing_policies--route_policy--outbound.md#section) |
| `where` | [where](data-sources--bgp--properties--where.md#section) |
| `where.site` | [where.site](data-sources--bgp--properties--where--site.md#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--bgp--properties--where--site--disable_internet_vip.md#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--bgp--properties--where--site--enable_internet_vip.md#section) |
| `where.site.network_type` | [where.site.network_type](data-sources--bgp--properties--where--site.md#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](data-sources--bgp--properties--where--site--ref.md#section) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--bgp--properties--where--site--ref.md#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--bgp--properties--where--site--ref.md#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--bgp--properties--where--site--ref.md#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--bgp--properties--where--site--ref.md#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--bgp--properties--where--site--ref.md#schema-where--site--ref--uid) |
| `where.virtual_site` | [where.virtual_site](data-sources--bgp--properties--where--virtual_site.md#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--bgp--properties--where--virtual_site--disable_internet_vip.md#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--bgp--properties--where--virtual_site--enable_internet_vip.md#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--bgp--properties--where--virtual_site.md#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--bgp--properties--where--virtual_site--ref.md#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--bgp--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--bgp--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--bgp--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--bgp--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--bgp--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--uid) |

## Next pages

- [bgp_parameters](data-sources--bgp--properties--bgp_parameters.md)
- [peers](data-sources--bgp--properties--peers.md)
- [where](data-sources--bgp--properties--where.md)
- [xcsh_bgp](../data-sources/bgp.md)
