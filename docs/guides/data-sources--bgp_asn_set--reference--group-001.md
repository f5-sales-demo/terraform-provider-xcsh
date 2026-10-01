---
page_title: "xcsh_bgp_asn_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set reference."
---

# xcsh_bgp_asn_set reference

<a id="canonical-f77d4e7e206bcb563c10db6a1d66e0f7227add89a6be618d75ba79917484a4eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96912a375544b39e205da98bb4b55486527e43df2cf3184721dc282ea401e885"></a>

## Property reference — Property reference / 7835e5bb41cb / 2

Breadcrumbs:

- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-117a5d23afbac64f2130fcb3963e2ecf60d3582aab779009e560eb05521c7e3f)
- Property reference

<a id="canonical-6cb04ec913d73fea8e51b82cab0a0ef3eff3ebddb3e7cd5a0ae67da5e2a446f2"></a>

## Direct properties — Property reference / 7835e5bb41cb / 3

<a id="canonical-ca79ee496c3f46bfa25bea5bd3bf56d9f0f0a485d0eb6e0eb744db83eeb1e732"></a>

<a id="canonical-412d4390b8a430b0001a05671d8f8154c4cd5ba8fcb61e314b0a7db1ffa54056"></a>

## annotations property — Property reference / 7835e5bb41cb / 4

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

<a id="canonical-50ee79ddd224deca0227868abbadd8bb78f7d99f5c6fd8b4b89ba718208a1bf2"></a>

<a id="canonical-7d32d8ce37a391b2b6857f37f598a20793a26e1f6b227fbaf738351db2da8bd6"></a>

## as_numbers property — Property reference / 7835e5bb41cb / 5

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or
blacklists for use in network policy or service policy.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or
blacklists for use in network policy or service policy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-88845489f66207d6981be7406d4de7871ac0631157ce2dce26332291b50867e0"></a>

<a id="canonical-770a6e56e6bd5baf95022f51b9ad6ba931926fad963b67d6c142f3049e9c1b9c"></a>

## description property — Property reference / 7835e5bb41cb / 6

Type: `"string"`. Computed.

Description of the BGPAsnSet.

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

<a id="canonical-2b2c49c2c6e87a9b8aa84a19b78168ce8081325248d8c5cc38fafa563108a6d7"></a>

<a id="canonical-9fbb7aecdc2ee2c49c26ca54f6d381fce67cf7e8ec4feebc204ddf2710e84bfd"></a>

## id property — Property reference / 7835e5bb41cb / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4bdefa98de03c8cb5c1bdc1f027be1c45ed6a7d891bdb0cfb05b6aa822fd565f"></a>

<a id="canonical-da2567a48ec7248df219a3602fae612e7429d3eb961a19877c9b41ac8e86d9b3"></a>

## labels property — Property reference / 7835e5bb41cb / 8

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

<a id="canonical-570b2437fb47634f1d285dd9409b1de99f0f07107d5488b9484df3fa233a7c1e"></a>

<a id="canonical-910c40d4c093d28b6e296e91e5a32bb1f6c199edf82fc776626ff090cf6c5296"></a>

## name property — Property reference / 7835e5bb41cb / 9

Type: `"string"`. Required.

Name of the BGPAsnSet.

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

<a id="canonical-7fde289fc451eabf30e44e7130e86594227efaa0f18d37d3c55d0082b095e156"></a>

<a id="canonical-3e1898fd22771f48cfb82aa5e644fbea4dbed2580cf64a00ac83ffdf7dd0023d"></a>

## namespace property — Property reference / 7835e5bb41cb / 10

Type: `"string"`. Required.

Namespace where the BGPAsnSet exists.

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

<a id="canonical-0699bd6ded980ef65a148c7935b9cea0be3f78fa4ef30fd312961e96e419d6a4"></a>

## All schema paths — Property reference / 7835e5bb41cb / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp_asn_set--reference--group-001.md#canonical-ca79ee496c3f46bfa25bea5bd3bf56d9f0f0a485d0eb6e0eb744db83eeb1e732) |
| `as_numbers` | [as_numbers](data-sources--bgp_asn_set--reference--group-001.md#canonical-50ee79ddd224deca0227868abbadd8bb78f7d99f5c6fd8b4b89ba718208a1bf2) |
| `description` | [description](data-sources--bgp_asn_set--reference--group-001.md#canonical-88845489f66207d6981be7406d4de7871ac0631157ce2dce26332291b50867e0) |
| `id` | [id](data-sources--bgp_asn_set--reference--group-001.md#canonical-2b2c49c2c6e87a9b8aa84a19b78168ce8081325248d8c5cc38fafa563108a6d7) |
| `labels` | [labels](data-sources--bgp_asn_set--reference--group-001.md#canonical-4bdefa98de03c8cb5c1bdc1f027be1c45ed6a7d891bdb0cfb05b6aa822fd565f) |
| `name` | [name](data-sources--bgp_asn_set--reference--group-001.md#canonical-570b2437fb47634f1d285dd9409b1de99f0f07107d5488b9484df3fa233a7c1e) |
| `namespace` | [namespace](data-sources--bgp_asn_set--reference--group-001.md#canonical-7fde289fc451eabf30e44e7130e86594227efaa0f18d37d3c55d0082b095e156) |

<a id="canonical-1d2573d2a313f0474ebe1a63ec1d7f5eec6616b610b602aa7a7e68e939d2f927"></a>

## Next pages — Property reference / 7835e5bb41cb / 12

- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-117a5d23afbac64f2130fcb3963e2ecf60d3582aab779009e560eb05521c7e3f)
