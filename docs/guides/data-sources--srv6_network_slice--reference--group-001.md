---
page_title: "xcsh_srv6_network_slice reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice reference."
---

# xcsh_srv6_network_slice reference

<a id="canonical-ceae08b891b45a658e6ab9b9d513560c18aeedd96150bf1affbd3770714da82d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92926db9861750eea1399181f1a4cab7164314b5bd2f56e6b83ab4eb5e4b9f49"></a>

## Property reference — Property reference / 0c8ba9bdc744 / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-01eb91def3bc993a7ab14854fcd32915db33e69c6c189dc0599a4ae5db267c6b)
- Property reference

<a id="canonical-42237379ede214e07ab6208217088fc1c97d712faadcaf69e7414815c419833e"></a>

## Direct properties — Property reference / 0c8ba9bdc744 / 3

<a id="canonical-a0083309ac4467ad73251ada5c22fbe541f2fbe2512d6468c51bd62c48498896"></a>

<a id="canonical-45bb9a44ede96d0884b6afe114b2494255354ede3be5d9a8aa6287168d8c1f1d"></a>

## annotations property — Property reference / 0c8ba9bdc744 / 4

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

<a id="canonical-b2bcbaaee3a4094bbbc5367c94927e1576f5074d86dd44ddf6434f60ef0d5cc2"></a>

<a id="canonical-9687c61060acfc1b118296a486898ffff934c7d030d5452430bc4cf5a411d792"></a>

## connect_to_access_networks property — Property reference / 0c8ba9bdc744 / 5

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

<a id="canonical-ede6a534caf0b8c7063185664689d4d3afe96fb2396ebc865dd4a39c736c0aa6"></a>

<a id="canonical-ac105c69fdb08948e3c587ed55f7ed651d87aaaa9eafe7799e852f823d76bd16"></a>

## connect_to_enterprise_networks property — Property reference / 0c8ba9bdc744 / 6

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

<a id="canonical-2d72739391ba981116180b5e5ff56acac8cd732601407e9081bdadae6e54eb21"></a>

<a id="canonical-b0c4fd32ae973038d4ca5e50e2bfc63866868d0f53ca9396bdc2512a088c6ea0"></a>

## connect_to_internet property — Property reference / 0c8ba9bdc744 / 7

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

<a id="canonical-7b2e921ebfd16fcbf1dd8910953d7d94ed7fd787466a8b30ff198959e1d20948"></a>

<a id="canonical-f6d9cdb03f10279cd33d6f712737e3efb3c6b53c809223319f3fa2f905369d3e"></a>

## description property — Property reference / 0c8ba9bdc744 / 8

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

<a id="canonical-65a2a307d48668fe3b1f50d12229957a00d1f5343da2fb7daae3c62ecc245289"></a>

<a id="canonical-ffdb03186b9582ecad81b0d02d3fef4343356fd3af1fe7c91e43919ed8535541"></a>

## id property — Property reference / 0c8ba9bdc744 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-b7ccbcde4c0ceeef31214117dae4369be0f1776c0ea367024184fbfb67cee043"></a>

<a id="canonical-7c2c3e174c9bb44af62eb15970820879d6b68f07c33029149084c75d6b951685"></a>

## labels property — Property reference / 0c8ba9bdc744 / 10

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

<a id="canonical-f680bbc89b43a366c1414dfe16d3041284cf586d60d494ac64c5c1d90276d26b"></a>

<a id="canonical-c821c258ebd922a68186828892c3261b4d45433b1503039dbd2d1c5acd389b8d"></a>

## name property — Property reference / 0c8ba9bdc744 / 11

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

<a id="canonical-9a92a2392602648dcbc200e8d8676106d2e29d142a76f72a1d9a5b5c678640da"></a>

<a id="canonical-52c6dae313522c0ee91b01bddf72f6a2eebd740add766224613f086285d2ddc7"></a>

## namespace property — Property reference / 0c8ba9bdc744 / 12

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

<a id="canonical-c167d2f2a53f663d91b9291e93c76ecc7979003d895dc7ae5ec1891bf0964cd7"></a>

<a id="canonical-9e1f00b9a1b88260590e28c5897dcf0654da32e393a9b0adcc939ddd524283ba"></a>

## sid_prefixes property — Property reference / 0c8ba9bdc744 / 13

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

<a id="canonical-a38ef41f733014c4ebdd6445440af9b165df273661dbb8990af5e579da66f560"></a>

## All schema paths — Property reference / 0c8ba9bdc744 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--srv6_network_slice--reference--group-001.md#canonical-a0083309ac4467ad73251ada5c22fbe541f2fbe2512d6468c51bd62c48498896) |
| `connect_to_access_networks` | [connect_to_access_networks](data-sources--srv6_network_slice--reference--group-001.md#canonical-b2bcbaaee3a4094bbbc5367c94927e1576f5074d86dd44ddf6434f60ef0d5cc2) |
| `connect_to_enterprise_networks` | [connect_to_enterprise_networks](data-sources--srv6_network_slice--reference--group-001.md#canonical-ede6a534caf0b8c7063185664689d4d3afe96fb2396ebc865dd4a39c736c0aa6) |
| `connect_to_internet` | [connect_to_internet](data-sources--srv6_network_slice--reference--group-001.md#canonical-2d72739391ba981116180b5e5ff56acac8cd732601407e9081bdadae6e54eb21) |
| `description` | [description](data-sources--srv6_network_slice--reference--group-001.md#canonical-7b2e921ebfd16fcbf1dd8910953d7d94ed7fd787466a8b30ff198959e1d20948) |
| `id` | [id](data-sources--srv6_network_slice--reference--group-001.md#canonical-65a2a307d48668fe3b1f50d12229957a00d1f5343da2fb7daae3c62ecc245289) |
| `labels` | [labels](data-sources--srv6_network_slice--reference--group-001.md#canonical-b7ccbcde4c0ceeef31214117dae4369be0f1776c0ea367024184fbfb67cee043) |
| `name` | [name](data-sources--srv6_network_slice--reference--group-001.md#canonical-f680bbc89b43a366c1414dfe16d3041284cf586d60d494ac64c5c1d90276d26b) |
| `namespace` | [namespace](data-sources--srv6_network_slice--reference--group-001.md#canonical-9a92a2392602648dcbc200e8d8676106d2e29d142a76f72a1d9a5b5c678640da) |
| `sid_prefixes` | [sid_prefixes](data-sources--srv6_network_slice--reference--group-001.md#canonical-c167d2f2a53f663d91b9291e93c76ecc7979003d895dc7ae5ec1891bf0964cd7) |

<a id="canonical-43c4c149aa8eab577854883a38c9bbe169c9942980dd69474b5316fb5d7c6856"></a>

## Next pages — Property reference / 0c8ba9bdc744 / 15

- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-01eb91def3bc993a7ab14854fcd32915db33e69c6c189dc0599a4ae5db267c6b)
