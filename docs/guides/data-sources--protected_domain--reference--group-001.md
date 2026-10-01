---
page_title: "xcsh_protected_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain reference."
---

# xcsh_protected_domain reference

<a id="canonical-109a5c5b1d87fd443e19544e636c3a201e0ebb9189d96c51395f44228e3d7f1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d574a36fbf25fb5bd9682d1a4cc5f47d3229c351aa0f48cfcbd8fda69d96e829"></a>

## Property reference — Property reference / c4b9a72f767b / 2

Breadcrumbs:

- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-3e543fadedff8ca807776d33c9261ab932e9a9940a7f7de944cf323d9229140f)
- Property reference

<a id="canonical-7ad722c03df32c97e0b1af8e7f541c88357ec52d35568a10facd12dc6e965004"></a>

## Direct properties — Property reference / c4b9a72f767b / 3

<a id="canonical-aec736ec8a5819d38dce4c0304a52007d910cb63a7e5b161548e1bcd7485944d"></a>

<a id="canonical-7f77232e3814c86b09f071075679bb76a9ecb0eed9a28e47c2cfb4e5f6360065"></a>

## annotations property — Property reference / c4b9a72f767b / 4

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

<a id="canonical-054a172c4225f026e0f4c00d1bbf8c7f4925b02971d7d0dccb84712d64450d1e"></a>

<a id="canonical-4f1a74a460d098701b30d53a33931e54a656cab9d88eb87020d58743ff900bac"></a>

## description property — Property reference / c4b9a72f767b / 5

Type: `"string"`. Computed.

Description of the ProtectedDomain.

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

<a id="canonical-71ab412a94d4172694543213a8b70f6dd92b6ed3f760e4f46e27fc393ae7365a"></a>

<a id="canonical-09d8c1ddd6d83db0998a54b22bf3a75c008a614eeeb73095bfbf3baed14ecb62"></a>

## id property — Property reference / c4b9a72f767b / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-27b1512e8a8689d2e2644082c48d74250604a7a502e332e20d9b181e9773a5fa"></a>

<a id="canonical-915f5e3f37bbcdb597f03b2d45ee72b8d4db60753acd66024a3f2ee0f52a000c"></a>

## labels property — Property reference / c4b9a72f767b / 7

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

<a id="canonical-76b56f37932150fff0d3532a96d7411d362879a7c556cb2dea39195a99168a97"></a>

<a id="canonical-3411638a1859c1dd376644defce494b3c32904b33a50ad19955e8c6aac8107c9"></a>

## name property — Property reference / c4b9a72f767b / 8

Type: `"string"`. Required.

Name of the ProtectedDomain.

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

<a id="canonical-57dfc410c2b2585518921cee421cd0b49bdcfdb4f26400edf56495c8fb4e67c0"></a>

<a id="canonical-a5130ba3789f0b62fc7cbc55c066f976c2d27b4678a984e00a225292f3487dab"></a>

## namespace property — Property reference / c4b9a72f767b / 9

Type: `"string"`. Required.

Namespace where the ProtectedDomain exists.

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

<a id="canonical-48fac7d89cdc41aa4cffe83240f74f21ec7eafbb28226caca670d91bd7e17256"></a>

<a id="canonical-57ef12f32b7093b7ab2c0cf5d0abe0e6805e03f0569e3fe1810826badf6ba9e2"></a>

## protected_domain property — Property reference / c4b9a72f767b / 10

Type: `"string"`. Computed.

For Client-Side Defense to work on the web pages where you injected the JS, you need to enter the
root domain below.

Upstream description:

For Client-Side Defense to work on the web pages where you injected the JS, you need to enter the
root domain below. Example: if you are adding Client-Side Defense JS on checkout.example.com, you
should enter example.com here.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.etld_plus_one": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.etld_plus_one": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-7314ac08f5d148bf8fac1dbe5725d2f177214290e003825236cde6da66b86629"></a>

## All schema paths — Property reference / c4b9a72f767b / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--protected_domain--reference--group-001.md#canonical-aec736ec8a5819d38dce4c0304a52007d910cb63a7e5b161548e1bcd7485944d) |
| `description` | [description](data-sources--protected_domain--reference--group-001.md#canonical-054a172c4225f026e0f4c00d1bbf8c7f4925b02971d7d0dccb84712d64450d1e) |
| `id` | [id](data-sources--protected_domain--reference--group-001.md#canonical-71ab412a94d4172694543213a8b70f6dd92b6ed3f760e4f46e27fc393ae7365a) |
| `labels` | [labels](data-sources--protected_domain--reference--group-001.md#canonical-27b1512e8a8689d2e2644082c48d74250604a7a502e332e20d9b181e9773a5fa) |
| `name` | [name](data-sources--protected_domain--reference--group-001.md#canonical-76b56f37932150fff0d3532a96d7411d362879a7c556cb2dea39195a99168a97) |
| `namespace` | [namespace](data-sources--protected_domain--reference--group-001.md#canonical-57dfc410c2b2585518921cee421cd0b49bdcfdb4f26400edf56495c8fb4e67c0) |
| `protected_domain` | [protected_domain](data-sources--protected_domain--reference--group-001.md#canonical-48fac7d89cdc41aa4cffe83240f74f21ec7eafbb28226caca670d91bd7e17256) |

<a id="canonical-4b811defa9f27874001a9aa692e3fc4d425450b4c892e293b7cfb5006db9c601"></a>

## Next pages — Property reference / c4b9a72f767b / 12

- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-3e543fadedff8ca807776d33c9261ab932e9a9940a7f7de944cf323d9229140f)
