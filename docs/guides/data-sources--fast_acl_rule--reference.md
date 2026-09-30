---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": [], "body_bytes": 11481, "body_sha256": "sha256:3be4980124a25a200819379c21debc6c82b60a319f69f2d12af206207e8aba0c", "canonical_id": "xcsh-docs:data-sources:fast_acl_rule:reference", "child_ids": ["xcsh-docs:data-sources:fast_acl_rule:properties:action", "xcsh-docs:data-sources:fast_acl_rule:properties:ip_prefix_set", "xcsh-docs:data-sources:fast_acl_rule:properties:port", "xcsh-docs:data-sources:fast_acl_rule:properties:prefix"], "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl_rule:reference", "parent_id": "xcsh-docs:data-sources:fast_acl_rule:fundamentals", "path": "docs/guides/data-sources--fast_acl_rule--reference.md", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_fast_acl_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
- Property reference

## Direct properties

- [action](data-sources--fast_acl_rule--properties--action.md): complete subsection reference.

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

Description of the FastACLRule.

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

- [ip_prefix_set](data-sources--fast_acl_rule--properties--ip_prefix_set.md): complete subsection reference.

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

Name of the FastACLRule.

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

Namespace where the FastACLRule exists.

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

- [port](data-sources--fast_acl_rule--properties--port.md): complete subsection reference.

- [prefix](data-sources--fast_acl_rule--properties--prefix.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--fast_acl_rule--properties--action.md#section) |
| `action.policer_action` | [action.policer_action](data-sources--fast_acl_rule--properties--action--policer_action.md#section) |
| `action.policer_action.ref` | [action.policer_action.ref](data-sources--fast_acl_rule--properties--action--policer_action--ref.md#section) |
| `action.policer_action.ref.kind` | [action.policer_action.ref.kind](data-sources--fast_acl_rule--properties--action--policer_action--ref.md#schema-action--policer_action--ref--kind) |
| `action.policer_action.ref.name` | [action.policer_action.ref.name](data-sources--fast_acl_rule--properties--action--policer_action--ref.md#schema-action--policer_action--ref--name) |
| `action.policer_action.ref.namespace` | [action.policer_action.ref.namespace](data-sources--fast_acl_rule--properties--action--policer_action--ref.md#schema-action--policer_action--ref--namespace) |
| `action.policer_action.ref.tenant` | [action.policer_action.ref.tenant](data-sources--fast_acl_rule--properties--action--policer_action--ref.md#schema-action--policer_action--ref--tenant) |
| `action.policer_action.ref.uid` | [action.policer_action.ref.uid](data-sources--fast_acl_rule--properties--action--policer_action--ref.md#schema-action--policer_action--ref--uid) |
| `action.protocol_policer_action` | [action.protocol_policer_action](data-sources--fast_acl_rule--properties--action--protocol_policer_action.md#section) |
| `action.protocol_policer_action.ref` | [action.protocol_policer_action.ref](data-sources--fast_acl_rule--properties--action--protocol_policer_action--ref.md#section) |
| `action.protocol_policer_action.ref.kind` | [action.protocol_policer_action.ref.kind](data-sources--fast_acl_rule--properties--action--protocol_policer_action--ref.md#schema-action--protocol_policer_action--ref--kind) |
| `action.protocol_policer_action.ref.name` | [action.protocol_policer_action.ref.name](data-sources--fast_acl_rule--properties--action--protocol_policer_action--ref.md#schema-action--protocol_policer_action--ref--name) |
| `action.protocol_policer_action.ref.namespace` | [action.protocol_policer_action.ref.namespace](data-sources--fast_acl_rule--properties--action--protocol_policer_action--ref.md#schema-action--protocol_policer_action--ref--namespace) |
| `action.protocol_policer_action.ref.tenant` | [action.protocol_policer_action.ref.tenant](data-sources--fast_acl_rule--properties--action--protocol_policer_action--ref.md#schema-action--protocol_policer_action--ref--tenant) |
| `action.protocol_policer_action.ref.uid` | [action.protocol_policer_action.ref.uid](data-sources--fast_acl_rule--properties--action--protocol_policer_action--ref.md#schema-action--protocol_policer_action--ref--uid) |
| `action.simple_action` | [action.simple_action](data-sources--fast_acl_rule--properties--action.md#schema-action--simple_action) |
| `annotations` | [annotations](data-sources--fast_acl_rule--reference.md#schema-annotations) |
| `description` | [description](data-sources--fast_acl_rule--reference.md#schema-description) |
| `id` | [id](data-sources--fast_acl_rule--reference.md#schema-id) |
| `ip_prefix_set` | [ip_prefix_set](data-sources--fast_acl_rule--properties--ip_prefix_set.md#section) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](data-sources--fast_acl_rule--properties--ip_prefix_set--ref.md#section) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](data-sources--fast_acl_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--kind) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](data-sources--fast_acl_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--name) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](data-sources--fast_acl_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--namespace) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](data-sources--fast_acl_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--tenant) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](data-sources--fast_acl_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--uid) |
| `labels` | [labels](data-sources--fast_acl_rule--reference.md#schema-labels) |
| `name` | [name](data-sources--fast_acl_rule--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--fast_acl_rule--reference.md#schema-namespace) |
| `port` | [port](data-sources--fast_acl_rule--properties--port.md#section) |
| `port.all` | [port.all](data-sources--fast_acl_rule--properties--port--all.md#section) |
| `port.dns` | [port.dns](data-sources--fast_acl_rule--properties--port--dns.md#section) |
| `port.user_defined` | [port.user_defined](data-sources--fast_acl_rule--properties--port.md#schema-port--user_defined) |
| `prefix` | [prefix](data-sources--fast_acl_rule--properties--prefix.md#section) |
| `prefix.prefix` | [prefix.prefix](data-sources--fast_acl_rule--properties--prefix.md#schema-prefix--prefix) |

## Next pages

- [action](data-sources--fast_acl_rule--properties--action.md)
- [ip_prefix_set](data-sources--fast_acl_rule--properties--ip_prefix_set.md)
- [port](data-sources--fast_acl_rule--properties--port.md)
- [prefix](data-sources--fast_acl_rule--properties--prefix.md)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
