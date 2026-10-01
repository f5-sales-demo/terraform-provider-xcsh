---
page_title: "xcsh_namespace reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace reference."
---

# xcsh_namespace reference

<a id="canonical-6f4493fe6a03e423b14ed16b610b1d61c98ab9ef17f2f98fa04f2fb20e9e9fd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ba11abdd27c1a3f3f717ad25bccc4c844520796da9ffb71dee508080cbd1d86"></a>

## Property reference — Property reference / 59f6c79ef2a5 / 2

Breadcrumbs:

- [xcsh_namespace](../data-sources/namespace.md#canonical-87303facfb5a86c45f03fd2edff3e82a1b37ef90f0d3be2e23ec03de74d32096)
- Property reference

<a id="canonical-90168069c34ab4cfde41a8cdc7e36082c44aea8fd57d0285c614a04c2189d7ab"></a>

## Direct properties — Property reference / 59f6c79ef2a5 / 3

<a id="canonical-a9478a823ddc583a80aeddb9e5deb393b577295e698bfc7350ca20f548d60cce"></a>

<a id="canonical-102d86638bafebc62fe16a6e041fc17c352ccbccd62c9aedfcaea96838e706d0"></a>

## annotations property — Property reference / 59f6c79ef2a5 / 4

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

<a id="canonical-4baa0cb519ec26fc4829c89d142eadca890d8a648fbf85a7a2ddbfc5e85b570b"></a>

<a id="canonical-4fe0ffaa55a832f35f3422689b256dd7dbadebaa7b3f307c01424a777373c02e"></a>

## description property — Property reference / 59f6c79ef2a5 / 5

Type: `"string"`. Computed.

Description of the Namespace.

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

<a id="canonical-0684d8eb2cc1e8d960568e160887f16065b3480a7a6829045671cf409a07bc3f"></a>

<a id="canonical-5b5369b89d2cc58f825eea143f221e72248d30e661c7d31a280780fd65ccb1db"></a>

## id property — Property reference / 59f6c79ef2a5 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-ee8d1dab28980832d9b87b02b0e9b5ca99971f1ff8db322aafa7bc3637df6a4c"></a>

<a id="canonical-5ae4d3681105850fdff63651af479db665248be1b0a9e9c1962312dab632201c"></a>

## labels property — Property reference / 59f6c79ef2a5 / 7

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

<a id="canonical-83c113185d3c80c084ce03a0598f661c5d0b4dbd26406e525697caddb22af838"></a>

<a id="canonical-b83d35455c64582a4e48648d0034d75daa5204fa7b59faa32b78911f3bb1e7bd"></a>

## name property — Property reference / 59f6c79ef2a5 / 8

Type: `"string"`. Required.

Name of the Namespace.

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

<a id="canonical-aeead717407198acc983cefb8cdc75c1c8087a86b6f3b3a737838687436371b9"></a>

<a id="canonical-aae6bed9f93eec551eec891b9ad060859f39f9ac9c8953ec7972684ba6c2a05b"></a>

## namespace property — Property reference / 59f6c79ef2a5 / 9

Type: `"string"`. Optional.

Namespaces are tenant-level objects. Omit this argument.

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

<a id="canonical-2d3b2ef440f6e20aa9b48cfdf086d134e62dfd20460dd6b55552c31cf67ee3f4"></a>

## All schema paths — Property reference / 59f6c79ef2a5 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--namespace--reference--group-001.md#canonical-a9478a823ddc583a80aeddb9e5deb393b577295e698bfc7350ca20f548d60cce) |
| `description` | [description](data-sources--namespace--reference--group-001.md#canonical-4baa0cb519ec26fc4829c89d142eadca890d8a648fbf85a7a2ddbfc5e85b570b) |
| `id` | [id](data-sources--namespace--reference--group-001.md#canonical-0684d8eb2cc1e8d960568e160887f16065b3480a7a6829045671cf409a07bc3f) |
| `labels` | [labels](data-sources--namespace--reference--group-001.md#canonical-ee8d1dab28980832d9b87b02b0e9b5ca99971f1ff8db322aafa7bc3637df6a4c) |
| `name` | [name](data-sources--namespace--reference--group-001.md#canonical-83c113185d3c80c084ce03a0598f661c5d0b4dbd26406e525697caddb22af838) |
| `namespace` | [namespace](data-sources--namespace--reference--group-001.md#canonical-aeead717407198acc983cefb8cdc75c1c8087a86b6f3b3a737838687436371b9) |

<a id="canonical-0c271f64c632f9da171e03e265e003a940997e28e987621b00eb8b8dc86acd76"></a>

## Next pages — Property reference / 59f6c79ef2a5 / 11

- [xcsh_namespace](../data-sources/namespace.md#canonical-87303facfb5a86c45f03fd2edff3e82a1b37ef90f0d3be2e23ec03de74d32096)
