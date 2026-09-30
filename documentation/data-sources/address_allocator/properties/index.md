---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_address_allocator."
xcsh_docs: {"aliases": [], "body_bytes": 10490, "body_sha256": "sha256:0a63e89bf54dd2fd8f797f70db7d1de6f4f55e23d1b5925ca571af0f3cd5d73d", "child_ids": ["xcsh-docs:data-sources:address_allocator:properties:address_allocation_scheme"], "collection_id": "xcsh-docs:data-sources:address_allocator:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:address_allocator:reference", "parent_id": "xcsh-docs:data-sources:address_allocator:fundamentals", "path": "documentation/data-sources/address_allocator/properties/index.md", "provider_name": "address_allocator", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/address_allocator/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_address_allocator.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/)
- Property reference

## Direct properties

- [address_allocation_scheme](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/address_allocation_scheme/): complete subsection reference.

<a id="schema-address_pool"></a>

### address_pool property

Type: `["list", "string"]`. Computed.

Address pool from which the allocator carves out subnets or addresses to its clients.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

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

Description of the AddressAllocator.

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

<a id="schema-mode"></a>

### mode property

Type: `"string"`. Computed.

\[Enum: LOCAL|GLOBAL\_PER\_SITE\_NODE\] Mode of the address allocator Address allocator is for VERs
within the local cluster or site Allocation is per site and then per node. Possible values are
\`LOCAL\`, \`GLOBAL\_PER\_SITE\_NODE\`. Defaults to \`LOCAL\`.

Upstream description:

Mode of the address allocator

Address allocator is for VERs within the local cluster or site Allocation is per site and then per
node.

Receipt-pinned upstream constraints:

```json
{
  "default": "LOCAL",
  "enum": [
    "LOCAL",
    "GLOBAL_PER_SITE_NODE"
  ],
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

Name of the AddressAllocator.

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

Namespace where the AddressAllocator exists.

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
| `address_allocation_scheme` | [address_allocation_scheme](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/address_allocation_scheme/#section) |
| `address_allocation_scheme.allocation_unit` | [address_allocation_scheme.allocation_unit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/address_allocation_scheme/#schema-address_allocation_scheme--allocation_unit) |
| `address_allocation_scheme.local_interface_address_offset` | [address_allocation_scheme.local_interface_address_offset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/address_allocation_scheme/#schema-address_allocation_scheme--local_interface_address_offset) |
| `address_allocation_scheme.local_interface_address_type` | [address_allocation_scheme.local_interface_address_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/address_allocation_scheme/#schema-address_allocation_scheme--local_interface_address_type) |
| `address_pool` | [address_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/#schema-address_pool) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/#schema-labels) |
| `mode` | [mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/#schema-mode) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/#schema-namespace) |

## Next pages

- [address_allocation_scheme](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/properties/address_allocation_scheme/)
- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/address_allocator/)
