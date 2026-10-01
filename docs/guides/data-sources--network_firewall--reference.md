---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 13221, "body_sha256": "sha256:ec7792454f98a6eccf7a5d921af449adaa765cd8a7f972b08d63f33f0b94585d", "canonical_id": "xcsh-docs:data-sources:network_firewall:reference", "child_ids": ["xcsh-docs:data-sources:network_firewall:properties:active_enhanced_firewall_policies", "xcsh-docs:data-sources:network_firewall:properties:active_fast_acls", "xcsh-docs:data-sources:network_firewall:properties:active_forward_proxy_policies", "xcsh-docs:data-sources:network_firewall:properties:active_network_policies", "xcsh-docs:data-sources:network_firewall:properties:disable_fast_acl", "xcsh-docs:data-sources:network_firewall:properties:disable_forward_proxy_policy", "xcsh-docs:data-sources:network_firewall:properties:disable_network_policy"], "collection_id": "xcsh-docs:data-sources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_firewall:reference", "parent_id": "xcsh-docs:data-sources:network_firewall:fundamentals", "path": "docs/guides/data-sources--network_firewall--reference.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_firewall/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md)
- Property reference

## Direct properties

- [active_enhanced_firewall_policies](data-sources--network_firewall--properties--active_enhanced_firewall_policies.md): complete subsection reference.

- [active_fast_acls](data-sources--network_firewall--properties--active_fast_acls.md): complete subsection reference.

- [active_forward_proxy_policies](data-sources--network_firewall--properties--active_forward_proxy_policies.md): complete subsection reference.

- [active_network_policies](data-sources--network_firewall--properties--active_network_policies.md): complete subsection reference.

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

Description of the NetworkFirewall.

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

- [disable_fast_acl](data-sources--network_firewall--properties--disable_fast_acl.md): complete subsection reference.

- [disable_forward_proxy_policy](data-sources--network_firewall--properties--disable_forward_proxy_policy.md): complete subsection reference.

- [disable_network_policy](data-sources--network_firewall--properties--disable_network_policy.md): complete subsection reference.

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

Name of the NetworkFirewall.

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

Namespace where the NetworkFirewall exists.

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
| `active_enhanced_firewall_policies` | [active_enhanced_firewall_policies](data-sources--network_firewall--properties--active_enhanced_firewall_policies.md#section) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies` | [active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--network_firewall--properties--active_enhanced_firewall_policies--enhanced_firewall_policies.md#section) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--network_firewall--properties--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--network_firewall--properties--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--network_firewall--properties--active_enhanced_firewall_policies--enhanced_firewall_policies.md#schema-active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `active_fast_acls` | [active_fast_acls](data-sources--network_firewall--properties--active_fast_acls.md#section) |
| `active_fast_acls.fast_acls` | [active_fast_acls.fast_acls](data-sources--network_firewall--properties--active_fast_acls--fast_acls.md#section) |
| `active_fast_acls.fast_acls.name` | [active_fast_acls.fast_acls.name](data-sources--network_firewall--properties--active_fast_acls--fast_acls.md#schema-active_fast_acls--fast_acls--name) |
| `active_fast_acls.fast_acls.namespace` | [active_fast_acls.fast_acls.namespace](data-sources--network_firewall--properties--active_fast_acls--fast_acls.md#schema-active_fast_acls--fast_acls--namespace) |
| `active_fast_acls.fast_acls.tenant` | [active_fast_acls.fast_acls.tenant](data-sources--network_firewall--properties--active_fast_acls--fast_acls.md#schema-active_fast_acls--fast_acls--tenant) |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](data-sources--network_firewall--properties--active_forward_proxy_policies.md#section) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](data-sources--network_firewall--properties--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](data-sources--network_firewall--properties--active_forward_proxy_policies--forward_proxy_policies.md#schema-active_forward_proxy_policies--forward_proxy_policies--name) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--network_firewall--properties--active_forward_proxy_policies--forward_proxy_policies.md#schema-active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--network_firewall--properties--active_forward_proxy_policies--forward_proxy_policies.md#schema-active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `active_network_policies` | [active_network_policies](data-sources--network_firewall--properties--active_network_policies.md#section) |
| `active_network_policies.network_policies` | [active_network_policies.network_policies](data-sources--network_firewall--properties--active_network_policies--network_policies.md#section) |
| `active_network_policies.network_policies.name` | [active_network_policies.network_policies.name](data-sources--network_firewall--properties--active_network_policies--network_policies.md#schema-active_network_policies--network_policies--name) |
| `active_network_policies.network_policies.namespace` | [active_network_policies.network_policies.namespace](data-sources--network_firewall--properties--active_network_policies--network_policies.md#schema-active_network_policies--network_policies--namespace) |
| `active_network_policies.network_policies.tenant` | [active_network_policies.network_policies.tenant](data-sources--network_firewall--properties--active_network_policies--network_policies.md#schema-active_network_policies--network_policies--tenant) |
| `annotations` | [annotations](data-sources--network_firewall--reference.md#schema-annotations) |
| `description` | [description](data-sources--network_firewall--reference.md#schema-description) |
| `disable_fast_acl` | [disable_fast_acl](data-sources--network_firewall--properties--disable_fast_acl.md#section) |
| `disable_forward_proxy_policy` | [disable_forward_proxy_policy](data-sources--network_firewall--properties--disable_forward_proxy_policy.md#section) |
| `disable_network_policy` | [disable_network_policy](data-sources--network_firewall--properties--disable_network_policy.md#section) |
| `id` | [id](data-sources--network_firewall--reference.md#schema-id) |
| `labels` | [labels](data-sources--network_firewall--reference.md#schema-labels) |
| `name` | [name](data-sources--network_firewall--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--network_firewall--reference.md#schema-namespace) |

## Next pages

- [active_enhanced_firewall_policies](data-sources--network_firewall--properties--active_enhanced_firewall_policies.md)
- [active_fast_acls](data-sources--network_firewall--properties--active_fast_acls.md)
- [active_forward_proxy_policies](data-sources--network_firewall--properties--active_forward_proxy_policies.md)
- [active_network_policies](data-sources--network_firewall--properties--active_network_policies.md)
- [disable_fast_acl](data-sources--network_firewall--properties--disable_fast_acl.md)
- [disable_forward_proxy_policy](data-sources--network_firewall--properties--disable_forward_proxy_policy.md)
- [disable_network_policy](data-sources--network_firewall--properties--disable_network_policy.md)
- [xcsh_network_firewall](../data-sources/network_firewall.md)
