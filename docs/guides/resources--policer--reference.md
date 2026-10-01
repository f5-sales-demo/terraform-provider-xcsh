---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_policer."
xcsh_docs: {"aliases": [], "body_bytes": 12922, "body_sha256": "sha256:1c831b99088c79a86d78cbd262af4f507463a538083377ea1641dda9243230c3", "canonical_id": "xcsh-docs:resources:policer:reference", "child_ids": ["xcsh-docs:resources:policer:properties:timeouts"], "collection_id": "xcsh-docs:resources:policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:policer:reference", "parent_id": "xcsh-docs:resources:policer:fundamentals", "path": "docs/guides/resources--policer--reference.md", "provider_name": "policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_policer](../resources/policer.md)
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

<a id="schema-burst_size"></a>

### burst_size property

Type: `"number"`. Required.

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

Upstream description:

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

<a id="schema-committed_information_rate"></a>

### committed_information_rate property

Type: `"number"`. Required.

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Upstream description:

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 10000000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10000000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10000000"
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

Name of the Policer. Must be unique within the namespace.

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

Namespace where the Policer is created.

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

<a id="schema-policer_mode"></a>

### policer_mode property

Type: `"string"`. Optional, Computed.

\[Enum: POLICER\_MODE\_NOT\_SHARED|POLICER\_MODE\_SHARED\] - POLICER\_MODE\_NOT\_SHARED: Not Shared
A separate policer instance is created for each reference to the policer - POLICER\_MODE\_SHARED:
Shared A common policer instance is used for for all references to the policer. Possible values are
\`POLICER\_MODE\_NOT\_SHARED\`, \`POLICER\_MODE\_SHARED\`. Defaults to
\`POLICER\_MODE\_NOT\_SHARED\`. Server applies default when omitted.

Upstream description:

&#8203;- POLICER\_MODE\_NOT\_SHARED: Not Shared

A separate policer instance is created for each reference to the policer &#8203;-
POLICER\_MODE\_SHARED: Shared

A common policer instance is used for for all references to the policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("POLICER_MODE_NOT_SHARED",
    "POLICER_MODE_SHARED"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "POLICER_MODE_NOT_SHARED",
  "enum": [
    "POLICER_MODE_NOT_SHARED",
    "POLICER_MODE_SHARED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-policer_type"></a>

### policer_type property

Type: `"string"`. Optional, Computed.

\[Enum: POLICER\_SINGLE\_RATE\_TWO\_COLOR\] Specifies the type of Policer Basic Single-Rate
Two-Color Policer. The only possible value is \`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Defaults to
\`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Server applies default when omitted.

Upstream description:

Specifies the type of Policer

Basic Single-Rate Two-Color Policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("POLICER_SINGLE_RATE_TWO_COLOR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "POLICER_SINGLE_RATE_TWO_COLOR",
  "enum": [
    "POLICER_SINGLE_RATE_TWO_COLOR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--policer--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--policer--reference.md#schema-annotations) |
| `burst_size` | [burst_size](resources--policer--reference.md#schema-burst_size) |
| `committed_information_rate` | [committed_information_rate](resources--policer--reference.md#schema-committed_information_rate) |
| `description` | [description](resources--policer--reference.md#schema-description) |
| `disable` | [disable](resources--policer--reference.md#schema-disable) |
| `id` | [id](resources--policer--reference.md#schema-id) |
| `labels` | [labels](resources--policer--reference.md#schema-labels) |
| `name` | [name](resources--policer--reference.md#schema-name) |
| `namespace` | [namespace](resources--policer--reference.md#schema-namespace) |
| `policer_mode` | [policer_mode](resources--policer--reference.md#schema-policer_mode) |
| `policer_type` | [policer_type](resources--policer--reference.md#schema-policer_type) |
| `timeouts` | [timeouts](resources--policer--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--policer--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--policer--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--policer--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--policer--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [timeouts](resources--policer--properties--timeouts.md)
- [xcsh_policer](../resources/policer.md)
