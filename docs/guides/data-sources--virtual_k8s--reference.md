---
page_title: "Property reference"
subcategory: "Container"
description: "Property reference for xcsh_virtual_k8s."
xcsh_docs: {"aliases": [], "body_bytes": 8513, "body_sha256": "sha256:03d89c8876f935f52ee331623edf1d3cc9d990e712033dd20707b62d6fd6e154", "canonical_id": "xcsh-docs:data-sources:virtual_k8s:reference", "child_ids": ["xcsh-docs:data-sources:virtual_k8s:properties:default_flavor_ref", "xcsh-docs:data-sources:virtual_k8s:properties:disabled", "xcsh-docs:data-sources:virtual_k8s:properties:isolated", "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs"], "collection_id": "xcsh-docs:data-sources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_k8s:reference", "parent_id": "xcsh-docs:data-sources:virtual_k8s:fundamentals", "path": "docs/guides/data-sources--virtual_k8s--reference.md", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_k8s/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_virtual_k8s.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md)
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

- [default_flavor_ref](data-sources--virtual_k8s--properties--default_flavor_ref.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the VirtualK8S.

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

- [disabled](data-sources--virtual_k8s--properties--disabled.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated](data-sources--virtual_k8s--properties--isolated.md): complete subsection reference.

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

Name of the VirtualK8S.

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

Namespace where the VirtualK8S exists.

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

- [vsite_refs](data-sources--virtual_k8s--properties--vsite_refs.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--virtual_k8s--reference.md#schema-annotations) |
| `default_flavor_ref` | [default_flavor_ref](data-sources--virtual_k8s--properties--default_flavor_ref.md#section) |
| `default_flavor_ref.name` | [default_flavor_ref.name](data-sources--virtual_k8s--properties--default_flavor_ref.md#schema-default_flavor_ref--name) |
| `default_flavor_ref.namespace` | [default_flavor_ref.namespace](data-sources--virtual_k8s--properties--default_flavor_ref.md#schema-default_flavor_ref--namespace) |
| `default_flavor_ref.tenant` | [default_flavor_ref.tenant](data-sources--virtual_k8s--properties--default_flavor_ref.md#schema-default_flavor_ref--tenant) |
| `description` | [description](data-sources--virtual_k8s--reference.md#schema-description) |
| `disabled` | [disabled](data-sources--virtual_k8s--properties--disabled.md#section) |
| `id` | [id](data-sources--virtual_k8s--reference.md#schema-id) |
| `isolated` | [isolated](data-sources--virtual_k8s--properties--isolated.md#section) |
| `labels` | [labels](data-sources--virtual_k8s--reference.md#schema-labels) |
| `name` | [name](data-sources--virtual_k8s--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--virtual_k8s--reference.md#schema-namespace) |
| `vsite_refs` | [vsite_refs](data-sources--virtual_k8s--properties--vsite_refs.md#section) |
| `vsite_refs.kind` | [vsite_refs.kind](data-sources--virtual_k8s--properties--vsite_refs.md#schema-vsite_refs--kind) |
| `vsite_refs.name` | [vsite_refs.name](data-sources--virtual_k8s--properties--vsite_refs.md#schema-vsite_refs--name) |
| `vsite_refs.namespace` | [vsite_refs.namespace](data-sources--virtual_k8s--properties--vsite_refs.md#schema-vsite_refs--namespace) |
| `vsite_refs.tenant` | [vsite_refs.tenant](data-sources--virtual_k8s--properties--vsite_refs.md#schema-vsite_refs--tenant) |
| `vsite_refs.uid` | [vsite_refs.uid](data-sources--virtual_k8s--properties--vsite_refs.md#schema-vsite_refs--uid) |

## Next pages

- [default_flavor_ref](data-sources--virtual_k8s--properties--default_flavor_ref.md)
- [disabled](data-sources--virtual_k8s--properties--disabled.md)
- [isolated](data-sources--virtual_k8s--properties--isolated.md)
- [vsite_refs](data-sources--virtual_k8s--properties--vsite_refs.md)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md)
