---
page_title: "xcsh_irule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule reference."
---

# xcsh_irule reference

<a id="canonical-86a3e148141db31b69bea2546604e814213c814fc5ca09e8be422e905b31ed84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6567959e69f35e078329a12aad6ef46dc35b6a643fde4fcd46bf0a1bec2070d"></a>

## Property reference — Property reference / 505bcd35ec6d / 2

Breadcrumbs:

- [xcsh_irule](../data-sources/irule.md#canonical-2d46463305002f54ff3a1ae4bb70fcb0aa618cbcb1816cb90bb6c2a6ed549962)
- Property reference

<a id="canonical-16ee54a2e26818475c3c1632edb5fa209697c1497ad2a63bee628d22e5156e1d"></a>

## Direct properties — Property reference / 505bcd35ec6d / 3

<a id="canonical-b0f829f81067aff59a7856629718e170982ea62839be6bab84af40f5972ea9a8"></a>

<a id="canonical-c689335c3e08d02ad9b69051aa69118fa8d01d864292a62fa879eafe26fe7f3a"></a>

## annotations property — Property reference / 505bcd35ec6d / 4

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

<a id="canonical-117526685d2b8dde78711cc5e3b55c7fa992891cb0c300748ec586076bbbce43"></a>

<a id="canonical-fa8b5becc193e1923814c18e90c34252e8c5fe9070107d2a3d28786c62d52534"></a>

## description property — Property reference / 505bcd35ec6d / 5

Type: `"string"`. Computed.

Description of the Irule.

Upstream description:

Specify Description for iRule.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-bf6952743bde5aebbc9cf14bd9a62c37152930c8e849a33ec7c65454c8261233"></a>

<a id="canonical-5574d118d66d233093176c403334c241f15274d9401c41723e5ca91ec407c6cc"></a>

## description_spec property — Property reference / 505bcd35ec6d / 6

Type: `"string"`. Computed.

Description for iRule. Specify Description for iRule.

<a id="canonical-0dc40e7ef5e27be6510d26577148206356f19facc8f232ca61353621ba7ca610"></a>

<a id="canonical-3afbd82a50b3a70a3d32e5826e54d10c4b28ba839d9944ac89aac3c686d3cc38"></a>

## id property — Property reference / 505bcd35ec6d / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-f3dcebe3569d4a90f009dd6a0faccf50bfc0b44ffa25ca7a171a07f144d4a81b"></a>

<a id="canonical-7e76822263948499709fe29e0539cac9fb8b80f3b6499fa4d50e280e7f364e60"></a>

## irule property — Property reference / 505bcd35ec6d / 8

Type: `"string"`. Computed.

www&#46;internal.example.f5.com')\} DNS::drop\} irule content.

Upstream description:

www&#46;internal.example.f5.com")\} DNS::drop\} irule content.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 24576,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 24576,
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
    "ves.io.schema.rules.string.max_len": "24576"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "24576"
  }
}
```

<a id="canonical-09c8f5bfac8f0c70692510e709f552af0ec87916e2f0ef016e0cf94ec5cbc4bc"></a>

<a id="canonical-9e92a41775a1242d0d06d6a9abb430644e48194783407db97266937572f98507"></a>

## labels property — Property reference / 505bcd35ec6d / 9

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

<a id="canonical-17a5b35d253a2f47f218ba2927710ad1a651ef5537710fd1d6ba92fccdc3c14c"></a>

<a id="canonical-c7ce12b4815f49dbb193db6795bc2345cdcd55f1290426f505f4679c5549fc17"></a>

## name property — Property reference / 505bcd35ec6d / 10

Type: `"string"`. Required.

Name of the Irule.

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

<a id="canonical-f45e0c2af38db2cfb5fc668abdd9db13930bbe5cf0dead7bc486405f2f8949db"></a>

<a id="canonical-e1149fd0b287d1220ffd10c2145a28b4faaa64d0565d2717030d04a55afa0d86"></a>

## namespace property — Property reference / 505bcd35ec6d / 11

Type: `"string"`. Required.

Namespace where the Irule exists.

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

<a id="canonical-884e51134f88f8fc05939c9e3d4bc02b33e5e23967ae693f46599efff6033de1"></a>

## All schema paths — Property reference / 505bcd35ec6d / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--irule--reference--group-001.md#canonical-b0f829f81067aff59a7856629718e170982ea62839be6bab84af40f5972ea9a8) |
| `description` | [description](data-sources--irule--reference--group-001.md#canonical-117526685d2b8dde78711cc5e3b55c7fa992891cb0c300748ec586076bbbce43) |
| `description_spec` | [description_spec](data-sources--irule--reference--group-001.md#canonical-bf6952743bde5aebbc9cf14bd9a62c37152930c8e849a33ec7c65454c8261233) |
| `id` | [id](data-sources--irule--reference--group-001.md#canonical-0dc40e7ef5e27be6510d26577148206356f19facc8f232ca61353621ba7ca610) |
| `irule` | [irule](data-sources--irule--reference--group-001.md#canonical-f3dcebe3569d4a90f009dd6a0faccf50bfc0b44ffa25ca7a171a07f144d4a81b) |
| `labels` | [labels](data-sources--irule--reference--group-001.md#canonical-09c8f5bfac8f0c70692510e709f552af0ec87916e2f0ef016e0cf94ec5cbc4bc) |
| `name` | [name](data-sources--irule--reference--group-001.md#canonical-17a5b35d253a2f47f218ba2927710ad1a651ef5537710fd1d6ba92fccdc3c14c) |
| `namespace` | [namespace](data-sources--irule--reference--group-001.md#canonical-f45e0c2af38db2cfb5fc668abdd9db13930bbe5cf0dead7bc486405f2f8949db) |

<a id="canonical-593be682184f3608dc2d9c9d74a6a10a6dadab00bda372edbb1eabb05aad3476"></a>

## Next pages — Property reference / 505bcd35ec6d / 13

- [xcsh_irule](../data-sources/irule.md#canonical-2d46463305002f54ff3a1ae4bb70fcb0aa618cbcb1816cb90bb6c2a6ed549962)
