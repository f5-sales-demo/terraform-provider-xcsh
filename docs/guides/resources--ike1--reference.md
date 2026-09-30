---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_ike1."
xcsh_docs: {"aliases": [], "body_bytes": 10261, "body_sha256": "sha256:fb8b8bfc76f5d0b7e30dff010b39b9ce53194e2960790447ebf9e207e24aac16", "canonical_id": "xcsh-docs:resources:ike1:reference", "child_ids": ["xcsh-docs:resources:ike1:properties:ike_keylifetime_hours", "xcsh-docs:resources:ike1:properties:ike_keylifetime_minutes", "xcsh-docs:resources:ike1:properties:reauth_disabled", "xcsh-docs:resources:ike1:properties:reauth_timeout_days", "xcsh-docs:resources:ike1:properties:reauth_timeout_hours", "xcsh-docs:resources:ike1:properties:timeouts", "xcsh-docs:resources:ike1:properties:use_default_keylifetime"], "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:reference", "parent_id": "xcsh-docs:resources:ike1:fundamentals", "path": "docs/guides/resources--ike1--reference.md", "provider_name": "ike1", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_ike1.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike1CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md)
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

- [ike_keylifetime_hours](resources--ike1--properties--ike_keylifetime_hours.md): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike1--properties--ike_keylifetime_minutes.md): complete subsection reference.

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

Name of the Ike1. Must be unique within the namespace.

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

Namespace where the Ike1 is created.

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

- [reauth_disabled](resources--ike1--properties--reauth_disabled.md): complete subsection reference.

- [reauth_timeout_days](resources--ike1--properties--reauth_timeout_days.md): complete subsection reference.

- [reauth_timeout_hours](resources--ike1--properties--reauth_timeout_hours.md): complete subsection reference.

- [timeouts](resources--ike1--properties--timeouts.md): complete subsection reference.

- [use_default_keylifetime](resources--ike1--properties--use_default_keylifetime.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike1--reference.md#schema-annotations) |
| `description` | [description](resources--ike1--reference.md#schema-description) |
| `disable` | [disable](resources--ike1--reference.md#schema-disable) |
| `id` | [id](resources--ike1--reference.md#schema-id) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike1--properties--ike_keylifetime_hours.md#section) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike1--properties--ike_keylifetime_hours.md#schema-ike_keylifetime_hours--duration) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike1--properties--ike_keylifetime_minutes.md#section) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike1--properties--ike_keylifetime_minutes.md#schema-ike_keylifetime_minutes--duration) |
| `labels` | [labels](resources--ike1--reference.md#schema-labels) |
| `name` | [name](resources--ike1--reference.md#schema-name) |
| `namespace` | [namespace](resources--ike1--reference.md#schema-namespace) |
| `reauth_disabled` | [reauth_disabled](resources--ike1--properties--reauth_disabled.md#section) |
| `reauth_timeout_days` | [reauth_timeout_days](resources--ike1--properties--reauth_timeout_days.md#section) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](resources--ike1--properties--reauth_timeout_days.md#schema-reauth_timeout_days--duration) |
| `reauth_timeout_hours` | [reauth_timeout_hours](resources--ike1--properties--reauth_timeout_hours.md#section) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](resources--ike1--properties--reauth_timeout_hours.md#schema-reauth_timeout_hours--duration) |
| `timeouts` | [timeouts](resources--ike1--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--ike1--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--ike1--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--ike1--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--ike1--properties--timeouts.md#schema-timeouts--update) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike1--properties--use_default_keylifetime.md#section) |

## Next pages

- [ike_keylifetime_hours](resources--ike1--properties--ike_keylifetime_hours.md)
- [ike_keylifetime_minutes](resources--ike1--properties--ike_keylifetime_minutes.md)
- [reauth_disabled](resources--ike1--properties--reauth_disabled.md)
- [reauth_timeout_days](resources--ike1--properties--reauth_timeout_days.md)
- [reauth_timeout_hours](resources--ike1--properties--reauth_timeout_hours.md)
- [timeouts](resources--ike1--properties--timeouts.md)
- [use_default_keylifetime](resources--ike1--properties--use_default_keylifetime.md)
- [xcsh_ike1](../resources/ike1.md)
