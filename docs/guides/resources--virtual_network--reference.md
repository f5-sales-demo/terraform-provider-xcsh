---
page_title: "Property reference"
subcategory: "Networking"
description: "Property reference for xcsh_virtual_network."
xcsh_docs: {"aliases": [], "body_bytes": 12189, "body_sha256": "sha256:c8a5036716dcaf285be38bc1bc5c97802ade26d1f5943148deebee920c0f5a0d", "canonical_id": "xcsh-docs:resources:virtual_network:reference", "child_ids": ["xcsh-docs:resources:virtual_network:properties:global_network", "xcsh-docs:resources:virtual_network:properties:site_local_inside_network", "xcsh-docs:resources:virtual_network:properties:site_local_network", "xcsh-docs:resources:virtual_network:properties:static_routes", "xcsh-docs:resources:virtual_network:properties:timeouts"], "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:reference", "parent_id": "xcsh-docs:resources:virtual_network:fundamentals", "path": "docs/guides/resources--virtual_network--reference.md", "provider_name": "virtual_network", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_virtual_network.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md)
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

- [global_network](resources--virtual_network--properties--global_network.md): complete subsection reference.

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

Name of the Virtual Network. Must be unique within the namespace.

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

Type: `"string"`. Optional, Computed.

Namespace for the Virtual Network. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [site_local_inside_network](resources--virtual_network--properties--site_local_inside_network.md): complete subsection reference.

- [site_local_network](resources--virtual_network--properties--site_local_network.md): complete subsection reference.

- [static_routes](resources--virtual_network--properties--static_routes.md): complete subsection reference.

- [timeouts](resources--virtual_network--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--virtual_network--reference.md#schema-annotations) |
| `description` | [description](resources--virtual_network--reference.md#schema-description) |
| `disable` | [disable](resources--virtual_network--reference.md#schema-disable) |
| `global_network` | [global_network](resources--virtual_network--properties--global_network.md#section) |
| `id` | [id](resources--virtual_network--reference.md#schema-id) |
| `labels` | [labels](resources--virtual_network--reference.md#schema-labels) |
| `name` | [name](resources--virtual_network--reference.md#schema-name) |
| `namespace` | [namespace](resources--virtual_network--reference.md#schema-namespace) |
| `site_local_inside_network` | [site_local_inside_network](resources--virtual_network--properties--site_local_inside_network.md#section) |
| `site_local_network` | [site_local_network](resources--virtual_network--properties--site_local_network.md#section) |
| `static_routes` | [static_routes](resources--virtual_network--properties--static_routes.md#section) |
| `static_routes.attrs` | [static_routes.attrs](resources--virtual_network--properties--static_routes.md#schema-static_routes--attrs) |
| `static_routes.default_gateway` | [static_routes.default_gateway](resources--virtual_network--properties--static_routes--default_gateway.md#section) |
| `static_routes.ip_address` | [static_routes.ip_address](resources--virtual_network--properties--static_routes.md#schema-static_routes--ip_address) |
| `static_routes.ip_prefixes` | [static_routes.ip_prefixes](resources--virtual_network--properties--static_routes.md#schema-static_routes--ip_prefixes) |
| `static_routes.node_interface` | [static_routes.node_interface](resources--virtual_network--properties--static_routes--node_interface.md#section) |
| `static_routes.node_interface.list` | [static_routes.node_interface.list](resources--virtual_network--properties--static_routes--node_interface--list.md#section) |
| `static_routes.node_interface.list.interface` | [static_routes.node_interface.list.interface](resources--virtual_network--properties--static_routes--node_interface--list--interface.md#section) |
| `static_routes.node_interface.list.interface.kind` | [static_routes.node_interface.list.interface.kind](resources--virtual_network--properties--static_routes--node_interface--list--interface.md#schema-static_routes--node_interface--list--interface--kind) |
| `static_routes.node_interface.list.interface.name` | [static_routes.node_interface.list.interface.name](resources--virtual_network--properties--static_routes--node_interface--list--interface.md#schema-static_routes--node_interface--list--interface--name) |
| `static_routes.node_interface.list.interface.namespace` | [static_routes.node_interface.list.interface.namespace](resources--virtual_network--properties--static_routes--node_interface--list--interface.md#schema-static_routes--node_interface--list--interface--namespace) |
| `static_routes.node_interface.list.interface.tenant` | [static_routes.node_interface.list.interface.tenant](resources--virtual_network--properties--static_routes--node_interface--list--interface.md#schema-static_routes--node_interface--list--interface--tenant) |
| `static_routes.node_interface.list.interface.uid` | [static_routes.node_interface.list.interface.uid](resources--virtual_network--properties--static_routes--node_interface--list--interface.md#schema-static_routes--node_interface--list--interface--uid) |
| `static_routes.node_interface.list.node` | [static_routes.node_interface.list.node](resources--virtual_network--properties--static_routes--node_interface--list.md#schema-static_routes--node_interface--list--node) |
| `timeouts` | [timeouts](resources--virtual_network--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--virtual_network--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_network--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--virtual_network--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--virtual_network--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [global_network](resources--virtual_network--properties--global_network.md)
- [site_local_inside_network](resources--virtual_network--properties--site_local_inside_network.md)
- [site_local_network](resources--virtual_network--properties--site_local_network.md)
- [static_routes](resources--virtual_network--properties--static_routes.md)
- [timeouts](resources--virtual_network--properties--timeouts.md)
- [xcsh_virtual_network](../resources/virtual_network.md)
