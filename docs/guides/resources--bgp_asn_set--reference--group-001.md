---
page_title: "xcsh_bgp_asn_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set reference."
---

# xcsh_bgp_asn_set reference

<a id="canonical-a9058fc90b7a6c3e30352fc706291f4deeabb26d142eebfe18a505b212bc5a60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4ca5b4929e4b0cad83d725e921ed414ed0f0287717ca9284fa7a625e3a4d9aa"></a>

## Property reference — Property reference / f04d3e5a990a / 2

Breadcrumbs:

- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653)
- Property reference

<a id="canonical-22b8a2743fe48f8567572851fbfd6b8b74cd7ce4a5b9c213028c2f4d50ce631d"></a>

## Direct properties — Property reference / f04d3e5a990a / 3

<a id="canonical-369a189b3e9f5cf59a55b3cf68ba481af53a8106d213eaa87ed1b5883d6a257b"></a>

<a id="canonical-4fb842294ef94454e56941bb6717991143438543ad38095f9fa4dab1ffad85a2"></a>

## annotations property — Property reference / f04d3e5a990a / 4

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

<a id="canonical-89f9ec5e16cf40a3737323ca97719975940c40aec6413812381cfd2185e47250"></a>

<a id="canonical-a72dabb9e21ab976a3ec9cc1bfe376a3f367a23d79a350b7385f3a2618287a31"></a>

## as_numbers property — Property reference / f04d3e5a990a / 5

Type: `["list", "number"]`. Required.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or
blacklists for use in network policy or service policy.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or
blacklists for use in network policy or service policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

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

<a id="canonical-7ab96e6fb533fabf3d0d8bcf7aeac89b23b35baaca06a1f9abc18191311753db"></a>

<a id="canonical-e03baf63aabdec1fd1cea33cff6dc08653453888f66e57378d76329f77bcd37f"></a>

## description property — Property reference / f04d3e5a990a / 6

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

<a id="canonical-0cd1e0c9022fe2a10a9c04ddb748204e7707f3b0ba107cfcbf26c03a59f04e6a"></a>

<a id="canonical-b93376e3beee5b14841155d31cdd36fac34f0fb7d04333340a1bee9f4465333c"></a>

## disable property — Property reference / f04d3e5a990a / 7

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

<a id="canonical-43cb6324a544b3e1404aa10e5bf85a3d20601c5c5365ef5853434106098774d5"></a>

<a id="canonical-7cb637f13874571e78fdbd394068dd1f0779fe1c2ec69e31d55fb7b8ab699248"></a>

## id property — Property reference / f04d3e5a990a / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-a170ee170dfd80a3db2b9f8bfbb8e6f3000c1b30176148d647ef802e5dd23a31"></a>

<a id="canonical-70e968de6c7d545c8cfbf3a947a974dac07baa4b133f0d7560bee6af18e3a055"></a>

## labels property — Property reference / f04d3e5a990a / 9

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

<a id="canonical-25eebd71d9d432b119e064ad163fcc17fc9c3bdb745b2470ccb81c5be418c20d"></a>

<a id="canonical-a9c84f1ebd7d2491ed78665193e672d3cc12452b7b9383a77c0ee4f8c5241b7f"></a>

## name property — Property reference / f04d3e5a990a / 10

Type: `"string"`. Required.

Name of the BGP Asn Set. Must be unique within the namespace.

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

<a id="canonical-0d209fd0260ff819fa914ccf96d307529e1c5850d8ddd803b20214d7e604836d"></a>

<a id="canonical-9364cdf034f3e54ad97526429aa56e15f9b653ae82ce5016a2f0e0e3ebaadd78"></a>

## namespace property — Property reference / f04d3e5a990a / 11

Type: `"string"`. Required.

