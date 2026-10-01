---
page_title: "xcsh_irule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule reference."
---

# xcsh_irule reference

<a id="canonical-4a843a98d007e27057ed2207e6455121fd9d1ff3073d3ba7dce67a1185a2a905"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d8699870777e06e4c04ed31cc9bb4e2db2e3d425ee79a1af1260c97117ae9ca"></a>

## Property reference — Property reference / 5afe32481ac9 / 2

Breadcrumbs:

- [xcsh_irule](../resources/irule.md#canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23)
- Property reference

<a id="canonical-26a96bc7c66c8fb6feed9d09b0f3dd4a0266cc98a79943d1c147e66c6aebd974"></a>

## Direct properties — Property reference / 5afe32481ac9 / 3

<a id="canonical-c673486e298bb0f7cdbd074ee410ff708fbaa3347a9d2ec0b8079360bcca8cec"></a>

<a id="canonical-675a4b06617810c9887bbca2bf686e24449aa949406b9a35300c847f3b2ddeee"></a>

## annotations property — Property reference / 5afe32481ac9 / 4

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

<a id="canonical-b8b600b7e21fd8ca0b9cb0fa5e74cfc22092aa22fb889ee69062bb73e7048b60"></a>

<a id="canonical-2787f480a00da53230385aaf766c441a08b1911941e75fb97c254ab5bb1bbb7f"></a>

## description property — Property reference / 5afe32481ac9 / 5

Type: `"string"`. Optional.

Human readable description for the object.

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

<a id="canonical-9137f0270a7c168245819f903e80a84c36dec23f005624205585f84dee611e7b"></a>

<a id="canonical-8128dc76845f49cdc2f63095ea8f573cd068b339cebeb4a0dde24aa74b42289f"></a>

## description_spec property — Property reference / 5afe32481ac9 / 6

Type: `"string"`. Required.

Description for iRule. Specify Description for iRule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-33c5a85c00ffa39497d42a46e723240fea8dbde8fe97a33a7a267f6ebbbe7c1b"></a>

<a id="canonical-ae8992748d8de1ab23c2b9a8cf1199e45eff057b4bc1520fbfdd21848700d892"></a>

## disable property — Property reference / 5afe32481ac9 / 7

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

<a id="canonical-251d32328725083966d913d525e0f81fe8952ecb5c85c3faddcaf7a549481faf"></a>

<a id="canonical-a4d34dce73d72afcf59a0c9e3a9a18bc5ab539f05f92808cc3e42df3abb7ee1c"></a>

## id property — Property reference / 5afe32481ac9 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-e1a0ef62ddd9d64411827b616a0a6c3f0d84cffe87ce10fb2941223ba6e10e3e"></a>

<a id="canonical-4a5b6193d154e8e66ddfe71ea932b8b947072eb9f6d3c6e5fa912313eb6549f6"></a>

## irule property — Property reference / 5afe32481ac9 / 9

Type: `"string"`. Required.

www&#46;internal.example.f5.com')\} DNS::drop\} irule content.

Upstream description:

www&#46;internal.example.f5.com")\} DNS::drop\} irule content.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(24576),
}
```

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

<a id="canonical-56dbb5e45a6b3991e34db2ee448a6ef1f19094384b4a7c0eff44a576f4ff92f9"></a>

<a id="canonical-4d54f047e4c3d9f88c9b7c07ee96dad7d465df46f9f7f7dece4de4fdd832bf53"></a>

## labels property — Property reference / 5afe32481ac9 / 10

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

<a id="canonical-d79afdfebe591d98eca6a522b4ba7f90035073d9c6a13a27ac08f3769bbb93e4"></a>

<a id="canonical-ab8e4c78f02842b90be0a48cfa3e53ebdd69a728077631d7144ad5560d2076bd"></a>

## name property — Property reference / 5afe32481ac9 / 11

Type: `"string"`. Required.

Name of the Irule. Must be unique within the namespace.

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

<a id="canonical-373f55d4f898a080937d3d77762d51deeb69a130272a999e5a582d2af33eedb7"></a>

<a id="canonical-49cfb136dbe4b8b298e9e6700d0a51f474a6b59c991f252b37f5018ef918dc2e"></a>

## namespace property — Property reference / 5afe32481ac9 / 12

Type: `"string"`. Required.

Namespace where the Irule is created.

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

- [timeouts](resources--irule--reference--group-001.md#canonical-0d5640696f2942248a862ccbf9dfbd347bd063781a8b24a4003391d454831526): complete subsection reference.

<a id="canonical-55059e7ebdc7cff4c45d1b99e06d1cf939d2d21cfd7f6431fd8816a31fedf662"></a>

## All schema paths — Property reference / 5afe32481ac9 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--irule--reference--group-001.md#canonical-c673486e298bb0f7cdbd074ee410ff708fbaa3347a9d2ec0b8079360bcca8cec) |
| `description` | [description](resources--irule--reference--group-001.md#canonical-b8b600b7e21fd8ca0b9cb0fa5e74cfc22092aa22fb889ee69062bb73e7048b60) |
| `description_spec` | [description_spec](resources--irule--reference--group-001.md#canonical-9137f0270a7c168245819f903e80a84c36dec23f005624205585f84dee611e7b) |
| `disable` | [disable](resources--irule--reference--group-001.md#canonical-33c5a85c00ffa39497d42a46e723240fea8dbde8fe97a33a7a267f6ebbbe7c1b) |
| `id` | [id](resources--irule--reference--group-001.md#canonical-251d32328725083966d913d525e0f81fe8952ecb5c85c3faddcaf7a549481faf) |
| `irule` | [irule](resources--irule--reference--group-001.md#canonical-e1a0ef62ddd9d64411827b616a0a6c3f0d84cffe87ce10fb2941223ba6e10e3e) |
| `labels` | [labels](resources--irule--reference--group-001.md#canonical-56dbb5e45a6b3991e34db2ee448a6ef1f19094384b4a7c0eff44a576f4ff92f9) |
| `name` | [name](resources--irule--reference--group-001.md#canonical-d79afdfebe591d98eca6a522b4ba7f90035073d9c6a13a27ac08f3769bbb93e4) |
| `namespace` | [namespace](resources--irule--reference--group-001.md#canonical-373f55d4f898a080937d3d77762d51deeb69a130272a999e5a582d2af33eedb7) |
| `timeouts` | [timeouts](resources--irule--reference--group-001.md#canonical-d0d36c726fa01fa074b7d754a1578f15148522226ec5f3a4f116adc3b1645804) |
| `timeouts.create` | [timeouts.create](resources--irule--reference--group-001.md#canonical-95e55608e1c147ec056399f7a0d5b4ab3f3f043b4c2651da9211a3f4e2b5e4bd) |
| `timeouts.delete` | [timeouts.delete](resources--irule--reference--group-001.md#canonical-8f39b8f0c88ce43adf4ae1906adbcd53ad3b2b6d8fd7bf560d5fc06c9b71c16b) |
| `timeouts.read` | [timeouts.read](resources--irule--reference--group-001.md#canonical-42b28b641cc0da51714ae0ae400a2dcd519d3c57fb64c675aa86136e3f66756b) |
| `timeouts.update` | [timeouts.update](resources--irule--reference--group-001.md#canonical-4852310f1dd05e03300748175d45eed3290623862afc5be50168bf0140dbc49b) |

<a id="canonical-40ccbbefcaf19995007afc93ced1ce7ce7649859b022f50164933562a283cf28"></a>

## Next pages — Property reference / 5afe32481ac9 / 14

- [timeouts](resources--irule--reference--group-001.md#canonical-0d5640696f2942248a862ccbf9dfbd347bd063781a8b24a4003391d454831526)
- [xcsh_irule](../resources/irule.md#canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23)

<a id="canonical-0d5640696f2942248a862ccbf9dfbd347bd063781a8b24a4003391d454831526"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73c7bc0a1a79bfa30a7baacfb696a362a3cf48dacd2d9831bcbf7bb5413f46b0"></a>

## timeouts — timeouts / ae1082760531 / 2

Breadcrumbs:

- [xcsh_irule](../resources/irule.md#canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23)
- [Property reference](resources--irule--reference--group-001.md#canonical-4a843a98d007e27057ed2207e6455121fd9d1ff3073d3ba7dce67a1185a2a905)
- timeouts

<a id="canonical-d0d36c726fa01fa074b7d754a1578f15148522226ec5f3a4f116adc3b1645804"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3ee0b0c0704a4c513d4f3a157988fbd3a5979a9ae484686f284f4472dfc7369"></a>

## Direct properties — timeouts / ae1082760531 / 3

<a id="canonical-95e55608e1c147ec056399f7a0d5b4ab3f3f043b4c2651da9211a3f4e2b5e4bd"></a>

<a id="canonical-edc67438b0becfe7e9a8a98e22ded62b49a56a8d104c150e1ea3574ee965ced1"></a>

## create property — timeouts / ae1082760531 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8f39b8f0c88ce43adf4ae1906adbcd53ad3b2b6d8fd7bf560d5fc06c9b71c16b"></a>

<a id="canonical-de61984a99ce27a3d721021a0a308933d647ebc1d8a6038dc6aeecb6e66150be"></a>

## delete property — timeouts / ae1082760531 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-42b28b641cc0da51714ae0ae400a2dcd519d3c57fb64c675aa86136e3f66756b"></a>

<a id="canonical-61e6b110cbb8dcaa6aa0a9fc346229cf8c62992372323762529f6cef40693d62"></a>

## read property — timeouts / ae1082760531 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-4852310f1dd05e03300748175d45eed3290623862afc5be50168bf0140dbc49b"></a>

<a id="canonical-da13aac07ea8a01e5c54a907b1f9c4dc500008bedb9af4b36f0ea9d2073dc867"></a>

## update property — timeouts / ae1082760531 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c3e4dafd7530976730b8be4f6d546db62096966c8f9996719106dedf257f26f5"></a>

## Next pages — timeouts / ae1082760531 / 8

- [Property reference](resources--irule--reference--group-001.md#canonical-4a843a98d007e27057ed2207e6455121fd9d1ff3073d3ba7dce67a1185a2a905)
- [xcsh_irule](../resources/irule.md#canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23)
