---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 21732, "body_sha256": "sha256:b56e3fe8ebd31210b7cc474e508a64ba8a8e81102309df68726f97cc359c07f9", "canonical_id": "xcsh-docs:data-sources:external_connector:reference", "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ce_site_reference", "xcsh-docs:data-sources:external_connector:properties:gre", "xcsh-docs:data-sources:external_connector:properties:ipsec"], "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:reference", "parent_id": "xcsh-docs:data-sources:external_connector:fundamentals", "path": "docs/guides/data-sources--external_connector--reference.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md)
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

- [ce_site_reference](data-sources--external_connector--properties--ce_site_reference.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the ExternalConnector.

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

- [gre](data-sources--external_connector--properties--gre.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipsec](data-sources--external_connector--properties--ipsec.md): complete subsection reference.

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

Name of the ExternalConnector.

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

Namespace where the ExternalConnector exists.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--external_connector--reference.md#schema-annotations) |
| `ce_site_reference` | [ce_site_reference](data-sources--external_connector--properties--ce_site_reference.md#section) |
| `ce_site_reference.name` | [ce_site_reference.name](data-sources--external_connector--properties--ce_site_reference.md#schema-ce_site_reference--name) |
| `ce_site_reference.namespace` | [ce_site_reference.namespace](data-sources--external_connector--properties--ce_site_reference.md#schema-ce_site_reference--namespace) |
| `ce_site_reference.tenant` | [ce_site_reference.tenant](data-sources--external_connector--properties--ce_site_reference.md#schema-ce_site_reference--tenant) |
| `description` | [description](data-sources--external_connector--reference.md#schema-description) |
| `gre` | [gre](data-sources--external_connector--properties--gre.md#section) |
| `gre.gre_parameters` | [gre.gre_parameters](data-sources--external_connector--properties--gre--gre_parameters.md#section) |
| `gre.gre_parameters.peer_ip_address` | [gre.gre_parameters.peer_ip_address](data-sources--external_connector--properties--gre--gre_parameters--peer_ip_address.md#section) |
| `gre.gre_parameters.peer_ip_address.addr` | [gre.gre_parameters.peer_ip_address.addr](data-sources--external_connector--properties--gre--gre_parameters--peer_ip_address.md#schema-gre--gre_parameters--peer_ip_address--addr) |
| `gre.gre_parameters.segment` | [gre.gre_parameters.segment](data-sources--external_connector--properties--gre--gre_parameters--segment.md#section) |
| `gre.gre_parameters.segment.refs` | [gre.gre_parameters.segment.refs](data-sources--external_connector--properties--gre--gre_parameters--segment--refs.md#section) |
| `gre.gre_parameters.segment.refs.kind` | [gre.gre_parameters.segment.refs.kind](data-sources--external_connector--properties--gre--gre_parameters--segment--refs.md#schema-gre--gre_parameters--segment--refs--kind) |
| `gre.gre_parameters.segment.refs.name` | [gre.gre_parameters.segment.refs.name](data-sources--external_connector--properties--gre--gre_parameters--segment--refs.md#schema-gre--gre_parameters--segment--refs--name) |
| `gre.gre_parameters.segment.refs.namespace` | [gre.gre_parameters.segment.refs.namespace](data-sources--external_connector--properties--gre--gre_parameters--segment--refs.md#schema-gre--gre_parameters--segment--refs--namespace) |
| `gre.gre_parameters.segment.refs.tenant` | [gre.gre_parameters.segment.refs.tenant](data-sources--external_connector--properties--gre--gre_parameters--segment--refs.md#schema-gre--gre_parameters--segment--refs--tenant) |
| `gre.gre_parameters.segment.refs.uid` | [gre.gre_parameters.segment.refs.uid](data-sources--external_connector--properties--gre--gre_parameters--segment--refs.md#schema-gre--gre_parameters--segment--refs--uid) |
| `gre.gre_parameters.site_local_inside_network` | [gre.gre_parameters.site_local_inside_network](data-sources--external_connector--properties--gre--gre_parameters--site_local_inside_network.md#section) |
| `gre.gre_parameters.site_local_network` | [gre.gre_parameters.site_local_network](data-sources--external_connector--properties--gre--gre_parameters--site_local_network.md#section) |
| `gre.gre_parameters.tunnel_eps` | [gre.gre_parameters.tunnel_eps](data-sources--external_connector--properties--gre--gre_parameters--tunnel_eps.md#section) |
| `gre.gre_parameters.tunnel_eps.interface` | [gre.gre_parameters.tunnel_eps.interface](data-sources--external_connector--properties--gre--gre_parameters--tunnel_eps.md#schema-gre--gre_parameters--tunnel_eps--interface) |
| `gre.gre_parameters.tunnel_eps.local_tunnel_ip` | [gre.gre_parameters.tunnel_eps.local_tunnel_ip](data-sources--external_connector--properties--gre--gre_parameters--tunnel_eps.md#schema-gre--gre_parameters--tunnel_eps--local_tunnel_ip) |
| `gre.gre_parameters.tunnel_eps.node` | [gre.gre_parameters.tunnel_eps.node](data-sources--external_connector--properties--gre--gre_parameters--tunnel_eps.md#schema-gre--gre_parameters--tunnel_eps--node) |
| `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` | [gre.gre_parameters.tunnel_eps.remote_tunnel_ip](data-sources--external_connector--properties--gre--gre_parameters--tunnel_eps.md#schema-gre--gre_parameters--tunnel_eps--remote_tunnel_ip) |
| `gre.gre_parameters.tunnel_mtu` | [gre.gre_parameters.tunnel_mtu](data-sources--external_connector--properties--gre--gre_parameters.md#schema-gre--gre_parameters--tunnel_mtu) |
| `id` | [id](data-sources--external_connector--reference.md#schema-id) |
| `ipsec` | [ipsec](data-sources--external_connector--properties--ipsec.md#section) |
| `ipsec.ike_parameters` | [ipsec.ike_parameters](data-sources--external_connector--properties--ipsec--ike_parameters.md#section) |
| `ipsec.ike_parameters.dpd_disabled` | [ipsec.ike_parameters.dpd_disabled](data-sources--external_connector--properties--ipsec--ike_parameters--dpd_disabled.md#section) |
| `ipsec.ike_parameters.dpd_keep_alive_timer` | [ipsec.ike_parameters.dpd_keep_alive_timer](data-sources--external_connector--properties--ipsec--ike_parameters--dpd_keep_alive_timer.md#section) |
| `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` | [ipsec.ike_parameters.dpd_keep_alive_timer.timeout](data-sources--external_connector--properties--ipsec--ike_parameters--dpd_keep_alive_timer.md#schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout) |
| `ipsec.ike_parameters.ike_phase1_profile` | [ipsec.ike_parameters.ike_phase1_profile](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase1_profile.md#section) |
| `ipsec.ike_parameters.ike_phase1_profile.name` | [ipsec.ike_parameters.ike_phase1_profile.name](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase1_profile.md#schema-ipsec--ike_parameters--ike_phase1_profile--name) |
| `ipsec.ike_parameters.ike_phase1_profile.namespace` | [ipsec.ike_parameters.ike_phase1_profile.namespace](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase1_profile.md#schema-ipsec--ike_parameters--ike_phase1_profile--namespace) |
| `ipsec.ike_parameters.ike_phase1_profile.tenant` | [ipsec.ike_parameters.ike_phase1_profile.tenant](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase1_profile.md#schema-ipsec--ike_parameters--ike_phase1_profile--tenant) |
| `ipsec.ike_parameters.ike_phase2_profile` | [ipsec.ike_parameters.ike_phase2_profile](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase2_profile.md#section) |
| `ipsec.ike_parameters.ike_phase2_profile.name` | [ipsec.ike_parameters.ike_phase2_profile.name](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase2_profile.md#schema-ipsec--ike_parameters--ike_phase2_profile--name) |
| `ipsec.ike_parameters.ike_phase2_profile.namespace` | [ipsec.ike_parameters.ike_phase2_profile.namespace](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase2_profile.md#schema-ipsec--ike_parameters--ike_phase2_profile--namespace) |
| `ipsec.ike_parameters.ike_phase2_profile.tenant` | [ipsec.ike_parameters.ike_phase2_profile.tenant](data-sources--external_connector--properties--ipsec--ike_parameters--ike_phase2_profile.md#schema-ipsec--ike_parameters--ike_phase2_profile--tenant) |
| `ipsec.ike_parameters.initiator` | [ipsec.ike_parameters.initiator](data-sources--external_connector--properties--ipsec--ike_parameters--initiator.md#section) |
| `ipsec.ike_parameters.responder` | [ipsec.ike_parameters.responder](data-sources--external_connector--properties--ipsec--ike_parameters--responder.md#section) |
| `ipsec.ike_parameters.rm_hostname` | [ipsec.ike_parameters.rm_hostname](data-sources--external_connector--properties--ipsec--ike_parameters.md#schema-ipsec--ike_parameters--rm_hostname) |
| `ipsec.ike_parameters.rm_ip_address` | [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address.md#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack` | [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack.md#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv4.md#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv4.md#schema-ipsec--ike_parameters--rm_ip_address--dual_stack--ipv4--addr) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6.md#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6.md#schema-ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6--addr) |
| `ipsec.ike_parameters.rm_ip_address.ipv4` | [ipsec.ike_parameters.rm_ip_address.ipv4](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv4.md#section) |
| `ipsec.ike_parameters.rm_ip_address.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.ipv4.addr](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv4.md#schema-ipsec--ike_parameters--rm_ip_address--ipv4--addr) |
| `ipsec.ike_parameters.rm_ip_address.ipv6` | [ipsec.ike_parameters.rm_ip_address.ipv6](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv6.md#section) |
| `ipsec.ike_parameters.rm_ip_address.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.ipv6.addr](data-sources--external_connector--properties--ipsec--ike_parameters--rm_ip_address--ipv6.md#schema-ipsec--ike_parameters--rm_ip_address--ipv6--addr) |
| `ipsec.ike_parameters.use_default_local_ike_id` | [ipsec.ike_parameters.use_default_local_ike_id](data-sources--external_connector--properties--ipsec--ike_parameters--use_default_local_ike_id.md#section) |
| `ipsec.ike_parameters.use_default_remote_ike_id` | [ipsec.ike_parameters.use_default_remote_ike_id](data-sources--external_connector--properties--ipsec--ike_parameters--use_default_remote_ike_id.md#section) |
| `ipsec.ipsec_tunnel_parameters` | [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters.md#section) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address` | [ipsec.ipsec_tunnel_parameters.peer_ip_address](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--peer_ip_address.md#section) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` | [ipsec.ipsec_tunnel_parameters.peer_ip_address.addr](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--peer_ip_address.md#schema-ipsec--ipsec_tunnel_parameters--peer_ip_address--addr) |
| `ipsec.ipsec_tunnel_parameters.psk` | [ipsec.ipsec_tunnel_parameters.psk](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters.md#schema-ipsec--ipsec_tunnel_parameters--psk) |
| `ipsec.ipsec_tunnel_parameters.segment` | [ipsec.ipsec_tunnel_parameters.segment](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment.md#section) |
| `ipsec.ipsec_tunnel_parameters.segment.refs` | [ipsec.ipsec_tunnel_parameters.segment.refs](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment--refs.md#section) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.kind` | [ipsec.ipsec_tunnel_parameters.segment.refs.kind](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment--refs.md#schema-ipsec--ipsec_tunnel_parameters--segment--refs--kind) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.name` | [ipsec.ipsec_tunnel_parameters.segment.refs.name](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment--refs.md#schema-ipsec--ipsec_tunnel_parameters--segment--refs--name) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` | [ipsec.ipsec_tunnel_parameters.segment.refs.namespace](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment--refs.md#schema-ipsec--ipsec_tunnel_parameters--segment--refs--namespace) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.tenant` | [ipsec.ipsec_tunnel_parameters.segment.refs.tenant](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment--refs.md#schema-ipsec--ipsec_tunnel_parameters--segment--refs--tenant) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.uid` | [ipsec.ipsec_tunnel_parameters.segment.refs.uid](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--segment--refs.md#schema-ipsec--ipsec_tunnel_parameters--segment--refs--uid) |
| `ipsec.ipsec_tunnel_parameters.site_local_inside_network` | [ipsec.ipsec_tunnel_parameters.site_local_inside_network](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--site_local_inside_network.md#section) |
| `ipsec.ipsec_tunnel_parameters.site_local_network` | [ipsec.ipsec_tunnel_parameters.site_local_network](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--site_local_network.md#section) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps` | [ipsec.ipsec_tunnel_parameters.tunnel_eps](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--tunnel_eps.md#section) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.interface](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--tunnel_eps.md#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--interface) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--tunnel_eps.md#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--local_tunnel_ip) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.node](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--tunnel_eps.md#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--node) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters--tunnel_eps.md#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--remote_tunnel_ip) |
| `ipsec.ipsec_tunnel_parameters.tunnel_mtu` | [ipsec.ipsec_tunnel_parameters.tunnel_mtu](data-sources--external_connector--properties--ipsec--ipsec_tunnel_parameters.md#schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu) |
| `labels` | [labels](data-sources--external_connector--reference.md#schema-labels) |
| `name` | [name](data-sources--external_connector--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--external_connector--reference.md#schema-namespace) |

## Next pages

- [ce_site_reference](data-sources--external_connector--properties--ce_site_reference.md)
- [gre](data-sources--external_connector--properties--gre.md)
- [ipsec](data-sources--external_connector--properties--ipsec.md)
- [xcsh_external_connector](../data-sources/external_connector.md)
