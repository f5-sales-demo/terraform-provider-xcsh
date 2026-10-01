---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cdn_purge_command."
xcsh_docs: {"aliases": [], "body_bytes": 13680, "body_sha256": "sha256:b3c7795d3c71a785f12feb0fc299e36886b721b219276531810ce20a371b85f2", "canonical_id": "xcsh-docs:resources:cdn_purge_command:reference", "child_ids": ["xcsh-docs:resources:cdn_purge_command:properties:hard_purge", "xcsh-docs:resources:cdn_purge_command:properties:purge_all", "xcsh-docs:resources:cdn_purge_command:properties:soft_purge", "xcsh-docs:resources:cdn_purge_command:properties:timeouts", "xcsh-docs:resources:cdn_purge_command:properties:virtual_host"], "collection_id": "xcsh-docs:resources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_purge_command:reference", "parent_id": "xcsh-docs:resources:cdn_purge_command:fundamentals", "path": "docs/guides/resources--cdn_purge_command--reference.md", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_purge_command/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cdn_purge_command.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md)
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

- [hard_purge](resources--cdn_purge_command--properties--hard_purge.md): complete subsection reference.

<a id="schema-hostname"></a>

### hostname property

Type: `"string"`. Optional, Computed.

\[OneOf: hostname, pattern, purge\_all, url\_path\] Exclusive with \[pattern purge\_all url\_path\]
Purge cached content by Hostname.

Upstream description:

Exclusive with \[pattern purge\_all url\_path\] Purge cached content by Hostname.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

OneOf alternatives in this subsection:

- [hostname](resources--cdn_purge_command--reference.md#schema-hostname)
- [pattern](resources--cdn_purge_command--reference.md#schema-pattern)
- [purge_all](resources--cdn_purge_command--properties--purge_all.md#section)
- [url_path](resources--cdn_purge_command--reference.md#schema-url_path)

Select alternatives according to the provider validators above.

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

Name of the CDN Purge Command. Must be unique within the namespace.

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

Namespace where the CDN Purge Command is created.

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

<a id="schema-pattern"></a>

### pattern property

Type: `"string"`. Optional, Computed.

Exclusive with \[hostname purge\_all url\_path\] Purge cached content using PCRE 1 compliant regular
expression.

Upstream description:

Exclusive with \[hostname purge\_all url\_path\] Purge cached content using PCRE 1 compliant regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [purge_all](resources--cdn_purge_command--properties--purge_all.md): complete subsection reference.

- [soft_purge](resources--cdn_purge_command--properties--soft_purge.md): complete subsection reference.

- [timeouts](resources--cdn_purge_command--properties--timeouts.md): complete subsection reference.

<a id="schema-url_path"></a>

### url_path property

Type: `"string"`. Optional, Computed.

Exclusive with \[hostname pattern purge\_all\] Purge cache by using a URL path.

Upstream description:

Exclusive with \[hostname pattern purge\_all\] Purge cache by using a URL path.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [virtual_host](resources--cdn_purge_command--properties--virtual_host.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cdn_purge_command--reference.md#schema-annotations) |
| `description` | [description](resources--cdn_purge_command--reference.md#schema-description) |
| `disable` | [disable](resources--cdn_purge_command--reference.md#schema-disable) |
| `hard_purge` | [hard_purge](resources--cdn_purge_command--properties--hard_purge.md#section) |
| `hostname` | [hostname](resources--cdn_purge_command--reference.md#schema-hostname) |
| `id` | [id](resources--cdn_purge_command--reference.md#schema-id) |
| `labels` | [labels](resources--cdn_purge_command--reference.md#schema-labels) |
| `name` | [name](resources--cdn_purge_command--reference.md#schema-name) |
| `namespace` | [namespace](resources--cdn_purge_command--reference.md#schema-namespace) |
| `pattern` | [pattern](resources--cdn_purge_command--reference.md#schema-pattern) |
| `purge_all` | [purge_all](resources--cdn_purge_command--properties--purge_all.md#section) |
| `soft_purge` | [soft_purge](resources--cdn_purge_command--properties--soft_purge.md#section) |
| `timeouts` | [timeouts](resources--cdn_purge_command--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--cdn_purge_command--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--cdn_purge_command--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--cdn_purge_command--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--cdn_purge_command--properties--timeouts.md#schema-timeouts--update) |
| `url_path` | [url_path](resources--cdn_purge_command--reference.md#schema-url_path) |
| `virtual_host` | [virtual_host](resources--cdn_purge_command--properties--virtual_host.md#section) |
| `virtual_host.name` | [virtual_host.name](resources--cdn_purge_command--properties--virtual_host.md#schema-virtual_host--name) |
| `virtual_host.namespace` | [virtual_host.namespace](resources--cdn_purge_command--properties--virtual_host.md#schema-virtual_host--namespace) |
| `virtual_host.tenant` | [virtual_host.tenant](resources--cdn_purge_command--properties--virtual_host.md#schema-virtual_host--tenant) |

## Next pages

- [hard_purge](resources--cdn_purge_command--properties--hard_purge.md)
- [purge_all](resources--cdn_purge_command--properties--purge_all.md)
- [soft_purge](resources--cdn_purge_command--properties--soft_purge.md)
- [timeouts](resources--cdn_purge_command--properties--timeouts.md)
- [virtual_host](resources--cdn_purge_command--properties--virtual_host.md)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md)
