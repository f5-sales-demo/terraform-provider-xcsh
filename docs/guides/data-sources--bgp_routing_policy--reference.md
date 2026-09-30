---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 9598, "body_sha256": "sha256:c2a10a1cccd9301a0cc17da1c505b8db173501949cf141950dcbe984c23b2b6f", "canonical_id": "xcsh-docs:data-sources:bgp_routing_policy:reference", "child_ids": ["xcsh-docs:data-sources:bgp_routing_policy:properties:rules"], "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:reference", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:fundamentals", "path": "docs/guides/data-sources--bgp_routing_policy--reference.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md)
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

Description of the BGPRoutingPolicy.

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

Name of the BGPRoutingPolicy.

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

Namespace where the BGPRoutingPolicy exists.

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

- [rules](data-sources--bgp_routing_policy--properties--rules.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp_routing_policy--reference.md#schema-annotations) |
| `description` | [description](data-sources--bgp_routing_policy--reference.md#schema-description) |
| `id` | [id](data-sources--bgp_routing_policy--reference.md#schema-id) |
| `labels` | [labels](data-sources--bgp_routing_policy--reference.md#schema-labels) |
| `name` | [name](data-sources--bgp_routing_policy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--bgp_routing_policy--reference.md#schema-namespace) |
| `rules` | [rules](data-sources--bgp_routing_policy--properties--rules.md#section) |
| `rules.action` | [rules.action](data-sources--bgp_routing_policy--properties--rules--action.md#section) |
| `rules.action.allow` | [rules.action.allow](data-sources--bgp_routing_policy--properties--rules--action--allow.md#section) |
| `rules.action.as_path` | [rules.action.as_path](data-sources--bgp_routing_policy--properties--rules--action.md#schema-rules--action--as_path) |
| `rules.action.community` | [rules.action.community](data-sources--bgp_routing_policy--properties--rules--action--community.md#section) |
| `rules.action.community.community` | [rules.action.community.community](data-sources--bgp_routing_policy--properties--rules--action--community.md#schema-rules--action--community--community) |
| `rules.action.deny` | [rules.action.deny](data-sources--bgp_routing_policy--properties--rules--action--deny.md#section) |
| `rules.action.local_preference` | [rules.action.local_preference](data-sources--bgp_routing_policy--properties--rules--action.md#schema-rules--action--local_preference) |
| `rules.action.metric` | [rules.action.metric](data-sources--bgp_routing_policy--properties--rules--action.md#schema-rules--action--metric) |
| `rules.match` | [rules.match](data-sources--bgp_routing_policy--properties--rules--match.md#section) |
| `rules.match.as_path` | [rules.match.as_path](data-sources--bgp_routing_policy--properties--rules--match.md#schema-rules--match--as_path) |
| `rules.match.community` | [rules.match.community](data-sources--bgp_routing_policy--properties--rules--match--community.md#section) |
| `rules.match.community.community` | [rules.match.community.community](data-sources--bgp_routing_policy--properties--rules--match--community.md#schema-rules--match--community--community) |
| `rules.match.ip_prefixes` | [rules.match.ip_prefixes](data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes.md#section) |
| `rules.match.ip_prefixes.prefixes` | [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes.md#section) |
| `rules.match.ip_prefixes.prefixes.equal_or_longer_than` | [rules.match.ip_prefixes.prefixes.equal_or_longer_than](data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes--equal_or_longer_than.md#section) |
| `rules.match.ip_prefixes.prefixes.exact_match` | [rules.match.ip_prefixes.prefixes.exact_match](data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes--exact_match.md#section) |
| `rules.match.ip_prefixes.prefixes.ip_prefixes` | [rules.match.ip_prefixes.prefixes.ip_prefixes](data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes.md#schema-rules--match--ip_prefixes--prefixes--ip_prefixes) |
| `rules.match.ip_prefixes.prefixes.longer_than` | [rules.match.ip_prefixes.prefixes.longer_than](data-sources--bgp_routing_policy--properties--rules--match--ip_prefixes--prefixes--longer_than.md#section) |

## Next pages

- [rules](data-sources--bgp_routing_policy--properties--rules.md)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md)
