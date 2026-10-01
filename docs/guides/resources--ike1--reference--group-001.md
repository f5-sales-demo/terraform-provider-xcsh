---
page_title: "xcsh_ike1 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 reference."
---

# xcsh_ike1 reference

<a id="canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-156908b79475009e0215d5d1d47891102afe3372e6140627635f9140afe1a466"></a>

## Property reference — Property reference / 88b49a1b230d / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- Property reference

<a id="canonical-245efefa7c49ca9bcafb4dd8b1eed654952f274eae16b1e781f0d9bc3c46399a"></a>

## Direct properties — Property reference / 88b49a1b230d / 3

<a id="canonical-90ca01bf25cb374b482e6bad68d224c0213048249ecccd742098b2ff89fdcde8"></a>

<a id="canonical-0c93293af0d845b2388eff4ba39a4c96d4aa3c71460ad3a7ee4c1d648ed5d08a"></a>

## annotations property — Property reference / 88b49a1b230d / 4

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

<a id="canonical-c98a1371f11bdcec54f52f3988a86754315c0550aa2ed59b3e0680cf4927275c"></a>

<a id="canonical-7c5461ff774e2f89f24b27a17a026aca4224f2d2ba4e428d8872b195aab9cf38"></a>

## description property — Property reference / 88b49a1b230d / 5

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

<a id="canonical-3906ec49aa34bb14f1fe640bd5adcf151f08faac24d9e633a859f33a9653e4c7"></a>

<a id="canonical-efbd945d2edc1fcd3dc5f0bef44b64eb9ae378423e704c6c629db70a35b6c6fd"></a>

## disable property — Property reference / 88b49a1b230d / 6

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

<a id="canonical-b01d2b97a5b2a96d9b73c7583abd72e1d80534d65510281a3a31e4cb03976cad"></a>

<a id="canonical-a8c44ca625ceddc22fb31d1942e0fe3c9403b4c41fbc2e3ca84bbdfefd87da06"></a>

## id property — Property reference / 88b49a1b230d / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](resources--ike1--reference--group-001.md#canonical-7d4cbd24f3c1a52cf3506946e4d2f7fdcd715d708313e4c23d50186e62215dfb): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike1--reference--group-001.md#canonical-7d3798eda718e8530f3f853e8b2413a730428ded9c151391023f8656a24da8d9): complete subsection reference.

<a id="canonical-e3720a1a4826c6c119f9a75a4f33a906033e218f7b72af18b7668688d11b44c8"></a>

<a id="canonical-e20eacf652a7638956ca5a154774a3fb0e2b7ae94e57f8a3666aa7e358352a5b"></a>

## labels property — Property reference / 88b49a1b230d / 8

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

<a id="canonical-18142fdbcd5d73012b3b1f15473c9e16b1a42deac4b41555b7fa43e6d89c1b8f"></a>

<a id="canonical-9f4b129179a442236a33f647b39dff0d18cfa83fcb9b3a4f3281db0d284406f1"></a>

## name property — Property reference / 88b49a1b230d / 9

Type: `"string"`. Required.

Name of the Ike1. Must be unique within the namespace.

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

<a id="canonical-d367d40d6f6d2836a566562121a0751360beeaf765afc6f3718489463335dc00"></a>

<a id="canonical-ca60675c1f07f0fc07ba39e6fde2cf66cd35b8228d22423b3e4a414290c5f7fc"></a>

## namespace property — Property reference / 88b49a1b230d / 10

Type: `"string"`. Required.

Namespace where the Ike1 is created.

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

