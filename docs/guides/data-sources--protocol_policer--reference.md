---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 9098, "body_sha256": "sha256:78416b068debda6bf935bab789fdad42ca39a35b853900450aecf4d8eeb49013", "canonical_id": "xcsh-docs:data-sources:protocol_policer:reference", "child_ids": ["xcsh-docs:data-sources:protocol_policer:properties:protocol_policer"], "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:reference", "parent_id": "xcsh-docs:data-sources:protocol_policer:fundamentals", "path": "docs/guides/data-sources--protocol_policer--reference.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md)
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

Description of the ProtocolPolicer.

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

Name of the ProtocolPolicer.

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

Type: `"string"`. Optional, Computed.

Namespace where the ProtocolPolicer exists.

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

- [protocol_policer](data-sources--protocol_policer--properties--protocol_policer.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--protocol_policer--reference.md#schema-annotations) |
| `description` | [description](data-sources--protocol_policer--reference.md#schema-description) |
| `id` | [id](data-sources--protocol_policer--reference.md#schema-id) |
| `labels` | [labels](data-sources--protocol_policer--reference.md#schema-labels) |
| `name` | [name](data-sources--protocol_policer--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--protocol_policer--reference.md#schema-namespace) |
| `protocol_policer` | [protocol_policer](data-sources--protocol_policer--properties--protocol_policer.md#section) |
| `protocol_policer.policer` | [protocol_policer.policer](data-sources--protocol_policer--properties--protocol_policer--policer.md#section) |
| `protocol_policer.policer.kind` | [protocol_policer.policer.kind](data-sources--protocol_policer--properties--protocol_policer--policer.md#schema-protocol_policer--policer--kind) |
| `protocol_policer.policer.name` | [protocol_policer.policer.name](data-sources--protocol_policer--properties--protocol_policer--policer.md#schema-protocol_policer--policer--name) |
| `protocol_policer.policer.namespace` | [protocol_policer.policer.namespace](data-sources--protocol_policer--properties--protocol_policer--policer.md#schema-protocol_policer--policer--namespace) |
| `protocol_policer.policer.tenant` | [protocol_policer.policer.tenant](data-sources--protocol_policer--properties--protocol_policer--policer.md#schema-protocol_policer--policer--tenant) |
| `protocol_policer.policer.uid` | [protocol_policer.policer.uid](data-sources--protocol_policer--properties--protocol_policer--policer.md#schema-protocol_policer--policer--uid) |
| `protocol_policer.protocol` | [protocol_policer.protocol](data-sources--protocol_policer--properties--protocol_policer--protocol.md#section) |
| `protocol_policer.protocol.dns` | [protocol_policer.protocol.dns](data-sources--protocol_policer--properties--protocol_policer--protocol--dns.md#section) |
| `protocol_policer.protocol.icmp` | [protocol_policer.protocol.icmp](data-sources--protocol_policer--properties--protocol_policer--protocol--icmp.md#section) |
| `protocol_policer.protocol.icmp.type` | [protocol_policer.protocol.icmp.type](data-sources--protocol_policer--properties--protocol_policer--protocol--icmp.md#schema-protocol_policer--protocol--icmp--type) |
| `protocol_policer.protocol.tcp` | [protocol_policer.protocol.tcp](data-sources--protocol_policer--properties--protocol_policer--protocol--tcp.md#section) |
| `protocol_policer.protocol.tcp.flags` | [protocol_policer.protocol.tcp.flags](data-sources--protocol_policer--properties--protocol_policer--protocol--tcp.md#schema-protocol_policer--protocol--tcp--flags) |
| `protocol_policer.protocol.udp` | [protocol_policer.protocol.udp](data-sources--protocol_policer--properties--protocol_policer--protocol--udp.md#section) |

## Next pages

- [protocol_policer](data-sources--protocol_policer--properties--protocol_policer.md)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md)
