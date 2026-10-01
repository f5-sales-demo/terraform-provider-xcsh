---
page_title: "xcsh_allowed_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain reference."
---

# xcsh_allowed_domain reference

<a id="canonical-814fabc4d196c20da75111723783348fc11b0934b815355a67be6505acee85a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fb84fd4f4be08a5a61c91ace974c068b95d77bbbdc7e59d40d94a1ce27ba176"></a>

## Property reference — Property reference / 158e3b0c7edc / 2

Breadcrumbs:

- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-cd1b8e8b9763ec41046ec5b23378de3e4eb43cba5cd5576863af43284e5075e8)
- Property reference

<a id="canonical-e3f752d452fd5dc5eeae79fc694dfef273dd890c16b53ec5e277017335dd3874"></a>

## Direct properties — Property reference / 158e3b0c7edc / 3

<a id="canonical-1a781d52292801abe8c60df8bf1b50d3fb0bc4737860c58cc1be91d41c9347f9"></a>

<a id="canonical-a28c7ba62b071ad856f64c35ef5607096b5c77f3bec75e10932c4da7d1c98b13"></a>

## allowed_domain property — Property reference / 158e3b0c7edc / 4

Type: `"string"`. Computed.

Enter root domain or domain to be entered to allow list below. Domains can be entered only one at a
time. In case of conflicting entries, the domain entry takes precedence over the root domain entry.

Upstream description:

Enter root domain or domain to be entered to allow list below. Domains can be entered only one at a
time. In case of conflicting entries, the domain entry takes precedence over the root domain entry.
Example: if you are adding Client-Side Defense JS on checkout.example.com, you should enter
example.com here.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-9c7c1883e3e58b234b77b8f09b7dde50033fb41d29e89013e2967140c2543c77"></a>

<a id="canonical-391de66970b53c694a9636acea25f69b5a77b320149f7e8875a0281b2c8dfb79"></a>

## annotations property — Property reference / 158e3b0c7edc / 5

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

<a id="canonical-195afaf8d2b6abb392e37545f05d99240514fcc3864d63b2e3a54632159b5fe3"></a>

<a id="canonical-cde7187dda80350e8cef48d227b291b46ceda1afc90d27ece7cc161ccf549737"></a>

## description property — Property reference / 158e3b0c7edc / 6

Type: `"string"`. Computed.

Description of the AllowedDomain.

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

<a id="canonical-9b2b46e2861c08e8eb76fa3bd5fe891213cae9ce393be33f4e8a4acb6e6f8a45"></a>

<a id="canonical-82210d74109673c1e491e1785c0206208eee05a22df90651974bc0f6761ceb85"></a>

## id property — Property reference / 158e3b0c7edc / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-61791b9e989fec4833e6c876691c058df381082efbe7efff110b5678d54492da"></a>

<a id="canonical-cde885ec9592b15d05ddc09a689fa499f94a4f2b10d82168bfaa81017381d635"></a>

## labels property — Property reference / 158e3b0c7edc / 8

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

<a id="canonical-fd9832827536ab295ae23abd7dbebb4261a5ec7793bdc8028b8a193f981d0d57"></a>

<a id="canonical-463afcd406333c19b5674a142192889f63193fd0644b1dc3864fa46ffcb48a12"></a>

## name property — Property reference / 158e3b0c7edc / 9

Type: `"string"`. Required.

Name of the AllowedDomain.

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

<a id="canonical-ac89a8e882f92711387d2d12baf8d355fee82cf0137099995aacad404e06ad6f"></a>

<a id="canonical-63263015ae40a2bcd36029bed4e0a94eff5afc063cc9509649a285cd5fb9fc79"></a>

## namespace property — Property reference / 158e3b0c7edc / 10

Type: `"string"`. Required.

Namespace where the AllowedDomain exists.

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

<a id="canonical-818fef52059ae0d39740ea9f8cf1e54c3910ab695f825375be473ac4905da5f6"></a>

## All schema paths — Property reference / 158e3b0c7edc / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowed_domain` | [allowed_domain](data-sources--allowed_domain--reference--group-001.md#canonical-1a781d52292801abe8c60df8bf1b50d3fb0bc4737860c58cc1be91d41c9347f9) |
| `annotations` | [annotations](data-sources--allowed_domain--reference--group-001.md#canonical-9c7c1883e3e58b234b77b8f09b7dde50033fb41d29e89013e2967140c2543c77) |
| `description` | [description](data-sources--allowed_domain--reference--group-001.md#canonical-195afaf8d2b6abb392e37545f05d99240514fcc3864d63b2e3a54632159b5fe3) |
| `id` | [id](data-sources--allowed_domain--reference--group-001.md#canonical-9b2b46e2861c08e8eb76fa3bd5fe891213cae9ce393be33f4e8a4acb6e6f8a45) |
| `labels` | [labels](data-sources--allowed_domain--reference--group-001.md#canonical-61791b9e989fec4833e6c876691c058df381082efbe7efff110b5678d54492da) |
| `name` | [name](data-sources--allowed_domain--reference--group-001.md#canonical-fd9832827536ab295ae23abd7dbebb4261a5ec7793bdc8028b8a193f981d0d57) |
| `namespace` | [namespace](data-sources--allowed_domain--reference--group-001.md#canonical-ac89a8e882f92711387d2d12baf8d355fee82cf0137099995aacad404e06ad6f) |

<a id="canonical-65d793379e4f8de798769e60099120f940e3f7f7d64bc72d7b4fe5fc07b04b77"></a>

## Next pages — Property reference / 158e3b0c7edc / 12

- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-cd1b8e8b9763ec41046ec5b23378de3e4eb43cba5cd5576863af43284e5075e8)
