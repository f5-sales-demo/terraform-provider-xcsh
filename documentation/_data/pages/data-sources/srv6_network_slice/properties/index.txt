---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": [], "body_bytes": 11110, "body_sha256": "sha256:4e9e60d5487ad0c08056f17d9c4daf23ef2a651b7e1ab0506d715b3e3a1b2be9", "child_ids": [], "collection_id": "xcsh-docs:data-sources:srv6_network_slice:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:srv6_network_slice:reference", "parent_id": "xcsh-docs:data-sources:srv6_network_slice:fundamentals", "path": "documentation/data-sources/srv6_network_slice/properties/index.md", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/srv6_network_slice/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_srv6_network_slice.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_srv6_network_slice](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/)
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

<a id="schema-connect_to_access_networks"></a>

### connect_to_access_networks property

Type: `"bool"`. Computed.

Connect all SRv6 Virtual Networks in this slice to their corresponding access networks by importing
route targets specified in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to their corresponding access networks by importing
route targets specified in the virtual network.

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

<a id="schema-connect_to_enterprise_networks"></a>

### connect_to_enterprise_networks property

Type: `"bool"`. Computed.

Connect all SRv6 Virtual Networks in this slice to their corresponding enterprise networks by
importing route targets specified in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to their corresponding enterprise networks by
importing route targets specified in the virtual network.

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

<a id="schema-connect_to_internet"></a>

### connect_to_internet property

Type: `"bool"`. Computed.

Connect all SRv6 Virtual Networks in this slice to the Internet by importing route targets specified
in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to the Internet by importing route targets specified
in the virtual network.

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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Srv6NetworkSlice.

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

Name of the Srv6NetworkSlice.

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

Namespace where the Srv6NetworkSlice exists.

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

<a id="schema-sid_prefixes"></a>

### sid_prefixes property

Type: `["list", "string"]`. Computed.

SID Locator from the prefix is allocated automatically for each node in each site.

Upstream description:

A SID Locator from the prefix is allocated automatically for each node in each site.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.items.string.min_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.items.string.min_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-annotations) |
| `connect_to_access_networks` | [connect_to_access_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-connect_to_access_networks) |
| `connect_to_enterprise_networks` | [connect_to_enterprise_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-connect_to_enterprise_networks) |
| `connect_to_internet` | [connect_to_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-connect_to_internet) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-namespace) |
| `sid_prefixes` | [sid_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/#schema-sid_prefixes) |

## Next pages

- [xcsh_srv6_network_slice](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/)
