---
page_title: "xcsh_namespace reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace reference."
---

# xcsh_namespace reference

<a id="canonical-09c5d1f37f639be8d72ab53f1082e4c1d53428b0ed5d288fca41f55d74dce8ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1e2c175f4f7bc9558ce585ab459e8ff3865ac57c6506fc205c54ac6bd61bab1"></a>

## Property reference — Property reference / 32edde2e15ca / 2

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md#canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4)
- Property reference

<a id="canonical-5576d7ad10b72ec38654c6807ba70ca109462b7868601af3d6271294da13fd22"></a>

## Direct properties — Property reference / 32edde2e15ca / 3

<a id="canonical-b8c9011c0983bc021bfb10c08479116a11ef5f04e49d00b76f876ad40743f5b9"></a>

<a id="canonical-afc20e1cd36b979bc6ed07e7fdb6a59eafe20710705d5106244cbc6339846fad"></a>

## annotations property — Property reference / 32edde2e15ca / 4

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

<a id="canonical-a4e65074ca0a8a008e3465a6f6617b11e62b25ab4f22ca6d68ea79cb2fc04806"></a>

<a id="canonical-643a20b30a494c2e6db34b1e4da0292f3c760cc8fee99d3edee78b6bc145fd2e"></a>

## description property — Property reference / 32edde2e15ca / 5

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

<a id="canonical-99c3db1d49b26b985d771952063e72290907e7e3e16be906bb3f25c2548efb03"></a>

<a id="canonical-6e971014665c0cf8449dd381866792726899b7116fb3eeeeb7e7eac8599510d9"></a>

## disable property — Property reference / 32edde2e15ca / 6

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

<a id="canonical-2b680dc0ebc1fbc707db218444c05d4163d7c46d3de4d74016feddd2c9d5ff53"></a>

<a id="canonical-1197a3eff23a279d976d50be99f700b5985f4bce1f3c03d60bde05a4e10168f4"></a>

## id property — Property reference / 32edde2e15ca / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-ba193f7bdd4b5d6eda8d915b97d5d09b9a5152b55433ebb6b13230d87b2475f4"></a>

<a id="canonical-b64f62b3ad92a5c8347e0ddd608a31ca37bfa1ab1746a5523898b88aac1cdde9"></a>

## labels property — Property reference / 32edde2e15ca / 8

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

<a id="canonical-8fa86c5921fe7a49f8349c5c005b5755469db5603991ca5c00b9de34c4214b94"></a>

<a id="canonical-3586c6a91a9b5ef7556611ca88e53ebc19296d5b766b12cb048a926682b7cbdc"></a>

## name property — Property reference / 32edde2e15ca / 9

Type: `"string"`. Required.

Name of the Namespace. Must be unique within the namespace.

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

<a id="canonical-2cec49d52c7513f91cb04c1b7aa14774ded61fa3f194ccf294d7bf6b29199003"></a>

<a id="canonical-1874cb15574a90eeb870207ff53c455c6bdd48f0713f0d3c5284a7aa2e3584d9"></a>

## namespace property — Property reference / 32edde2e15ca / 10

Type: `"string"`. Optional, Computed.

Namespace for the Namespace. For this resource type, namespace should be empty or omitted.

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

