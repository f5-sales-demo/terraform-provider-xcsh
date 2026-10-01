---
page_title: "xcsh_authorization_server reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server reference."
---

# xcsh_authorization_server reference

<a id="canonical-4fe633a5d8eadcbda1f3928dac2979247f71fce9af1db4ad727fc19ccb660f8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99b0cbc77b39c13a9596f1059a39863354afa87be832d5859c0b32a2b62ed569"></a>

## Property reference — Property reference / 8abe85ca8460 / 2

Breadcrumbs:

- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-c9a064e11c392b1729b64c4abb5a0c14e6107920f082e94b97489dd17471f27a)
- Property reference

<a id="canonical-191abfd05b4a53a73671e2de59800399eb27b6804b5ecc0ce2a7635f3e624eee"></a>

## Direct properties — Property reference / 8abe85ca8460 / 3

<a id="canonical-f4288896e7a0fbaf6ffe9bbb3bbf5ad03600d8a48e66ff096c36f90b95e80766"></a>

<a id="canonical-e1278fb2ddc9afefa897e8f70656cf0a3228ad982d784c9307631dd021c48921"></a>

## annotations property — Property reference / 8abe85ca8460 / 4

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

<a id="canonical-879d0074366c559c80ab500ea56056469da67001d00b24a762228c14e484a1ea"></a>

<a id="canonical-583affbdcf0cf681661e63ed60618cb02db236d410282ed9602f1d9e27dfb7f4"></a>

## description property — Property reference / 8abe85ca8460 / 5

Type: `"string"`. Computed.

Description of the AuthorizationServer.

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

<a id="canonical-947b283d86ca18e4ada6e67477929d426e324be400800b635c8eb4efd65d20c0"></a>

<a id="canonical-e247b1fecbcbb37821fca85056435dfa8fd7f99ee5911c87a641af10b66fc141"></a>

## id property — Property reference / 8abe85ca8460 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-c4409087b95aba94e9bf17d6fa24b6517078ec18eb84d59aeceb9432c5bb3680"></a>

<a id="canonical-f6db6729ffc0d5c611a5d2766bef846c2bfe111758b96f9fab3c382f464bc461"></a>

## jwks_uri property — Property reference / 8abe85ca8460 / 7

Type: `"string"`. Computed.

X-textBlockContent: Automatic fetching of JWKS will happen once daily. You can also do it manually
from the list of Authorization Servers at any time.

Upstream description:

X-textBlockContent: Automatic fetching of JWKS will happen once daily. You can also do it manually
from the list of Authorization Servers at any time.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.uri": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.uri": "true"
  }
}
```

<a id="canonical-fbbf36718b247c206eff5e57e09a0c757fbbefb9926b3c819304d3f540ab29a4"></a>

<a id="canonical-e119deb9f41ae0de4067bc00c32b2465b8b4d6d0e8091afa38f94f152ac27adb"></a>

## labels property — Property reference / 8abe85ca8460 / 8

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

<a id="canonical-7d80b6d58b868c83c1263ce572f84361d7f55bf6362ef55933c250300a61115e"></a>

<a id="canonical-9099bf736272850c606ccada43a47058c4939c3abbc46d42c534b988d6caa1f1"></a>

## name property — Property reference / 8abe85ca8460 / 9

Type: `"string"`. Required.

Name of the AuthorizationServer.

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

<a id="canonical-6c912e3071145e5d2377191432122cbcf445f1f196ad35e0468839628f70d24e"></a>

<a id="canonical-98ca363c6bc470f02be7e1910eb8c5e308d9ba4346d83cdde5b0e0357dd8833d"></a>

## namespace property — Property reference / 8abe85ca8460 / 10

Type: `"string"`. Required.

Namespace where the AuthorizationServer exists.

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

<a id="canonical-f6bcfae9e8390c27544a7a681b5ad8637c91c342e21cc6440616e3d7403bb54d"></a>

## All schema paths — Property reference / 8abe85ca8460 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--authorization_server--reference--group-001.md#canonical-f4288896e7a0fbaf6ffe9bbb3bbf5ad03600d8a48e66ff096c36f90b95e80766) |
| `description` | [description](data-sources--authorization_server--reference--group-001.md#canonical-879d0074366c559c80ab500ea56056469da67001d00b24a762228c14e484a1ea) |
| `id` | [id](data-sources--authorization_server--reference--group-001.md#canonical-947b283d86ca18e4ada6e67477929d426e324be400800b635c8eb4efd65d20c0) |
| `jwks_uri` | [jwks_uri](data-sources--authorization_server--reference--group-001.md#canonical-c4409087b95aba94e9bf17d6fa24b6517078ec18eb84d59aeceb9432c5bb3680) |
| `labels` | [labels](data-sources--authorization_server--reference--group-001.md#canonical-fbbf36718b247c206eff5e57e09a0c757fbbefb9926b3c819304d3f540ab29a4) |
| `name` | [name](data-sources--authorization_server--reference--group-001.md#canonical-7d80b6d58b868c83c1263ce572f84361d7f55bf6362ef55933c250300a61115e) |
| `namespace` | [namespace](data-sources--authorization_server--reference--group-001.md#canonical-6c912e3071145e5d2377191432122cbcf445f1f196ad35e0468839628f70d24e) |

<a id="canonical-53087cf63852ef20196e65154ff2ac17a89f04216eb7ede1d55d38ce9e90efd7"></a>

## Next pages — Property reference / 8abe85ca8460 / 12

- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-c9a064e11c392b1729b64c4abb5a0c14e6107920f082e94b97489dd17471f27a)
