---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 15673, "body_sha256": "sha256:97a434ac8b3344b900997516e3a17e6422a5c4d4a273166be4c015bd654d0a77", "canonical_id": "xcsh-docs:data-sources:tunnel:reference", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip", "xcsh-docs:data-sources:tunnel:properties:params", "xcsh-docs:data-sources:tunnel:properties:remote_ip"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:reference", "parent_id": "xcsh-docs:data-sources:tunnel:fundamentals", "path": "docs/guides/data-sources--tunnel--reference.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
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

Description of the Tunnel.

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

- [local_ip](data-sources--tunnel--properties--local_ip.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Tunnel.

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

Namespace where the Tunnel exists.

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

- [params](data-sources--tunnel--properties--params.md): complete subsection reference.

- [remote_ip](data-sources--tunnel--properties--remote_ip.md): complete subsection reference.

<a id="schema-tunnel_type"></a>

### tunnel_type property

Type: `"string"`. Computed.

\[Enum: IPSEC\_PSK|GRE\] Supported tunnel types are IPsec IPsec tunnel type with PSK GRE tunnel
type. Possible values are \`IPSEC\_PSK\`, \`GRE\`. Defaults to \`IPSEC\_PSK\`.

Upstream description:

Supported tunnel types are IPsec

IPsec tunnel type with PSK GRE tunnel type.

Receipt-pinned upstream constraints:

```json
{
  "default": "IPSEC_PSK",
  "enum": [
    "IPSEC_PSK",
    "GRE"
  ],
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
| `annotations` | [annotations](data-sources--tunnel--reference.md#schema-annotations) |
| `description` | [description](data-sources--tunnel--reference.md#schema-description) |
| `id` | [id](data-sources--tunnel--reference.md#schema-id) |
| `labels` | [labels](data-sources--tunnel--reference.md#schema-labels) |
| `local_ip` | [local_ip](data-sources--tunnel--properties--local_ip.md#section) |
| `local_ip.intf` | [local_ip.intf](data-sources--tunnel--properties--local_ip--intf.md#section) |
| `local_ip.intf.local_intf` | [local_ip.intf.local_intf](data-sources--tunnel--properties--local_ip--intf--local_intf.md#section) |
| `local_ip.intf.local_intf.kind` | [local_ip.intf.local_intf.kind](data-sources--tunnel--properties--local_ip--intf--local_intf.md#schema-local_ip--intf--local_intf--kind) |
| `local_ip.intf.local_intf.name` | [local_ip.intf.local_intf.name](data-sources--tunnel--properties--local_ip--intf--local_intf.md#schema-local_ip--intf--local_intf--name) |
| `local_ip.intf.local_intf.namespace` | [local_ip.intf.local_intf.namespace](data-sources--tunnel--properties--local_ip--intf--local_intf.md#schema-local_ip--intf--local_intf--namespace) |
| `local_ip.intf.local_intf.tenant` | [local_ip.intf.local_intf.tenant](data-sources--tunnel--properties--local_ip--intf--local_intf.md#schema-local_ip--intf--local_intf--tenant) |
| `local_ip.intf.local_intf.uid` | [local_ip.intf.local_intf.uid](data-sources--tunnel--properties--local_ip--intf--local_intf.md#schema-local_ip--intf--local_intf--uid) |
| `local_ip.ip_address` | [local_ip.ip_address](data-sources--tunnel--properties--local_ip--ip_address.md#section) |
| `local_ip.ip_address.auto` | [local_ip.ip_address.auto](data-sources--tunnel--properties--local_ip--ip_address--auto.md#section) |
| `local_ip.ip_address.ip_address` | [local_ip.ip_address.ip_address](data-sources--tunnel--properties--local_ip--ip_address--ip_address.md#section) |
| `local_ip.ip_address.ip_address.dual_stack` | [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md#section) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4` | [local_ip.ip_address.ip_address.dual_stack.ipv4](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv4.md#section) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv4.addr](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv4.md#schema-local_ip--ip_address--ip_address--dual_stack--ipv4--addr) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6` | [local_ip.ip_address.ip_address.dual_stack.ipv6](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv6.md#section) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv6.addr](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv6.md#schema-local_ip--ip_address--ip_address--dual_stack--ipv6--addr) |
| `local_ip.ip_address.ip_address.ipv4` | [local_ip.ip_address.ip_address.ipv4](data-sources--tunnel--properties--local_ip--ip_address--ip_address--ipv4.md#section) |
| `local_ip.ip_address.ip_address.ipv4.addr` | [local_ip.ip_address.ip_address.ipv4.addr](data-sources--tunnel--properties--local_ip--ip_address--ip_address--ipv4.md#schema-local_ip--ip_address--ip_address--ipv4--addr) |
| `local_ip.ip_address.ip_address.ipv6` | [local_ip.ip_address.ip_address.ipv6](data-sources--tunnel--properties--local_ip--ip_address--ip_address--ipv6.md#section) |
| `local_ip.ip_address.ip_address.ipv6.addr` | [local_ip.ip_address.ip_address.ipv6.addr](data-sources--tunnel--properties--local_ip--ip_address--ip_address--ipv6.md#schema-local_ip--ip_address--ip_address--ipv6--addr) |
| `local_ip.ip_address.virtual_network_type` | [local_ip.ip_address.virtual_network_type](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type.md#section) |
| `local_ip.ip_address.virtual_network_type.public` | [local_ip.ip_address.virtual_network_type.public](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--public.md#section) |
| `local_ip.ip_address.virtual_network_type.site_local` | [local_ip.ip_address.virtual_network_type.site_local](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--site_local.md#section) |
| `local_ip.ip_address.virtual_network_type.site_local_inside` | [local_ip.ip_address.virtual_network_type.site_local_inside](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--site_local_inside.md#section) |
| `name` | [name](data-sources--tunnel--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--tunnel--reference.md#schema-namespace) |
| `params` | [params](data-sources--tunnel--properties--params.md#section) |
| `params.ipsec` | [params.ipsec](data-sources--tunnel--properties--params--ipsec.md#section) |
| `params.ipsec.ipsec_psk` | [params.ipsec.ipsec_psk](data-sources--tunnel--properties--params--ipsec--ipsec_psk.md#section) |
| `params.ipsec.ipsec_psk.blindfold_secret_info` | [params.ipsec.ipsec_psk.blindfold_secret_info](data-sources--tunnel--properties--params--ipsec--ipsec_psk--blindfold_secret_info.md#section) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider](data-sources--tunnel--properties--params--ipsec--ipsec_psk--blindfold_secret_info.md#schema-params--ipsec--ipsec_psk--blindfold_secret_info--decryption_provider) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.location` | [params.ipsec.ipsec_psk.blindfold_secret_info.location](data-sources--tunnel--properties--params--ipsec--ipsec_psk--blindfold_secret_info.md#schema-params--ipsec--ipsec_psk--blindfold_secret_info--location) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.store_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.store_provider](data-sources--tunnel--properties--params--ipsec--ipsec_psk--blindfold_secret_info.md#schema-params--ipsec--ipsec_psk--blindfold_secret_info--store_provider) |
| `params.ipsec.ipsec_psk.clear_secret_info` | [params.ipsec.ipsec_psk.clear_secret_info](data-sources--tunnel--properties--params--ipsec--ipsec_psk--clear_secret_info.md#section) |
| `params.ipsec.ipsec_psk.clear_secret_info.provider_ref` | [params.ipsec.ipsec_psk.clear_secret_info.provider_ref](data-sources--tunnel--properties--params--ipsec--ipsec_psk--clear_secret_info.md#schema-params--ipsec--ipsec_psk--clear_secret_info--provider_ref) |
| `params.ipsec.ipsec_psk.clear_secret_info.url` | [params.ipsec.ipsec_psk.clear_secret_info.url](data-sources--tunnel--properties--params--ipsec--ipsec_psk--clear_secret_info.md#schema-params--ipsec--ipsec_psk--clear_secret_info--url) |
| `remote_ip` | [remote_ip](data-sources--tunnel--properties--remote_ip.md#section) |
| `remote_ip.endpoints` | [remote_ip.endpoints](data-sources--tunnel--properties--remote_ip--endpoints.md#section) |
| `remote_ip.endpoints.endpoints` | [remote_ip.endpoints.endpoints](data-sources--tunnel--properties--remote_ip--endpoints--endpoints.md#section) |
| `remote_ip.ip` | [remote_ip.ip](data-sources--tunnel--properties--remote_ip--ip.md#section) |
| `remote_ip.ip.dual_stack` | [remote_ip.ip.dual_stack](data-sources--tunnel--properties--remote_ip--ip--dual_stack.md#section) |
| `remote_ip.ip.dual_stack.ipv4` | [remote_ip.ip.dual_stack.ipv4](data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv4.md#section) |
| `remote_ip.ip.dual_stack.ipv4.addr` | [remote_ip.ip.dual_stack.ipv4.addr](data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv4.md#schema-remote_ip--ip--dual_stack--ipv4--addr) |
| `remote_ip.ip.dual_stack.ipv6` | [remote_ip.ip.dual_stack.ipv6](data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv6.md#section) |
| `remote_ip.ip.dual_stack.ipv6.addr` | [remote_ip.ip.dual_stack.ipv6.addr](data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv6.md#schema-remote_ip--ip--dual_stack--ipv6--addr) |
| `remote_ip.ip.ipv4` | [remote_ip.ip.ipv4](data-sources--tunnel--properties--remote_ip--ip--ipv4.md#section) |
| `remote_ip.ip.ipv4.addr` | [remote_ip.ip.ipv4.addr](data-sources--tunnel--properties--remote_ip--ip--ipv4.md#schema-remote_ip--ip--ipv4--addr) |
| `remote_ip.ip.ipv6` | [remote_ip.ip.ipv6](data-sources--tunnel--properties--remote_ip--ip--ipv6.md#section) |
| `remote_ip.ip.ipv6.addr` | [remote_ip.ip.ipv6.addr](data-sources--tunnel--properties--remote_ip--ip--ipv6.md#schema-remote_ip--ip--ipv6--addr) |
| `tunnel_type` | [tunnel_type](data-sources--tunnel--reference.md#schema-tunnel_type) |

## Next pages

- [local_ip](data-sources--tunnel--properties--local_ip.md)
- [params](data-sources--tunnel--properties--params.md)
- [remote_ip](data-sources--tunnel--properties--remote_ip.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