- [timeouts](resources--namespace--reference--group-001.md#canonical-23bd8a9dae1a4578dc94cd474bccdd7bc1a3eb7cad30aecc87ac00c02a7c8360): complete subsection reference.

<a id="canonical-55eda9d8e1455036f6e47c2b530991360a59a8aeb8ed6abba461e9766d1a78c0"></a>

## All schema paths — Property reference / 32edde2e15ca / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--namespace--reference--group-001.md#canonical-b8c9011c0983bc021bfb10c08479116a11ef5f04e49d00b76f876ad40743f5b9) |
| `description` | [description](resources--namespace--reference--group-001.md#canonical-a4e65074ca0a8a008e3465a6f6617b11e62b25ab4f22ca6d68ea79cb2fc04806) |
| `disable` | [disable](resources--namespace--reference--group-001.md#canonical-99c3db1d49b26b985d771952063e72290907e7e3e16be906bb3f25c2548efb03) |
| `id` | [id](resources--namespace--reference--group-001.md#canonical-2b680dc0ebc1fbc707db218444c05d4163d7c46d3de4d74016feddd2c9d5ff53) |
| `labels` | [labels](resources--namespace--reference--group-001.md#canonical-ba193f7bdd4b5d6eda8d915b97d5d09b9a5152b55433ebb6b13230d87b2475f4) |
| `name` | [name](resources--namespace--reference--group-001.md#canonical-8fa86c5921fe7a49f8349c5c005b5755469db5603991ca5c00b9de34c4214b94) |
| `namespace` | [namespace](resources--namespace--reference--group-001.md#canonical-2cec49d52c7513f91cb04c1b7aa14774ded61fa3f194ccf294d7bf6b29199003) |
| `timeouts` | [timeouts](resources--namespace--reference--group-001.md#canonical-7bc8234b432bcfa6ac068161373c669e351b79cab1808244ec2dd666bccd4264) |
| `timeouts.create` | [timeouts.create](resources--namespace--reference--group-001.md#canonical-a4c5ec411bbbf7b26f3c453812e1c46edacd31ca2389e85618ef4ca9234a7371) |
| `timeouts.delete` | [timeouts.delete](resources--namespace--reference--group-001.md#canonical-a91de8ae2858da1e820a38a797d2215aeb25645a91af3db750431f6fec085743) |
| `timeouts.read` | [timeouts.read](resources--namespace--reference--group-001.md#canonical-321bea21a5809f5bca30411310bb5b06f073660124f76006f65d10cef49b14ef) |
| `timeouts.update` | [timeouts.update](resources--namespace--reference--group-001.md#canonical-fb43e42c0ece31f06c8ea65334cbb202d2a16cef058805bb99a306754b5b8fc7) |

<a id="canonical-766aaa3a489092004252480e9fb4b1de16ad118a0a10374003de8def6dbb126c"></a>

## Next pages — Property reference / 32edde2e15ca / 12

- [timeouts](resources--namespace--reference--group-001.md#canonical-23bd8a9dae1a4578dc94cd474bccdd7bc1a3eb7cad30aecc87ac00c02a7c8360)
- [xcsh_namespace](../resources/namespace.md#canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4)

<a id="canonical-23bd8a9dae1a4578dc94cd474bccdd7bc1a3eb7cad30aecc87ac00c02a7c8360"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01b12635058a2ca3157ff4543d62571770ede36504cc6dc1216a22b2986c9062"></a>

## timeouts — timeouts / 4a02e8f4fb1a / 2

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md#canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4)
- [Property reference](resources--namespace--reference--group-001.md#canonical-09c5d1f37f639be8d72ab53f1082e4c1d53428b0ed5d288fca41f55d74dce8ac)
- timeouts

<a id="canonical-7bc8234b432bcfa6ac068161373c669e351b79cab1808244ec2dd666bccd4264"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-c36a06bcc52f9f38ba59adb7759ed72630e4e3dc5a06237836f45041b5ac4e4a"></a>

## Direct properties — timeouts / 4a02e8f4fb1a / 3

<a id="canonical-a4c5ec411bbbf7b26f3c453812e1c46edacd31ca2389e85618ef4ca9234a7371"></a>

<a id="canonical-8cdb7e28e45569510b06c0eaef7cb46e93dcd7e9eda01d7a1e215c701b13b0f8"></a>

## create property — timeouts / 4a02e8f4fb1a / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a91de8ae2858da1e820a38a797d2215aeb25645a91af3db750431f6fec085743"></a>

<a id="canonical-874ccbdee228ffdcf2c7dc902167f7e8d803bcd1670547627ed8d6db3fbb37ec"></a>

## delete property — timeouts / 4a02e8f4fb1a / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-321bea21a5809f5bca30411310bb5b06f073660124f76006f65d10cef49b14ef"></a>

<a id="canonical-bb38191c634f504a7da6075f8db5e7af70b16014a12aef1abc643f9380706767"></a>

## read property — timeouts / 4a02e8f4fb1a / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-fb43e42c0ece31f06c8ea65334cbb202d2a16cef058805bb99a306754b5b8fc7"></a>

<a id="canonical-18d74f74f73d66baa6f1f85e7c65c0f44e71024b58c712ba228731b5f17e5849"></a>

## update property — timeouts / 4a02e8f4fb1a / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-ada117af477ea1434b1982d33cafbcf0bb677c738aa5e625d6fcd00fd37fe078"></a>

## Next pages — timeouts / 4a02e8f4fb1a / 8

- [Property reference](resources--namespace--reference--group-001.md#canonical-09c5d1f37f639be8d72ab53f1082e4c1d53428b0ed5d288fca41f55d74dce8ac)
- [xcsh_namespace](../resources/namespace.md#canonical-5d4d79984896dc184c2d678bc72c561a372fa7660188fc55cde1c40e575c98b4)