Namespace where the BGP Asn Set is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [timeouts](resources--bgp_asn_set--reference--group-001.md#canonical-424554f44ff05aa9a415dd5694b67924fffed93f2b211c7ed3ba8d835f815c45): complete subsection reference.

<a id="canonical-6782b58747213a0195da04b0ed81473def91451d28e8b4418d683b8fa520629d"></a>

## All schema paths — Property reference / f04d3e5a990a / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bgp_asn_set--reference--group-001.md#canonical-369a189b3e9f5cf59a55b3cf68ba481af53a8106d213eaa87ed1b5883d6a257b) |
| `as_numbers` | [as_numbers](resources--bgp_asn_set--reference--group-001.md#canonical-89f9ec5e16cf40a3737323ca97719975940c40aec6413812381cfd2185e47250) |
| `description` | [description](resources--bgp_asn_set--reference--group-001.md#canonical-7ab96e6fb533fabf3d0d8bcf7aeac89b23b35baaca06a1f9abc18191311753db) |
| `disable` | [disable](resources--bgp_asn_set--reference--group-001.md#canonical-0cd1e0c9022fe2a10a9c04ddb748204e7707f3b0ba107cfcbf26c03a59f04e6a) |
| `id` | [id](resources--bgp_asn_set--reference--group-001.md#canonical-43cb6324a544b3e1404aa10e5bf85a3d20601c5c5365ef5853434106098774d5) |
| `labels` | [labels](resources--bgp_asn_set--reference--group-001.md#canonical-a170ee170dfd80a3db2b9f8bfbb8e6f3000c1b30176148d647ef802e5dd23a31) |
| `name` | [name](resources--bgp_asn_set--reference--group-001.md#canonical-25eebd71d9d432b119e064ad163fcc17fc9c3bdb745b2470ccb81c5be418c20d) |
| `namespace` | [namespace](resources--bgp_asn_set--reference--group-001.md#canonical-0d209fd0260ff819fa914ccf96d307529e1c5850d8ddd803b20214d7e604836d) |
| `timeouts` | [timeouts](resources--bgp_asn_set--reference--group-001.md#canonical-4cea5914207a4cc6cbd3adbb88a6f5de2177664c6707c9452a93173b35f2494c) |
| `timeouts.create` | [timeouts.create](resources--bgp_asn_set--reference--group-001.md#canonical-fbada06337fff69f597fadc786faab0754130bc5984fa3fbdaeeb9560c678064) |
| `timeouts.delete` | [timeouts.delete](resources--bgp_asn_set--reference--group-001.md#canonical-4f4b38d2700b862525b3f52ccc0519b88b0b074115b93b74de0bcabc6a756bba) |
| `timeouts.read` | [timeouts.read](resources--bgp_asn_set--reference--group-001.md#canonical-665b69247e04114add7bb044d35dfe6cfb3749df626917e8a1ff58d89387abed) |
| `timeouts.update` | [timeouts.update](resources--bgp_asn_set--reference--group-001.md#canonical-81632b124c658aa99bf99760cecf92c9cc344b7735b3b16aabbf55f67d5955a9) |

<a id="canonical-ee5444beb470feb0e5dc9e841609dd1b7eab85238fb29220fb70242802cfb814"></a>

## Next pages — Property reference / f04d3e5a990a / 13

- [timeouts](resources--bgp_asn_set--reference--group-001.md#canonical-424554f44ff05aa9a415dd5694b67924fffed93f2b211c7ed3ba8d835f815c45)
- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653)

<a id="canonical-424554f44ff05aa9a415dd5694b67924fffed93f2b211c7ed3ba8d835f815c45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aa0d4cc59543c1f0078e5ee5b49537909dafc24496cd3336be7a639a21fc4db"></a>

## timeouts — timeouts / 65ab8aec9ba2 / 2

Breadcrumbs:

- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653)
- [Property reference](resources--bgp_asn_set--reference--group-001.md#canonical-a9058fc90b7a6c3e30352fc706291f4deeabb26d142eebfe18a505b212bc5a60)
- timeouts

<a id="canonical-4cea5914207a4cc6cbd3adbb88a6f5de2177664c6707c9452a93173b35f2494c"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea50ad7c3ebcde3e854e11ed4a7db12cb5531377e42de63a5c48da5dce3d3b21"></a>

## Direct properties — timeouts / 65ab8aec9ba2 / 3

<a id="canonical-fbada06337fff69f597fadc786faab0754130bc5984fa3fbdaeeb9560c678064"></a>

<a id="canonical-07e21053ed2c4b4f2b0a6d8153e884e4b9aefc672156fbed735e5d07ce38e378"></a>

## create property — timeouts / 65ab8aec9ba2 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-4f4b38d2700b862525b3f52ccc0519b88b0b074115b93b74de0bcabc6a756bba"></a>

<a id="canonical-72093eb6d29f1c89140568e1744e4545b7facc8e95f5c7717d1141a91f490ca4"></a>

## delete property — timeouts / 65ab8aec9ba2 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-665b69247e04114add7bb044d35dfe6cfb3749df626917e8a1ff58d89387abed"></a>

<a id="canonical-62bf099e8ae2e1845e6fb2da3aff5deb50927f30f241fffef822f272f37621ef"></a>

## read property — timeouts / 65ab8aec9ba2 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-81632b124c658aa99bf99760cecf92c9cc344b7735b3b16aabbf55f67d5955a9"></a>

<a id="canonical-234b495fc0fbd98534b6b0ba6b8352b9d66d7f624d2820a73875b9e754bd2152"></a>

## update property — timeouts / 65ab8aec9ba2 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1a5b1bb707321e365d91e9f7eaa012580762f25d0aed45ab58995fa236ab9a5d"></a>

## Next pages — timeouts / 65ab8aec9ba2 / 8

- [Property reference](resources--bgp_asn_set--reference--group-001.md#canonical-a9058fc90b7a6c3e30352fc706291f4deeabb26d142eebfe18a505b212bc5a60)
- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653)