- [reauth_disabled](resources--ike1--reference--group-001.md#canonical-1c9c3252e94bdce1d8f2ceaf2112a3194837b6fa3e357b51cd7e30dafaeb13fb): complete subsection reference.

- [reauth_timeout_days](resources--ike1--reference--group-001.md#canonical-28ec9dd74f10070328d2534849806fb4014ac5e2411ee398446bd7cca6125243): complete subsection reference.

- [reauth_timeout_hours](resources--ike1--reference--group-001.md#canonical-785b57a6062fc7efc0291a0f76804955c40312ee74622fd9d57a0ac1512d5eab): complete subsection reference.

- [timeouts](resources--ike1--reference--group-001.md#canonical-7aa42090bf98365f40fe565d3ec676465ee43839225fccbd230199ab9e519b63): complete subsection reference.

- [use_default_keylifetime](resources--ike1--reference--group-001.md#canonical-569eb53763e9de12349109ca5b3ae50d5c7611171213f16c212e02de29dbde5a): complete subsection reference.

<a id="canonical-e5463a63c3694e9772fa9b24f1479b15b087aba99c3171d166c0f268db2ddc8c"></a>

## All schema paths — Property reference / 88b49a1b230d / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike1--reference--group-001.md#canonical-90ca01bf25cb374b482e6bad68d224c0213048249ecccd742098b2ff89fdcde8) |
| `description` | [description](resources--ike1--reference--group-001.md#canonical-c98a1371f11bdcec54f52f3988a86754315c0550aa2ed59b3e0680cf4927275c) |
| `disable` | [disable](resources--ike1--reference--group-001.md#canonical-3906ec49aa34bb14f1fe640bd5adcf151f08faac24d9e633a859f33a9653e4c7) |
| `id` | [id](resources--ike1--reference--group-001.md#canonical-b01d2b97a5b2a96d9b73c7583abd72e1d80534d65510281a3a31e4cb03976cad) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike1--reference--group-001.md#canonical-8838edc5a392ca3731dddfc5cc1c7088a2a7630524a8292ba39fda525568efe2) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike1--reference--group-001.md#canonical-617976262c296bfc7fba40ff9e97208294c3dd95149b76569a0e344be4e78833) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike1--reference--group-001.md#canonical-239ba7408ee6bf495aeba07a33077fc287982fb0eaedf7980e019ad8f2c00e8a) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike1--reference--group-001.md#canonical-9243fd72359fe6cf4975366650b04a9b4a642166ab5eb065528ac0567ceee362) |
| `labels` | [labels](resources--ike1--reference--group-001.md#canonical-e3720a1a4826c6c119f9a75a4f33a906033e218f7b72af18b7668688d11b44c8) |
| `name` | [name](resources--ike1--reference--group-001.md#canonical-18142fdbcd5d73012b3b1f15473c9e16b1a42deac4b41555b7fa43e6d89c1b8f) |
| `namespace` | [namespace](resources--ike1--reference--group-001.md#canonical-d367d40d6f6d2836a566562121a0751360beeaf765afc6f3718489463335dc00) |
| `reauth_disabled` | [reauth_disabled](resources--ike1--reference--group-001.md#canonical-5b902b2e80dfdd58ba7aa5e7057bb033b551fa791b72c8f23fd12ced2aeb083f) |
| `reauth_timeout_days` | [reauth_timeout_days](resources--ike1--reference--group-001.md#canonical-7cf486015e540508203c8da73536de393f1a6ef5e399999214976584b0072113) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](resources--ike1--reference--group-001.md#canonical-7a0414095a51c0b3c5bcdd55175737532fb3013f18953cb9fdf1be2dc2fb7c91) |
| `reauth_timeout_hours` | [reauth_timeout_hours](resources--ike1--reference--group-001.md#canonical-0176b749808280709e325eb9540bed4a73730637fb297dbbc63c42d80a72eea4) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](resources--ike1--reference--group-001.md#canonical-db1b9c74a252c04daae1cf639f34767fad792a1dafb1787d32a9c986aa24bf88) |
| `timeouts` | [timeouts](resources--ike1--reference--group-001.md#canonical-cd35e8026d0296afdf5936e82ba2a383184801628635e295a35ae14079383f9f) |
| `timeouts.create` | [timeouts.create](resources--ike1--reference--group-001.md#canonical-2760f86740612208cd56a772f6d4e36fd5911509d5c03067b185d2dcbb315e11) |
| `timeouts.delete` | [timeouts.delete](resources--ike1--reference--group-001.md#canonical-4ecc6afe86853b0885ac8c31746a51db83ad6e5a0396e81edd4837e080ee335e) |
| `timeouts.read` | [timeouts.read](resources--ike1--reference--group-001.md#canonical-698bfb16f00a5cb71547142357c404d0364674d4752e6c97f9fec898ff9c4218) |
| `timeouts.update` | [timeouts.update](resources--ike1--reference--group-001.md#canonical-5b78f63cc4b9adeb8fd82905e83c2202e1295eea46bf0b95211460c782b0efc0) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike1--reference--group-001.md#canonical-0b8d77af5b4630e948290c3f72b50f9e83515d8ac625179fac9ab98cde4c679b) |

<a id="canonical-10c4467616e4265281c0a2c06176e0941c395b568e28d3e834576242fb8d5bb8"></a>

## Next pages — Property reference / 88b49a1b230d / 12

- [ike_keylifetime_hours](resources--ike1--reference--group-001.md#canonical-7d4cbd24f3c1a52cf3506946e4d2f7fdcd715d708313e4c23d50186e62215dfb)
- [ike_keylifetime_minutes](resources--ike1--reference--group-001.md#canonical-7d3798eda718e8530f3f853e8b2413a730428ded9c151391023f8656a24da8d9)
- [reauth_disabled](resources--ike1--reference--group-001.md#canonical-1c9c3252e94bdce1d8f2ceaf2112a3194837b6fa3e357b51cd7e30dafaeb13fb)
- [reauth_timeout_days](resources--ike1--reference--group-001.md#canonical-28ec9dd74f10070328d2534849806fb4014ac5e2411ee398446bd7cca6125243)
- [reauth_timeout_hours](resources--ike1--reference--group-001.md#canonical-785b57a6062fc7efc0291a0f76804955c40312ee74622fd9d57a0ac1512d5eab)
- [timeouts](resources--ike1--reference--group-001.md#canonical-7aa42090bf98365f40fe565d3ec676465ee43839225fccbd230199ab9e519b63)
- [use_default_keylifetime](resources--ike1--reference--group-001.md#canonical-569eb53763e9de12349109ca5b3ae50d5c7611171213f16c212e02de29dbde5a)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)

<a id="canonical-7d4cbd24f3c1a52cf3506946e4d2f7fdcd715d708313e4c23d50186e62215dfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b950ed389989bd7af0a9ade7a0c71aa8b08720c36876bcaf36abfd23dc1cde11"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / bddc6831db2c / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- ike_keylifetime_hours

<a id="canonical-8838edc5a392ca3731dddfc5cc1c7088a2a7630524a8292ba39fda525568efe2"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
```

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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](resources--ike1--reference--group-001.md#canonical-8838edc5a392ca3731dddfc5cc1c7088a2a7630524a8292ba39fda525568efe2)
- [ike_keylifetime_minutes](resources--ike1--reference--group-001.md#canonical-239ba7408ee6bf495aeba07a33077fc287982fb0eaedf7980e019ad8f2c00e8a)
- [use_default_keylifetime](resources--ike1--reference--group-001.md#canonical-0b8d77af5b4630e948290c3f72b50f9e83515d8ac625179fac9ab98cde4c679b)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-83fa7dd3940aed60b4167a7177508f57f903846aca1f40db0501fd5126ff3984"></a>

## Direct properties — ike_keylifetime_hours / bddc6831db2c / 3

<a id="canonical-617976262c296bfc7fba40ff9e97208294c3dd95149b76569a0e344be4e78833"></a>

<a id="canonical-c6c81f5be300809e3b04b14c4008a6581db4c43ee50c87d5c337c661de3ae853"></a>

## duration property — ike_keylifetime_hours / bddc6831db2c / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-445ab6bfee2b540b1e8e4c8a1903a435a9bd032de3b5bfe0cdb283d39735e3c1"></a>

## Next pages — ike_keylifetime_hours / bddc6831db2c / 5

- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)

<a id="canonical-7d3798eda718e8530f3f853e8b2413a730428ded9c151391023f8656a24da8d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f85cfc12a12bc4937a4f40c04237318a99f83302e74ea34e4ebf2a27053fb6ef"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / ac34a2c116ca / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- ike_keylifetime_minutes

<a id="canonical-239ba7408ee6bf495aeba07a33077fc287982fb0eaedf7980e019ad8f2c00e8a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
```

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

Terraform syntax:

```terraform
ike_keylifetime_minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0e14a60d1b62447c8291f088a1fc250a779c3d3c2682930ad96e496154719a78"></a>

## Direct properties — ike_keylifetime_minutes / ac34a2c116ca / 3

<a id="canonical-9243fd72359fe6cf4975366650b04a9b4a642166ab5eb065528ac0567ceee362"></a>

<a id="canonical-22b2cffb68c3c57273a576c02b6c4471d295139dfabdfdbb7929fbf6827ed925"></a>

## duration property — ike_keylifetime_minutes / ac34a2c116ca / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(10, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-1e7ce8410ec27b9e0fe6f6b39edae6d5ea393df12cd4e1a98987bfac139dcd06"></a>

## Next pages — ike_keylifetime_minutes / ac34a2c116ca / 5

- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)

<a id="canonical-1c9c3252e94bdce1d8f2ceaf2112a3194837b6fa3e357b51cd7e30dafaeb13fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54693daee5d718f2398431f705ab0a886e7f57c43c29bbba676f520401fbca74"></a>

## reauth_disabled — reauth_disabled / 3113e77de78d / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- reauth_disabled

<a id="canonical-5b902b2e80dfdd58ba7aa5e7057bb033b551fa791b72c8f23fd12ced2aeb083f"></a>

Type: `["object", {}]`. Optional.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [reauth_disabled](resources--ike1--reference--group-001.md#canonical-5b902b2e80dfdd58ba7aa5e7057bb033b551fa791b72c8f23fd12ced2aeb083f)
- [reauth_timeout_days](resources--ike1--reference--group-001.md#canonical-7cf486015e540508203c8da73536de393f1a6ef5e399999214976584b0072113)
- [reauth_timeout_hours](resources--ike1--reference--group-001.md#canonical-0176b749808280709e325eb9540bed4a73730637fb297dbbc63c42d80a72eea4)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
reauth_disabled = {}
```

<a id="canonical-69f1bf6e04fc94b7be0fb0003bd53991c4491efacd8837b200b06a26a95fa3fc"></a>

## Direct properties — reauth_disabled / 3113e77de78d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53ae5e9912bae534f9321a38461c1adb5f2f9322fc85d159cea708ddc312f43e"></a>

## Next pages — reauth_disabled / 3113e77de78d / 4

- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)

<a id="canonical-28ec9dd74f10070328d2534849806fb4014ac5e2411ee398446bd7cca6125243"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdba17be6de0c54fade501d84efeac9da200c65a03f979453cb148ce7ccc0103"></a>

## reauth_timeout_days — reauth_timeout_days / 00e1a5f67110 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- reauth_timeout_days

<a id="canonical-7cf486015e540508203c8da73536de393f1a6ef5e399999214976584b0072113"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
```

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

Terraform syntax:

```terraform
reauth_timeout_days {
  # Configure direct properties listed below.
}
```

<a id="canonical-ecefc264dcf3209e6985bd0a7aeba77a40180e05073e3d8f9311a26730fb274f"></a>

## Direct properties — reauth_timeout_days / 00e1a5f67110 / 3

<a id="canonical-7a0414095a51c0b3c5bcdd55175737532fb3013f18953cb9fdf1be2dc2fb7c91"></a>

<a id="canonical-066f212a52c375cd2fee2c450e0fc357e7cfee8072c88545bf5a345573dbc06e"></a>

## duration property — reauth_timeout_days / 00e1a5f67110 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-8582287680522eaff2e16de16e977779d9fb24605aa3f8ce66de883a744971e9"></a>

## Next pages — reauth_timeout_days / 00e1a5f67110 / 5

- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)

<a id="canonical-785b57a6062fc7efc0291a0f76804955c40312ee74622fd9d57a0ac1512d5eab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-965adf6bcd1d1dfa1531611e51d757cad512fc6df32c9a15d3f035f477ae5f9a"></a>

## reauth_timeout_hours — reauth_timeout_hours / 6c349afce735 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- reauth_timeout_hours

<a id="canonical-0176b749808280709e325eb9540bed4a73730637fb297dbbc63c42d80a72eea4"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout hours.

Upstream description:

Input Hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
```

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

Terraform syntax:

```terraform
reauth_timeout_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-55be485d7713289ea585b7c419044b16be83bbb0901898c755393367ae2bc779"></a>

## Direct properties — reauth_timeout_hours / 6c349afce735 / 3

<a id="canonical-db1b9c74a252c04daae1cf639f34767fad792a1dafb1787d32a9c986aa24bf88"></a>

<a id="canonical-9830f3b928332d91ee67898ae637e7e726a7e8197383e4c5622c9e0c99fd09c9"></a>

## duration property — reauth_timeout_hours / 6c349afce735 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-09451b5fee2f34c077217bdf40e7a4e7db16dce4754dc0711ddac16e02ca0e5e"></a>

## Next pages — reauth_timeout_hours / 6c349afce735 / 5

- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)

<a id="canonical-7aa42090bf98365f40fe565d3ec676465ee43839225fccbd230199ab9e519b63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b43aeb00c29ca7a788cd8b79d38b759252aff71042a291f940cd193b4ef5002c"></a>

## timeouts — timeouts / bbf5a5932e54 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- timeouts

<a id="canonical-cd35e8026d0296afdf5936e82ba2a383184801628635e295a35ae14079383f9f"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7ec16f9f7f698559f68649ba85cbea916d354982ef912d09ab2cd6f26f603b3"></a>

## Direct properties — timeouts / bbf5a5932e54 / 3

<a id="canonical-2760f86740612208cd56a772f6d4e36fd5911509d5c03067b185d2dcbb315e11"></a>

<a id="canonical-27ae2188725a4698b6b38dc5fae018b1b6a1e4b91b320a1d4d34150363b9b162"></a>

## create property — timeouts / bbf5a5932e54 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-4ecc6afe86853b0885ac8c31746a51db83ad6e5a0396e81edd4837e080ee335e"></a>

<a id="canonical-32a3de60cafca465e3e4878033b1654f84ace0311dff938a874068bef7dbd888"></a>

## delete property — timeouts / bbf5a5932e54 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-698bfb16f00a5cb71547142357c404d0364674d4752e6c97f9fec898ff9c4218"></a>

<a id="canonical-903d14994ff8078a56509c5c2a2f7ac73c54fbb6e9afcc47d4c296e81c550ed0"></a>

## read property — timeouts / bbf5a5932e54 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-5b78f63cc4b9adeb8fd82905e83c2202e1295eea46bf0b95211460c782b0efc0"></a>

<a id="canonical-7dda23e76e5449b853447fcea8869808b5a1320a80aaf44f3706d2fbe035d5de"></a>

## update property — timeouts / bbf5a5932e54 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-61e1c4689a1c32936b6b6143ed511bb52f9644145d337d49b54cebb4b4c097ff"></a>

## Next pages — timeouts / bbf5a5932e54 / 8

- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)

<a id="canonical-569eb53763e9de12349109ca5b3ae50d5c7611171213f16c212e02de29dbde5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20ae053914361694aa32911157b02272edff1fc7f6d69d11f45e8a528da57e37"></a>

## use_default_keylifetime — use_default_keylifetime / 0fe26c0c932d / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- use_default_keylifetime

<a id="canonical-0b8d77af5b4630e948290c3f72b50f9e83515d8ac625179fac9ab98cde4c679b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use default keylifetime.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
use_default_keylifetime = {}
```

<a id="canonical-b636a6b1e1db97e0234f2a3880dd48f2e4cddc436f85dd9ceda4f6ab5181ad2c"></a>

## Direct properties — use_default_keylifetime / 0fe26c0c932d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54a4038e9c210f87d0216e4385c2f716556a807138edb9fc3b4bd3561b6db3b8"></a>

## Next pages — use_default_keylifetime / 0fe26c0c932d / 4

- [Property reference](resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
