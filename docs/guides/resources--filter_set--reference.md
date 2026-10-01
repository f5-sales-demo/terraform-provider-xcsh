---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 10949, "body_sha256": "sha256:da5d8562fe94a39762fa70cd3db9bca5d40806354f17284a33d8a7bb4ceab39e", "canonical_id": "xcsh-docs:resources:filter_set:reference", "child_ids": ["xcsh-docs:resources:filter_set:properties:filter_fields", "xcsh-docs:resources:filter_set:properties:timeouts"], "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:reference", "parent_id": "xcsh-docs:resources:filter_set:fundamentals", "path": "docs/guides/resources--filter_set--reference.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md)
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

<a id="schema-context_key"></a>

### context_key property

Type: `"string"`. Required.

Indexable context key that identifies a page or page type for which the FilterSet is applicable.

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

- [filter_fields](resources--filter_set--properties--filter_fields.md): complete subsection reference.

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

Name of the Filter Set. Must be unique within the namespace.

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

Namespace where the Filter Set is created.

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

- [timeouts](resources--filter_set--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--filter_set--reference.md#schema-annotations) |
| `context_key` | [context_key](resources--filter_set--reference.md#schema-context_key) |
| `description` | [description](resources--filter_set--reference.md#schema-description) |
| `disable` | [disable](resources--filter_set--reference.md#schema-disable) |
| `filter_fields` | [filter_fields](resources--filter_set--properties--filter_fields.md#section) |
| `filter_fields.date_field` | [filter_fields.date_field](resources--filter_set--properties--filter_fields--date_field.md#section) |
| `filter_fields.date_field.absolute` | [filter_fields.date_field.absolute](resources--filter_set--properties--filter_fields--date_field--absolute.md#section) |
| `filter_fields.date_field.absolute.end_date` | [filter_fields.date_field.absolute.end_date](resources--filter_set--properties--filter_fields--date_field--absolute.md#schema-filter_fields--date_field--absolute--end_date) |
| `filter_fields.date_field.absolute.start_date` | [filter_fields.date_field.absolute.start_date](resources--filter_set--properties--filter_fields--date_field--absolute.md#schema-filter_fields--date_field--absolute--start_date) |
| `filter_fields.date_field.relative` | [filter_fields.date_field.relative](resources--filter_set--properties--filter_fields--date_field.md#schema-filter_fields--date_field--relative) |
| `filter_fields.field_id` | [filter_fields.field_id](resources--filter_set--properties--filter_fields.md#schema-filter_fields--field_id) |
| `filter_fields.filter_expression_field` | [filter_fields.filter_expression_field](resources--filter_set--properties--filter_fields--filter_expression_field.md#section) |
| `filter_fields.filter_expression_field.expression` | [filter_fields.filter_expression_field.expression](resources--filter_set--properties--filter_fields--filter_expression_field.md#schema-filter_fields--filter_expression_field--expression) |
| `filter_fields.string_field` | [filter_fields.string_field](resources--filter_set--properties--filter_fields--string_field.md#section) |
| `filter_fields.string_field.field_values` | [filter_fields.string_field.field_values](resources--filter_set--properties--filter_fields--string_field.md#schema-filter_fields--string_field--field_values) |
| `id` | [id](resources--filter_set--reference.md#schema-id) |
| `labels` | [labels](resources--filter_set--reference.md#schema-labels) |
| `name` | [name](resources--filter_set--reference.md#schema-name) |
| `namespace` | [namespace](resources--filter_set--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--filter_set--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--filter_set--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--filter_set--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--filter_set--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--filter_set--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [filter_fields](resources--filter_set--properties--filter_fields.md)
- [timeouts](resources--filter_set--properties--timeouts.md)
- [xcsh_filter_set](../resources/filter_set.md)
